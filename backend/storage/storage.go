package storage

import "time"

type CrossConsultation struct {
	ID        string    `db:"id"`
	Prompt    string    `db:"prompt"`
	Content   string    `db:"content"`
	Status    int16     `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Storage interface {
	Close() error

	GetCrossConsultationByID(id string) (*CrossConsultation, error)
	GetCrossConsultationsByStatus(status int) ([]*CrossConsultation, error)
	GetCrossConsultationsByTimeRange(startTime, endTime time.Time) ([]*CrossConsultation, error)
	UpdateCrossConsultation(consultation *CrossConsultation) error
	UpdateCrossConsultationById(id string, updater func(old CrossConsultation) (CrossConsultation, error)) error
	DeleteCrossConsultation(id string) error
	CreateCrossConsultation(consultation *CrossConsultation) error
}
