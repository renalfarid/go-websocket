package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api/database"
	"rest-api/helper"

	"github.com/go-chi/chi/v5"
)

type schoolsResource struct{}

type School struct {
	NamaSekolah      string             `json:"nama_sekolah"`
	NPSN             int64              `json:"npsn"`
	Alamat           string             `json:"alamat"`
	BentukPendidikan string             `json:"bentuk_pendidikan"`
	StatusSekolah    string             `json:"status_sekolah"`
	Desa             string             `json:"desa"`
	Kecamatan        string             `json:"kecamatan"`
	KabKota          string             `json:"kab_kota"`
	Propinsi         string             `json:"propinsi"`
	Kodepos          helper.NullFloat64 `json:"kode_pos"`
	Lintang          helper.NullFloat64 `json:"lintang"`
	Bujur            helper.NullFloat64 `json:"bujur"`
}

func (rs schoolsResource) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	r.Get("/", rs.Schools) // GET /users - read a list of users

	return r
}

func (rs schoolsResource) Schools(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT * FROM sekolah")
	if err != nil {
		http.Error(w, "Error executing query", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var schools []School

	for rows.Next() {
		var school School

		if err := rows.Scan(&school.NamaSekolah, &school.NPSN, &school.Alamat,
			&school.BentukPendidikan, &school.StatusSekolah, &school.Desa, &school.Kecamatan,
			&school.KabKota, &school.Propinsi, &school.Kodepos, &school.Lintang, &school.Bujur); err != nil {
			fmt.Println("error schools query:", err)
			http.Error(w, "Error scanning row", http.StatusInternalServerError)
			return
		}

		// Append the struct to the slice.
		schools = append(schools, school)
	}

	// Marshal the results into JSON.
	responseJSON, err := json.Marshal(schools)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json.
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response.
	w.Write(responseJSON)

}
