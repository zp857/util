package kafka

type Config struct {
	Consumer struct {
		Urls []string `yaml:"urls" json:"urls"`
	} `yaml:"consumer" json:"consumer"`
	Producer struct {
		Urls []string `yaml:"urls" json:"urls"`
	} `yaml:"producer" json:"producer"`
}
