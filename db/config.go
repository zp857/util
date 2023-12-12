package db

type Config struct {
	DBType    string `yaml:"dbType" json:"dbType"`
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

type GeneralDB struct {
	Path         string `yaml:"path" json:"path"`                      // 服务器地址
	Port         string `yaml:"port" json:"port"`                      // 端口
	Config       string `yaml:"config" json:"config"`                  // 高级配置
	DBName       string `yaml:"dbName" json:"dbName"`                  // 数据库名
	Username     string `yaml:"username" json:"username"`              // 数据库用户名
	Password     string `yaml:"password" json:"password"`              // 数据库密码
	Prefix       string `yaml:"prefix" json:"prefix"`                  // 全局表前缀，单独定义TableName则不生效
	Singular     bool   `yaml:"singular" json:"singular"`              // 是否开启全局禁用复数，true表示开启
	Engine       string `yaml:"engine" json:"engine" default:"InnoDB"` // 数据库引擎，默认InnoDB
	MaxIdleConns int    `yaml:"maxIdleConns" json:"maxIdleConns"`      // 空闲中的最大连接数
	MaxOpenConns int    `yaml:"maxOpenConns" json:"maxOpenConns"`      // 打开到数据库的最大连接数
	LogMode      string `yaml:"logMode" json:"logMode"`                // 是否开启Gorm全局日志
	LogZap       bool   `yaml:"logZap" json:"logZap"`                  // 是否通过zap写入日志文件
}

func (c *Config) GetLogMode() string {
	return c.LogMode
}

func (c *Config) DSN() string {
	switch c.DBType {
	case "mysql":
		return c.Username + ":" + c.Password + "@tcp(" + c.Path + ":" + c.Port + ")/" + c.DBName + "?" + c.Config
	case "pgsql":
		return "host=" + c.Path + " user=" + c.Username + " password=" + c.Password + " dbname=" + c.DBName + " port=" + c.Port + " " + c.Config
	default:
		return ""
	}
}
