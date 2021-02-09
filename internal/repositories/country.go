package model_resources

import (
    "context"
	"database/sql"
    "time"

    "github.com/luchacomics/ccdata-server/internal/models"
)

// CountryRepo implements models.CountryRepository
type CountryRepo struct{
    db *sql.DB
}

// Constructor
func NewCountryRepo(db *sql.DB) *CountryRepo {
    return &CountryRepo{
        db: db,
    }
}

func (cr *CountryRepo) Insert(ctx context.Context, m *models.Country) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "INSERT INTO countries (id, code, name) VALUES ($1, $2, $3)"

	stmt, err := cr.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		m.Id,
		m.Code,
		m.Name,
	)
	return err
}

func (cr *CountryRepo) Update(ctx context.Context, m *models.Country) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "UPDATE countries SET name = $1, code = $2 WHERE id = $3"
	stmt, err := cr.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		m.Name,
		m.Code,
		m.Id,
	)
	return err
}


func (r *CountryRepo) GetById(ctx context.Context, id uint64) (*models.Country, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	m := new(models.Country)

	query := "SELECT id, code, name FROM countries WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&m.Id,
		&m.Code,
		&m.Name,
	)
	if err != nil {
		// CASE 1 OF 2: Cannot find record with that email.
		if err == sql.ErrNoRows {
			return nil, nil
		} else { // CASE 2 OF 2: All other errors.
			return nil, err
		}
	}
	return m, nil
}

func (cr *CountryRepo) InsertOrUpdate(ctx context.Context, m *models.Country) error {
    m, err := cr.GetById(ctx, m.Id)
    if err != nil {
        return err
    }

    if m == nil {
        return cr.Insert(ctx, m)
    }
    return cr.Update(ctx, m)
}
