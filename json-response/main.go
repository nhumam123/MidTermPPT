package main

import (
	"log"
	"mahasiswa"
	"net/http"
)

func main() {
	http.HandleFunc("/nama", mahasiswa.CreateMahasiswa)
	http.HandleFunc("/semuadata", mahasiswa.ViewAllMahasiswa)

	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
