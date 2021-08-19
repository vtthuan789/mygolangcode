package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

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
	reqBody := "first_name=Thuan"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Vo")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=vtthuan@example.com")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=999-9999-99999")

	req, _ := http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
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
	reqBody = "first_name=T"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Vo")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=vtthuan@example.com")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=999-9999-99999")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
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
	reqBody = "first_name=Thuan"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Vo")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=vtthuan@example.com")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=999-9999-99999")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
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
	reqBody = "first_name=Thuan"
	reqBody = fmt.Sprintf("%s&%s", reqBody, "last_name=Vo")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "email=vtthuan@example.com")
	reqBody = fmt.Sprintf("%s&%s", reqBody, "phone=999-9999-99999")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
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

func getCtx(req *http.Request) context.Context {
	ctx, err := session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}
	return ctx
}
