package config

type Config struct {
	MongoURI string `env:"MONGO_URI" envDefault:"mongodb://localhost:27017"`
	Port     string `env:"PORT" envDefault:":8080"`
}
