package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"media_management_go/backend/common"
	"media_management_go/backend/database"
	"media_management_go/backend/handlers"
)

func main() {
	common.MustLoadConfig()
	common.LoadLogger()

	cfg := common.GetConfig()
	slog.Info("Config loaded")
	slog.Info("Logger loaded", slog.String("env", cfg.ENV))
	slog.Debug("Debug logs enabled")

	database.MustOpen(cfg.DB_PATH)
	defer database.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /login", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/login" {
			http.NotFound(w, r)
			slog.Info("Login endpoint not processed", slog.String("expected", "/login"), slog.String("received", r.URL.Path))
			return
		}

		slog.Info("Processing GET login request")
		handlers.HandleGetLogin(w, r)
	})

	mux.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/login" {
			http.NotFound(w, r)
			slog.Info("Login endpoint not processed", slog.String("expected", "/login"), slog.String("received", r.URL.Path))
			return
		}

		slog.Info("Processing POST login request")
		handlers.HandlePostLogin(w, r)
	})

	mux.HandleFunc("GET /link", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/link" {
			http.NotFound(w, r)
			slog.Info("Link endpoint not processed", slog.String("expected", "/link"), slog.String("received", r.URL.Path))
			return
		}

		slog.Info("Processing GET link request")
		handlers.HandleGetLink(w, r)
	})

	mux.HandleFunc("POST /link", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/link" {
			http.NotFound(w, r)
			slog.Info("Link endpoint not processed", slog.String("expected", "/link"), slog.String("received", r.URL.Path))
			return
		}

		slog.Info("Processing POST link request")
		handlers.HandlePostLink(w, r)
	})

	mux.HandleFunc("GET /note", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/note" {
			http.NotFound(w, r)
			slog.Info("Note endpoint not processed", slog.String("expected", "/note"), slog.String("received", r.URL.Path))
			return
		}

		slog.Info("Processing GET note request")
		handlers.HandleGetNote(w, r)
	})

	mux.HandleFunc("POST /note", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/note" {
			http.NotFound(w, r)
			slog.Info("Note endpoint not processed", slog.String("expected", "/note"), slog.String("received", r.URL.Path))
			return
		}

		slog.Info("Processing POST note request")
		handlers.HandlePostNote(w, r)
	})

	addr := fmt.Sprintf("%s:%s", cfg.ADDR, cfg.PORT)
	slog.Info("Server is running.", slog.String("addr", "http://"+addr))
	http.ListenAndServe(addr, mux)
}
