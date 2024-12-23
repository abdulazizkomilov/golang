package service

import (
	"context"
	"user-service/internal/domain"
	"user-service/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, req domain.CreateUserRequest) (*domain.User, error)
	GetUser(ctx context.Context, id uint) (*domain.User, error)
	UpdateUser(ctx context.Context, id uint, req domain.UpdateUserRequest) (*domain.User, error)
	DeleteUser(ctx context.Context, id uint) error
	ListUsers(ctx context.Context) ([]domain.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, req domain.CreateUserRequest) (*domain.User, error) {
	user := &domain.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUser(ctx context.Context, id uint) (*domain.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *userService) UpdateUser(ctx context.Context, id uint, req domain.UpdateUserRequest) (*domain.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Email != "" {
		user.Email = req.Email
	}

	if err := s.repo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUser(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *userService) ListUsers(ctx context.Context) ([]domain.User, error) {
	return s.repo.List(ctx)
}
