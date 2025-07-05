package auth

import (
	"HttpServer/configs"
	"HttpServer/pkg/req"
	"HttpServer/pkg/res"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
}

type AuthHandlerDeps struct {
	*configs.Config
}

func NewAuth(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /register", handler.Register())
	router.HandleFunc("POST /login", handler.Login())
}

func (auth *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := req.HandleBody[LoginRequest](&w, r)
		data := LoginResponse{Token: "123"}
		if err != nil {
			return
		}
		res.Json(w, 201, data)
	}
}

func (auth *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := req.HandleBody[RegisterRequest](&w, r)
		data := RegisterResponse{Token: "234"}
		if err != nil {
			return
		}
		res.Json(w, 201, data)
	}
}
