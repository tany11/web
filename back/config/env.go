package config

type DatabaseEnv struct {
	Name      string `env:"DB_NAME"`
	User      string `env:"DB_USER"`
	Password  string `env:"DB_USER_PASSWORD"`
	Address   string `env:"DB_ADDRESS"`
	Collation string `env:"DB_COLLATION"`
}
