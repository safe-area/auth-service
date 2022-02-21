package service

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/safe-area/auth-service/config"
	"github.com/safe-area/auth-service/internal/repository"
)

type Service interface {
	Auth(token string) (string, error)
	SignIn(name, pass string) (string, error)
	SignUp(name, pass string) error
}

func New(cfg *config.Config, repo repository.Repository) Service {
	return &service{
		cfg:  cfg,
		repo: repo,
	}
}

type service struct {
	cfg  *config.Config
	repo repository.Repository
}

func (s *service) Auth(token string) (string, error) {
	return s.parseToken(token)
}

func (s *service) SignIn(name, pass string) (string, error) {
	h := sha1.New()
	h.Write([]byte(pass))
	pass = hex.EncodeToString(h.Sum(nil))
	userId, err := s.repo.SignIn(name, pass)
	if err != nil {
		return "", err
	}
	return s.createToken(userId)
}

func (s *service) SignUp(name, pass string) error {
	h := sha1.New()
	h.Write([]byte(pass))
	pass = hex.EncodeToString(h.Sum(nil))
	return s.repo.SignUp(name, pass)
}
