package repositories

import (
    "context"
	"database/sql"
    "time"

    "github.com/luchacomics/ccdata-server/internal/models"
)

// SeriesRepo implements models.SeriesRepository
type SeriesRepo struct{
    db *sql.DB
}

// Constructor
func NewSeriesRepo(db *sql.DB) *SeriesRepo {
    return &SeriesRepo{
        db: db,
    }
}

func (cr *SeriesRepo) Insert(ctx context.Context, m *models.Series) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
    INSERT INTO series (
        id, name, sort_name, format, year_began, year_began_uncertain,
        year_ended, year_ended_uncertain, publication_dates, first_issue_id,
        last_issue_id, is_current, publisher_id, country_id, language_id,
        tracking_notes, notes, has_gallery, issue_count, deleted,
        has_indicia_frequency, has_isbn, has_barcode, has_issue_title, has_volume,
        is_comics_publication, color, dimensions, paper_stock, binding,
        publishing_format, has_rating, publication_type_id, is_singleton,
        has_about_comics, has_indicia_printer
    ) VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16,
        $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30,
        $31, $32, $33, $34, $35, $36
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
        m.SortName,
        m.Format,
        m.YearBegan,
        m.YearBeganUncertain,
        m.YearEnded,
        m.YearEndedUncertain,
        m.PublicationDates,
        m.FirstIssueId,
        m.LastIssueId,
        m.IsCurrent,
        m.PublisherId,
        m.CountryId,
        m.LanguageId,
        m.TrackingNotes,
        m.Notes,
        m.HasGallery,
        m.IssueCount,
        m.Deleted,
        m.HasIndiciaFrequency,
        m.HasIsbn,
        m.HasBarcode,
        m.HasIssueTitle,
        m.HasVolume,
        m.IsComicsPublication,
        m.Color,
        m.Dimensions,
        m.PaperStock,
        m.Binding,
        m.PublishingFormat,
        m.HasRating,
        m.PublicationTypeId,
        m.IsSingleton,
        m.HasAboutComics,
        m.HasIndiciaPrinter,
	)
	return err
}

func (cr *SeriesRepo) Update(ctx context.Context, m *models.Series) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
    UPDATE series SET
        name = $1, sort_name = $2, format = $3, year_began = $4,
        year_began_uncertain = $5, year_ended = $6, year_ended_uncertain = $7,
        publication_dates = $8, first_issue_id = $9, last_issue_id = $10,
        is_current = $11, publisher_id = $12, country_id = $13, language_id = $14,
        tracking_notes = $15, notes = $16, has_gallery = $17, issue_count = $18,
        deleted = $19, has_indicia_frequency = $20, has_isbn = $21, has_barcode = $22,
        has_issue_title = $23, has_volume = $24, is_comics_publication = $25,
        color = $26, dimensions = $27, paper_stock = $28, binding = $29,
        publishing_format = $30, has_rating = $31, publication_type_id = $32,
        is_singleton = $33, has_about_comics = $34, has_indicia_printer = $35
    WHERE
        id = $36`
	stmt, err := cr.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
        m.Name,
        m.SortName,
        m.Format,
        m.YearBegan,
        m.YearBeganUncertain,
        m.YearEnded,
        m.YearEndedUncertain,
        m.PublicationDates,
        m.FirstIssueId,
        m.LastIssueId,
        m.IsCurrent,
        m.PublisherId,
        m.CountryId,
        m.LanguageId,
        m.TrackingNotes,
        m.Notes,
        m.HasGallery,
        m.IssueCount,
        m.Deleted,
        m.HasIndiciaFrequency,
        m.HasIsbn,
        m.HasBarcode,
        m.HasIssueTitle,
        m.HasVolume,
        m.IsComicsPublication,
        m.Color,
        m.Dimensions,
        m.PaperStock,
        m.Binding,
        m.PublishingFormat,
        m.HasRating,
        m.PublicationTypeId,
        m.IsSingleton,
        m.HasAboutComics,
        m.HasIndiciaPrinter,
        m.Id,
	)
	return err
}


func (r *SeriesRepo) GetById(ctx context.Context, seriesId uint64) (*models.Series, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	m := new(models.Series)

	query := `
    SELECT
        id, name, sort_name, format, year_began, year_began_uncertain,
        year_ended, year_ended_uncertain, publication_dates, first_issue_id,
        last_issue_id, is_current, publisher_id, country_id, language_id,
        tracking_notes, notes, has_gallery, issue_count, deleted,
        has_indicia_frequency, has_isbn, has_barcode, has_issue_title, has_volume,
        is_comics_publication, color, dimensions, paper_stock, binding,
        publishing_format, has_rating, publication_type_id, is_singleton,
        has_about_comics, has_indicia_printer
    FROM series WHERE
        id = $1
    `

	err := r.db.QueryRowContext(ctx, query, seriesId).Scan(
		&m.Id,
        &m.Name,
        &m.SortName,
        &m.Format,
        &m.YearBegan,
        &m.YearBeganUncertain,
        &m.YearEnded,
        &m.YearEndedUncertain,
        &m.PublicationDates,
        &m.FirstIssueId,
        &m.LastIssueId,
        &m.IsCurrent,
        &m.PublisherId,
        &m.CountryId,
        &m.LanguageId,
        &m.TrackingNotes,
        &m.Notes,
        &m.HasGallery,
        &m.IssueCount,
        &m.Deleted,
        &m.HasIndiciaFrequency,
        &m.HasIsbn,
        &m.HasBarcode,
        &m.HasIssueTitle,
        &m.HasVolume,
        &m.IsComicsPublication,
        &m.Color,
        &m.Dimensions,
        &m.PaperStock,
        &m.Binding,
        &m.PublishingFormat,
        &m.HasRating,
        &m.PublicationTypeId,
        &m.IsSingleton,
        &m.HasAboutComics,
        &m.HasIndiciaPrinter,
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

func (r *SeriesRepo) CheckIfExistsById(ctx context.Context, id uint64) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

    var exists bool

    query := `SELECT 1 FROM series WHERE id = $1;`

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

func (r *SeriesRepo) InsertOrUpdate(ctx context.Context, m *models.Series) error {
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

func (r *SeriesRepo) List(ctx context.Context, pageToken uint64, pageSize uint64) ([]*models.SeriesLite, uint64) {
    // Run the go routine for fetching records data.
    dataCh := make(chan []*models.SeriesLite)
    go func() {
        data, _ := r.listDataRoutine(ctx, pageToken, pageSize)
        dataCh <- data[:]
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


func (r *SeriesRepo) listDataRoutine(ctx context.Context, pageToken uint64, pageSize uint64) ([]*models.SeriesLite, error) {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "SELECT id, name, year_began, year_ended, first_issue_id, last_issue_id, publisher_id, country_id, language_id, issue_count FROM series WHERE id > $1 ORDER BY id ASC LIMIT $2"
	rows, err := r.db.QueryContext(ctx, query, pageToken, pageSize)
	if err != nil {
		return nil, err
	}

	var arr []*models.SeriesLite
	defer rows.Close()
	for rows.Next() {
		m := new(models.SeriesLite)
		err = rows.Scan(
            &m.Id,
            &m.Name,
            &m.YearBegan,
            &m.YearEnded,
            &m.FirstIssueId,
            &m.LastIssueId,
            &m.PublisherId,
            &m.CountryId,
            &m.LanguageId,
            &m.IssueCount,
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


func (r *SeriesRepo) listCountRoutine(ctx context.Context) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `SELECT COUNT(id) FROM series`
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
