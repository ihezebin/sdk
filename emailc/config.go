package emailc

const (
	HostQQMail = "smtp.qq.com"
	HostExmail = "smtp.exmail.qq.com"
	PortQQMail = 25
	PortExmail = 465
)

type Config struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
}

type Option func(config *Config)

func newConfig(options ...Option) Config {
	config := Config{}
	for _, option := range options {
		option(&config)
	}
	return config
}

func WithHost(host string) Option {
	return func(config *Config) {
		config.Host = host
	}
}

func WithPort(port int) Option {
	return func(config *Config) {
		config.Port = port
	}
}

func WithUsername(username string) Option {
	return func(config *Config) {
		config.Username = username
	}
}

func WithPassword(password string) Option {
	return func(config *Config) {
		config.Password = password
	}
}
