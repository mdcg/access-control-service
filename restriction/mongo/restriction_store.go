package mongo

import (
	"context"

	"github.com/mdcg/access-control-service/config"                // Importa o pacote de configuração do sistema
	"github.com/mdcg/access-control-service/restriction"           // Importa o pacote do domínio com a entidade Restriction
	"github.com/mdcg/access-control-service/restriction/mongo/dto" // Importa os DTOs usados para persistência em MongoDB
	"go.mongodb.org/mongo-driver/mongo"                            // Importa o driver oficial do MongoDB
)

// RestrictionStore representa o repositório que interage com a coleção de restrições no MongoDB.
type RestrictionStore struct {
	ctx        context.Context
	collection *mongo.Collection // Referência à coleção do MongoDB usada para armazenar restrições
}

// NewRestrictionStore cria uma nova instância de RestrictionStore,
// usando a coleção definida nas variáveis de ambiente (via config).
func NewRestrictionStore(ctx context.Context, db *mongo.Database) *RestrictionStore {
	return &RestrictionStore{
		ctx:        ctx,
		collection: db.Collection(config.GetEnvVars().MongoDBRestrictionCollection),
	}
}

// Create insere uma nova restrição no banco de dados.
// Recebe um ponteiro para a entidade Restriction (do domínio).
func (rs *RestrictionStore) Create(r *restriction.Restriction) error {
	// Cria um DTO para mapear a entidade para o formato esperado pelo MongoDB
	rdto := &dto.RestrictionDTO{
		Key:   r.Key,
		Value: r.Value,
		Rules: map[string]dto.TimeRangeDTO{},
	}

	// Converte o mapa de regras do domínio para o formato DTO (TimeRangeDTO)
	for service, tr := range r.Rules {
		rdto.Rules[service] = dto.TimeRangeDTO{
			StartDate: tr.StartDate,
			EndDate:   tr.EndDate,
		}
	}

	// Insere o DTO convertido na coleção MongoDB
	_, err := rs.collection.InsertOne(rs.ctx, rdto)
	return err // Retorna qualquer erro ocorrido na operação
}
