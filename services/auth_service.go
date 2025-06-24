package services

import (
	"errors"
	"fmt"
	"os"
	"time"
	"votes/models"
	"votes/requests"
	"votes/utils/errs"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userSvc *UserService
}

type TokenType string

const (
	TokenTypeAccess  TokenType = "access_token"
	TokenTypeRefresh TokenType = "refresh_token"
)

type TokenTTL time.Duration

const (
	AccessTokenTTL  TokenTTL = TokenTTL(15 * time.Minute)
	RefreshTokenTTL TokenTTL = TokenTTL(30 * 24 * time.Hour)
)

type Claim struct {
	TokenType TokenType
	jwt.RegisteredClaims
}

var SecretKey = []byte(os.Getenv("SECRET_KEY"))

func NewAuthService(userSvc *UserService) *AuthService {
	return &AuthService{
		userSvc: userSvc,
	}
}

func (a *AuthService) Register(req *requests.RegisterRequest) (*models.User, error) {
	if _, err := a.userSvc.GetByEmail(req.Email); err == nil {
		return nil, errs.ErrEmailAlreadyExist
	} else if !errors.Is(err, errs.ErrDataNotFound) {
		return nil, fmt.Errorf("check existing email: %w", err)
	}

	if _, err := a.userSvc.GetByUsername(req.Username); err == nil {
		return nil, errs.ErrUsernameAlreadyExist
	} else if !errors.Is(err, errs.ErrDataNotFound) {
		return nil, fmt.Errorf("check existing username: %w", err)
	}

	register, err := a.userSvc.Create(req)
	if err != nil {
		return nil, fmt.Errorf("register: %w", err)
	}

	newUser, err := a.userSvc.GetById(register.ID)
	if err != nil {
		return nil, fmt.Errorf("fetch registered user: %w", err)
	}

	return newUser, nil
}

func (a *AuthService) generateToken(userId uuid.UUID, t TokenType, d TokenTTL) (string, error) {
	now := time.Now()
	exp := now.Add(time.Duration(d))

	claims := &Claim{
		TokenType: t,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userId.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", fmt.Errorf("generate token: %w", err)
	}

	return tokenString, nil
}

// return access_token, refresh_token, error
func (a *AuthService) Login(req *requests.LoginRequest) (string, string, error) {
	user, err := a.userSvc.GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, errs.ErrDataNotFound) {
			return "", "", errs.ErrInvalidLogin
		}

		return "", "", fmt.Errorf("check existing email: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", "", errs.ErrInvalidLogin
		}

		return "", "", fmt.Errorf("compare password: %w", err)
	}

	accToken, err := a.generateToken(user.ID, TokenTypeAccess, AccessTokenTTL)
	if err != nil {
		return "", "", err
	}

	refToken, err := a.generateToken(user.ID, TokenTypeRefresh, RefreshTokenTTL)
	if err != nil {
		return "", "", err
	}

	return accToken, refToken, nil
}

func (a *AuthService) ValidateToken(token string) (*Claim, error) {
	validateToken, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return &Claim{}, errs.ErrInvalidToken
	}

	claims, ok := validateToken.Claims.(*Claim)
	if !ok {
		return &Claim{}, errs.ErrInvalidClaim
	}

	return claims, nil
}

// return new access token
func (a *AuthService) RefreshToken(req *requests.RefreshTokenRequest) (string, error) {
	token, err := a.ValidateToken(req.RefreshToken)
	if err != nil {
		return "", errs.ErrInvalidToken
	}

	if token.TokenType != TokenTypeRefresh {
		return "", errs.ErrInvalidToken
	}

	userId, err := uuid.Parse(token.Subject)
	if err != nil {
		return "", fmt.Errorf("parse user_id")
	}

	newAccToken, err := a.generateToken(userId, TokenTypeAccess, AccessTokenTTL)
	if err != nil {
		return "", err
	}

	return newAccToken, nil
}
