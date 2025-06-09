package restriction

type Service struct {
	repo RestrictionRepository
}

func NewService(r RestrictionRepository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateRestriction(restriction *Restriction) error {
	if err := restriction.ValidateDates(); err != nil {
		return err
	}
	return s.repo.Create(restriction)
}
