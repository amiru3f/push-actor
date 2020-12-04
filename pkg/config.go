package pkg

type Config struct {
	RabbitConfig RabbitConfig `yaml:"rabbit"`

	DatabaseConfig struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"database"`

	GmailConfig   GmailConfig   `yaml:"gmailSender"`
	OutlookConfig OutlookConfig `yaml:"outlookSender"`
}

type GmailConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type OutlookConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type RabbitConfig struct {
	Port         int    `yaml:"port"`
	Host         string `yaml:"host"`
	User         string `yaml:"username"`
	Password     string `yaml:"password"`
	RetrySeconds int    `yaml:"retrySeconds"`
}
