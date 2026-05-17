package config

type Config struct {
	App AppConfig
}

type AppConfig struct {
	Name string
	Env  string
	Port string
}
