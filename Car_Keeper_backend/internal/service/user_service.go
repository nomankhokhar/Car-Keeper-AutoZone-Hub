package service

import (
	"Car_Keeper/internal/dto"
	"Car_Keeper/internal/models"
	"Car_Keeper/internal/repository"
	"Car_Keeper/pkg/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(req *dto.RegisterRequest) (*models.User, error)
	Login(req *dto.LoginRequest) (*models.User, string, error)
	GetByID(id uint) (*models.User, error)
	Update(id uint, req *dto.UpdateUserRequest) (*models.User, error)
	Delete(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(req *dto.RegisterRequest) (*models.User, error) {
	// Check if user exists
	existing, _ := s.repo.FindByEmail(req.Email)
	if existing != nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Name:     req.Name,
		Phone:    req.Phone,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Login(req *dto.LoginRequest) (*models.User, string, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *userService) GetByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) Update(id uint, req *dto.UpdateUserRequest) (*models.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Delete(id uint) error {
	return s.repo.Delete(id)
}
