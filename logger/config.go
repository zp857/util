package logger

// Config 配置
type Config struct {
	Director     string `json:"director" yaml:"director"`
	Format       string `json:"format" yaml:"format"`
	Level        string `json:"level" yaml:"level"`
	MaxAge       int    `json:"maxAge" yaml:"maxAge"`
	Compress     bool   `json:"compress" yaml:"compress"`
	LogInConsole bool   `json:"logInConsole" yaml:"logInConsole"`
}

var (
	DefaultConfig = &Config{
		Director:     "logs",
		Format:       "console",
		Level:        "debug",
		MaxAge:       30,
		Compress:     true,
		LogInConsole: true,
	}
)
