package db

import (
	"go.uber.org/zap"
)

const (
	InitDataExist   = "%v 表的初始数据已存在"
	InitDataFailed  = "%v 表初始化数据失败: %v"
	InitDataSuccess = "%v 表初始化数据成功"
)

type InitData interface {
	TableName() string
	Initialize() (err error)
	CheckDataExist() bool
}

func InitTableData(inits ...InitData) error {
	logger := zap.L().Sugar()
	for i := 0; i < len(inits); i++ {
		if inits[i].CheckDataExist() {
			logger.Infof(InitDataExist, inits[i].TableName())
			continue
		}
		if err := inits[i].Initialize(); err != nil {
			logger.Infof(InitDataFailed, inits[i].TableName(), err)
		}
		logger.Infof(InitDataSuccess, inits[i].TableName())
	}
	return nil
}
