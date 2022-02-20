package service

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/safe-area/auth-service/config"
	"github.com/safe-area/auth-service/internal/repository"
)

type Service interface {
	Auth(token string) error
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

func (s *service) Auth(token string) error {
	return nil
}

func (s *service) SignIn(name, pass string) (string, error) {
	h := sha1.New()
	h.Write([]byte(pass))
	pass = hex.EncodeToString(h.Sum(nil))
	// TODO return token against uuid
	return s.repo.SignIn(name, pass)
}

func (s *service) SignUp(name, pass string) error {
	h := sha1.New()
	h.Write([]byte(pass))
	pass = hex.EncodeToString(h.Sum(nil))
	return s.repo.SignUp(name, pass)
}
