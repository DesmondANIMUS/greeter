package greetdb

import (
	"github.com/DesmondANIMUS/greeter/greetpackages/greetmodel"
	"gopkg.in/mgo.v2"
)

// AddUser function will add a user in the database
func AddUser(person greetmodel.PersonModel, session *mgo.Session) {
	c := session.DB(greetmodel.DatabaseString).C(greetmodel.PersonModelString)
	c.Insert(person)
}
