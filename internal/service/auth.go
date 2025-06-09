package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/mohit838/inventory-managements-golang/dtos"
	"github.com/mohit838/inventory-managements-golang/internal/repository"
	"github.com/mohit838/inventory-managements-golang/models"
	"github.com/mohit838/inventory-managements-golang/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, dto *dtos.RegisterRequestDTO) (*dtos.AuthResponseDTO, error)
	Login(ctx context.Context, dto *dtos.LoginRequestDTO) (*dtos.AuthResponseDTO, error)
	RefreshToken(ctx context.Context, dto *dtos.RefreshTokenRequest) (*dtos.TokenResponse, error)
}

type authService struct {
	repo       repository.UserRepository
	jwtService *auth.JWTService
}

func NewAuthService(repo repository.UserRepository, jwt *auth.JWTService) AuthService {
	return &authService{
		repo:       repo,
		jwtService: jwt,
	}
}

// Register a new user
func (s *authService) Register(ctx context.Context, dto *dtos.RegisterRequestDTO) (*dtos.AuthResponseDTO, error) {
	existing, _ := s.repo.GetByEmail(ctx, dto.Email)
	if existing != nil {
		return nil, errors.New("email already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: string(hashed),
		RoleID:   dto.RoleID,
		BaseModel: models.BaseModel{
			IsActive: true,
		},
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return s.generateTokens(user.ID)
}

// Login user
func (s *authService) Login(ctx context.Context, dto *dtos.LoginRequestDTO) (*dtos.AuthResponseDTO, error) {
	user, err := s.repo.GetByEmail(ctx, dto.Email)

	if err != nil || user == nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return s.generateTokens(user.ID)
}

func (s *authService) RefreshToken(ctx context.Context, dto *dtos.RefreshTokenRequest) (*dtos.TokenResponse, error) {
	claims, err := s.jwtService.VerifyToken(dto.RefreshToken, auth.TokenTypeRefresh)
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %w", err)
	}

	access, err := s.jwtService.GenerateToken(claims.UserID, auth.TokenTypeAccess)
	if err != nil {
		return nil, err
	}

	refresh, err := s.jwtService.GenerateToken(claims.UserID, auth.TokenTypeRefresh)
	if err != nil {
		return nil, err
	}

	return &dtos.TokenResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

// Common: create access & refresh tokens
func (s *authService) generateTokens(userID int64) (*dtos.AuthResponseDTO, error) {
	access, err := s.jwtService.GenerateToken(userID, auth.TokenTypeAccess)
	if err != nil {
		return nil, err
	}

	refresh, err := s.jwtService.GenerateToken(userID, auth.TokenTypeRefresh)
	if err != nil {
		return nil, err
	}

	return &dtos.AuthResponseDTO{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}
