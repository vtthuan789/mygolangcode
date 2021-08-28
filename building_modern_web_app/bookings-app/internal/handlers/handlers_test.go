package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/vtthuan789/mygolangcode/building_modern_web_app/bookings-app/internal/driver"
	"github.com/vtthuan789/mygolangcode/building_modern_web_app/bookings-app/internal/models"
)

// type postData struct {
// 	key   string
// 	value string
// }

var theTests = []struct {
	name               string
	url                string
	method             string
	expectedStatusCode int
}{
	{"home", "/", "GET", http.StatusOK},
	{"about", "/about", "GET", http.StatusOK},
	{"generals-quarters", "/generals-quarters", "GET", http.StatusOK},
	{"majors-suite", "/majors-suite", "GET", http.StatusOK},
	{"search-availability", "/search-availability", "GET", http.StatusOK},
	{"contact", "/contact", "GET", http.StatusOK},
	{"non-existent", "/green/eggs/and/ham", "GET", http.StatusNotFound},
	{"choose-room", "/choose-room/1", "GET", http.StatusOK},
	{"book-room", "/book-room", "GET", http.StatusOK},
	{"make-reservation", "/make-reservation", "GET", http.StatusOK},
	{"reservation-summary", "/reservation-summary", "GET", http.StatusOK},
	{"user login", "/user/login", "GET", http.StatusOK},
	{"user logout", "/user/logout", "GET", http.StatusOK},
	{"dashboard", "/admin/dashboard", "GET", http.StatusOK},
	{"reservations-new", "/admin/reservations-new", "GET", http.StatusOK},
	{"reservations-all", "/admin/reservations-all", "GET", http.StatusOK},
	{"reservations-calendar", "/admin/reservations-calendar", "GET", http.StatusOK},
	{"reservations-calendar with query params", "/admin/reservations-calendar?y=2021&m=8", "GET", http.StatusOK},
	{"process-reservation", "/admin/process-reservation/new/1/do", "GET", http.StatusOK},
	{"process-reservation with query params", "/admin/process-reservation/new/1/do?y=2021&m=8", "GET", http.StatusOK},
	{"delete-reservation", "/admin/delete-reservation/new/1/do", "GET", http.StatusOK},
	{"delete-reservation with query params", "/admin/delete-reservation/new/1/do?y=2021&m=8", "GET", http.StatusOK},
	{"show reservation", "/admin/reservations/{src}/1/show", "GET", http.StatusOK},
	// {"make-res", "/make-reservation", "GET", []postData{}, http.StatusOK},
	// {"post-search-avail", "/search-availability", "POST", []postData{
	// 	{key: "start", value: "2020-01-01"},
	// 	{key: "end", value: "2020-01-07"},
	// }, http.StatusOK},
	// {"post-search-avail-json", "/search-availability-json", "POST", []postData{
	// 	{key: "start", value: "2020-01-01"},
	// 	{key: "end", value: "2020-01-07"},
	// }, http.StatusOK},
	// {"post-make-reservation", "/search-availability-json", "POST", []postData{
	// 	{key: "first_name", value: "Tony"},
	// 	{key: "last_name", value: "Stark"},
	// 	{key: "email", value: "tonystark@gmail.com"},
	// 	{key: "phone", value: "999-999-9999"},
	// }, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()

	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		// if e.method == "GET" {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}
		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("for testing %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}
		// } else {
		// 	values := url.Values{}
		// 	for _, x := range e.params {
		// 		values.Add(x.key, x.value)
		// 	}
		// 	resp, err := ts.Client().PostForm(ts.URL+e.url, values)
		// 	if err != nil {
		// 		t.Log(err)
		// 		t.Fatal(err)
		// 	}
		// 	if resp.StatusCode != e.expectedStatusCode {
		// 		t.Errorf("for testing %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		// 	}
		// }
	}

}

func TestRepository_Reservation(t *testing.T) {
	reservation := models.Reservation{
		RoomID: 1,
		Room: models.Room{
			ID:       1,
			RoomName: "General's Quarters",
		},
	}

	req, _ := http.NewRequest("GET", "/make-reservation", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)

	handler := http.HandlerFunc(Repo.Reservation)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Reservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusOK)
	}

	// test where the reservation is not in session
	req, _ = http.NewRequest("GET", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Reservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test non-existent room
	req, _ = http.NewRequest("GET", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()
	reservation.RoomID = 100
	session.Put(ctx, "reservation", reservation)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Reservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}
}

