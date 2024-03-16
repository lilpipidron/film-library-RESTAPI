package server

import (
	"database/sql"
	"github.com/lilpipidron/vk-godeveloper-task/api/server/handler/login"
	"log"
	"net/http"
)

func Start(db *sql.DB) error {

	mux := http.NewServeMux()

	app := new(login.Application)
	app.Auth.Username = "admin"
	app.Auth.Password = "admin"
	mux.HandleFunc("/admin/", app.AdminAuth(db))
	mux.HandleFunc("/user/", app.UserAuth(db))
	log.Println("server listening on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return err
	}

	return nil
}
