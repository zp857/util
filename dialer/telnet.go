package dialer

import (
	"fmt"
	"github.com/zp857/util/dialer/telnet"
	"strings"
)

// TelnetCli telnet 连接结构体
type TelnetCli struct {
	client *telnet.Conn
	user   string
	ip     string
}

func (t *TelnetCli) GetIP() string {
	return t.ip
}

func (t *TelnetCli) CreateClient(cli *BaseCli) (err error) {
	conn, err := telnet.DialTimeout("tcp", fmt.Sprintf("%v:%v", cli.IP, cli.Port), timeout)
	if err != nil {
		return
	}
	if err = telnet.Login(conn, cli.User, cli.Password); err == nil {
		t.client = conn
		t.user = cli.User
		t.ip = cli.IP
	} else {
		err = fmt.Errorf("登录失败")
	}
	return
}

func (t *TelnetCli) ExecCmd(cmd string, charsetOptional ...string) (result string, err error) {
	telnet.Send(t.client, cmd)
	result, err = telnet.Read(t.client, []string{t.user + "@"})
	result = strings.TrimPrefix(result, cmd)
	result = strings.TrimSpace(result)
	resultLines := strings.Split(result, "\n")
	result = strings.Join(resultLines[:len(resultLines)-1], "\n")
	return
}
