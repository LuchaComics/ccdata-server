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
        id, number, volume, no_volume, display_volume_with_number, series_id,
        indicia_publisher_id, indicia_pub_not_printed, brand_id, no_brand,
        publication_date, key_date, sort_code, price, page_count,
        page_count_uncertain, indicia_frequency, no_indicia_frequency,
        editing, no_editing, notes, deleted, is_indexed, isbn, valid_isbn,
        no_isbn, variant_of_id, variant_name, barcode, no_barcode, title,
        no_title, on_sale_date, on_sale_date_uncertain, rating, no_rating,
        volume_not_printed, no_indicia_printer
    ) VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16,
        $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30,
        $31, $32, $33, $34, $35, $36, $37, $38
    )
    `

	stmt, err := cr.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
    _, err = stmt.ExecContext(
		ctx,
		m.Id, // 1
		m.Number,
		m.Volume,
		m.NoVolume,
		m.DisplayVolumeWithNumber,
		m.SeriesId,
		m.IndiciaPublisherId,
		m.IndiciaPubNotPrinted,
		m.BrandId,
        m.NoBrand, // 10
        m.PublicationDate,
        m.KeyDate,
        m.SortCode,
        m.Price,
        m.PageCount,
        m.PageCountUncertain,
        m.IndiciaFrequency,
        m.NoIndiciaFrequency,
        m.Editing,
        m.NoEditing, // 20
        m.Notes,
        m.Deleted,
        m.IsIndexed,
        m.Isbn,
        m.ValidIsbn,
        m.NoIsbn,
        m.VariantOfId,
        m.VariantName,
        m.Barcode,
        m.NoBarcode, // 30
        m.Title,
        m.NoTitle,
        m.OnSaleDate,
        m.OnSaleDateUncertain,
        m.Rating,
        m.NoRating,
        m.VolumeNotPrinted,
        m.NoIndiciaPrinter, // 38
	)
	return err
}

func (cr *IssueRepo) Update(ctx context.Context, m *models.Issue) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
    UPDATE issues SET
        number = $1, volume = $2, no_volume = $3, display_volume_with_number = $4,
        series_id = $5, indicia_publisher_id = $6, indicia_pub_not_printed = $7,
        brand_id = $8, no_brand = $9, publication_date = $10, key_date = $11,
        sort_code = $12, price = $13, page_count = $14, page_count_uncertain = $15,
        indicia_frequency = $16, no_indicia_frequency = $17, editing = $18,
        no_editing = $19, notes = $20, deleted = $21, is_indexed = $22,
        isbn = $23, valid_isbn = $24, no_isbn = $25, variant_of_id = $26,
        variant_name = $27, barcode = $28, no_barcode = $29, title = $30,
        no_title = $31, on_sale_date = $32, on_sale_date_uncertain = $33,
        rating = $34, no_rating = $35, volume_not_printed = $36,
        no_indicia_printer = $37
    WHERE
        id = $38`
	stmt, err := cr.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
        m.Number,
		m.Volume,
		m.NoVolume,
		m.DisplayVolumeWithNumber,
		m.SeriesId,
		m.IndiciaPublisherId,
		m.IndiciaPubNotPrinted,
		m.BrandId,
        m.NoBrand,
        m.PublicationDate,
        m.KeyDate,
        m.SortCode,
        m.Price,
        m.PageCount,
        m.PageCountUncertain,
        m.IndiciaFrequency,
        m.NoIndiciaFrequency,
        m.Editing,
        m.NoEditing,
        m.Notes,
        m.Deleted,
        m.IsIndexed,
        m.Isbn,
        m.ValidIsbn,
        m.NoIsbn,
        m.VariantOfId,
        m.VariantName,
        m.Barcode,
        m.NoBarcode,
        m.Title,
        m.NoTitle,
        m.OnSaleDate,
        m.OnSaleDateUncertain,
        m.Rating,
        m.NoRating,
        m.VolumeNotPrinted,
        m.NoIndiciaPrinter,
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
        id, number, volume, no_volume, display_volume_with_number, series_id,
        indicia_publisher_id, indicia_pub_not_printed, brand_id, no_brand,
        publication_date, key_date, sort_code, price, page_count,
        page_count_uncertain, indicia_frequency, no_indicia_frequency,
        editing, no_editing, notes, deleted, is_indexed, isbn, valid_isbn,
        no_isbn, variant_of_id, variant_name, barcode, no_barcode, title,
        no_title, on_sale_date, on_sale_date_uncertain, rating, no_rating,
        volume_not_printed, no_indicia_printer
    FROM issues WHERE
        id = $1
    `

	err := r.db.QueryRowContext(ctx, query, issueId).Scan(
		&m.Id,
		&m.Number,
        &m.Volume,
		&m.NoVolume,
		&m.DisplayVolumeWithNumber,
		&m.SeriesId,
		&m.IndiciaPublisherId,
		&m.IndiciaPubNotPrinted,
		&m.BrandId,
        &m.NoBrand,
        &m.PublicationDate,
        &m.KeyDate,
        &m.SortCode,
        &m.Price,
        &m.PageCount,
        &m.PageCountUncertain,
        &m.IndiciaFrequency,
        &m.NoIndiciaFrequency,
        &m.Editing,
        &m.NoEditing,
        &m.Notes,
        &m.Deleted,
        &m.IsIndexed,
        &m.Isbn,
        &m.ValidIsbn,
        &m.NoIsbn,
        &m.VariantOfId,
        &m.VariantName,
        &m.Barcode,
        &m.NoBarcode,
        &m.Title,
        &m.NoTitle,
        &m.OnSaleDate,
        &m.OnSaleDateUncertain,
        &m.Rating,
        &m.NoRating,
        &m.VolumeNotPrinted,
        &m.NoIndiciaPrinter,
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

func (r *IssueRepo) CheckIfExistsById(ctx context.Context, id uint64) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

    var exists bool

    query := `SELECT 1 FROM issues WHERE id = $1;`

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

func (r *IssueRepo) InsertOrUpdate(ctx context.Context, m *models.Issue) error {
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

	query := "SELECT id, number, series_id, indicia_publisher_id FROM issues WHERE id > $1 ORDER BY id ASC LIMIT $2"
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
            &m.SeriesId,
            &m.IndiciaPublisherId,
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
