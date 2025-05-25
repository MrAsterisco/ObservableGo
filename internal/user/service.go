package user

import (
	"context"
	"regexp"

	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

var emailRegex = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)

func (s *Service) Create(ctx context.Context, email string) (*User, error) {
	if !emailRegex.MatchString(email) {
		return nil, ErrInvalidEmail
	}
	id := uuid.New()
	u := &User{ID: id, Email: email}
	if err := s.repo.Create(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Service) GetAll(ctx context.Context) ([]User, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) GetByID(ctx context.Context, id string) (*User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, email string) error {
	if !emailRegex.MatchString(email) {
		return ErrInvalidEmail
	}
	u := &User{ID: id, Email: email}
	return s.repo.Update(ctx, u)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

var ErrInvalidEmail = &ServiceError{"invalid email address"}

type ServiceError struct {
	Msg string
}

func (e *ServiceError) Error() string { return e.Msg }
