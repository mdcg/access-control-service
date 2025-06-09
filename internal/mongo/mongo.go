package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB realiza a conexão com o MongoDB utilizando a URI fornecida
// e retorna uma referência ao banco de dados especificado.
func ConnectDB(uri, dbName string) *mongo.Database {
	// Define um contexto com timeout de 10 segundos para a operação de conexão
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Garante que o cancelamento será chamado ao final da função

	// Define as opções do cliente MongoDB usando a URI fornecida
	clientOptions := options.Client().ApplyURI(uri)

	// Cria uma nova conexão com o MongoDB usando as opções especificadas
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err) // Encerra o programa se a conexão falhar
	}

	// Testa a conexão com o MongoDB usando o método Ping
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB ping failed:", err) // Encerra o programa se o ping falhar
	}

	log.Println("Connected to MongoDB!") // Confirma que a conexão foi bem-sucedida

	// Retorna a instância do banco de dados especificado
	return client.Database(dbName)
}

// GetCollection retorna uma referência a uma coleção específica dentro do banco de dados fornecido.
func GetCollection(db *mongo.Database, collectionName string) *mongo.Collection {
	return db.Collection(collectionName)
}
