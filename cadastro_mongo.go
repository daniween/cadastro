package main

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//UsersMONGO serve and database
type UsersMONGO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	//COLLECTION nome da collection do mongo
	COLLECTION = "users"
)

//Connect para conectar com o servidor
func (u *UsersMONGO) Connect() {
	session, err := mgo.Dial(u.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(u.Database)
}

func (u *UsersMONGO) getAllUsers() ([]User, error) {
	var users []User
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

func (u *UsersMONGO) getUser(email string) (User, error) {
	var user User
	err := db.C(COLLECTION).Find(bson.M{"email": email}).One(&user)
	return user, err
}

func (u *UsersMONGO) save(user User) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}

func (u *UsersMONGO) deleteUser(user User) error {
	err := db.C(COLLECTION).Remove(&user)
	return err
}

func (u *UsersMONGO) updateUser(user User, name string) error {
	err := db.C(COLLECTION).Update(bson.M{"email": user.Email}, &user)
	return err
}

func (u *UsersMONGO) deleteAll() error { // para os testes
	_, err := db.C(COLLECTION).RemoveAll(nil)
	return err
}
