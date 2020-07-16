package apiserver

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/shal/hsa-2020/04/pkg/cache"

	"github.com/shal/hsa-2020/04/pkg/model"
	"github.com/shal/hsa-2020/04/pkg/store"
)

type server struct {
	router  *http.ServeMux
	store   store.Store
	cache   cache.Cache
	enabled bool
}

func New(store store.Store, cache cache.Cache, enabled bool) *server {
	s := server{
		router:  http.NewServeMux(),
		store:   store,
		cache:   cache,
		enabled: enabled,
	}

	rand.Seed(time.Now().UnixNano())
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

// TODO: refactor!
func (s *server) Transactions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if s.enabled {
			result, err := s.cache.Transaction().Get(r.Context())
			if err == cache.ErrNotFound {
				transactions, err := s.store.Transaction().All(r.Context())
				if err != nil {
					http.Error(w, "system.db_failed", http.StatusInternalServerError)
					return
				}

				result = &model.Result{}
				for _, tx := range transactions {
					result.TotalCount++
					result.TotalSum += tx.Amount
				}

				err = s.cache.Transaction().Set(r.Context(), result)
				if err != nil {
					http.Error(w, "system.cache_failed_set", http.StatusInternalServerError)
					return
				}
			} else if err != nil {
				log.Println(err)
				http.Error(w, "system.cache_failed_get", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(result); err != nil {
				w.Header().Del("Content-Type")
				http.Error(w, "system.unhealthy", http.StatusInternalServerError)
				return
			}
		} else {
			transactions, err := s.store.Transaction().All(r.Context())
			if err != nil {
				http.Error(w, "system.db_failed", http.StatusInternalServerError)
				return
			}

			var result model.Result
			for _, tx := range transactions {
				result.TotalCount++
				result.TotalSum += tx.Amount
			}

			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(result); err != nil {
				w.Header().Del("Content-Type")
				http.Error(w, "system.unhealthy", http.StatusInternalServerError)
				return
			}
		}
	default:
		http.Error(w, "request.not_allowed", http.StatusMethodNotAllowed)
	}
}

func (s *server) Transaction(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var transaction model.Transaction

		transaction.Amount = rand.Float64() * 10
		transaction.Time = time.Now()

		err := s.store.Transaction().Create(r.Context(), &transaction)
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
