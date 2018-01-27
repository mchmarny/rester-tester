package ping

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/mchmarny/rester-tester/util"
)

var (
	logger        = log.New(os.Stdout, "[ping] ", log.Lshortfile|log.Ldate|log.Ltime)
	stateInstance *state
)

// LoadRouts loads routes to the passed in router
func LoadRouts(router *mux.Router) {

	host, err := os.Hostname()
	if err != nil {
		host = "undefined"
	}

	stateInstance = &state{
		ID:   util.GetUUIDv4(),
		Host: host,
		Ts:   time.Now().UTC().String(),
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
	ID   string `json:"service_id"`
	Host string `json:"host_name"`
	Ts   string `json:"started_at"`
}
