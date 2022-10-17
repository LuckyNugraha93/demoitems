package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/LuckyNugraha93/demoitems/api/auth"
	"github.com/LuckyNugraha93/demoitems/api/models"
	"github.com/LuckyNugraha93/demoitems/api/responses"
	"github.com/LuckyNugraha93/demoitems/api/utils/formaterror"
)

func (server *Server) CreateTransaction(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	transaction := models.Transaction{}
	err = json.Unmarshal(body, &transaction)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	transaction.Prepare()
	err = transaction.Validate()
	
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	for i, _ := range transaction.TransactionDetails {
		transaction.TransactionDetails[i].Prepare()
		err = transaction.TransactionDetails[i].Validate()
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
	}

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	transactionCreated, err := transaction.SaveTransaction(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, transactionCreated.ID))
	responses.JSON(w, http.StatusCreated, transactionCreated)
}

func (server *Server) GetTransactions(w http.ResponseWriter, r *http.Request) {

	transaction := models.Transaction{}

	transactions, err := transaction.FindAllTransactions(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, transactions)
}

func (server *Server) GetTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	transaction := models.Transaction{}

	transactionReceived, err := transaction.FindTransactionByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, transactionReceived)
}

func (server *Server) UpdateTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the transaction id is valid
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//CHeck if the auth token is valid and  get the user id from it
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the transaction exist
	transaction := models.Transaction{}
	err = server.DB.Debug().Model(models.Transaction{}).Where("id = ?", pid).Take(&transaction).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Transaction not found"))
		return
	}

	// Read the data
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	transactionUpdate := models.Transaction{}
	err = json.Unmarshal(body, &transactionUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	transactionUpdate.Prepare()
	err = transactionUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	for i, _ := range transactionUpdate.TransactionDetails {
		transactionUpdate.TransactionDetails[i].Prepare()
		err = transactionUpdate.TransactionDetails[i].Validate()
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
	}

	transactionUpdate.ID = transaction.ID 

	transactionUpdated, err := transactionUpdate.UpdateATransaction(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, transactionUpdated)
}

func (server *Server) DeleteTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid transaction id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Is this user authenticated?
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the transaction exist
	transaction := models.Transaction{}
	err = server.DB.Debug().Model(models.Transaction{}).Where("id = ?", pid).Take(&transaction).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	transactionUpdated, err := transactionUpdate.DeleteATransaction(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, transactionUpdated)
}