func TestRepository_PostReservation(t *testing.T) {
	reservation := models.Reservation{
		RoomID: 1,
		Room: models.Room{
			ID:       1,
			RoomName: "General's Quarters",
		},
	}
	// 2021-01-01 -- 01/02 03:04:05PM '06 -0700
	layout := "2006-01-02"
	startDate, err := time.Parse(layout, "2099-01-01")
	if err != nil {
		t.Error(err)
		return
	}
	endDate, err := time.Parse(layout, "2099-01-02")
	if err != nil {
		t.Error(err)
		return
	}
	reservation.StartDate = startDate
	reservation.EndDate = endDate
	// reqBody := "first_name=Thuan"
	// reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Vo")
	// reqBody = fmt.Sprintf("%s&%s", reqBody, "email=vtthuan@example.com")
	// reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=999-9999-99999")
	postedData := url.Values{}
	postedData.Add("first_name", "Thuan")
	postedData.Add("last_name", "Vo")
	postedData.Add("email", "vtthuan@example.com")
	postedData.Add("phone", "999-9999-99999")

	req, _ := http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode()))
	ctx := getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)

	handler := http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Errorf("PostReservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
	}

	// test where the reservation is not in session
	req, _ = http.NewRequest("POST", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler failed when testing where the reservation is not in session: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test for missing post body
	req, _ = http.NewRequest("POST", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)

	handler = http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler failed when testing for missing post body: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test for invalid form data, first_name is less than 3 character long
	// reqBody = "first_name=T"
	// reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Vo")
	// reqBody = fmt.Sprintf("%s&%s", reqBody, "email=vtthuan@example.com")
	// reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=999-9999-99999")
	postedData = url.Values{}
	postedData.Add("first_name", "T")
	postedData.Add("last_name", "Vo")
	postedData.Add("email", "vtthuan@example.com")
	postedData.Add("phone", "999-9999-99999")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode()))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)

	handler = http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Errorf("PostReservation handler failed when testing for invalid form data: got %d, wanted %d", rr.Code, http.StatusSeeOther)
	}

	// test for failure for insert reservation into database
	// reqBody = "first_name=Thuan"
	// reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Vo")
	// reqBody = fmt.Sprintf("%s&%s", reqBody, "email=vtthuan@example.com")
	// reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=999-9999-99999")
	postedData = url.Values{}
	postedData.Add("first_name", "Thuan")
	postedData.Add("last_name", "Vo")
	postedData.Add("email", "vtthuan@example.com")
	postedData.Add("phone", "999-9999-99999")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode()))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	reservation.RoomID = 2
	session.Put(ctx, "reservation", reservation)

	handler = http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler failed when testing for fail inserting reservation: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test for failure for insert restriction into database
	// reqBody = "first_name=Thuan"
	// reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Vo")
	// reqBody = fmt.Sprintf("%s&%s", reqBody, "email=vtthuan@example.com")
	// reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=999-9999-99999")
	postedData = url.Values{}
	postedData.Add("first_name", "Thuan")
	postedData.Add("last_name", "Vo")
	postedData.Add("email", "vtthuan@example.com")
	postedData.Add("phone", "999-9999-99999")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(postedData.Encode()))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	reservation.RoomID = 1000
	session.Put(ctx, "reservation", reservation)

	handler = http.HandlerFunc(Repo.PostReservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostReservation handler failed when testing for fail inserting restriction: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}
}

