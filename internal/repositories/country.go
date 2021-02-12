package repositories

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

func (r *CountryRepo) CheckIfExistsById(ctx context.Context, id uint64) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

    var exists bool

    query := `SELECT 1 FROM countries WHERE id = $1;`

	err := r.db.QueryRowContext(ctx, query, id).Scan(&exists)
	if err != nil {
		// CASE 1 OF 2: Cannot find record with that email.
		if err == sql.ErrNoRows {
			return false, nil
		} else { // CASE 2 OF 2: All other errors.
			return false, err
		}
	}
	return exists, nil
}

func (r *CountryRepo) InsertOrUpdate(ctx context.Context, m *models.Country) error {
    if m.Id == 0 {
        return r.Insert(ctx, m)
    }

    doesExist, err := r.CheckIfExistsById(ctx, m.Id)
    if err != nil {
        return err
    }

    if doesExist == false {
        return r.Insert(ctx, m)
    }
    return r.Update(ctx, m)
}

func (r *CountryRepo) List(ctx context.Context, pageToken uint64, pageSize uint64) ([]*models.Country, uint64) {
    // Run the go routine for fetching records data.
    dataCh := make(chan []*models.Country)
    go func() {
        data, _ := r.listDataRoutine(ctx, pageToken, pageSize)
        dataCh <- data
    }()

    // Run the go routine for fetching records count.
    countCh := make(chan uint64)
    go func() {
        count, _ := r.listCountRoutine(ctx)
        countCh <- count
    }()

    // Block function until goroutines finish.
    dataVal, countVal := <- dataCh, <- countCh

    return dataVal, countVal
}


func (r *CountryRepo) listDataRoutine(ctx context.Context, pageToken uint64, pageSize uint64) ([]*models.Country, error) {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "SELECT id, code, name FROM countries WHERE id > $1 ORDER BY id ASC LIMIT $2"
	rows, err := r.db.QueryContext(ctx, query, pageToken, pageSize)
	if err != nil {
		return nil, err
	}

	var arr []*models.Country
	defer rows.Close()
	for rows.Next() {
		m := new(models.Country)
		err = rows.Scan(
            &m.Id,
    		&m.Code,
    		&m.Name,
		)
		if err != nil {
			return nil, err
		}
		arr = append(arr, m)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return arr, err
}


func (r *CountryRepo) listCountRoutine(ctx context.Context) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `SELECT COUNT(id) FROM countries`
	var count uint64
	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, err
	}
	if err != nil {
		return 0, err
	}
	return count, nil
}
