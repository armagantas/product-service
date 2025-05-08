package authentication

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/armagantas/ecommerce-microservice/product-service/internal/application/clients"
	domain "github.com/armagantas/ecommerce-microservice/product-service/internal/domain/models"
)

func GetUserInfoFromRequest(r *http.Request, client clients.UserServiceClient) (*domain.UserInfo, error) {
	token := extractBearerToken(r.Header.Get("Authorization"))
	if token == "" {
		return &domain.UserInfo{}, errors.New("unauthorized: no token")
	}

	userInfo, err := client.GetUserInfo(token)
	if err != nil {
		return &domain.UserInfo{}, fmt.Errorf("auth failed: %w", err)
	}

	return userInfo, nil
}

func extractBearerToken(authHeader string) string {
	const prefix = "Bearer "

	if len(authHeader) < len(prefix) || authHeader[:len(prefix)] != prefix {
		return ""
	}

	return authHeader[len(prefix):]
}
