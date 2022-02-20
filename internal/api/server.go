package api

import (
	"database/sql"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/lab259/cors"
	"github.com/safe-area/auth-service/internal/models"
	"github.com/safe-area/auth-service/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
	"time"
)

type Server struct {
	r    *fasthttprouter.Router
	serv *fasthttp.Server
	svc  service.Service
	port string
}

func New(svc service.Service, port string) *Server {
	innerRouter := fasthttprouter.New()
	innerHandler := innerRouter.Handler
	s := &Server{
		innerRouter,
		&fasthttp.Server{
			ReadTimeout:  time.Duration(10) * time.Second,
			WriteTimeout: time.Duration(10) * time.Second,
			IdleTimeout:  time.Duration(10) * time.Second,
			Handler:      cors.AllowAll().Handler(innerHandler),
		},
		svc,
		port,
	}

	s.r.POST("/api/v1/sign-in", s.SignInHandler)
	s.r.POST("/api/v1/sign-up", s.SignUpHandler)
	//s.r.POST("/api/v1/auth", s.TestHandler)

	return s
}

func (s *Server) SignInHandler(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	body := ctx.PostBody()
	var userData models.UserData
	err := jsoniter.Unmarshal(body, &userData)
	if err != nil {
		logrus.Errorf("TestHandler: error while unmarshalling request: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	uuid, err := s.svc.SignIn(userData.Name, userData.Password)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	bs, err := jsoniter.Marshal(models.TokenData{Token: uuid})
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	} else if err == sql.ErrNoRows {
		ctx.SetStatusCode(fasthttp.StatusNoContent)
		return
	}
	ctx.Write(bs)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (s *Server) SignUpHandler(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	body := ctx.PostBody()
	var userData models.UserData
	err := jsoniter.Unmarshal(body, &userData)
	if err != nil {
		logrus.Errorf("SignUpHandler: error while unmarshalling request: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	err = s.svc.SignUp(userData.Name, userData.Password)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (s *Server) Start() error {
	return fmt.Errorf("server start: %s", s.serv.ListenAndServe(s.port))
}
func (s *Server) Shutdown() error {
	return s.serv.Shutdown()
}