func TestRepository_ReservationSummary(t *testing.T) {
	reservation := models.Reservation{
		RoomID: 1,
		Room: models.Room{
			ID:       1,
			RoomName: "General's Quarters",
		},
		FirstName: "abc",
		LastName:  "def",
		Email:     "abc@ex.com",
	}
	// 2021-01-01 -- 01/02 03:04:05PM '06 -0700
	layout := "2006-01-02"
	startDate, err := time.Parse(layout, "2099-01-01")
	if err != nil {
		t.Error(err)
		return
	}
	endDate, err := time.Parse(layout, "2099-01-02")
	if err != nil {
		t.Error(err)
		return
	}
	reservation.StartDate = startDate
	reservation.EndDate = endDate

	req, _ := http.NewRequest("GET", "/reservation-summary", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)

	handler := http.HandlerFunc(Repo.ReservationSummary)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("ReservationSummary handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusOK)
	}

	// test where the reservation is not in session
	req, _ = http.NewRequest("GET", "/reservation-summary", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("ReservationSummary handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}
}

func TestRepository_PostAvailability(t *testing.T) {
	reqBody := "start=2021-09-09"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=2021-09-10")

	req, _ := http.NewRequest("POST", "/search-availability", strings.NewReader(reqBody))
	ctx := getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Repo.PostAvailability)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("PostAvailability handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusOK)
	}

	// test for missing post body
	req, _ = http.NewRequest("POST", "/search-availability", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostAvailability)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostAvailability handler failed when testing for missing post body: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test for invalid form data, start field is not a date
	reqBody = "start=invalid"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=2021-09-10")
	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(reqBody))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostAvailability)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostAvailability handler failed when testing for missing post body: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test for invalid form data, end field is not a date
	reqBody = "start=2021-09-09"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=invalid")
	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(reqBody))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostAvailability)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostAvailability handler failed when testing for missing post body: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test for no room available
	reqBody = "start=2050-01-01"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=2050-01-02")
	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(reqBody))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostAvailability)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Errorf("PostAvailability handler failed when testing for missing post body: got %d, wanted %d", rr.Code, http.StatusSeeOther)
	}

	// test for database error
	reqBody = "start=2060-01-01"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=2050-01-02")
	req, _ = http.NewRequest("POST", "/search-availability", strings.NewReader(reqBody))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostAvailability)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("PostAvailability handler failed when testing for missing post body: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}
}

