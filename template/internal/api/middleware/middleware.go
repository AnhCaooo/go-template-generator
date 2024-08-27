// Created by Anh Cao on 27.08.2024.

package middleware

import (
	"net/http"

	"github.com/AnhCaooo/go-web-server-template-generator/template/internal/logger"
	"go.uber.org/zap"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Logger.Info("request received", zap.String("method", r.Method), zap.String("endpoint", r.URL.Path))
		next.ServeHTTP(w, r)
	})
}
