package repository

import (
	"time"

	"github.com/IbrahimMohammedi/Bookings/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (int, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
}
