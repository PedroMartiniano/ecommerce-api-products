package middlewares

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/PedroMartiniano/ecommerce-api-products/internal/configs"
	"github.com/gin-gonic/gin"
)

type request struct {
	Success bool `json:"success"`
	Data    User  `json:"data"`
}

type User struct {
	ID        string `json:"id"`
	RoleID    string `json:"role_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CPF       string `json:"cpf"`
	BirthDate string `json:"birth_date"`
	Status    bool   `json:"status"`
	CreatedAt string `json:"created_at"`
}

func VerifyToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Authorization token is required"})
		c.Abort()
		return
	}

	url := configs.GetEnv("AUTH_SERVICE_URL")

	fullURL := fmt.Sprintf("%s/users/me", url)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error to create request"})
		c.Abort()
		return
	}
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error to send request"})
		c.Abort()
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error to read response body"})
		c.Abort()
		return
	}

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"success": false, "message": "Invalid token"})
		c.Abort()
		return
	}

	var request request

	err = json.Unmarshal(body, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error to parse response body"})
		c.Abort()
		return
	}

	c.Set("user", request.Data)

	c.Next()
}
