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
    PageCount string `db:"page_count" json:"page_count,omitempty"`
    PageCountUncertain bool `db:"page_count_uncertain" json:"page_count_uncertain,omitempty"`
    IndiciaFrequency string `db:"indicia_frequency" json:"indicia_frequency,omitempty"`
    NoIndiciaFrequency bool `db:"no_indicia_frequency" json:"no_indicia_frequency,omitempty"`
    Editing string `db:"editing" json:"editing,omitempty"`
    NoEditing bool `db:"no_editing" json:"no_editing,omitempty"`
    Notes string `db:"notes" json:"notes,omitempty"`
    // // created
    // // modified
    Deleted bool `db:"deleted" json:"deleted,omitempty"`
    IsIndexed bool `db:"is_indexed" json:"is_indexed,omitempty"`
    Isbn string `db:"isbn" json:"isbn,omitempty"`
    ValidIsbn bool `db:"valid_isbn" json:"valid_isbn,omitempty"`
    NoIsbn bool `db:"no_isbn" json:"no_isbn,omitempty"`
    VariantOfId uint64 `db:"variant_of_id" json:"variant_of_id,omitempty"`
    VariantName string `db:"variant_name" json:"variant_name,omitempty"`
    Barcode string `db:"barcode" json:"barcode,omitempty"`
    NoBarcode bool `db:"no_barcode" json:"no_barcode,omitempty"`
    Title string `db:"title" json:"title,omitempty"`
    NoTitle bool `db:"no_title" json:"no_title,omitempty"`
    OnSaleDate string `db:"on_sale_date" json:"on_sale_date,omitempty"`
    OnSaleDateUncertain bool `db:"on_sale_date_uncertain" json:"on_sale_date_uncertain,omitempty"`
    Rating string `db:"rating" json:"rating,omitempty"`
    NoRating bool `db:"no_rating" json:"no_rating,omitempty"`
    VolumeNotPrinted bool `db:"volume_not_printed" json:"volume_not_printed,omitempty"`
    NoIndiciaPrinter bool `db:"no_indicia_printer" json:"no_indicia_printer,omitempty"`
}

type IssueLite struct {
    Id uint64 `db:"id" json:"id,omitempty"`
    Number string `db:"number" json:"number,omitempty"`
    SeriesId uint64 `db:"series_id" json:"series_id,omitempty"`
    IndiciaPublisherId uint64 `db:"indicia_publisher_id" json:"indicia_publisher_id,omitempty"`
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
