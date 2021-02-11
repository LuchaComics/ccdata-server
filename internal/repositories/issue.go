package repositories

import (
    // "log"
    "context"
	"database/sql"
    "time"

    "github.com/luchacomics/ccdata-server/internal/models"
)

// IssueRepo implements models.IssueRepository
type IssueRepo struct{
    db *sql.DB
}

// Constructor
func NewIssueRepo(db *sql.DB) *IssueRepo {
    return &IssueRepo{
        db: db,
    }
}

func (cr *IssueRepo) Insert(ctx context.Context, m *models.Issue) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
    INSERT INTO issues (
        id, number
    ) VALUES (
        $1, $2
    )
    `

	stmt, err := cr.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
    _, err = stmt.ExecContext(
		ctx,
		m.Id,
		m.Number,
	)
	return err
}

func (cr *IssueRepo) Update(ctx context.Context, m *models.Issue) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
    UPDATE issues SET
        number = $1
    WHERE
        id = $2`
	stmt, err := cr.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
        m.Number,
		m.Id,
	)
	return err
}


func (r *IssueRepo) GetById(ctx context.Context, issueId uint64) (*models.Issue, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	m := new(models.Issue)

	query := `
    SELECT
        id, number
    FROM issues WHERE
        id = $1
    `

	err := r.db.QueryRowContext(ctx, query, issueId).Scan(
		&m.Id,
		&m.Number,
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

func (r *IssueRepo) InsertOrUpdate(ctx context.Context, m *models.Issue) error {
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

func (r *IssueRepo) List(ctx context.Context, pageToken uint64, pageSize uint64) ([]*models.IssueLite, uint64) {
    // Run the go routine for fetching records data.
    dataCh := make(chan []*models.IssueLite)
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


func (r *IssueRepo) listDataRoutine(ctx context.Context, pageToken uint64, pageSize uint64) ([]*models.IssueLite, error) {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "SELECT id, number FROM issues WHERE id > $1 ORDER BY id ASC LIMIT $2"
	rows, err := r.db.QueryContext(ctx, query, pageToken, pageSize)
	if err != nil {
		return nil, err
	}

	var arr []*models.IssueLite
	defer rows.Close()
	for rows.Next() {
		m := new(models.IssueLite)
		err = rows.Scan(
            &m.Id,
            &m.Number,
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


func (r *IssueRepo) listCountRoutine(ctx context.Context) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `SELECT COUNT(id) FROM issues`
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
