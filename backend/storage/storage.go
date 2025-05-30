package storage

import "time"

const (
	StatusFetching  int = 1 // 正在抓取
	StatusDraft     int = 2 // 草稿
	StatusCompleted int = 3 // 生成完成
	StatusPublished int = 4 // 已发布
)

type CrossConsultation struct {
	ID        string    `db:"id" json:"id"`
	Prompt    string    `db:"prompt" json:"prompt"`
	Content   string    `db:"content" json:"content"`
	Status    int       `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}

type CrossConsultationFilter struct {
	Status    int       `json:"status"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Offset    int       `json:"offset"`
	Limit     int       `json:"limit"`
}

type Storage interface {
	Close() error

	GetCrossConsultationByID(id string) (*CrossConsultation, error)
	GetCrossConsultationsByStatus(status int) ([]*CrossConsultation, error)
	GetCrossConsultationsByTimeRange(startTime, endTime time.Time) ([]*CrossConsultation, error)
	ListCrossConsultationsByFilter(filter *CrossConsultationFilter) ([]*CrossConsultation, int, error)
	UpdateCrossConsultation(consultation *CrossConsultation) error
	UpdateCrossConsultationById(id string, updater func(old CrossConsultation) (CrossConsultation, error)) error
	DeleteCrossConsultation(id string) error
	CreateCrossConsultation(consultation *CrossConsultation) error
}
