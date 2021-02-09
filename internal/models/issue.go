package models

//TODO: IMPL.

type Issue struct {
    ID uint64
    Number string
    Volume string
    NoVolume bool
    DisplayVolumeWithNumber bool
    SeriesID uint64
    IndiciaPublisherID uint64
    IndiciaPubNotPrinted bool
    BrandID uint64
    NoBrand bool
    PublicationDate string
    KeyDate string
    SortCode string
    Price string
    PageCount string
    PageCountUncertain bool
    IndiciaFrequency string
    NoIndiciaFrequency bool
    Editing string
    NoEditing bool
    Notes string
    // created
    // modified
    Deleted bool
    IsIndexed bool
    ISBN string
    ValidISBN bool
    VariantOfID uint64
    VariantName string
    Barcode string
    NoBarcode bool
    Title string
    NoTitle bool
    OnSaleDate string
    OnSaleDateUncertain bool
    Rating string
    NoRating bool
    VolumeNotPrinted bool
}
