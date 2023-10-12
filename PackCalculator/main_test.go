package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHandleCalculatePacks(t *testing.T) {
	tests := []struct {
		orderQuantity int
		expectedPacks map[int]int
	}{
		{
			orderQuantity: 1,
			expectedPacks: map[int]int{250: 1},
		},
		{
			orderQuantity: 251,
			expectedPacks: map[int]int{500: 1},
		},
		{
			orderQuantity: 500,
			expectedPacks: map[int]int{500: 1},
		},
		{
			orderQuantity: 501,
			expectedPacks: map[int]int{500: 1, 250: 1},
		},
		{
			orderQuantity: 12001,
			expectedPacks: map[int]int{5000: 2, 2000: 1, 250: 1},
		},
	}

	for _, test := range tests {
		reqBody := map[string]int{"orderQuantity": test.orderQuantity}
		reqBodyJSON, _ := json.Marshal(reqBody)
		req, err := http.NewRequest("POST", "/calculate-packs", bytes.NewBuffer(reqBodyJSON))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		http.HandlerFunc(handleCalculatePacks).ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("For orderQuantity %d, expected status code %d but got %d",
				test.orderQuantity, http.StatusOK, rr.Code)
		}

		var response struct {
			PacksNeeded map[int]int `json:"packsNeeded"`
		}
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(response.PacksNeeded, test.expectedPacks) {
			t.Errorf("For orderQuantity %d, expected response %v but got %v",
				test.orderQuantity, test.expectedPacks, response.PacksNeeded)
		}
	}
}
