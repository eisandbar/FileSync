package endpoints

import (
	"fmt"
	"log"
	"io"
	"errors"
	"encoding/json"
	"net/http"
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
	db "eisandbar/filesync/src/server/db"
	utils "eisandbar/filesync/src/server/utils"
)


func RoomsPOST(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating new room.")

	pgdb, err := gorm.Open("postgres", utils.DB_URI)
	if err != nil {
		log.Fatalf("Error connecting to postgres: %v", err)
	}
	defer pgdb.Close()

	var room db.Room

	err = json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	res := pgdb.Create(&room)

	io.WriteString(w, "Room created")

	fmt.Println("Room created.\n", res)
}

func RoomsDEL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := utils.Atoi(params["id"])
	if id < 0 {
		fmt.Fprintf(w, "Bad id")
		return
	}

	fmt.Println("Deleting room ", id)

	pgdb, err := gorm.Open("postgres", utils.DB_URI)
	if err != nil {
		log.Fatalf("Error connecting to postgres: %v", err)
	}
	defer pgdb.Close()
  
	pgdb.Delete(&db.Room{}, id)
	
	io.WriteString(w, "Room deleted\n")
	fmt.Println("Room deleted ", id)
}

func RoomsGET(w http.ResponseWriter, r *http.Request) {
	var room db.Room
	params := mux.Vars(r)
	id := utils.Atoi(params["id"])
	if id < 0 {
		fmt.Fprintf(w, "Bad id")
		return
	}

	fmt.Println("Finding room ", id)

	pgdb, err := gorm.Open("postgres", utils.DB_URI)
	if err != nil {
		log.Fatalf("Error connecting to postgres: %v", err)
	}
	defer pgdb.Close()

	err = pgdb.First(&room, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Fprintln(w, "Room not found")
		return
	}

	resJSON, err := json.Marshal(room)
	if err != nil {
		log.Fatalf("Error converting to json")
	}

	fmt.Fprintf(w, string(resJSON) + "\n")
	fmt.Println("Found room ", id)
}

func RoomsUsersGET(w http.ResponseWriter, r *http.Request) {
	var users []db.User
	params := mux.Vars(r)
	id := utils.Atoi(params["id"])
	if id < 0 {
		fmt.Fprintf(w, "Bad id")
		return
	}

	fmt.Println("Finding users in room ", id)
	
	pgdb, err := gorm.Open("postgres", utils.DB_URI)
	if err != nil {
		log.Fatalf("Error connecting to postgres: %v", err)
	}
	defer pgdb.Close()

	// limit 20, clicking 'see all' will trigger a different handle
	pgdb.Limit(20).Where(&db.User{RoomId: id}).Find(&users)

	resJSON, err := json.Marshal(users)
	if err != nil {
		log.Fatalf("Error converting to json")
	}

	fmt.Fprintf(w, string(resJSON) + "\n")
	fmt.Println("Found users in room  ", id)
}

func RoomsFilesGET(w http.ResponseWriter, r *http.Request) {
	var files []db.File
	params := mux.Vars(r)
	id := utils.Atoi(params["id"])
	if id < 0 {
		fmt.Fprintf(w, "Bad id")
		return
	}

	fmt.Println("Finding files in room ", id)
	
	pgdb, err := gorm.Open("postgres", utils.DB_URI)
	if err != nil {
		log.Fatalf("Error connecting to postgres: %v", err)
	}
	defer pgdb.Close()

	// limit 20, clicking 'see all' will trigger a different handle
	pgdb.Limit(20).Where(&db.File{RoomId: id}).Find(&files)

	resJSON, err := json.Marshal(files)
	if err != nil {
		log.Fatalf("Error converting to json")
	}

	fmt.Fprintf(w, string(resJSON) + "\n")
	fmt.Println("Found files in room  ", id)
}