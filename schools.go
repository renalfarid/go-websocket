package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api/database"
	"rest-api/helper"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type schoolsResource struct{}

type School struct {
	Id               int64              `json:"id"`
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

	r.Get("/", rs.Schools)
	r.Delete("/{id}", rs.DeleteSchool)

	return r
}

func (rs schoolsResource) Schools(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT * FROM sekolah ORDER BY id")
	if err != nil {
		http.Error(w, "Error executing query", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var schools []School

	for rows.Next() {
		var school School

		if err := rows.Scan(&school.Id, &school.NamaSekolah, &school.NPSN, &school.Alamat,
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

func (rs schoolsResource) DeleteSchool(w http.ResponseWriter, r *http.Request) {
	// Extract school ID from URL parameters
	schoolIDStr := chi.URLParam(r, "id")
	schoolID, err := strconv.Atoi(schoolIDStr)
	if err != nil {
		http.Error(w, "Invalid school ID", http.StatusBadRequest)
		return
	}

	// Perform deletion operation in the database
	err = deleteSchoolByID(schoolID)
	if err != nil {
		http.Error(w, "Error deleting school", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("School deleted successfully"))
}

func deleteSchoolByID(schoolID int) error {
	// Implement the logic to delete the school by its ID in the database
	// You should replace this with your actual database deletion logic

	// Example: Delete school from a hypothetical 'schools' table
	_, err := database.DB.Exec("DELETE FROM sekolah WHERE id = ?", schoolID)
	if err != nil {
		fmt.Println("Error deleting school:", err)
		return err
	}

	return nil
}
