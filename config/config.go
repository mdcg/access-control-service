package config

import (
	"log"

	"github.com/spf13/viper" // Biblioteca usada para carregar variáveis de ambiente e arquivos de configuração
)

var (
	envVars *Environments // Variável global que armazenará as configurações carregadas
)

// Struct que define os campos esperados nas variáveis de ambiente, com tags indicando os nomes das variáveis
type Environments struct {
	MongoDBURI                   string `mapstructure:"MONGODB_URI"`
	MongoDBDatabase              string `mapstructure:"MONGODB_DATABASE"`
	MongoDBPermissionCollection  string `mapstructure:"MONGODB_PERMISSION_COLLECTION"`
	MongoDBRestrictionCollection string `mapstructure:"MONGODB_RESTRICTION_COLLECTION"`
	RedisHost                    string `mapstructure:"REDIS_HOST"`
	RedisPort                    string `mapstructure:"REDIS_PORT"`
	RabbitMQHost                 string `mapstructure:"RABBITMQ_HOST"`
	RabbitMQPort                 string `mapstructure:"RABBITMQ_PORT"`
	RabbitMQUsername             string `mapstructure:"RABBITMQ_USERNAME"`
	RabbitMQPassword             string `mapstructure:"RABBITMQ_PASSWORD"`
	OtelExporterOtlpEndpoint     string `mapstructure:"OTEL_EXPORTER_OTLP_ENDPOINT"`
	OtelExporterOtlpInsecure     string `mapstructure:"OTEL_EXPORTER_OTLP_INSECURE"`
	OtelServiceName              string `mapstructure:"OTEL_SERVICE_NAME"`
}

// Função responsável por carregar as variáveis de ambiente
func LoadEnvVars() *Environments {
	// Define que o arquivo de configuração a ser lido é o ".env"
	viper.SetConfigFile(".env")

	// Define valores padrão caso não estejam presentes no .env ou nas variáveis de ambiente
	viper.SetDefault("OTEL_EXPORTER_OTLP_ENDPOINT", "http://localhost:4318")
	viper.SetDefault("OTEL_EXPORTER_OTLP_INSECURE", "true")
	viper.SetDefault("OTEL_SERVICE_NAME", "access-control-service")
	viper.SetDefault("MONGODB_URI", "mongodb://localhost:27017")
	viper.SetDefault("MONGODB_DATABASE", "access-control")
	viper.SetDefault("MONGODB_PERMISSION_COLLECTION", "permission")
	viper.SetDefault("MONGODB_RESTRICTION_COLLECTION", "restriction")
	viper.SetDefault("REDIS_HOST", "localhost")
	viper.SetDefault("REDIS_PORT", "6379")
	viper.SetDefault("RABBITMQ_HOST", "localhost")
	viper.SetDefault("RABBITMQ_PORT", "5672")
	viper.SetDefault("RABBITMQ_USERNAME", "admin")
	viper.SetDefault("RABBITMQ_PASSWORD", "admin123")

	// Permite que o Viper também leia variáveis de ambiente do sistema operacional
	viper.AutomaticEnv()

	// Lê o arquivo de configuração definido anteriormente (".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("unable to find or read config file: %s", err)
	}

	// Mapeia os dados lidos para a struct Environments
	if err := viper.Unmarshal(&envVars); err != nil {
		log.Fatalf("unable to unmarshal config from environments variables: %s", err)
	}

	// Retorna o ponteiro para a struct com as configurações carregadas
	return envVars
}

func GetEnvVars() *Environments {
	return envVars
}
