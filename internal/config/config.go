package config

type Config struct {
	Mysql struct {
		DataSource string
	}

	Redis struct {
		Host string
		Type string
	}
}
