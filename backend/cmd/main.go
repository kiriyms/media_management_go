package main

import (
	"fmt"
	"log/slog"
	"media_management_go/backend/common"
	"net/http"
)

func main() {
	common.MustLoadConfig()
	common.LoadLogger()

	cfg := common.GetConfig()
	slog.Info("Config loaded")
	slog.Info("Logger loaded", slog.String("env", cfg.ENV))
	slog.Debug("Debug logs enabled.")

	mux := http.NewServeMux()

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/login" {
			http.NotFound(w, r)
			slog.Info("Login endpoint not processed", slog.String("expected", "/login"), slog.String("received", r.URL.Path))
			return
		}

		switch r.Method {
		case http.MethodPost:
			slog.Info("Processing POST login request")
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			slog.Warn("Login endpoint received unsupported method", slog.String("method", r.Method))
		}
	})

	mux.HandleFunc("/link", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/link" {
			http.NotFound(w, r)
			slog.Info("Link endpoint not processed", slog.String("expected", "/link"), slog.String("received", r.URL.Path))
			return
		}

		switch r.Method {
		case http.MethodGet:
			slog.Info("Processing GET link request")
		case http.MethodPost:
			slog.Info("Processing POST link request")
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			slog.Warn("Link endpoint received unsupported method", slog.String("method", r.Method))
		}
	})

	mux.HandleFunc("/note", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/note" {
			http.NotFound(w, r)
			slog.Info("Note endpoint not processed", slog.String("expected", "/note"), slog.String("received", r.URL.Path))
			return
		}

		switch r.Method {
		case http.MethodGet:
			slog.Info("Processing GET note request")
		case http.MethodPost:
			slog.Info("Processing POST note request")
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			slog.Warn("Note endpoint received unsupported method", slog.String("method", r.Method))
		}
	})

	addr := fmt.Sprintf("%s:%s", cfg.ADDR, cfg.PORT)
	slog.Info("Server is running.", slog.String("addr", "http://"+addr))
	http.ListenAndServe(addr, mux)
}
