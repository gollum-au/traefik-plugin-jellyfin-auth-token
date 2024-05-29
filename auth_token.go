package main

import (
    "context"
    "net/http"

    "github.com/traefik/traefik/v2/pkg/middlewares"
    "github.com/traefik/traefik/v2/pkg/middlewares/plugin"
    "github.com/traefik/traefik/v2/pkg/types"
)

type Config struct {}

func CreateConfig() *Config {
    return &Config{}
}

type AuthToken struct {
    next http.Handler
    name string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
    return &AuthToken{
        next: next,
        name: name,
    }, nil
}

func (a *AuthToken) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
    authHeader := req.Header.Get("Authorization")
    if authHeader != "" {
        if req.Header.Get("X-Emby-Token") == "" {
            req.Header.Set("X-Emby-Token", authHeader)
        }
    }
    a.next.ServeHTTP(rw, req)
}

