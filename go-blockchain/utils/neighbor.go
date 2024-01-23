package utils

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
)

func IsFoundHost(guessHost string, port uint16) bool {

	return true
}

var PATTERN = regexp.MustCompile(`((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?\.){3})(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`)

func FindNeighbors(myHost string, myPort uint16, startIp uint8, endIp uint8, startPort uint16, endPort uint16) []string {
	address := fmt.Sprintf("%s:%d", myHost, myPort)

	m := PATTERN.FindStringSubmatch(myHost)
	if m == nil {
		return nil
	}
	prefixHost := m[1]
	lastIp, _ := strconv.Atoi(m[len(m)-1])
	neighbors := make([]string, 0)

	for port := startPort; port <= endPort; port += 1 {
		for ip := startIp; ip <= endIp; ip += 1 {
			guessHost := fmt.Sprintf("%s%d", prefixHost, lastIp+int(ip))
			guessTarget := fmt.Sprintf("%s:%d", guessHost, port)
			if guessTarget != address && IsFoundHost(guessHost, port) {
				neighbors = append(neighbors, guessTarget)
			}
		}
	}
	return neighbors
}

func GetHost() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "127.0.0.1"
	}
	address, err := net.LookupHost(hostname)
	if err != nil {
		return "127.0.0.1"
	}
	return address[0]
}

// TODO: Get this code to work ... 4:06:00 was last time working
