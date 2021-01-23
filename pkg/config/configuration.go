package config

type Config struct {
	Server struct {
		Port string `yaml:"port" env:"SERVER_PORT" env-default:"8080"`
		Host string `yaml:"host" env:"SERVER_HOST" env-default:"0.0.0.0"`
	} `yaml:"server"`
}
