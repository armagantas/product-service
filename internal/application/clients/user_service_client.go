package clients

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	domain "github.com/armagantas/ecommerce-microservice/product-service/internal/domain/models"
)

type UserServiceClient interface {
	GetUserInfo(token string) (*domain.UserInfo, error)
}

type userServiceClient struct {
	userServiceURL string
}

func NewUserServiceClient(userServiceURL string) UserServiceClient {
	return &userServiceClient{userServiceURL: userServiceURL}
}

func (c *userServiceClient) GetUserInfo(token string) (*domain.UserInfo, error) {
	req, err := http.NewRequest("GET", c.userServiceURL+"/api/users/profile", nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Error getting user info:", resp.Status)
		return nil, errors.New("error getting user info")
	}

	var response struct {
		Success bool            `json:"success"`
		Data    domain.UserInfo `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Println("Error decoding user info:", err)
		return nil, err
	}

	if !response.Success {
		return nil, errors.New("user service returned unsuccessful response")
	}

	// Username yoksa email'den oluştur (eski kullanıcılar için geçici çözüm)
	if response.Data.Username == "" && response.Data.Email != "" {
		parts := strings.Split(response.Data.Email, "@")
		response.Data.Username = parts[0]
	}

	log.Printf("User info retrieved: ID=%s, Username=%s", response.Data.ID, response.Data.Username)

	return &response.Data, nil
}
