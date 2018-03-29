package greetlinker

import (
	"github.com/DesmondANIMUS/greeter/greetpackages/greetdb"
	"github.com/DesmondANIMUS/greeter/greetpackages/greethelper"
	"github.com/DesmondANIMUS/greeter/greetpackages/greetmodel"
)

// AddUserLinker function links to the PersModel and then adds a new user
func AddUserLinker(person greetmodel.PersonModel) {
	session := greethelper.Session.Clone()
	defer session.Close()
	greetdb.AddUser(person, session)
}
