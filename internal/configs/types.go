package configs

type (
	Config struct {
		Service  Service  `mapstructure:"service"`
		Database Database `mapstructure:"database"`
	}
	Service struct {
		Port         string `mapstructure:"port"`
		SecretJWT   string `mapstructure:"secretJWT"`
	}
	Database struct {
		// Host     string `mapstructure:"host"`
		// Port     string `mapstructure:"port"`
		// Username string `mapstructure:"username"`
		// Password string `mapstructure:"password"`
		DataSourceName string `mapstructure:"dataSourceName"`
	}
)
