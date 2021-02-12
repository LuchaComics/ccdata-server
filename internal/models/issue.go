package models

import (
    "context"
)

type Issue struct {
    Id uint64 `db:"id" json:"id,omitempty"`
    Number string `db:"number" json:"number,omitempty"`
    Volume string `db:"volume" json:"volume,omitempty"`
    NoVolume bool `db:"no_volume" json:"no_volume,omitempty"`
    DisplayVolumeWithNumber bool `db:"display_volume_with_number" json:"display_volume_with_number,omitempty"`
    SeriesId uint64 `db:"series_id" json:"series_id,omitempty"`
    IndiciaPublisherId uint64 `db:"indicia_publisher_id" json:"indicia_publisher_id,omitempty"`
    IndiciaPubNotPrinted bool `db:"indicia_pub_not_printed" json:"indicia_pub_not_printed,omitempty"`
    BrandId uint64 `db:"brand_id" json:"brand_id,omitempty"`    
    NoBrand bool `db:"no_brand" json:"no_brand,omitempty"`
    PublicationDate string `db:"publication_date" json:"publication_date,omitempty"`
    KeyDate string `db:"key_date" json:"key_date,omitempty"`
    SortCode string `db:"sort_code" json:"sort_code,omitempty"`
    Price string `db:"price" json:"price,omitempty"`

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
    GetById(ctx context.Context, id uint64) (*Issue, error)
    CheckIfExistsById(ctx context.Context, id uint64) (bool, error)
    InsertOrUpdate(ctx context.Context, c *Issue) error
    List(ctx context.Context, page_token uint64, page_size uint64) ([]*IssueLite, uint64)
}
