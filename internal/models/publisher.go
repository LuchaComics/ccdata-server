package models

//TODO: IMPL.

type Publisher struct {
    ID uint64
    Name string
    CountryID uint64
    YearBegan uint16
    YearBeganUncertain uint16
    YearEnded uint16
    YearEndedUncertain bool
    Notes string
    URL string
    BrandCount uint16
    IndiciaPublisherCount uint16
    SeriesCount uint16
    // created
    // modified
    IssueCount uint16
    Deleted bool
}
