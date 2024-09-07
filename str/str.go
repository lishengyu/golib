package str

import (
	"encoding/base64"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

func CheckNull(str string) bool {
	if str == "" {
		return false
	}
	return true
}

func CheckPort(str string) bool {
	if str == "" {
		return false
	}

	num, err := strconv.Atoi(str)
	if err != nil {
		log.Printf("strconv.Atoi(%s) failed: %v\n", str, err)
		return false
	}

	if num < 1 && num > 65535 {
		return false
	}

	return true
}

func CheckIp(str string) bool {
	if str == "" {
		return false
	}

	if net.ParseIP(str) == nil {
		log.Printf("net.ParseIP(%s) failed\n", str)
		return false
	}

	return true
}

func CheckTime(str string) bool {
	if str == "" {
		return false
	}

	_, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		log.Printf("time.Parse(%s) failed: %v\n", str, err)
		return false
	}

	return true
}

func CheckInt64(str string) bool {
	if str == "" {
		return false
	}

	_, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Printf("strconv.ParseInt(%s) failed: %v\n", str, err)
		return false
	}

	return true
}

func CheckBase64(str string) bool {
	if str == "" {
		return false
	}

	fields := strings.Split(str, ",")
	for _, s := range fields {
		_, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			log.Printf("base64.StdEncoding.DecodeString(%s) failed: %v\n", s, err)
			return false
		}
	}

	return true
}
