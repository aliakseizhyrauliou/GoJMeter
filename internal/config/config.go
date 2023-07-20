package config

type AppConfig struct {
	Port   string
	DBHost string
	DBPort string
	DBUser string
	DBPass string
}

func LoadConfig() *AppConfig {

	return &AppConfig{
		Port: ":8080",
	}
}
