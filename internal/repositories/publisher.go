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

	query := `
    INSERT INTO publishers (
        id, name, country_id, year_began, year_began_uncertain, year_ended,
        year_ended_uncertain, notes, url, brand_count, indicia_publisher_count,
        series_count, issue_count, deleted, year_overall_began, year_overall_began_uncertain,
        year_overall_ended, year_overall_ended_uncertain
    ) VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16,
        $17, $18
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
		m.Name,
        m.CountryId,
        m.YearBegan,
        m.YearBeganUncertain,
        m.YearEnded,
        m.YearEndedUncertain,
        m.Notes,
        m.Url,
        m.BrandCount,
        m.IndiciaPublisherCount,
        m.SeriesCount,
        m.IssueCount,
        m.Deleted,
        m.YearOverallBegan,
        m.YearOverallBeganUncertain,
        m.YearOverallEnded,
        m.YearOverallEndedUncertain,
	)
	return err
}

func (cr *PublisherRepo) Update(ctx context.Context, m *models.Publisher) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
    UPDATE publishers SET
        name = $1, country_id = $2, year_began = $3, year_began_uncertain = $4,
        year_ended = $5, year_ended_uncertain = $6, notes = $7, url = $8,
        brand_count = $9, indicia_publisher_count = $10, series_count = $11,
        issue_count = $12, deleted = $13, year_overall_began = $14,
        year_overall_began_uncertain = $15, year_overall_ended = $16,
        year_overall_ended_uncertain = $17
    WHERE
        id = $18`
	stmt, err := cr.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
        m.Name,
        m.CountryId,
        m.YearBegan,
        m.YearBeganUncertain,
        m.YearEnded,
        m.YearEndedUncertain,
        m.Notes,
        m.Url,
        m.BrandCount,
        m.IndiciaPublisherCount,
        m.SeriesCount,
        m.IssueCount,
        m.Deleted,
        m.YearOverallBegan,
        m.YearOverallBeganUncertain,
        m.YearOverallEnded,
        m.YearOverallEndedUncertain,
		m.Id,
	)
	return err
}


func (r *PublisherRepo) GetById(ctx context.Context, publisherId uint64) (*models.Publisher, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	m := new(models.Publisher)

	query := `
    SELECT
        id, name, country_id, year_began, year_began_uncertain, year_ended,
        year_ended_uncertain, notes, url, brand_count, indicia_publisher_count,
        series_count, issue_count, deleted, year_overall_began,
        year_overall_began_uncertain, year_overall_ended,
        year_overall_ended_uncertain
    FROM publishers WHERE
        id = $1
    `

	err := r.db.QueryRowContext(ctx, query, publisherId).Scan(
		&m.Id,
		&m.Name,
        &m.CountryId,
        &m.YearBegan,
        &m.YearBeganUncertain,
        &m.YearEnded,
        &m.YearEndedUncertain,
        &m.Notes,
        &m.Url,
        &m.BrandCount,
        &m.IndiciaPublisherCount,
        &m.SeriesCount,
        &m.IssueCount,
        &m.Deleted,
        &m.YearOverallBegan,
        &m.YearOverallBeganUncertain,
        &m.YearOverallEnded,
        &m.YearOverallEndedUncertain,
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