func TestRepository_AvailabilityJSON(t *testing.T) {
	// first case - rooms are not available
	reqBody := "start=2050-01-01"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=2050-01-02")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

	// create request
	req, _ := http.NewRequest("POST", "/search-availability-json", strings.NewReader(reqBody))

	// get context with session
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	// set the request header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// make handler function
	handler := http.HandlerFunc(Repo.AvailabilityJSON)

	// get response recorder
	rr := httptest.NewRecorder()

	// make request to our handler
	handler.ServeHTTP(rr, req)

	var j jsonResponse
	err := json.Unmarshal(rr.Body.Bytes(), &j)
	if err != nil {
		t.Error("failed to parse json")
	}

	if j.OK || j.Message != "" {
		t.Error("AvailabilityJSON handler failed when testing where rooms are not available")
	}

	// second case - room is available
	reqBody = "start=2049-12-30"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=2049-12-31")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

	// create request
	req, _ = http.NewRequest("POST", "/search-availability-json", strings.NewReader(reqBody))

	// get context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	// set the request header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// make handler function
	handler = http.HandlerFunc(Repo.AvailabilityJSON)

	// get response recorder
	rr = httptest.NewRecorder()

	// make request to our handler
	handler.ServeHTTP(rr, req)

	err = json.Unmarshal(rr.Body.Bytes(), &j)
	if err != nil {
		t.Error("failed to parse json")
	}

	if !j.OK {
		t.Error("AvailabilityJSON handler failed when testing where room is available")
	}

	// third case - cannot parse the form
	// create request with no body
	req, _ = http.NewRequest("POST", "/search-availability-json", nil)

	// get context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	// set the request header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// make handler function
	handler = http.HandlerFunc(Repo.AvailabilityJSON)

	// get response recorder
	rr = httptest.NewRecorder()

	// make request to our handler
	handler.ServeHTTP(rr, req)

	err = json.Unmarshal(rr.Body.Bytes(), &j)
	if err != nil {
		t.Error("failed to parse json")
	}

	if j.OK || j.Message != "Cannot parse the form" {
		t.Error("Got availability or wrong message when request body was empty")
	}

	// fourth case - invalid start field
	reqBody = "start=invalid"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=2049-12-31")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

	// create request
	req, _ = http.NewRequest("POST", "/search-availability-json", strings.NewReader(reqBody))

	// get context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	// set the request header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// make handler function
	handler = http.HandlerFunc(Repo.AvailabilityJSON)

	// get response recorder
	rr = httptest.NewRecorder()

	// make request to our handler
	handler.ServeHTTP(rr, req)

	err = json.Unmarshal(rr.Body.Bytes(), &j)
	if err != nil {
		t.Error("failed to parse json")
	}

	if j.OK || j.Message != "Cannot parse the start field to the time value" {
		t.Error("Got availability or wrong message when start field is invalid")
	}

	// fifth case - invalid end field
	reqBody = "start=2049-12-30"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=invalid")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=1")

	// create request
	req, _ = http.NewRequest("POST", "/search-availability-json", strings.NewReader(reqBody))

	// get context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	// set the request header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// make handler function
	handler = http.HandlerFunc(Repo.AvailabilityJSON)

	// get response recorder
	rr = httptest.NewRecorder()

	// make request to our handler
	handler.ServeHTTP(rr, req)

	err = json.Unmarshal(rr.Body.Bytes(), &j)
	if err != nil {
		t.Error("failed to parse json")
	}

	if j.OK || j.Message != "Cannot parse the end field to the time value" {
		t.Error("Got availability or wrong message when end field is invalid")
	}

	// sixth case - room id is not a number
	reqBody = "start=2049-12-30"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=2049-12-31")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=invalid")

	// create request
	req, _ = http.NewRequest("POST", "/search-availability-json", strings.NewReader(reqBody))

	// get context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	// set the request header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// make handler function
	handler = http.HandlerFunc(Repo.AvailabilityJSON)

	// get response recorder
	rr = httptest.NewRecorder()

	// make request to our handler
	handler.ServeHTTP(rr, req)

	err = json.Unmarshal(rr.Body.Bytes(), &j)
	if err != nil {
		t.Error("failed to parse json")
	}

	if j.OK || j.Message != "Cannot convert room id field to integer type" {
		t.Error("Got availability or wrong message when room id is not a number")
	}

	// seventh case - database error
	reqBody = "start=2049-12-30"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "end=2049-12-31")
	// set room_id equal to 99 to trap database error
	reqBody = fmt.Sprintf("%s&%s", reqBody, "room_id=99")

	// create request
	req, _ = http.NewRequest("POST", "/search-availability-json", strings.NewReader(reqBody))

	// get context with session
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	// set the request header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// make handler function
	handler = http.HandlerFunc(Repo.AvailabilityJSON)

	// get response recorder
	rr = httptest.NewRecorder()

	// make request to our handler
	handler.ServeHTTP(rr, req)

	err = json.Unmarshal(rr.Body.Bytes(), &j)
	if err != nil {
		t.Error("failed to parse json")
	}

	if j.OK || j.Message != "Error querying database" {
		t.Error("Got availability or wrong message when expecting database error")
	}
}

