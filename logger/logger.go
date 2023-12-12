package logger

import (
	"github.com/zp857/util/fileutil"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

func Init(config *Config) {
	if ok := fileutil.PathExists(config.Director); !ok {
		_ = os.Mkdir(config.Director, os.ModePerm)
	}
	// 获取日志写入位置
	writeSyncer, err := getLogWriter(config.Director, config.MaxAge, config.LogInConsole)
	if err != nil {
		log.Fatal("日志级别配置异常")
	}
	// 获取日志编码格式
	encoder := getEncoder(config.Format)
	// 获取日志最低等级, 即>=该等级, 才会被写入
	var l = new(zapcore.Level)
	if err := l.UnmarshalText([]byte(config.Level)); err != nil {
		log.Fatal("日志级别配置异常")
	}
	// 创建一个将日志写入 WriteSyncer 的核心
	core := zapcore.NewCore(encoder, writeSyncer, l)
	logger := zap.New(core, zap.AddCaller())
	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(logger)
}
