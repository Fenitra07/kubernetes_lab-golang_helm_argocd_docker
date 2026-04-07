package configs

type Configuration struct {
	API    APIConfig `yaml:"API"`
	JWT    JWTConfig `yaml:"JWT"`
	DB_DSN string    `yaml:"DB_DSN"`
}

type APIConfig struct {
	Host         string   `yaml:"HOST"`
	Port         int      `yaml:"PORT"`
	AllowOrigins []string `yaml:"ALLOW_ORIGINS"`
}

type JWTConfig struct {
	SecretKey   string `yaml:"SECRET_KEY"`
	TokenExpire string `yaml:"TOKEN_EXPIRE"`
}
