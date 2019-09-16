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

	srv := newAPI()
	srv.db = Memory{mockDatabase}

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

	srv := newAPI()
	srv.db = Memory{mockDatabase}

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

	if len(srv.db.Items) != 2 {
		t.Fatalf("the item has not be inserted in database")
	}

	if srv.db.Items[1].Name != "BottleMilk" {
		t.Fatalf("The item name expected to be BottleMilk but we got %s ", srv.db.Items[1].Name)
	}

}

func TestHandleOpen(t *testing.T) {

	srv := newAPI()
	srv.db = Memory{mockDatabase}

	req, err := http.NewRequest("GET", "/open?id=1&open=false", nil)
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

	if srv.db.Items[0].IsOpen {
		t.Fatalf("expected item open state to be false but we got %v", srv.db.Items[0].IsOpen)
	}

}

func TestHandleDelete(t *testing.T) {

	srv := newAPI()
	srv.db = Memory{mockDatabase}

	req, err := http.NewRequest("GET", "/del?id=1", nil)
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

	if len(srv.db.Items) != 0 {
		t.Fatalf("there should be 0 Items in database but there are %d", len(srv.db.Items))
	}
}
