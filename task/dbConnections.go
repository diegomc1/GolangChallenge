package task

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct{}

const SERVER = "localhost:27017"

//AddTask inserts a Task in the DB, returns boolean
func (r Repository) AddTask(task Task) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	// Creates ID for task, then inserts task
	task.ID = bson.NewObjectId()
	session.DB(task.Database).C(task.User).Insert(task)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

//UpdateTask updates an Task in the DB, returns boolean
func (r Repository) UpdateTask(task Task) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	//Removes old task using ID
	session.DB(task.Database).C(task.User).RemoveId(task.ID)
	time.Sleep(100 * time.Millisecond)
	//Updates task inserting a new ID
	task.ID = bson.NewObjectId()
	session.DB(task.Database).C(task.User).Insert(task)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

//DeleteTask deletes a Task in the DB, returns boolean
func (r Repository) DeleteTask(task Task) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	//Deletes the task using ID
	session.DB(task.Database).C(task.User).RemoveId(task.ID)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

//DropCollection deletes the collection, returns boolean
func (r Repository) DropCollection(task Task) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	//Drops Collection
	session.DB(task.Database).C(task.User).DropCollection()
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

//DropDatabase deletes the Database, returns boolean
func (r Repository) DropDatabase(drop Drop) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	//Drops Database
	session.DB(drop.Database).DropDatabase()
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
