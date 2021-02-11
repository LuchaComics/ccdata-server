package repositories

import (
    "context"
	"database/sql"
    "time"

    "github.com/luchacomics/ccdata-server/internal/models"
)

// PublisherRepo implements models.PublisherRepository
type PublisherRepo struct{
    db *sql.DB
}

// Constructor
func NewPublisherRepo(db *sql.DB) *PublisherRepo {
    return &PublisherRepo{
        db: db,
    }
}

func (cr *PublisherRepo) Insert(ctx context.Context, m *models.Publisher) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "INSERT INTO publishers (id, name) VALUES ($1, $2)"

	stmt, err := cr.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
    _, err = stmt.ExecContext(
		ctx,
		m.Id,
		m.Name,
	)
	return err
}

func (cr *PublisherRepo) Update(ctx context.Context, m *models.Publisher) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "UPDATE publishers SET name = $1 WHERE id = $2"
	stmt, err := cr.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		m.Name,
		m.Id,
	)
	return err
}


func (r *PublisherRepo) GetById(ctx context.Context, id uint64) (*models.Publisher, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	m := new(models.Publisher)

	query := "SELECT id, name FROM publishers WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&m.Id,
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

func (r *PublisherRepo) InsertOrUpdate(ctx context.Context, m *models.Publisher) error {
    if m.Id == 0 {
        return r.Insert(ctx, m)
    }

    found, err := r.GetById(ctx, m.Id)
    if err != nil {
        return err
    }

    if found == nil {
        return r.Insert(ctx, m)
    }
    return r.Update(ctx, m)
}

func (r *PublisherRepo) List(ctx context.Context, pageToken uint64, pageSize uint64) ([]*models.PublisherLite, uint64) {
    // Run the go routine for fetching records data.
    dataCh := make(chan []*models.PublisherLite)
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


func (r *PublisherRepo) listDataRoutine(ctx context.Context, pageToken uint64, pageSize uint64) ([]*models.PublisherLite, error) {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "SELECT id, name FROM publishers WHERE id > $1 ORDER BY id ASC LIMIT $2"
	rows, err := r.db.QueryContext(ctx, query, pageToken, pageSize)
	if err != nil {
		return nil, err
	}

	var arr []*models.PublisherLite
	defer rows.Close()
	for rows.Next() {
		m := new(models.PublisherLite)
		err = rows.Scan(
            &m.Id,
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


func (r *PublisherRepo) listCountRoutine(ctx context.Context) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `SELECT COUNT(id) FROM publishers`
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
