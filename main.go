package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sangianpatrick/go-rest-echo/model"
)

func main() {
	server := echo.New()
	server.GET("/", func(c echo.Context) error {
		u := &model.User{
			UserID:    "a346bcccd75765f86c665f555",
			FirstName: "Patrick",
			LastName:  "Maurits",
			Email:     "patricksangian@gmail.com",
		}

		return c.JSON(http.StatusOK, u)
	})
	server.Logger.Fatal(server.Start(":9000"))
}
