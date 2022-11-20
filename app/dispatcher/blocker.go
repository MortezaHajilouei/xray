package dispatcher

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

var blockedIpList string
var hasBlockedIp bool

func getOsArgValue(s []string, flags ...string) string {
	for i, v := range s {
		for _, flagVal := range flags {
			if v == flagVal {
				return s[i+1]
			}
		}
	}

	return ""
}

func getBanFilePath() string {
	fileName := getOsArgValue(os.Args, "-ban")

	if len(fileName) == 0 {
		return ""
	}

	dirBan, err := os.Getwd()
	if err != nil {
		return ""
	}

	banFilePath := filepath.Join(dirBan, fileName)
	return banFilePath
}

func readFileString(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		return ""
	}

	return string(b)
}

func readDB() {
	path := getBanFilePath()
	if len(path) <= 0 {
		return
	}

	for {
		blockedIpList = readFileString(path)
		if len(blockedIpList) > 0 {
			hasBlockedIp = true
		} else {
			hasBlockedIp = false
		}
		time.Sleep(30 * time.Second)
	}
}

func checkIP(ip string) bool {
	if hasBlockedIp {
		if strings.Contains(blockedIpList, ip) {
			return true
		}
	}
	return false
}

func initBlocker() {
	go readDB()
}
