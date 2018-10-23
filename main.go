package main

import "fmt"

//User usuarios
type User struct {
	Name  string
	Email string
}

//Users usuarios
var Users = []User{}

//UserDB mongo
var UserDB *UsersMONGO

func init() {

	UserDB = &UsersMONGO{}
	UserDB.Connect()
}

//ListUsers lista os usuarios
func ListUsers() ([]User, error) {
	return UserDB.getAllUsers()
}

//GetUserByEmail lista o usuario pelo email
func GetUserByEmail(email string) (User, error) {
	return UserDB.getUser(email)
}

//AddUser adiciona usuarios
func AddUser(user User) error {
	_, err := GetUserByEmail(user.Email)
	if err == nil { // se o erro for igual nil, entao ele encontrou um usuario. Entao nao deixa adicionar
		return fmt.Errorf("User already exists")
	}
	return UserDB.save(user)
}

//RemoveUser remove usuarios
func RemoveUser(user User) error {
	return UserDB.deleteUser(user)
}

//UpdateUserName atualiza o nome
func UpdateUserName(user User, name string) error {
	user = User{name, user.Email}
	return UserDB.updateUser(user, name)
}
