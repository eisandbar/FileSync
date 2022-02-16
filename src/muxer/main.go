package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"os/exec"
    "io/ioutil"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

// to convert from mp4 to hls with ffmpeg
// ffmpeg -i filename.mp4 -codec: copy -start_number 0 -hls_time 10 -hls_list_size 0 -f hls filename.m3u8

func main() {
	port := 8080
	router := mux.NewRouter()
	router.HandleFunc("/upload", uploadFile).Methods("POST", "GET")

	handler := cors.Default().Handler(router)

	// serve and log errors
	fmt.Printf("Starting upload server on %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), handler))
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
    fmt.Println("File Upload Endpoint Hit")

    // Parse our multipart form, 10 << 20 specifies a maximum
    // upload of 10 MB files.
    r.ParseMultipartForm(10 << 20)
    // FormFile returns the first file for the given key `myFile`
    // it also returns the FileHeader so we can get the Filename,
    // the Header and the size of the file
    file, handler, err := r.FormFile("myFile")
    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }
    defer file.Close()
    fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

    // read all of the contents of our uploaded file into a
    // byte array
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }

	// Create a temporary file within our temp-images directory
	tempPath := fmt.Sprintf("temp-files/%s", handler.Filename)
    err = os.WriteFile(tempPath, fileBytes, 0644)
	        if err != nil {
        fmt.Println(err)
    }
    defer os.Remove(tempPath)

	MP4ToHLS(handler.Filename)

    // return that we have successfully uploaded our file!
    fmt.Fprintf(w, "Successfully Uploaded File\n")

}

func MP4ToHLS(filename string) {
	fmt.Println("Processing", filename)

	args := ProcessPath(filename)
	fmt.Println("ffmpeg args", args)
	cmd := exec.Command("ffmpeg", strings.Split(args, " ")...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error running ffmpeg")
	}
}

func ProcessPath(filename string) (args string) {
	name := strings.Split(filename, ".mp4")[0]
	outputDir := fmt.Sprintf("../app/files/%s", name)

	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating dir for files")
	}
	fmt.Printf("Made dir %s for file %s\n", outputDir, filename)

	args = fmt.Sprintf("-i temp-files/%s -codec: copy -start_number 0 -hls_time 10 -hls_list_size 0 -f hls ../app/files/%s/%s.m3u8", filename, name, name)
	return
}