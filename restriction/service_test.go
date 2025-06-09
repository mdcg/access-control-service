package restriction

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Mock do RestrictionRepository
type MockRestrictionRepo struct {
	mock.Mock
}

func (m *MockRestrictionRepo) Create(r *Restriction) error {
	args := m.Called(r)
	return args.Error(0)
}

func TestCreateRestriction(t *testing.T) {
	now := time.Now()

	t.Run("sucesso na criação", func(t *testing.T) {
		repo := new(MockRestrictionRepo)
		service := NewService(repo)

		restriction := &Restriction{
			Key:   "user",
			Value: "123",
			Rules: map[string]TimeRange{
				"service-a": {
					StartDate: now,
					EndDate:   now.Add(1 * time.Hour),
				},
			},
		}

		repo.On("Create", restriction).Return(nil)

		err := service.CreateRestriction(restriction)

		require.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("erro de datas inválidas", func(t *testing.T) {
		repo := new(MockRestrictionRepo)
		service := NewService(repo)

		restriction := &Restriction{
			Key:   "user",
			Value: "123",
			Rules: map[string]TimeRange{
				"service-a": {
					StartDate: now,
					EndDate:   now.Add(-1 * time.Hour), // inválido
				},
			},
		}

		err := service.CreateRestriction(restriction)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "start_date must be before end_date")
		repo.AssertNotCalled(t, "Create")
	})

	t.Run("erro no repositório", func(t *testing.T) {
		repo := new(MockRestrictionRepo)
		service := NewService(repo)

		restriction := &Restriction{
			Key:   "user",
			Value: "123",
			Rules: map[string]TimeRange{
				"service-a": {
					StartDate: now,
					EndDate:   now.Add(1 * time.Hour),
				},
			},
		}

		repo.On("Create", restriction).Return(errors.New("db error"))

		err := service.CreateRestriction(restriction)

		require.Error(t, err)
		assert.EqualError(t, err, "db error")
		repo.AssertExpectations(t)
	})
}
