package sql

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/oxidnova/go-kit/x/errorx"
	"github.com/oxidnova/novadm/backend/driver/schema/code"
	"github.com/oxidnova/novadm/backend/storage"
)

// getCrossConsultation retrieves a consultation by ID, supports both transaction and non-transaction operations
func (s *Storage) getCrossConsultation(tx *sqlx.Tx, id string) (*storage.CrossConsultation, error) {
	var consultation storage.CrossConsultation
	query := `SELECT * FROM cross_consultations WHERE id = $1`

	var err error
	if tx != nil {
		err = tx.Get(&consultation, query, id)
	} else {
		err = s.db.Get(&consultation, query, id)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errorx.Errorf(code.NotFound, "consultation not found with id: %s", id)
		}
		return nil, errorx.Errorf(code.Internal, "failed to query consultation: %w", err)
	}
	return &consultation, nil
}

// updateCrossConsultation updates a consultation, supports both transaction and non-transaction operations
func (s *Storage) updateCrossConsultation(tx *sqlx.Tx, consultation *storage.CrossConsultation) error {
	query := `
		UPDATE cross_consultations
		SET prompt = $1, content = $2, status = $3, updated_at = $4
		WHERE id = $5
	`

	var result sql.Result
	var err error
	if tx != nil {
		result, err = tx.Exec(query,
			consultation.Prompt,
			consultation.Content,
			consultation.Status,
			time.Now(),
			consultation.ID,
		)
	} else {
		result, err = s.db.Exec(query,
			consultation.Prompt,
			consultation.Content,
			consultation.Status,
			time.Now(),
			consultation.ID,
		)
	}

	if err != nil {
		return errorx.Errorf(code.Internal, "failed to update consultation: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errorx.Errorf(code.Internal, "failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errorx.Errorf(code.NotFound, "consultation not found with id: %s", consultation.ID)
	}

	return nil
}

// GetCrossConsultationByID retrieves a record by ID
func (s *Storage) GetCrossConsultationByID(id string) (*storage.CrossConsultation, error) {
	return s.getCrossConsultation(nil, id)
}

// UpdateCrossConsultation updates a cross consultation record by ID
func (s *Storage) UpdateCrossConsultation(consultation *storage.CrossConsultation) error {
	return s.updateCrossConsultation(nil, consultation)
}

// UpdateCrossConsultationById updates a cross consultation record by ID using an updater function within a transaction
func (s *Storage) UpdateCrossConsultationById(id string, updater func(old storage.CrossConsultation) (storage.CrossConsultation, error)) error {
	// Start a transaction
	tx, err := s.db.Beginx()
	if err != nil {
		return errorx.Errorf(code.Internal, "failed to begin transaction: %w", err)
	}
	defer tx.Rollback() // Rollback if not committed

	// Get the existing consultation within transaction
	old, err := s.getCrossConsultation(tx, id)
	if err != nil {
		return err
	}

	// Apply the updater function
	new, err := updater(*old)
	if err != nil {
		return errorx.Errorf(code.Internal, "failed to update consultation: %w", err)
	}

	// Update the consultation in database within transaction
	if err := s.updateCrossConsultation(tx, &new); err != nil {
		return err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return errorx.Errorf(code.Internal, "failed to commit transaction: %w", err)
	}

	return nil
}

// GetCrossConsultationsByStatus retrieves records by status
func (s *Storage) GetCrossConsultationsByStatus(status int) ([]*storage.CrossConsultation, error) {
	var consultations []*storage.CrossConsultation
	query := `SELECT * FROM cross_consultations WHERE status = $1`
	err := s.db.Select(&consultations, query, status)
	if err != nil {
		return nil, errorx.Errorf(code.Internal, "failed to query consultations by status: %w", err)
	}
	return consultations, nil
}

// GetCrossConsultationsByTimeRange retrieves records within a time range
func (s *Storage) GetCrossConsultationsByTimeRange(startTime, endTime time.Time) ([]*storage.CrossConsultation, error) {
	var consultations []*storage.CrossConsultation
	query := `SELECT * FROM cross_consultations WHERE created_at BETWEEN $1 AND $2`
	err := s.db.Select(&consultations, query, startTime, endTime)
	if err != nil {
		return nil, errorx.Errorf(code.Internal, "failed to query consultations by time range: %w", err)
	}
	return consultations, nil
}

// ListCrossConsultationsByFilter retrieves records with filter and pagination
func (s *Storage) ListCrossConsultationsByFilter(filter *storage.CrossConsultationFilter) ([]*storage.CrossConsultation, int, error) {
	// Start a transaction
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, 0, errorx.Errorf(code.Internal, "failed to begin transaction: %w", err)
	}
	defer tx.Rollback() // Rollback if not committed

	// Build query conditions
	conditions := make([]string, 0)
	args := make([]interface{}, 0)
	argPos := 1

	if filter.Status > 0 {
		conditions = append(conditions, fmt.Sprintf("status = $%d", argPos))
		args = append(args, filter.Status)
		argPos++
	}

	if !filter.StartTime.IsZero() {
		conditions = append(conditions, fmt.Sprintf("created_at >= $%d", argPos))
		args = append(args, filter.StartTime)
		argPos++
	}

	if !filter.EndTime.IsZero() {
		conditions = append(conditions, fmt.Sprintf("created_at <= $%d", argPos))
		args = append(args, filter.EndTime)
		argPos++
	}

	// Build WHERE clause
	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	// Get total count
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM cross_consultations %s", whereClause)
	var total int
	err = tx.Get(&total, countQuery, args...)
	if err != nil {
		return nil, 0, errorx.Errorf(code.Internal, "failed to get total count: %w", err)
	}

	// Get paginated records
	query := fmt.Sprintf("SELECT * FROM cross_consultations %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d",
		whereClause, argPos, argPos+1)
	args = append(args, filter.Limit, filter.Offset)

	var records []*storage.CrossConsultation
	err = tx.Select(&records, query, args...)
	if err != nil {
		return nil, 0, errorx.Errorf(code.Internal, "failed to query consultations: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, 0, errorx.Errorf(code.Internal, "failed to commit transaction: %w", err)
	}

	return records, total, nil
}

// DeleteCrossConsultation deletes a cross consultation record by ID
func (s *Storage) DeleteCrossConsultation(id string) error {
	query := `DELETE FROM cross_consultations WHERE id = $1`
	result, err := s.db.Exec(query, id)
	if err != nil {
		return errorx.Errorf(code.Internal, "failed to delete consultation: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errorx.Errorf(code.Internal, "failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errorx.Errorf(code.NotFound, "consultation not found with id: %s", id)
	}

	return nil
}

// CreateCrossConsultation creates a new cross consultation record
func (s *Storage) CreateCrossConsultation(consultation *storage.CrossConsultation) error {
	query := `
		INSERT INTO cross_consultations (prompt, content, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	now := time.Now()
	if consultation.CreatedAt.IsZero() {
		consultation.CreatedAt = now
	} else {
		consultation.UpdatedAt = consultation.CreatedAt
	}
	err := s.db.QueryRow(query,
		consultation.Prompt,
		consultation.Content,
		consultation.Status,
		consultation.CreatedAt,
		consultation.UpdatedAt,
	).Scan(&consultation.ID)

	if err != nil {
		return errorx.Errorf(code.Internal, "failed to create consultation: %w", err)
	}

	return nil
}
