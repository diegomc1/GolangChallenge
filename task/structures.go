package task

import "gopkg.in/mgo.v2/bson"

//Task structure, which the user is going to fill, represents a task
type Task struct {
	ID          bson.ObjectId `bson:"_id"`
	User        string        `json:"user"`
	Database    string        `json:"database"`
	Task        string        `json:"task"`
	Description string        `json:"description"`
	Date        string        `json:"date"`
}

//Tasks is an array of task
type Tasks []Task

//Structure used for dropping the database
type Drop struct {
	Database string `json:"database"`
}

//Structure that gets Database and User
type Keys struct {
	Database string `json:"database"`
	User     string `json:"user"`
}
