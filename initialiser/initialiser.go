package initialiser

import(
"encoding/json"
"fmt"
)

type SampleServer struct {
	Server Server 		`"json" : "Server"`
	Database Database 	`"json": "Database"`
	Logger Logger 		`"json": "logger"`
}

type Server struct{
	port int 	`"json": "port"`
	host string `"json": "host"`

}

type Database struct {
	name string 	`"json": "name"`
	username string `"json": "username"`
	password string `"json": "password"`
	keyspace string `"json": "keyspace"`
	url string 		`"json": "url"`

}

type Logger struct {
	path string 	`"json": "path"`
	loglevel string `"json": "info"`

}

var ServerObject SampleServer