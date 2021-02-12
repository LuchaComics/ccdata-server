package models

import (
    "context"
)

type Country struct {
    Id uint64
    Code string
    Name string
}

type CountryRepository interface {
    Insert(ctx context.Context, c *Country) error
    Update(ctx context.Context, c *Country) error
    GetById(ctx context.Context, id uint64) (Country, error)
    CheckIfExistsById(ctx context.Context, id uint64) (Country, error)
    InsertOrUpdate(ctx context.Context, c *Country) error
    List(ctx context.Context, page_token uint64, page_size uint64) ([]*Country, uint64)
}

type CountryListRequest struct {
    Count uint64 `json:"count"`
    Results []*Country `json:"results"`
}
