package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"
	"fmt"

	"github.com/gorilla/mux"
	. "github.com/D3n0Duz/db-service/config"
	. "github.com/D3n0Duz/db-service/dao"
	. "github.com/D3n0Duz/db-service/models"
)

var config = Config{}
var dao = ClientTransactionDAO{}

// GET list of clientTransactions
func AllClientTransactionsEndPoint(w http.ResponseWriter, r *http.Request) {
	clientTransactions, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, clientTransactions)
}

// GET a clientTransaction by its ID
func FindClientTransactionEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	clientTransaction, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ClientTransaction ID")
		return
	}
	respondWithJson(w, http.StatusOK, clientTransaction)
}

// POST a new clientTransaction
func CreateClientTransactionEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var clientTransaction ClientTransaction
	if err := json.NewDecoder(r.Body).Decode(&clientTransaction); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	clientTransaction.ID = bson.NewObjectId()
	if err := dao.Insert(clientTransaction); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, clientTransaction)
}

// PUT update an existing clientTransaction
func UpdateClientTransactionEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var clientTransaction ClientTransaction
	if err := json.NewDecoder(r.Body).Decode(&clientTransaction); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(clientTransaction); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing clientTransaction
func DeleteClientTransactionEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var clientTransaction ClientTransaction
	if err := json.NewDecoder(r.Body).Decode(&clientTransaction); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(clientTransaction); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database

	time.Sleep(30 * time.Second)
	dao.Connect()
}

// Define HTTP request routes
func main() {
	log.Printf("Hello go")
	r := mux.NewRouter()
	r.HandleFunc("/clientTransactions", AllClientTransactionsEndPoint).Methods("GET")
	r.HandleFunc("/clientTransactions", CreateClientTransactionEndPoint).Methods("POST")
	r.HandleFunc("/clientTransactions", UpdateClientTransactionEndPoint).Methods("PUT")
	r.HandleFunc("/clientTransactions", DeleteClientTransactionEndPoint).Methods("DELETE")
	r.HandleFunc("/clientTransactions/{id}", FindClientTransactionEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
