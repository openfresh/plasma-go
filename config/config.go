package config

type Config struct {
	Redis Redis
	Type  string
}

type Redis struct {
	Addr     string
	Password string
	DB       int
	Channel  string
}
