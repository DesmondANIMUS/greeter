package greethelper

import (
	"github.com/DesmondANIMUS/greeter/greetpackages/greetmodel"
	mgo "gopkg.in/mgo.v2"
)

// Session global variable for mongodb global session
var Session *mgo.Session

// ConnectDB function creates a global connection to the database
func ConnectDB() *mgo.Session {
	session, err := mgo.Dial(greetmodel.ConnectionString)
	if err != nil {
		panic(err)
	}

	return session
}
