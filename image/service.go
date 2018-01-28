package image

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/mchmarny/rester-tester/util"
)

const (
	defaultThumbnailWidth  = 200
	defaultThumbnailHeight = 200
	minThumbnailWidth      = 50
	maxThumbnailWidth      = 500
	minThumbnailHeight     = 50
	maxThumbnailHeight     = 500
)

var (
	logger = log.New(os.Stdout, "[image] ", log.Lshortfile|log.Ldate|log.Ltime)
)

// RequestStatusType holds the current request status
type RequestStatusType int

func (s RequestStatusType) String() string {
	names := [...]string{
		"Undefined",
		"Processed",
		"Failed"}

	if s < StatusUndefined || s > StatusFailed {
		return "Unknown"
	}

	return names[s]
}

// IsSet indicates whether the status was actually set
func (s RequestStatusType) IsSet() bool {
	switch s {
	case StatusProcessed, StatusFailed:
		return true
	default:
		return false
	}
}

const (
	// StatusUndefined is the deafult state of status (not set)
	StatusUndefined RequestStatusType = iota // 0
	StatusProcessed                          // 1
	StatusFailed                             // 2
)

// LoadRouts loads routes to the passed in router
func LoadRouts(router *mux.Router) {
	router.HandleFunc("/image", getImageProcessEndpoint).Methods("POST")
	router.HandleFunc("/image", getImageProcessEndpointInfo).Methods("GET")
}

func getImageProcessEndpointInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&struct {
		Code    string
		Message string
	}{
		"UnsupportedMethod",
		"GET not supported, use POST with video URL as `src` parameter to get a thumbnail",
	})
}

func getImageProcessEndpoint(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var req ThumbnailRequest
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &req)

	state := &ThumbnailResponse{
		ID:      util.GetUUIDv4(),
		Ts:      time.Now().UTC().String(),
		Status:  fmt.Sprint(StatusFailed),
		Request: req,
		Message: "Processing...",
	}

	if state.Request.Src == "" {
		state.Message = "Required parameter missing: src"
		logger.Print(state.Message)
		json.NewEncoder(w).Encode(state)
		return
	}

	thumbPath, err := makeThumbnail(state.ID, &state.Request)
	if err != nil {
		logger.Printf("Error while creating thumbnail: %v -> %v", state, err)
		json.NewEncoder(w).Encode(state)
		return
	}

	// update status
	state.Status = fmt.Sprint(StatusProcessed)
	state.ThumbnailURL = thumbPath //TODO: replace with object store URL

	// TODO: Post to object store
	logger.Printf("Thumbnail: %s", thumbPath)
	json.NewEncoder(w).Encode(state)

}

// ThumbnailRequest represents service request
type ThumbnailRequest struct {
	Src    string `json:"src"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func (r *ThumbnailRequest) String() string {
	return fmt.Sprintf("Src:%s Width:%d Height:%d", r.Src, r.Width, r.Height)
}

// ThumbnailResponse represents service result message
type ThumbnailResponse struct {
	ID           string           `json:"request_id"`
	Ts           string           `json:"created_at"`
	Status       string           `json:"status"`
	Request      ThumbnailRequest `json:"req"`
	Message      string           `json:"message"`
	ThumbnailURL string           `json:"thumbnail"`
}

func (r *ThumbnailResponse) String() string {
	return fmt.Sprintf("ID:%s Ts:%s Status:%v Request:%v", r.ID, r.Ts, r.Status, r.Request)
}
