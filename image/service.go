package image

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/mchmarny/rester-tester/util"
)

// RequestStatusType holds the current request status
type RequestStatusType int

func (s RequestStatusType) String() string {
	names := [...]string{
		"Undefined",
		"Submitted",
		"Processing",
		"Ready",
		"Failed"}

	if s < StatusUndefined || s > StatusFailed {
		return "Unknown"
	}

	return names[s]
}

// IsSet indicates whether the status was actually set
func (s RequestStatusType) IsSet() bool {
	switch s {
	case StatusSubmitted, StatusProcessing, StatusReady, StatusFailed:
		return true
	default:
		return false
	}
}

const (
	// StatusUndefined is the deafult state of status (not set)
	StatusUndefined  RequestStatusType = iota // 0
	StatusSubmitted                           // 1
	StatusProcessing                          // 2
	StatusReady                               // 3
	StatusFailed                              // 4
)

var (
	logger = log.New(os.Stdout, "[image] ", log.Lshortfile|log.Ldate|log.Ltime)
)

// LoadRouts loads routes to the passed in router
func LoadRouts(router *mux.Router) {
	router.HandleFunc("/image", getImageProcessEndpoint).Methods("POST")
}

func getImageProcessEndpoint(w http.ResponseWriter, r *http.Request) {

	state := &resp{
		ID:     util.GetUUIDv4(),
		Ts:     time.Now().UTC().String(),
		Status: "submitted",
	}

	//TODO: capture submitted URL and take it from there

	json.NewEncoder(w).Encode(state)
}

type resp struct {
	ID     string `json:"request_id"`
	Ts     string `json:"created_at"`
	Status string `json:"status"`
}
