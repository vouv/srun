package utils

import (
	"fmt"
	"strings"
)

func FormatFlux(byte int64) string {
	const tb = 1024 * 1024 * 1024 * 1024
	const gb = tb / 1024
	const mb = gb / 1024
	const kb = 1024
	if byte > tb {
		return fmt.Sprintf("%.2fTB", float64(byte)/tb)
	}
	if byte > gb {
		return fmt.Sprintf("%.2fGB", float64(byte)/gb)
	}
	if byte > mb {
		return fmt.Sprintf("%.1fMB", float64(byte)/mb)
	}
	if byte > kb {
		return fmt.Sprintf("%dKB", byte/kb)
	}
	return fmt.Sprintf("%dB", byte)
}

func FormatTime(sec int64) string {
	h := sec / 3600
	sec %= 3600
	m := sec / 60
	sec %= 60
	s := sec
	out := strings.Builder{}
	if h < 10 {
		out.WriteString(fmt.Sprint("0", h, "时"))
	} else {
		out.WriteString(fmt.Sprint(h, "时"))
	}
	if m < 10 {
		out.WriteString(fmt.Sprint("0", m, "分"))
	} else {
		out.WriteString(fmt.Sprint(m, "分"))
	}
	if s < 10 {
		out.WriteString(fmt.Sprint("0", s, "秒"))
	} else {
		out.WriteString(fmt.Sprint(s, "秒"))
	}
	return out.String()
}
