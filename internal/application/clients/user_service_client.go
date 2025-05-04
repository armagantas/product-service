package clients

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

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

	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/user/info", nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+ token)

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

	var userInfo domain.UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		log.Println("Error decoding user info:", err)
		return nil, err
	}

	return &userInfo, nil
}
