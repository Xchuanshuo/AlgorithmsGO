package test468

import (
	"strconv"
	"strings"
)

func validIPAddress(queryIP string) string {
	if strs := strings.Split(queryIP, "."); len(strs) == 4 {
		for _, s := range strs {
			if len(s) == 0  || (len(s) > 1 && s[0] == '0') {
				return "Neither"
			}
			if v, err := strconv.Atoi(s); err != nil || v < 0 || v > 255 {
				return "Neither"
			}
		}
		return "IPv4"
	}
	if strs := strings.Split(queryIP, ":"); len(strs) == 8 {
		for _, s := range strs {
			if len(s) > 4 || len(s) == 0 {
				return "Neither"
			}
			v, err := strconv.ParseInt(s, 16, 64)
			if err != nil || v < 0 || v > 0xffff {
				return "Neither"
			}
		}
		return "IPv6"
	}
	return "Neither"
}