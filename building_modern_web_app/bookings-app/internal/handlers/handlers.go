package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/vtthuan789/mygolangcode/building_modern_web_app/bookings-app/internal/config"
	"github.com/vtthuan789/mygolangcode/building_modern_web_app/bookings-app/internal/driver"
	"github.com/vtthuan789/mygolangcode/building_modern_web_app/bookings-app/internal/forms"
	"github.com/vtthuan789/mygolangcode/building_modern_web_app/bookings-app/internal/helpers"
	"github.com/vtthuan789/mygolangcode/building_modern_web_app/bookings-app/internal/models"
	"github.com/vtthuan789/mygolangcode/building_modern_web_app/bookings-app/internal/render"
	"github.com/vtthuan789/mygolangcode/building_modern_web_app/bookings-app/internal/repository"
	"github.com/vtthuan789/mygolangcode/building_modern_web_app/bookings-app/internal/repository/dbrepo"
)

// Repo the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo creates new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

// NewTestRepo creates test repository
func NewTestRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewTestingRepo(a),
	}
}

// NewHandlers sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "about.page.tmpl", &models.TemplateData{})
}

// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	res, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.Session.Put(r.Context(), "error", "failed to cast to models.Reservation")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	room, err := m.DB.GetRoomByID(res.RoomID)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	res.Room.RoomName = room.RoomName

	m.App.Session.Put(r.Context(), "reservation", res)

	sd := res.StartDate.Format("2006-01-02")
	ed := res.EndDate.Format("2006-01-02")

	stringMap := make(map[string]string)
	stringMap["start_date"] = sd
	stringMap["end_date"] = ed

	data := make(map[string]interface{})
	data["reservation"] = res

	render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form:      forms.New(nil),
		Data:      data,
		StringMap: stringMap,
	})
}

// PostReservation handles the post request of reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.Session.Put(r.Context(), "error", "failed to cast to models.Reservation")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	err := r.ParseForm()
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't parse form!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	reservation.FirstName = r.Form.Get("first_name")
	reservation.LastName = r.Form.Get("last_name")
	reservation.Email = r.Form.Get("email")
	reservation.Phone = r.Form.Get("phone")

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		http.Error(w, "my own error message for testing", http.StatusSeeOther)
		render.Template(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	newReservationID, err := m.DB.InsertReservation(reservation)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't insert reservation!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	restriction := models.RoomRestriction{
		StartDate:     reservation.StartDate,
		EndDate:       reservation.EndDate,
		RoomID:        reservation.RoomID,
		ReservationID: newReservationID,
		RestrictionID: 1,
	}

	err = m.DB.InsertRoomRestriction(restriction)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't insert room restriction!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)

}

// Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "generals.page.tmpl", &models.TemplateData{})
}

// Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't parse form!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	start := r.Form.Get("start")
	end := r.Form.Get("end")

	// 2021-01-01 -- 01/02 03:04:05PM '06 -0700
	layout := "2006-01-02"
	startDate, err := time.Parse(layout, start)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't parse start date!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	endDate, err := time.Parse(layout, end)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't parse end date!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	rooms, err := m.DB.SearchAvailabilityForAllRooms(startDate, endDate)
	if err != nil {
		m.App.Session.Put(r.Context(), "error", "can't get availability for rooms")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	if len(rooms) == 0 {
		m.App.Session.Put(r.Context(), "error", "No room available")
		http.Redirect(w, r, "/search-availability", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["rooms"] = rooms

	res := models.Reservation{
		StartDate: startDate,
		EndDate:   endDate,
	}
	m.App.Session.Put(r.Context(), "reservation", res)

	render.Template(w, r, "choose-room.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

type jsonResponse struct {
	OK        bool   `json:"ok"`
	Message   string `json:"message"`
	RoomID    string `json:"room_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

// AvailabilityJSON handles request for availability and sends JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	// need to parse request body
	err := r.ParseForm()
	if err != nil {
		// return approriate json
		resp := jsonResponse{
			OK:      false,
			Message: "Cannot parse the form",
		}
		out, _ := json.MarshalIndent(resp, "", "      ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		return
	}

	start := r.Form.Get("start")
	end := r.Form.Get("end")

	// 2021-01-01 -- 01/02 03:04:05PM '06 -0700
	layout := "2006-01-02"
	startDate, err := time.Parse(layout, start)
	if err != nil {
		resp := jsonResponse{
			OK:      false,
			Message: "Cannot parse the start field to the time value",
		}
		out, _ := json.MarshalIndent(resp, "", "      ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		return
	}
	endDate, err := time.Parse(layout, end)
	if err != nil {
		resp := jsonResponse{
			OK:      false,
			Message: "Cannot parse the end field to the time value",
		}
		out, _ := json.MarshalIndent(resp, "", "      ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		return
	}
	roomIDString := r.Form.Get("room_id")
	roomID, err := strconv.Atoi(roomIDString)
	if err != nil {
		resp := jsonResponse{
			OK:      false,
			Message: "Cannot convert room id field to integer type",
		}
		out, _ := json.MarshalIndent(resp, "", "      ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		return
	}

	available, err := m.DB.SearchAvailabilityByDatesAndRoomID(startDate, endDate, roomID)
	if err != nil {
		resp := jsonResponse{
			OK:      false,
			Message: "Error querying database",
		}
		out, _ := json.MarshalIndent(resp, "", "      ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		return
	}
	log.Println("available:", available)
	resp := jsonResponse{
		OK:        available,
		Message:   "",
		StartDate: start,
		EndDate:   end,
		RoomID:    roomIDString,
	}

	// Remove the error check otherwise it's a never ending chain
	out, _ := json.MarshalIndent(resp, "", "     ")

	fmt.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// ReservationSummary renders the reservation summary page
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.Session.Put(r.Context(), "error", "Cannot get reservation from session!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation

	sd := reservation.StartDate.Format("2006-01-02")
	ed := reservation.EndDate.Format("2006-01-02")

	stringMap := make(map[string]string)
	stringMap["start_date"] = sd
	stringMap["end_date"] = ed

	render.Template(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data:      data,
		StringMap: stringMap,
	})
}

// ChooseRoom displays list of available rooms
func (m *Repository) ChooseRoom(w http.ResponseWriter, r *http.Request) {
	roomID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	res, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(w, errors.New("failed to cast to models.Reservation"))
		return
	}

	res.RoomID = roomID

	m.App.Session.Put(r.Context(), "reservation", res)

	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)
}

// BookRoom builds a session variable from URL params and redriect to make reservation screen
func (m *Repository) BookRoom(w http.ResponseWriter, r *http.Request) {
	// id, s, e
	roomID, _ := strconv.Atoi(r.URL.Query().Get("id"))
	sd := r.URL.Query().Get("s")
	ed := r.URL.Query().Get("e")

	var res models.Reservation

	// 2021-01-01 -- 01/02 03:04:05PM '06 -0700
	layout := "2006-01-02"
	startDate, err := time.Parse(layout, sd)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	endDate, err := time.Parse(layout, ed)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	room, err := m.DB.GetRoomByID(roomID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	res.RoomID = roomID
	res.StartDate = startDate
	res.EndDate = endDate
	res.Room.RoomName = room.RoomName

	m.App.Session.Put(r.Context(), "reservation", res)
	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)
}
