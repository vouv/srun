#!/bin/bash
while true; do
  ./srun login
  sleep 3600  # 每一小时执行一次
done
