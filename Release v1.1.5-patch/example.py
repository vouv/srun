# 构建srun_source_code.py时的参考代码
#!/usr/bin/env python3
# -*-coding:utf-8-*-

import base64
import hashlib
import hmac
import json
import os
import time
import re
import requests


def char_code_at(stri, index):
    return 0 if index >= len(stri) else ord(stri[index])


def xencode(msg: str, key):
    '''
    xEncode
    '''
    def s(a: str, b: bool):
        c = len(a)
        v = []
        for i in range(0, c, 4):
            v.append(char_code_at(a, i) | (char_code_at(a, i+1) << 8) |
                     (char_code_at(a, i+2) << 16) | (char_code_at(a, i+3) << 24))
        if b:
            v.append(c)
        return v

    def l(a, b):
        d = len(a)
        c = (d-1) << 2
        if b:
            m = a[d-1]
            if (m < c-3) or (m > c):
                return None
            c = m
        for i in range(0, d):
            a[i] = ''.join([chr(a[i] & 0xff), chr((a[i] >> 8) & 0xff), chr(
                (a[i] >> 16) & 0xff), chr((a[i] >> 24) & 0xff)])
        if b:
            return (''.join(a))[0:c]
        else:
            return ''.join(a)

    if msg == "":
        return ""
    v = s(msg, True)
    k = s(key, False)
    # print(v)
    # print(k)
    n = len(v) - 1
    z = v[n]
    y = v[0]
    c = 0x86014019 | 0x183639A0
    m = 0
    e = 0
    p = 0
    q = 6 + 52 // (n + 1)
    d = 0
    while 0 < q:
        q -= 1
        d = d + c & (0x8CE0D9BF | 0x731F2640)
        e = d >> 2 & 3
        for p in range(0, n):
            y = v[p+1]
            m = z >> 5 ^ y << 2
            m += (y >> 3 ^ z << 4) ^ (d ^ y)
            m += k[(p & 3) ^ e] ^ z
            z = v[p] = v[p] + m & (0xEFB8D130 | 0x10472ECF)
        y = v[0]
        m = z >> 5 ^ y << 2
        m += (y >> 3 ^ z << 4) ^ (d ^ y)
        m += k[(n & 3) ^ e] ^ z
        z = v[n] = v[n] + m & (0xBB390742 | 0x44C6F8BD)
    # print(v)
    return l(v, False)


def get_json(url, data):
    '''Http GET, return json
    '''
    callback = "jsonp%s" % int(time.time()*1000)
    data["callback"] = callback

    response = requests.get(url, data)
    response_content = response.content.decode('utf-8')[len(callback)+1:-1]
    response_json = json.loads(response_content)

    return response_json


def srun_login(username, password=None, action='login'):
    '''srun login and logout
    Args:
        username: username
        password: password
        action: 'login' or 'logout'
    Returns:
        a json object.
    '''
    def data_info(get_data, token):
        if get_data['action'] == 'login':
            x_encode_json = {
                "username": get_data['username'],
                "password": get_data['password'],
                "ip": get_data['ip'],
                "acid": get_data['ac_id'],
                "enc_ver": enc
            }
        else:
            x_encode_json = {
                "username": get_data['username'],
                "ip": get_data['ip'],
                "acid": get_data['ac_id'],
                "enc_ver": enc
            }

        x_encode_str = json.dumps(x_encode_json, separators=(',', ':'))
        x_encode_key = token
        x_encode_res = xencode(x_encode_str, x_encode_key)
        # print("x_encode('%s', '%s')" % (x_encode_str, x_encode_key))
        # print('x_encode_res(len: %s): %s' % (len(x_encode_res), x_encode_res))
        # print("x_encode_res unicode:", [ord(s) for s in x_encode_res])

        # base64_encode
        mapping = dict(zip("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=",
                           "LVoJPiCN2R8G90yg+hmFHuacZ1OWMnrsSTXkYpUq/3dlbfKwv6xztjI7DeBE45QA="))
        b64_res = base64.b64encode(
            bytes([ord(s) for s in x_encode_res])).decode()
        base64_encode_res = ''.join([mapping[b] for b in b64_res])
        # print('base64 encode res(len: %s): %s' % (len(base64_encode_res), base64_encode_res))

        return "{SRBX1}" + base64_encode_res

    def pwd_hmd5(password, token):
        hmac_key = token.encode('utf-8')
        hmac_msg = password.encode('utf-8')
        hmd5 = hmac.new(hmac_key, hmac_msg, digestmod='MD5').hexdigest()
        # print(hmd5)
        return '{MD5}' + hmd5

    def checksum(get_data, token):
        if get_data['action'] == 'login':
            str_list = ['', get_data['username'], get_data['password'][5:],
                        get_data['ac_id'], get_data['ip'], str(n), str(type_), get_data['info']]
        else:
            str_list = ['', get_data['username'], get_data['ac_id'],
                        get_data['ip'], str(n), str(type_), get_data['info']]
        chksum_str = token.join(str_list)
        chksum = hashlib.sha1(chksum_str.encode('utf-8')).hexdigest()
        return chksum

    if action not in ['login', 'logout']:
        print('action must be "login" or "logout".')
        return
    enc = "srun_bx1"
    n = 200
    type_ = 1
    get_challenge_url = "http://10.0.0.55/cgi-bin/get_challenge"
    srun_portal_url = "http://10.0.0.55/cgi-bin/srun_portal"
    url = 'http://detectportal.firefox.com/success.txt'
    #Check if Redirect, when not, set to default
    try:
        r = requests.get(url, timeout=0.1)
        ac_id=re.findall(r'index_(\d*).html',r.url)[0]
    except requests.exceptions.Timeout:
            ac_id=1
    if not ac_id:
        ac_id=1
    ac_id=str(ac_id)
    if action == 'login':
        get_data = {
            "action": action,
            "username": username,
            "password": password,
            "ac_id": ac_id,
            "ip": '',
            "info": '',
            "chksum": '',
            "n": n,
            "type": type_
        }
    else:
        get_data = {
            "action": action,
            "username": username,
            # "password": password, # logout,
            "ac_id": ac_id,
            "ip": '',
            "info": '',
            "chksum": '',
            "n": n,
            "type": type_
        }
    # get token
    challenge_json = get_json(
        get_challenge_url, {"username": get_data['username']})
    if challenge_json['res'] != "ok":
        print('Error getting challenge. %s failed.' % action)
        print('Server response:\n%s' % json.dumps(challenge_json, indent=4))
        return
    token = challenge_json['challenge']
    get_data['ip'] = challenge_json['client_ip']
    get_data['info'] = data_info(get_data, token)

    if action == 'login':
        get_data['password'] = pwd_hmd5('', token)
        # get_data['password'] = pwd_hmd5(get_data['password'], token) # srun's bug

    get_data['chksum'] = checksum(get_data, token)

    # print('get data: %s' % json.dumps(get_data, indent=4))
    res = get_json(srun_portal_url, get_data)
    # print("Server response: %s" % json.dumps(res, indent=4))

    if res['error'] == 'ok':
        print('%s success.' % action)
    else:
        print("%s failed.\n%s %s" % (action, res['error'], res['error_msg']))

    return res


if __name__ == "__main__":
    username = "你的校园网账号"
    password = "你的校园网密码"

    srun_login(username, password)
#srun_login(username, action="logout")
