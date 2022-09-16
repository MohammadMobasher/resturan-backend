package models

type Configuration struct {
	Environment        string `mapstructure:"ENV"`
	MongoServer        string `mapstructure:"MONGO_SERVER"`
	MongoDatabase      string `mapstructure:"MONGO_DATABASE"`
	MongoUsername      string `mapstructure:"MONGO_USERNAME"`
	MongoPassword      string `mapstructure:"MONGO_PASSWORD"`
	LaravelBaseUrl     string `mapstructure:"LARAVEL_BASE_URL"`
	LaravelUserName    string `mapstructure:"LARAVEL_USERNAME"`
	LaravelPassword    string `mapstructure:"LARAVEL_PASSWORD"`
	UseTestNet         bool   `mapstructure:"USE_TEST_NET"`
	StreamsPerListener int64  `mapstructure:"NUM_OF_STREAMS_PER_LISTENER"`
}
