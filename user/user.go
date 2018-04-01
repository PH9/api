package user

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Example get user
func GetUser(c echo.Context) error {
	u := &User{"AnuchitO", "Nong@gmail.com"}
	return c.JSON(http.StatusOK, u)
}
