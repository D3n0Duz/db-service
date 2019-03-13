package dao

import (
	"log"

	. "github.com/D3n0Duz/db-service/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ClientTransactionDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "clientTransaction"
)

// Establish a connection to database
func (m *ClientTransactionDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of clientTransactions
func (m *ClientTransactionDAO) FindAll() ([]ClientTransaction, error) {
	var clientTransactions []ClientTransaction
	err := db.C(COLLECTION).Find(bson.M{}).All(&clientTransactions)
	return clientTransactions, err
}

// Find a clientTransaction by its id
func (m *ClientTransactionDAO) FindById(id string) (ClientTransaction, error) {
	var clientTransaction ClientTransaction
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&clientTransaction)
	return clientTransaction, err
}

// Insert a clientTransaction into database
func (m *ClientTransactionDAO) Insert(clientTransaction ClientTransaction) error {
	err := db.C(COLLECTION).Insert(&clientTransaction)
	return err
}

// Delete an existing clientTransaction
func (m *ClientTransactionDAO) Delete(clientTransaction ClientTransaction) error {
	err := db.C(COLLECTION).Remove(&clientTransaction)
	return err
}

// Update an existing clientTransaction
func (m *ClientTransactionDAO) Update(clientTransaction ClientTransaction) error {
	err := db.C(COLLECTION).UpdateId(clientTransaction.ID, &clientTransaction)
	return err
}