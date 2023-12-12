package dialer

import "time"

// BaseCli 基础客户端机构体
type BaseCli struct {
	IP         string // IP地址
	Port       int    // 端口
	User       string // 账号
	Password   string // 密码
	PrivateKey string // ssh 私钥
}

// Conninter 登录连接接口
type Conninter interface {
	CreateClient(cli *BaseCli) (err error)
	ExecCmd(cmd string, charsetOptional ...string) (string, error)
	GetIP() string
}

var (
	portMap = map[string]string{
		"ssh":    "22",   // ssh 协议默认端口
		"telnet": "23",   // telnet 协议默认端口
		"winrm":  "5985", // winrm 协议默认端口
	}

	conninterMap = map[string]Conninter{
		"ssh":    &SSHCli{},
		"telnet": &TelnetCli{},
		"winrm":  &WinrmCli{},
	}

	timeout = 15 * time.Second
)

func GetConninter(connectType string) (conn Conninter) {
	switch connectType {
	case "ssh":
		conn = conninterMap["ssh"]
	case "telnet":
		conn = conninterMap["telnet"]
	case "winrm":
		conn = conninterMap["winrm"]
	default:
		conn = conninterMap["ssh"]
	}
	return conn
}

func GetPort(connectType string) (port string) {
	switch connectType {
	case "ssh":
		port = portMap["ssh"]
	case "telnet":
		port = portMap["telnet"]
	case "winrm":
		port = portMap["winrm"]
	default:
		port = portMap["ssh"]
	}
	return
}
