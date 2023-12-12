package dialer

import (
	"context"
	"strings"

	"github.com/masterzen/winrm"
	"github.com/zp857/util/convert"
)

type WinrmCli struct {
	client *winrm.Client
	ip     string
}

func (w *WinrmCli) GetIP() string {
	return w.ip
}

func (w *WinrmCli) CreateClient(cli *BaseCli) (err error) {
	endpoint := winrm.NewEndpoint(cli.IP, cli.Port, false, false, nil, nil, nil, timeout)
	var client *winrm.Client
	client, err = winrm.NewClient(endpoint, cli.User, cli.Password)
	if err != nil {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var code int
	_, _, code, err = client.RunWithContextWithString(ctx, "echo ok", "")
	if err != nil && code != 0 {
		return
	}
	w.client = client
	w.ip = cli.IP
	return
}

func (w *WinrmCli) ExecCmd(cmd string, charsetOptional ...string) (string, error) {
	charset := "UTF-8"
	if len(charsetOptional) > 0 {
		charset = charsetOptional[0]
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	output, outputerr, _, err := w.client.RunWithContextWithString(ctx, cmd, "")

	if err != nil {
		output += outputerr
	}
	output = convert.Byte2String([]byte(output), convert.Charset(charset))
	return strings.TrimSpace(output), err
}
