package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var mockDatabase = []Item{
	{
		ID:      1,
		Created: timestamp{time.Date(2019, 9, 05, 00, 00, 00, 00, time.UTC)},
		Good: Good{
			Name:         "CannedOnion",
			Manufactured: timestamp{time.Date(2019, 7, 02, 00, 00, 00, 00, time.UTC)},
			ExpDate:      timestamp{time.Date(2019, 10, 04, 00, 00, 00, 00, time.UTC)},
			ExpOpen:      20,
		},
		IsOpen: true,
		Opened: timestamp{time.Date(2019, 9, 03, 00, 00, 00, 00, time.UTC)},
	},
}

func TestHandleLists(t *testing.T) {

	srv := api{
		router: http.NewServeMux(),
		db:     Memory{mockDatabase},
	}

	srv.routes()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status OK; got %v", res.Status)
	}

	data := []Item{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		t.Fatalf("could not decode the response: %v", err)
	}

	if data[0].Name != mockDatabase[0].Name {
		t.Fatalf("Expected %v name and got %v name", mockDatabase[0].Name, data[0].Name)
	}

}

func TestHandleAdd(t *testing.T) {
	newItem := []byte(`[{
		"id":2,
		"name":"BottleMilk",
		"manufactured":"03-08-2019",
		"expdate":"02-11-2019",
		"expopen":10,
		"isopen":false
		}]`)

	srv := api{
		router: http.NewServeMux(),
		db:     Memory{mockDatabase},
	}

	srv.routes()
	req, err := http.NewRequest("POST", "/add", bytes.NewBuffer(newItem))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

}

func TestHandleOpen(t *testing.T) {
}

func TestHandleDelete(t *testing.T) {
}
