package prime

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	// maxNumber to check for primality - increase for a longer test
	defaultMaxNumber = 50000000
)

var (
	logger = log.New(os.Stdout, "[math] ", log.Lshortfile|log.Ldate|log.Ltime)
)

// LoadRouts loads routes to the passed in router
func LoadRouts(router *mux.Router) {
	router.HandleFunc("/prime", getPrimeEndpoint).Methods("GET")
	router.HandleFunc("/prime/{max:[0-9]+}", getPrimeArgEndpoint).Methods("GET")
}

func getPrimeEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&resp{Prime: *getPrimeResp(defaultMaxNumber)})
}

func getPrimeArgEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	i, err := strconv.Atoi(vars["max"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&resp{Prime: *getPrimeResp(i)})
	}
}

func getPrimeResp(maxNumber int) *prime {

	var x, y, n int
	nsqrt := math.Sqrt(float64(maxNumber))

	isPrime := make([]bool, maxNumber)

	for x = 1; float64(x) <= nsqrt; x++ {
		for y = 1; float64(y) <= nsqrt; y++ {
			n = 4*(x*x) + y*y
			if n <= maxNumber && (n%12 == 1 || n%12 == 5) {
				isPrime[n] = !isPrime[n]
			}
			n = 3*(x*x) + y*y
			if n <= maxNumber && n%12 == 7 {
				isPrime[n] = !isPrime[n]
			}
			n = 3*(x*x) - y*y
			if x > y && n <= maxNumber && n%12 == 11 {
				isPrime[n] = !isPrime[n]
			}
		}
	}

	for n = 5; float64(n) <= nsqrt; n++ {
		if isPrime[n] {
			for y = n * n; y < maxNumber; y += n * n {
				isPrime[y] = false
			}
		}
	}

	isPrime[2] = true
	isPrime[3] = true

	primes := make([]int, 0, 1270606)
	for x = 0; x < len(isPrime)-1; x++ {
		if isPrime[x] {
			primes = append(primes, x)
		}
	}

	return &prime{Max: maxNumber, Value: primes[len(primes)-1]}
}

type resp struct {
	Prime prime `json:"prime"`
}

type prime struct {
	Max   int `json:"max"`
	Value int `json:"val"`
}
