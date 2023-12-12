package telnet

import (
	"strings"
)

func Login(conn *Conn, user, password string) (err error) {
	_, err = Read(conn, []string{"login: "})
	if err != nil {
		return
	}
	Send(conn, user)
	_, err = Read(conn, []string{"Password: "})
	if err != nil {
		return
	}
	Send(conn, password)
	_, err = Read(conn, []string{user + "@"})
	return
}

func Read(conn *Conn, stopWords []string) (result string, err error) {
	for {
		var buf = make([]byte, 4096)
		var n int
		n, err = conn.Read(buf[0:])
		if err != nil {
			break
		}
		result += string(buf[:n])
		for _, word := range stopWords {
			if strings.Contains(result, word) {
				return
			}
		}
	}
	return
}

func Send(conn *Conn, s string) {
	buf := make([]byte, len(s)+1)
	copy(buf, s)
	buf[len(s)] = '\n'
	_, err := conn.Write(buf)
	if err != nil {
		return
	}
}
