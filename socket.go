package main

import (
	"encoding/json"
	"fmt"
	"log"
	"rest-api/database"
	"time"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

type SchoolRealtime struct {
	NamaSekolah string `json:"nama_sekolah"`
	NPSN        int64  `json:"npsn"`
	Alamat      string `json:"alamat"`
}

func newServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleScools(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client to schools feed:", ws.RemoteAddr())
	defer ws.Close()

	for {
		rows, err := database.DB.Query("SELECT nama_sekolah, npsn, alamat FROM sekolah")
		if err != nil {
			log.Println("Error executing query:", err)
			return
		}
		defer rows.Close()

		var schools []SchoolRealtime

		for rows.Next() {
			var school SchoolRealtime

			if err := rows.Scan(&school.NamaSekolah, &school.NPSN, &school.Alamat); err != nil {
				log.Println("Error scanning row:", err)
				return
			}

			// Append the struct to the slice.
			schools = append(schools, school)
		}

		// Marshal the results into JSON.
		responseJSON, err := json.Marshal(schools)
		if err != nil {
			log.Println("Error encoding JSON:", err)
			return
		}

		// Write the JSON response to the WebSocket connection.
		if err := websocket.Message.Send(ws, string(responseJSON)); err != nil {
			log.Println("Error writing to WebSocket:", err)
			return
		}

		time.Sleep(time.Second * 2)
	}
}
