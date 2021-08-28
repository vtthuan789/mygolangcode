package dbrepo

import (
	"errors"
	"log"
	"time"

	"github.com/vtthuan789/mygolangcode/building_modern_web_app/bookings-app/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	if res.RoomID == 2 {
		return 0, errors.New("invalid room ID")
	}
	return 1, nil
}

// InsertRoomRestriction inserts a room restriction into database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	if r.RoomID == 1000 {
		return errors.New("invalid room ID")
	}
	return nil
}

// SearchAvailabilityByDatesAndRoomID returns true if availability exists for roomID, otherwise false
func (m *testDBRepo) SearchAvailabilityByDatesAndRoomID(
	start, end time.Time, roomID int) (bool, error) {
	// set up a test time
	layout := "2006-01-02"
	str := "2049-12-31"
	t, err := time.Parse(layout, str)
	if err != nil {
		log.Println(err)
	}

	// this is our test to fail the query -- specify 99 as roomID
	if roomID == 99 {
		return false, errors.New("database error")
	}

	// if the start date is after 2049-12-31, then return false,
	// indicating no availability;
	if start.After(t) {
		return false, nil
	}

	// otherwise, we have availability
	return true, nil
}

// SearchAvailabilityForAllRooms returns a slice of available rooms, if any for a given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	// set up a test time
	layout := "2006-01-02"
	str := "2049-12-31"
	t, err := time.Parse(layout, str)
	if err != nil {
		log.Println(err)
	}

	// this is our test to fail the query -- specify 2060-01-01 as start
	testDateToFail, err := time.Parse(layout, "2060-01-01")
	if err != nil {
		log.Println(err)
	}

	if start == testDateToFail {
		return rooms, errors.New("some error")
	}

	if start.After(t) {
		return rooms, nil
	}

	rooms = append(rooms, models.Room{})
	return rooms, nil
}

// GetRoomByID gets a room by ID
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("invalid ID")
	}
	return room, nil
}

// GetUserByID gets a user by ID
func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	var u models.User

	return u, nil
}

// UpdateUser updates a user in the database
func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}

// Authenticate authenticates a user
func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	if email == "admin@admin.com" {
		return 1, "", nil
	}
	return 0, "", errors.New("invalid email or password")
}

func (m *testDBRepo) AllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation
	return reservations, nil
}

// AllNewReservations gets all new reservations
func (m *testDBRepo) AllNewReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation
	return reservations, nil
}

// GetReservationByID gets a reservation by ID
func (m *testDBRepo) GetReservationByID(id int) (models.Reservation, error) {
	var res models.Reservation
	return res, nil
}

// UpdateReservation updates a reservation in the database
func (m *testDBRepo) UpdateReservation(r models.Reservation) error {
	return nil
}

// DeleteReservation deletes a reservation from the database
func (m *testDBRepo) DeleteReservation(id int) error {
	return nil
}

// UpdateProcessedForReservation updates processed for a reservation by id
func (m *testDBRepo) UpdateProcessedForReservation(id, processed int) error {
	return nil
}

// AllRooms gets all rooms
func (m *testDBRepo) AllRooms() ([]models.Room, error) {
	rooms := []models.Room{
		{
			ID:        1,
			RoomName:  "General's Quarter",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			RoomName:  "Major's Suite",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	return rooms, nil
}

// GetRestrictionsForRoomByDate returns restrictions for a room by date range
func (m *testDBRepo) GetRestrictionsForRoomByDate(roomID int, startDate, endDate time.Time) ([]models.RoomRestriction, error) {
	restrictions := []models.RoomRestriction{
		{
			ID:            1,
			StartDate:     time.Now(),
			EndDate:       time.Now(),
			RoomID:        1,
			ReservationID: 1,
			RestrictionID: 1,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			ID:            2,
			StartDate:     time.Now(),
			EndDate:       time.Now(),
			RoomID:        2,
			ReservationID: 0,
			RestrictionID: 2,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
	}
	return restrictions, nil
}

// InsertBlockForRoom inserts a room restriction
func (m *testDBRepo) InsertBlockForRoom(id int, startDate time.Time) error {
	return nil
}

// DeleteBlockByID deletes a block from room restrictions
func (m *testDBRepo) DeleteBlockByID(id int) error {
	return nil
}
