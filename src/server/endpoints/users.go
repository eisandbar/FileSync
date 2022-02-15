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


func UsersPOST(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating user")

	pgdb, err := gorm.Open("postgres", utils.DB_URI)
	if err != nil {
		log.Fatalf("Error connecting to postgres: %v", err)
	}
	defer pgdb.Close()

	var user db.User

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	res := pgdb.Create(&user)

	io.WriteString(w, "User created")
	fmt.Println("User created.\n", res)
}

func UsersDEL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := utils.Atoi(params["id"])
	if id < 0 {
		fmt.Fprintf(w, "Bad id")
		return
	}

	fmt.Println("Deleting user ", id)

	pgdb, err := gorm.Open("postgres", utils.DB_URI)
	if err != nil {
		log.Fatalf("Error connecting to postgres: %v", err)
	}
	defer pgdb.Close()
  
	pgdb.Delete(&db.User{}, id)
	
	io.WriteString(w, "User deleted\n")
	fmt.Println("User deleted ", id)
}

func UsersGET(w http.ResponseWriter, r *http.Request) {
	var user db.User
	params := mux.Vars(r)
	id := utils.Atoi(params["id"])
	if id < 0 {
		fmt.Fprintf(w, "Bad id")
		return
	}

	fmt.Println("Finding user ", id)

	pgdb, err := gorm.Open("postgres", utils.DB_URI)
	if err != nil {
		log.Fatalf("Error connecting to postgres: %v", err)
	}
	defer pgdb.Close()

	err = pgdb.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Fprintln(w, "User not found")
		return
	}

	resJSON, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Error converting to json")
	}

	fmt.Fprintf(w, string(resJSON) + "\n")
	fmt.Println("Found user ", id)
}

func UsersPUT(w http.ResponseWriter, r *http.Request) {
	var user db.User
	params := mux.Vars(r)
	id := utils.Atoi(params["id"])
	if id < 0 {
		fmt.Fprintf(w, "Bad id")
		return
	}

	var roomId int64
	err := json.NewDecoder(r.Body).Decode(&roomId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	fmt.Println("Changing user room ", id)

	pgdb, err := gorm.Open("postgres", utils.DB_URI)
	if err != nil {
		log.Fatalf("Error connecting to postgres: %v", err)
	}
	defer pgdb.Close()

	err = pgdb.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Fprintln(w, "User not found")
		return
	}

	user.RoomId = roomId
	pgdb.Save(&user)

	resJSON, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Error converting to json")
	}

	fmt.Fprintf(w, string(resJSON) + "\n")
	fmt.Println("Changed user room ", id)
}