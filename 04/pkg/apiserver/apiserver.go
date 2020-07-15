package apiserver

import (
	"encoding/json"
	"github.com/shal/hsa-2020/04/pkg/cache"
	"net/http"
	"strings"
	"time"

	"github.com/shal/hsa-2020/04/pkg/model"
	"github.com/shal/hsa-2020/04/pkg/store"
)

type server struct {
	router *http.ServeMux
	store  store.Store
	cache  cache.Cache
}

func New(store store.Store, cache cache.Cache) *server {
	s := server{
		router: http.NewServeMux(),
		store:  store,
		cache:  cache,
	}

	s.configureRoutes()

	return &s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRoutes() {
	s.router.HandleFunc("/api/v1/04/transactions", s.Transactions)
	s.router.HandleFunc("/api/v1/04/transaction/", s.Transaction)
	s.router.HandleFunc("/api/v1/04/transaction", s.Transaction)
}

func (s *server) Transactions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		transactions, err := s.store.Transaction().All(r.Context())
		if err != nil {
			http.Error(w, "system.unhealthy", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(transactions); err != nil {
			w.Header().Del("Content-Type")
			http.Error(w, "system.unhealthy", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "request.not_allowed", http.StatusMethodNotAllowed)
	}
}

func (s *server) Transaction(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var transaction model.Transaction
		err := json.NewDecoder(r.Body).Decode(&transaction)
		if err != nil {
			http.Error(w, "request.invalid_body", http.StatusBadRequest)
			return
		}

		transaction.Time = time.Now()

		err = s.store.Transaction().Create(r.Context(), &transaction)
		if err != nil {
			http.Error(w, "system.unhealthy", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(transaction); err != nil {
			w.Header().Del("Content-Type")
			http.Error(w, "system.unhealthy", http.StatusInternalServerError)
			return
		}
	case http.MethodGet:
		p := strings.Split(r.URL.Path, "/")
		if len(p) < 6 {
			http.Error(w, "request.invalid_id", http.StatusBadRequest)
			return
		}

		id := p[5]

		transaction, err := s.store.Transaction().FindByID(r.Context(), id)
		if err != nil {
			http.Error(w, "system.unhealthy", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(transaction); err != nil {
			w.Header().Del("Content-Type")
			http.Error(w, "system.unhealthy", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "request.not_allowed", http.StatusMethodNotAllowed)
	}
}
