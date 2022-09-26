package dbrepo

import (
	"context"
	"time"

	"github.com/hamza-starcevic/bookings/pkg/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

func (m *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int

	statement := `insert into reservations (first_name, last_name, email, phone, start_date,
				  end_date, room_id, created_at, updated_at)
				  values ($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id`

	err := m.DB.QueryRowContext(ctx, statement,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

// * InsertRoomRestriction inserts a room restriction into the database
func (m *postgresDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	statement := `INSERT into room_restrictions (start_date, end_date, room_id,
				  restriction_id, created_at, updated_at, reservation_id)
				  values ($1,$2,$3,$4,$5,$6,$7)`
	_, err := m.DB.ExecContext(ctx, statement,
		r.StartDate,
		r.EndDate,
		r.RoomID,
		r.RestrictionID,
		time.Now(),
		time.Now(),
		r.ReservationID,
	)
	if err != nil {
		return err
	}
	return nil
}

// * SearchAvailabilityByDates returns true if availability exists for roomID, and false if no availability
func (m *postgresDBRepo) SearchAvailabilityByDates(start, end time.Time, roomID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var numRows int

	query := `
			SELECT
				count(id) 
			FROM
				room_restrictions
			WHERE
				room_id = $1
				AND $1 < end_date AND $2 > start_date`

	row := m.DB.QueryRowContext(ctx, query, start, end)
	err := row.Scan(&numRows)
	if err != nil {
		return false, err
	}

	if numRows == 0 {
		return true, nil
	}
	return false, nil
}
