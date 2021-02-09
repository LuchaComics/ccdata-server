package models

//TODO: IMPL.

type Series struct {
    ID uint64
    Name string
    SortName string
    Format string
    YearBegan uint16
    YearBeganUncertain bool
    YearEnded uint16
    YearEndedUncertain bool
    PublicationDates string
    FirstIssueID uint64
    LastIssueID uint64
    IsCurrent bool
    PublisherID uint64
    CountryID uint64
    LanguageID uint64
    TrackingNotes string
    Notes string
    HasGallery bool
    IssueCount uint16
    // created
    // modified
    Deleted bool
    HasIndiciaFrequency bool
    HasISBN bool
    HasBarcode bool
    HasIssueTitle bool
    HasVolume bool
    IsComicsPublication bool
    Color string
    Dimensions string
    PaperStock string
    Binding string
    PublishingFormat string
    HasRating bool
    PublicationTypeID uint64
    IsSingleton bool
}
