package sftputil

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"path"
	"path/filepath"
	"time"
)

type Options struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

type Client struct {
	sftpClient *sftp.Client
}

func NewClient(options Options) (client *Client, err error) {
	config := &ssh.ClientConfig{
		User:            options.Username,
		Auth:            []ssh.AuthMethod{ssh.Password(options.Password)},
		Timeout:         10 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	// connect to ssh
	addr := fmt.Sprintf("%v:%v", options.Host, options.Port)
	conn, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return
	}
	sftpClient, err := sftp.NewClient(conn)
	if err != nil {
		return
	}
	client = &Client{
		sftpClient: sftpClient,
	}
	return
}

func (c *Client) Upload(localFile, remoteFile string) (err error) {
	srcFile, err := os.Open(localFile)
	if err != nil {
		return
	}
	defer srcFile.Close()

	remoteDir := filepath.Dir(remoteFile)
	remoteDir = filepath.ToSlash(remoteDir)
	err = c.sftpClient.MkdirAll(remoteDir)
	if err != nil {
		return
	}

	dstFile, err := c.sftpClient.Create(remoteFile)
	if err != nil {
		return
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return
}

func (c *Client) UploadDirectory(localDir, remoteDir string) (err error) {
	localFiles, err := os.ReadDir(localDir)
	if err != nil {
		return
	}

	for _, dir := range localFiles {
		localFilePath := path.Join(localDir, dir.Name())
		remoteFilePath := path.Join(remoteDir, dir.Name())
		remoteFilePath = filepath.ToSlash(remoteFilePath)
		if dir.IsDir() {
			err = c.sftpClient.MkdirAll(remoteFilePath)
			if err != nil {
				return
			}
			err = c.Upload(localFilePath, remoteFilePath)
			if err != nil {
				return
			}
		} else {
			err = c.Upload(path.Join(localDir, dir.Name()), remoteFilePath)
			if err != nil {
				return
			}
		}
	}
	return
}
