package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/lilpipidron/vk-godeveloper-task/server/handler/login"
)

func Start(db *sql.DB) error {
	mux := http.NewServeMux()

	app := new(login.Application)
	app.Auth.Username = "admin"
	app.Auth.Password = "admin"
	mux.HandleFunc("/admin/", app.AdminAuth(db, mux))
	mux.HandleFunc("/user/", app.UserAuth(db, mux))
	log.Println("server listening on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return err
	}

	return nil
}
