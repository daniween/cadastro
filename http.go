package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	e.GET("/", getAllUsers)
	e.GET("/:email", getUser)             // curl localhost:1323/manu@gmail.com --> mostra o usuario manuela, manu@gmail.com
	e.PUT("/users/:email", updateUser)    // curl -X PUT -F "name=JOE" localhost:1323/users/joe@gmail.com
	e.POST("/save", save)                 // curl -F "name=Joe Smith" -F "email=joe@labstack.com" http://localhost:1323/save
	e.DELETE("/users/:email", deleteUser) // curl -X DELETE localhost:1323/users/manu@gmail.com

	e.Logger.Fatal(e.Start(":8000"))
}

func getAllUsers(c echo.Context) error { //curl localhost:1323/ --> lista todos os usuarios
	list, _ := ListUsers()
	return c.JSON(http.StatusOK, list)
}

func getUser(c echo.Context) error {
	email := c.Param("email")
	user, err := GetUserByEmail(email)
	if err != nil {
		return c.String(http.StatusNotFound, "Not found")
	}
	return c.JSON(http.StatusOK, user)
}

func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	err := AddUser(User{Name: name, Email: email})
	if err != nil {
		return c.String(http.StatusNotFound, "User already exists\n")
	}
	return c.JSON(http.StatusCreated, "name:"+name+", email:"+email)
}

func updateUser(c echo.Context) error {

	newName := c.FormValue("name")
	email := c.Param("email")

	user, err := GetUserByEmail(email)
	if err != nil {
		return c.String(http.StatusNotFound, "Not found")
	}

	erro := UpdateUserName(user, newName)
	if erro != nil {
		return c.String(http.StatusNotFound, "Not found")
	}
	return c.JSON(http.StatusCreated, User{newName, email})
}

func deleteUser(c echo.Context) error {

	email := c.Param("email")
	user, err := GetUserByEmail(email)
	if err != nil {
		return err
	}

	erro := RemoveUser(user)
	if erro != nil {
		return c.String(http.StatusNotFound, "Not found")
	}
	return c.JSON(http.StatusOK, Users)
}
