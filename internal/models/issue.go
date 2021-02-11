package models

import (
    "context"
)

type Issue struct {
    Id uint64 `db:"id" json:"id,omitempty"`
    Number string `db:"number" json:"number,omitempty"`
    // Volume string
    // NoVolume bool
    // DisplayVolumeWithNumber bool
    // IssueID uint64
    // IndiciaPublisherID uint64
    // IndiciaPubNotPrinted bool
    // BrandID uint64
    // NoBrand bool
    // PublicationDate string
    // KeyDate string
    // SortCode string
    // Price string
    // PageCount string
    // PageCountUncertain bool
    // IndiciaFrequency string
    // NoIndiciaFrequency bool
    // Editing string
    // NoEditing bool
    // Notes string
    // // created
    // // modified
    // Deleted bool
    // IsIndexed bool
    // ISBN string
    // ValidISBN bool
    // VariantOfID uint64
    // VariantName string
    // Barcode string
    // NoBarcode bool
    // Title string
    // NoTitle bool
    // OnSaleDate string
    // OnSaleDateUncertain bool
    // Rating string
    // NoRating bool
    // VolumeNotPrinted bool
}

type IssueLite struct {
    Id uint64 `db:"id" json:"id,omitempty"`
    Number string `db:"number" json:"number,omitempty"`
}

type IssueListRequest struct {
    Count uint64 `json:"count"`
    Results []*IssueLite `json:"results"`
}

type IssueRepository interface {
    Insert(ctx context.Context, c *Issue) error
    Update(ctx context.Context, c *Issue) error
    GetById(ctx context.Context, id uint64) (Issue, error)
    InsertOrUpdate(ctx context.Context, c *Issue) error
    List(ctx context.Context, page_token uint64, page_size uint64) ([]*IssueLite, uint64)
}
