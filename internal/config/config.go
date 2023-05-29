package config

type Config struct {
	ConnectionString string `env:"CONNECTION_STRING" envDefault:"mongodb://localhost:27017"`
	PortGin          string `env:"PORT_GIN" envDefault:":8080"`
}
