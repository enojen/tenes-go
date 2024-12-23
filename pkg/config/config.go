package config

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Port string
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: ":8080",
		},
	}
}
