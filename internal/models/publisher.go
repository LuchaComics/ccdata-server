package models

import (
    "context"
)

type Publisher struct {
    Id uint64
    Name string
    // CountryID uint64
    // YearBegan uint16
    // YearBeganUncertain uint16
    // YearEnded uint16
    // YearEndedUncertain bool
    // Notes string
    // URL string
    // BrandCount uint16
    // IndiciaPublisherCount uint16
    // SeriesCount uint16
    // // created
    // // modified
    // IssueCount uint16
    // Deleted bool
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
