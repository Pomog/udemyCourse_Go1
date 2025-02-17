package dbrepo

import (
	"errors"
	"fmt"
	"time"
	"udemyCourse1/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// InserReservation inserts a reservation into the database
func (m *testDBRepo) InserReservation(res models.Reservation) (int, error) {
	// if the room id is 2, then fail; otherwise, pass
	if res.RoomID == 2 {
		return 0, errors.New("InserReservation error")
	}
	return 1, nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	return nil
}

// SearchAvailabilityByDatesByRoomID returns true if availability exists for roomID, and false if no availability exists
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

// SearchAvailabilityForAllRooms returns a slice of available rooms, if any, for given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

// GetRoomById gets a room by ID
func (m *testDBRepo) GetRoomById(id int) (models.Room, error) {
	fmt.Println("testDBRepo: GetRoomById: ID = ", id)
	var room models.Room
	if id > 2 {
		fmt.Println("GetRoomById Error - returned")
		return room, errors.New("GetRoomById Error")
	}
	return room, nil
}

func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	return models.User{}, nil
}
func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}
func (m *testDBRepo) Autenticate(email, testPassword string) (int, string, error) {
	if email == "me@here.com" {
		return 1, "", nil
	}
	return 0, "", errors.New("invalid Credentials")
}

// AllReservations returns a slice of all reservations
func (m *testDBRepo) AllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation

	return reservations, nil
}

// AllNewReservations returns a slice of all NON processed reservations
func (m *testDBRepo) AllNewReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation

	return reservations, nil
}

// GetReservationByID returns reservation by ID
func (m *testDBRepo) GetReservationByID(id int) (models.Reservation, error) {
	var i models.Reservation

	return i, nil
}

// UpdateReservation updates the reservation by reservation model in the database
func (m *testDBRepo) UpdateReservation(r models.Reservation) error {
	return nil
}

// DeleteReservation delete the reservation by ID
func (m *testDBRepo) DeleteReservation(id int) error {
	return nil
}

// UpdateProcessedForReservation updates processed for a reservation by ID
func (m *testDBRepo) UpdateProcessedForReservation(id, processed int) error {
	return nil
}

func (m *testDBRepo) AllRooms() ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

// GetRestictionsForRoomByDate returns restrictions for a room by date range
func (m *testDBRepo) GetRestictionsForRoomByDate(roomID int, start, end time.Time) ([]models.RoomRestriction, error) {
	var restrictions []models.RoomRestriction
	return restrictions, nil
}

// insert Room restriction
func (m *testDBRepo) InsertBlockForRoom(id int, startDate time.Time) error {
	return nil
}

// insert Room restriction
func (m *testDBRepo) DeleteBlockForRoom(id int) error {
	return nil
}
