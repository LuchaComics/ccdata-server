package models

import (
    "context"
)

type Publisher struct {
    Id uint64 `db:"id" json:"id,omitempty"`
    Name string `db:"name" json:"name,omitempty"`
    CountryId uint64 `db:"country_id" json:"country_id,omitempty"`
    YearBegan int64 `db:"year_began" json:"year_began,omitempty"`
    YearBeganUncertain bool `db:"year_began_uncertain" json:"year_began_uncertain,omitempty"`
    YearEnded int64 `db:"year_ended" json:"year_ended,omitempty"`
    YearEndedUncertain bool `db:"year_ended_uncertain" json:"year_ended_uncertain,omitempty"`
    Notes string `db:"notes" json:"notes,omitempty"`
    Url string `db:"url" json:"url,omitempty"`
    BrandCount int64 `db:"brand_count" json:"brand_count,omitempty"`
    IndiciaPublisherCount int64 `db:"indicia_publisher_count" json:"indicia_publisher_count,omitempty"`
    SeriesCount int64 `db:"series_count" json:"series_count,omitempty"`
    // // created
    // // modified
    IssueCount int64 `db:"issue_count" json:"issue_count,omitempty"`
    Deleted bool `db:"deleted" json:"deleted,omitempty"`
    YearOverallBegan int64 `db:"year_overall_began" json:"year_overall_began,omitempty"`
    YearOverallBeganUncertain bool `db:"year_overall_began_uncertain" json:"year_overall_began_uncertain,omitempty"`
    YearOverallEnded int64 `db:"year_overall_ended" json:"year_overall_ended,omitempty"`
    YearOverallEndedUncertain bool `db:"year_overall_ended_uncertain" json:"year_overall_ended_uncertain,omitempty"`
}

type PublisherLite struct {
    Id uint64
    Name string
}

type PublisherListRequest struct {
    Count uint64 `json:"count"`
    Results []*PublisherLite `json:"results"`
}

type PublisherRepository interface {
    Insert(ctx context.Context, c *Publisher) error
    Update(ctx context.Context, c *Publisher) error
    GetById(ctx context.Context, id uint64) (Publisher, error)
    InsertOrUpdate(ctx context.Context, c *Publisher) error
    List(ctx context.Context, page_token uint64, page_size uint64) ([]*PublisherLite, uint64)
}
