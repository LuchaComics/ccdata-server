package models

import (
    "context"
)

type Series struct {
    Id uint64 `db:"id" json:"id,omitempty"`
    Name string `db:"name" json:"name,omitempty"`
    SortName string `db:"sort_name" json:"sort_name,omitempty"`
    Format string `db:"format" json:"format,omitempty"`
    YearBegan int64 `db:"year_began" json:"year_began,omitempty"`
    YearBeganUncertain bool `db:"year_began_uncertain" json:"year_began_uncertain,omitempty"`
    YearEnded int64 `db:"year_ended" json:"year_ended,omitempty"`
    YearEndedUncertain bool `db:"year_ended_uncertain" json:"year_ended_uncertain,omitempty"`
    PublicationDates string `db:"publication_dates" json:"publication_dates,omitempty"`
    FirstIssueId uint64 `db:"first_issue_id" json:"first_issue_id,omitempty"`
    LastIssueId uint64 `db:"last_issue_id" json:"last_issue_id,omitempty"`
    IsCurrent bool `db:"is_current" json:"is_current,omitempty"`
    PublisherId uint64 `db:"publisher_id" json:"publisher_id,omitempty"`
    CountryId uint64 `db:"country_id" json:"country_id,omitempty"`
    LanguageId uint64 `db:"language_id" json:"language_id,omitempty"`
    TrackingNotes string `db:"tracking_notes" json:"tracking_notes,omitempty"`
    Notes string `db:"notes" json:"notes,omitempty"`
    HasGallery bool `db:"has_gallery" json:"has_gallery,omitempty"`
    IssueCount int64 `db:"issue_count" json:"issue_count,omitempty"`
    // created
    // modified
    Deleted bool `db:"deleted" json:"deleted,omitempty"`
    HasIndiciaFrequency bool `db:"has_indicia_frequency" json:"has_indicia_frequency,omitempty"`
    HasIsbn bool `db:"has_isbn" json:"has_isbn,omitempty"`
    HasBarcode bool `db:"has_barcode" json:"has_barcode,omitempty"`
    HasIssueTitle bool `db:"has_issue_title" json:"has_issue_title,omitempty"`
    HasVolume bool `db:"has_volume" json:"has_volume,omitempty"`
    IsComicsPublication bool `db:"is_comics_publication" json:"is_comics_publication,omitempty"`
    Color string `db:"color" json:"color,omitempty"`
    Dimensions string `db:"dimensions" json:"dimensions,omitempty"`
    PaperStock string `db:"paper_stock" json:"paper_stock,omitempty"`
    Binding string `db:"binding" json:"binding,omitempty"`
    PublishingFormat string `db:"publishing_format" json:"publishing_format,omitempty"`
    HasRating bool `db:"has_rating" json:"has_rating,omitempty"`
    PublicationTypeId uint64 `db:"publication_type_id" json:"publication_type_id,omitempty"`
    IsSingleton bool `db:"is_singleton" json:"is_singleton,omitempty"`
    HasAboutComics bool `db:"has_about_comics" json:"has_about_comics,omitempty"`
    HasIndiciaPrinter bool `db:"has_indicia_printer" json:"has_indicia_printer,omitempty"`
}

type SeriesLite struct {
    Id uint64 `db:"id" json:"id,omitempty"`
    Name string `db:"name" json:"name,omitempty"`
    YearBegan int64 `db:"year_began" json:"year_began,omitempty`
    YearEnded int64 `db:"year_ended" json:"year_ended,omitempty"`
    FirstIssueId uint64 `db:"first_issue_id" json:"first_issue_id,omitempty"`
    LastIssueId uint64 `db:"last_issue_id" json:"last_issue_id,omitempty"`
    PublisherId uint64 `db:"publisher_id" json:"publisher_id,omitempty"`
    CountryId uint64 `db:"country_id" json:"country_id,omitempty"`
    LanguageId uint64 `db:"language_id" json:"language_id,omitempty"`
    IssueCount int64 `db:"issue_count" json:"issue_count,omitempty"`    
}

type SeriesListRequest struct {
    Count uint64 `json:"count"`
    Results []*SeriesLite `json:"results"`
}

type SeriesRepository interface {
    Insert(ctx context.Context, c *Series) error
    Update(ctx context.Context, c *Series) error
    GetById(ctx context.Context, id uint64) (*Series, error)
    CheckIfExistsById(ctx context.Context, id uint64) (bool, error)
    InsertOrUpdate(ctx context.Context, c *Series) error
    List(ctx context.Context, page_token uint64, page_size uint64) ([]*SeriesLite, uint64)
}
