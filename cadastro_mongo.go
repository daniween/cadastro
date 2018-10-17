package main

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UsersMONGO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "users"
)

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
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(email)).One(&user)
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
	err := db.C(COLLECTION).UpdateId(user.Email, &user)
	return err
}
