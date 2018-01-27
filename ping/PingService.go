package ping

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var (
	logger        = log.New(os.Stdout, "[ping] ", log.Lshortfile|log.Ldate|log.Ltime)
	stateInstance *state
)

// LoadRouts loads routes to the passed in router
func LoadRouts(router *mux.Router) {

	id, _ := uuid.NewRandom()
	stateInstance = &state{
		ID: id.String(),
		TS: time.Now().UTC().String(),
	}

	router.HandleFunc("/ping", getEndpoint).Methods("GET")
}

func getEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(&resp{Content: *stateInstance})
}

type resp struct {
	Content state `json:"pong"`
}

type state struct {
	ID string `json:"service_id"`
	TS string `json:"created_at"`
}
