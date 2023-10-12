package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"

	"github.com/gorilla/mux"
)

var (
	packSizes []int
	packLock  sync.RWMutex
)

func loadPackSizes() {
	packLock.Lock()
	defer packLock.Unlock()

	// Simulate loading pack sizes from a configuration file or database.
	packSizes = []int{250, 500, 1000, 2000, 5000}
}

func calculatePacks(orderQuantity int) map[int]int {
	packLock.RLock()
	defer packLock.RUnlock()

	packCounts := make(map[int]int)

	//Sort the pack sizes in descending order to start with larger packs
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	for _, packSize := range packSizes {
		if orderQuantity <= 0 {
			break
		}

		packsNeeded := orderQuantity / packSize

		if packsNeeded > 0 {
			packCounts[packSize] = packsNeeded
			orderQuantity -= packsNeeded * packSize
		}
	}

	// Check if the remaining quantity can be fulfilled with a single pack
	if orderQuantity > 0 {
		// If the remaining quantity is larger than the smallest pack, use the next larger pack
		if orderQuantity >= packSizes[len(packSizes)-1] {
			nextLargerPack := getNextLargerPackSize(orderQuantity)
			packCounts[nextLargerPack]++
		} else {
			// Otherwise, use the smallest pack
			smallestPackSize := packSizes[len(packSizes)-1]
			packCounts[smallestPackSize]++
		}
	}
	return packCounts
}

func getNextLargerPackSize(quantity int) int {
	for i := len(packSizes) - 1; i >= 0; i-- {
		if packSizes[i] > quantity {
			return packSizes[i]
		}
	}
	return packSizes[0]
}

func handleCalculatePacks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var orderQuantity int
	err := json.NewDecoder(r.Body).Decode(&orderQuantity)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	packCounts := calculatePacks(orderQuantity)

	response := struct {
		PacksNeeded map[int]int `json:"packsNeeded"`
	}{
		PacksNeeded: packCounts,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func main() {
	loadPackSizes()

	r := mux.NewRouter()

	// Serve static files (HTML, CSS, JS)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.HandleFunc("/calculate-packs", handleCalculatePacks)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})

	http.Handle("/", r)

	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
