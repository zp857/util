package dialer

import (
	"strings"

	"github.com/zp857/util/convert"

	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
)

// SSHCli ssh 连接结构体
type SSHCli struct {
	client *goph.Client
	ip     string
}

func (s *SSHCli) GetIP() string {
	return s.ip
}

func (s *SSHCli) CreateClient(cli *BaseCli) (err error) {
	var auth goph.Auth
	if cli.PrivateKey != "" {
		var signer ssh.Signer
		signer, err = ssh.ParsePrivateKey([]byte(cli.PrivateKey))
		if err != nil {
			return
		}
		auth = goph.Auth{
			ssh.PublicKeys(signer),
		}
	} else {
		auth = goph.Password(strings.TrimSpace(cli.Password))
	}
	var client *goph.Client
	client, err = goph.NewConn(&goph.Config{
		User:     cli.User,
		Addr:     cli.IP,
		Port:     uint(cli.Port),
		Auth:     auth,
		Callback: ssh.InsecureIgnoreHostKey(),
		Timeout:  timeout,
	})
	if err != nil {
		return
	}
	s.client = client
	s.ip = cli.IP
	return
}

func (s *SSHCli) ExecCmd(cmd string, charsetOptional ...string) (result string, err error) {
	var (
		session *ssh.Session
		buf     []byte
	)
	charset := "UTF-8"
	if len(charsetOptional) > 0 {
		charset = charsetOptional[0]
	}
	if session, err = s.client.NewSession(); err != nil {
		return
	}
	defer session.Close()
	if buf, err = session.CombinedOutput(cmd); err != nil {
		return
	}

	result = convert.Byte2String(buf, convert.Charset(charset))

	return
}
