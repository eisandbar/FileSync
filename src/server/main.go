package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
	db "eisandbar/filesync/src/server/db"
	utils "eisandbar/filesync/src/server/utils"
	ep "eisandbar/filesync/src/server/endpoints"
)

func main() {

    db.InitDB()

	router := mux.NewRouter()

    router.HandleFunc("/rooms", ep.RoomsPOST).Methods("POST") // create room
	router.HandleFunc("/rooms", ep.RoomsDEL).Methods("DELETE") // delete room
    router.HandleFunc("/rooms/{id}", ep.RoomsGET).Methods("GET") // get room info
    router.HandleFunc("/rooms/{id}/users", ep.RoomsUsersGET).Methods("GET") // get users in room
	router.HandleFunc("/rooms/{id}/files", ep.RoomsFilesGET).Methods("GET") // get files in room

	router.HandleFunc("/files", ep.FilesPOST).Methods("POST") // upload a file
	router.HandleFunc("/files/{id}", ep.FilesDEL).Methods("DELETE") // remove a file
	router.HandleFunc("/files/{id}", ep.FilesGET).Methods("GET") // get file info (not data)

	router.HandleFunc("/users", ep.UsersPOST).Methods("POST") // create user
	router.HandleFunc("/users/{id}", ep.UsersDEL).Methods("DELETE") // delete user
	router.HandleFunc("/users/{id}", ep.UsersGET).Methods("GET") // get user info
	router.HandleFunc("/users/{id}/changeRoom", ep.UsersPUT).Methods("PUT") // change the users room

    handler := cors.Default().Handler(router)
    
    fmt.Println("Listening on port:", utils.Server_PORT)

    log.Fatal(http.ListenAndServe(":" + utils.Server_PORT, handler))
}