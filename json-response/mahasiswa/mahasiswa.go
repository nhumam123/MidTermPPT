package mahasiswa

import (
	"encoding/json"
	"net/http"
)

type Mahasiswa struct {
	Nama   string `json:"Nama"`
	Nim    string `json:"Nim"`
	Alamat string `json:"Alamat"`
}

var dataMahasiswa []Mahasiswa

func CreateMahasiswa(w http.ResponseWriter, r *http.Request) {
	var newMhs Mahasiswa

	nama := r.FormValue("Nama")
	nim := r.FormValue("Nim")
	alamat := r.FormValue("Alamat")

	newMhs = Mahasiswa{
		Nama:   nama,
		Nim:    nim,
		Alamat: alamat,
	}

	dataMahasiswa = append(dataMahasiswa, newMhs)

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMhs)
}

func ViewAllMahasiswa(w http.ResponseWriter, r *http.Request) {
	jsonData, _ := json.Marshal(dataMahasiswa)

	w.Header().Set("Content-Type", "Application/json")
	w.Write(jsonData)
}
