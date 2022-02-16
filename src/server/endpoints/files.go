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


func FilesPOST(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adding file")

	pgdb, err := gorm.Open("postgres", utils.DB_URI)
	if err != nil {
		log.Fatalf("Error connecting to postgres: %v", err)
	}
	defer pgdb.Close()

	var file db.File

	err = json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	res := pgdb.Create(&file)

	io.WriteString(w, "File added")
	fmt.Println("File added.\n", res)
}

func FilesDEL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := utils.Atoi(params["id"])
	if id < 0 {
		fmt.Fprintf(w, "Bad id")
		return
	}

	fmt.Println("Deleting file ", id)

	pgdb, err := gorm.Open("postgres", utils.DB_URI)
	if err != nil {
		log.Fatalf("Error connecting to postgres: %v", err)
	}
	defer pgdb.Close()
  
	pgdb.Delete(&db.File{}, id)
	
	io.WriteString(w, "File deleted\n")
	fmt.Println("File deleted ", id)
}

func FilesGET(w http.ResponseWriter, r *http.Request) {
	var file db.File
	params := mux.Vars(r)
	id := utils.Atoi(params["id"])
	if id < 0 {
		fmt.Fprintf(w, "Bad id")
		return
	}

	fmt.Println("Finding file ", id)

	pgdb, err := gorm.Open("postgres", utils.DB_URI)
	if err != nil {
		log.Fatalf("Error connecting to postgres: %v", err)
	}
	defer pgdb.Close()

	err = pgdb.First(&file, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Fprintln(w, "File not found")
		return
	}

	resJSON, err := json.Marshal(file)
	if err != nil {
		log.Fatalf("Error converting to json")
	}

	fmt.Fprintf(w, string(resJSON) + "\n")
	fmt.Println("Found file ", id)
}
