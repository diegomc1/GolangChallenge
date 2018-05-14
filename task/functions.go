package task

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"
)

//Controller struct
type Controller struct {
	Repository Repository
}

//AddTask POST
func (c *Controller) AddTask(w http.ResponseWriter, r *http.Request) {
	//Create task structure, then read body
	var task Task
	body, err := ioutil.ReadAll(io.Reader(r.Body))
	//Valudations
	if err != nil {
		log.Fatalln("Error AddTask", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddTask", err)
	}
	if err := json.Unmarshal(body, &task); err != nil { // unmarshall body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddTask unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	//On Success, run AddTask
	success := c.Repository.AddTask(task)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}

//UpdateTask PUT
func (c *Controller) UpdateTask(w http.ResponseWriter, r *http.Request) {
	//Create task structure, then read body
	var task Task
	body, err := ioutil.ReadAll(io.Reader(r.Body))
	//Validations
	if err != nil {
		log.Fatalln("Error UpdateTask", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddaUpdateTask", err)
	}
	if err := json.Unmarshal(body, &task); err != nil {
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error UpdateTask unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	//On success, run UpdateTask
	success := c.Repository.UpdateTask(task)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

//DeleteTask DELETE
func (c *Controller) DeleteTask(w http.ResponseWriter, r *http.Request) {
	//Create task structure, then read body
	var task Task
	body, err := ioutil.ReadAll(io.Reader(r.Body))
	//Validations
	if err != nil {
		log.Fatalln("Error DeleteTask", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error DeleteTask", err)
	}
	if err := json.Unmarshal(body, &task); err != nil {
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error UpdateTask unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	//On success, run DeleteTask
	success := c.Repository.DeleteTask(task)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

//DropDatabase POST
func (c *Controller) DropDatabase(w http.ResponseWriter, r *http.Request) {
	//Create Drop structure, then read body
	var drop Drop
	body, err := ioutil.ReadAll(io.Reader(r.Body))
	//Validations
	if err != nil {
		log.Fatalln("Error DropDatabase", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error DropDatabase", err)
	}
	if err := json.Unmarshal(body, &drop); err != nil {
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error UpdateDropDatabase unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	//On success, run DropDatabase
	success := c.Repository.DropDatabase(drop)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

//DropCollection POST
func (c *Controller) DropCollection(w http.ResponseWriter, r *http.Request) {
	//Create task structure, then read body
	var task Task
	body, err := ioutil.ReadAll(io.Reader(r.Body))
	//Validations
	if err != nil {
		log.Fatalln("Error DropCollection", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error DropCollection", err)
	}
	if err := json.Unmarshal(body, &task); err != nil { // unmarshall body contents as a type Candidate
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error UpdateTask unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	//On success, run function DropCollection
	success := c.Repository.DropCollection(task)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

//Echo POST
func (c *Controller) Echo(w http.ResponseWriter, r *http.Request) {
	//Create keys structure, decode JSON and create results structure
	keys := Keys{}
	err := json.NewDecoder(r.Body).Decode(&keys)
	results := Tasks{}
	//Connect to DB
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	//Query DB to find all results
	resultJson := session.DB(keys.Database).C(keys.User)
	if err := resultJson.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	fmt.Println(results)
	//Parse results
	jsonParse, err := json.Marshal(results)
	if err != nil {
		fmt.Println("error:", err)
	}
	//Set headers and write results
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonParse)
	return

}