func TestRepository_BookRoom(t *testing.T) {
	// first case --- normal case
	req, _ := http.NewRequest("GET", "/book-room?s=2050-01-01&e=2050-01-02&id=1", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Repo.BookRoom)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Errorf("BookRoom handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
	}

	// second case --- start date cannot be parsed
	req, _ = http.NewRequest("GET", "/book-room?s=invalid&e=2050-01-02&id=1", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.BookRoom)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("BookRoom handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// third case --- end date cannot be parsed
	req, _ = http.NewRequest("GET", "/book-room?s=2050-01-01&e=invalid&id=1", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.BookRoom)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("BookRoom handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// fourth case --- database error
	req, _ = http.NewRequest("GET", "/book-room?s=2050-01-01&e=2050-01-02&id=99", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)

	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.BookRoom)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("BookRoom handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}
}

func TestRepository_ChooseRoom(t *testing.T) {
	// first case --- normal case
	reservation := models.Reservation{}
	req, _ := http.NewRequest("GET", "/choose-room/1", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)
	// set the RequestURI on the request so that we can grab the ID
	// from the URL
	req.RequestURI = "/choose-room/1"

	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)

	handler := http.HandlerFunc(Repo.ChooseRoom)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Errorf("ChooseRoom handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
	}

	// second case --- room id is not an integer
	req, _ = http.NewRequest("GET", "/choose-room/invalidRoomID", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.RequestURI = "/choose-room/invalidRoomID"

	rr = httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)

	handler = http.HandlerFunc(Repo.ChooseRoom)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("ChooseRoom handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// third case --- reservation is not in session
	req, _ = http.NewRequest("GET", "/choose-room/1", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.RequestURI = "/choose-room/1"

	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.ChooseRoom)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("ChooseRoom handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}
}

func TestRepository_NewRepo(t *testing.T) {
	var db driver.DB
	testRepo := NewRepo(&app, &db)

	if reflect.TypeOf(testRepo).String() != "*handlers.Repository" {
		t.Errorf("Did not get correct type from NewRepo: got %s, wanted *Repository", reflect.TypeOf(testRepo).String())
	}
}

var loginTests = []struct {
	name               string
	email              string
	expectedStatusCode int
	expectedHTML       string
	expectedLocation   string
}{
	{
		"valid-credentials",
		"admin@admin.com",
		http.StatusSeeOther,
		"",
		"/",
	},
	{
		"invalid-credentials",
		"me@example.com",
		http.StatusSeeOther,
		"",
		"/user/login",
	},
	{
		"invalid-data",
		"wrongemail",
		http.StatusOK,
		`action="/user/login"`,
		"",
	},
}

func TestLogin(t *testing.T) {
	// range through all tests
	for _, e := range loginTests {
		postedData := url.Values{}
		postedData.Add("email", e.email)
		postedData.Add("password", "password")

		req, _ := http.NewRequest("POST", "/user/login", strings.NewReader(postedData.Encode()))
		ctx := getCtx(req)
		req = req.WithContext(ctx)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(Repo.PostShowLogin)
		handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedStatusCode {
			t.Errorf("failed %s: got code %d, wanted code %d", e.name, rr.Code, e.expectedStatusCode)
		}

		if e.expectedLocation != "" {
			// get the URL from the test
			actualLoc, _ := rr.Result().Location()
			if actualLoc.String() != e.expectedLocation {
				t.Errorf("failed %s: got URL %s, wanted URL %s", e.name, actualLoc.String(), e.expectedLocation)
			}
		}

		if e.expectedHTML != "" {
			// read the response body into a string
			actualHTML := rr.Body.String()
			if !strings.Contains(actualHTML, e.expectedHTML) {
				t.Errorf("failed %s: expected to find %s, but did not", e.name, e.expectedHTML)
			}
		}
	}
}

func TestRepository_AdminPostShowReservation(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("first_name", "Thuan")
	postedData.Add("last_name", "Vo")
	postedData.Add("email", "admin@admin.com")
	postedData.Add("phone", "999-9999-99999")
	postedData.Add("year", "2021")
	postedData.Add("month", "8")

	req, _ := http.NewRequest("POST", "/admin/reservations/{src}/1", strings.NewReader(postedData.Encode()))
	ctx := getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Repo.AdminPostShowReservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Errorf("AdminPostShowReservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
	}
	// second case - postedDate does not contain year
	postedData = url.Values{}
	postedData.Add("first_name", "Thuan")
	postedData.Add("last_name", "Vo")
	postedData.Add("email", "admin@admin.com")
	postedData.Add("phone", "999-9999-99999")

	req, _ = http.NewRequest("POST", "/admin/reservations/{src}/1", strings.NewReader(postedData.Encode()))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.AdminPostShowReservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Errorf("AdminPostShowReservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
	}
}

func TestRepository_AdminPostReservationsCalendar(t *testing.T) {
	blockMap_1 := make(map[string]int)
	blockMap_1["2021-08-28"] = 0
	blockMap_2 := make(map[string]int)
	blockMap_2["2021-08-28"] = 1
	postedData := url.Values{}
	postedData.Add("first_name", "Thuan")
	postedData.Add("last_name", "Vo")
	postedData.Add("email", "admin@admin.com")
	postedData.Add("phone", "999-9999-99999")
	postedData.Add("year", "2021")
	postedData.Add("month", "8")
	postedData.Add("add_block_1_2021-08-29", "1")

	req, _ := http.NewRequest("POST", "/admin/process-reservation/{src}/1/do", strings.NewReader(postedData.Encode()))
	ctx := getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	session.Put(ctx, "block_map_1", blockMap_1)
	session.Put(ctx, "block_map_2", blockMap_2)

	handler := http.HandlerFunc(Repo.AdminPostReservationsCalendar)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Errorf("AdminPostReservationsCalendar handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
	}

}

func getCtx(req *http.Request) context.Context {
	ctx, err := session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}
	return ctx
}
