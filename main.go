package main

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
}

//Users usuarios
var Users = []User{}

func dois() {

	AddUser(User{Name: "Stephanie Augusta", Email: "stephanie@gmail.com"})
	AddUser(User{Name: "manuela", Email: "manu@gmail.com"})
	UpdateUserName(User{Name: "manuela", Email: "manu@gmail.com"}, "Manuela Augusta")
	RemoveUser(User{Name: "Stephanie Augusta", Email: "stephanie@gmail.com"})
	list, _ := ListUsers()
	fmt.Println(list)

}

//ListUsers lista os usuarios
func ListUsers() ([]User, error) {

	return Users, nil
}

//GetUserByEmail lista o usuario pelo email
func GetUserByEmail(email string) (User, error) {

	for _, u := range Users {

		if u.Email == email {
			fmt.Println("User:", u)
			return u, nil
		}
	}

	return User{}, fmt.Errorf("Email not found")
}

//AddUser adiciona usuarios
func AddUser(user User) error {

	for _, u := range Users {

		if user.Email == u.Email {
			return fmt.Errorf("User already exists")
		}
	}

	Users = append(Users, user)
	return nil
}

//RemoveUser remove usuarios
func RemoveUser(user User) error {

	for i, u := range Users {

		if u == user {
			Users = append(Users[:i], Users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User não existe")
}

//UpdateUserName atualiza o nome
func UpdateUserName(user User, name string) error {

	//runtime.Breakpoint()
	for i, u := range Users {

		if u == user {
			copy(Users[i:], Users[i+1:])
			Users[len(Users)-1] = User{}
			Users = Users[:len(Users)-1]

			newUserName := User{Name: name, Email: user.Email}
			AddUser(newUserName)
			return nil
		}
	}
	return fmt.Errorf("Usuário não existe")
}
