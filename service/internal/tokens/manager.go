package tokens

import (
	"TrafficPolice/internal/domain"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"math/rand"
	"time"
)

type TokenInfo struct {
	UserID   string
	UserRole domain.Role
}

type TokenManager interface {
	NewJWT(tokenInfo TokenInfo, ttl time.Duration) (string, error)
	Parse(accessToken string) (TokenInfo, error)
	NewRefreshToken() (string, error)
}

type manager struct {
	signingKey string
}

func NewTokenManager(signingKey string) (TokenManager, error) {
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}

	return &manager{signingKey: signingKey}, nil
}

func (m *manager) NewJWT(tokenInfo TokenInfo, ttl time.Duration) (string, error) {
	if tokenInfo.UserID == "" || tokenInfo.UserRole == "" {
		return "", fmt.Errorf("userID or user role is empty")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  time.Now().Add(ttl).Unix(),
		"sub":  tokenInfo.UserID,
		"role": tokenInfo.UserRole,
	})

	return token.SignedString([]byte(m.signingKey))
}

func (m *manager) Parse(accessToken string) (TokenInfo, error) {
	keyFunc := func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.signingKey), nil
	}
	token, err := jwt.Parse(accessToken, keyFunc)
	if err != nil {
		return TokenInfo{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return TokenInfo{}, fmt.Errorf("error get user claims from token")
	}

	return TokenInfo{
		UserID:   claims["sub"].(string),
		UserRole: domain.Role(claims["role"].(string)),
	}, nil
}

func (m *manager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}
