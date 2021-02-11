package controllers

import (
    "net/http"

    "github.com/luchacomics/ccdata-server/internal/repositories"
	"github.com/luchacomics/ccdata-server/internal/session"
)

type BaseHandler struct {
    SecretSigningKeyBin []byte
    UserRepo *repositories.UserRepo
    CountryRepo *repositories.CountryRepo
    PublisherRepo *repositories.PublisherRepo
    SessionManager *session.SessionManager
}

func NewBaseHandler(keyBin []byte, ur *repositories.UserRepo, cr *repositories.CountryRepo, pr *repositories.PublisherRepo, sm *session.SessionManager) (*BaseHandler) {
    return &BaseHandler{
        SecretSigningKeyBin: keyBin,
        UserRepo: ur,
        CountryRepo: cr,
        PublisherRepo: pr,
        SessionManager: sm,
    }
}

func (h *BaseHandler) HandleRequests(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Get our URL paths which are slash-seperated.
    ctx := r.Context()
    p := ctx.Value("url_split").([]string)
    n := len(p)

    // Get our authorization information.
    isAuthsorized, ok := ctx.Value("is_authorized").(bool)

    switch {
    case n == 1 && p[0] == "version" && r.Method == http.MethodGet:
        if ok && isAuthsorized {
            h.getAuthenticatedVersion(w, r)
        } else {
            h.getVersion(w, r)
        }
    case n == 1 && p[0] == "countries" && r.Method == http.MethodGet:
        if ok && isAuthsorized {
            h.getCountries(w, r)
        } else {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
        }
    case n == 1 && p[0] == "publishers" && r.Method == http.MethodGet:
        if ok && isAuthsorized {
            h.getPublishers(w, r)
        } else {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
        }
    default:
        http.NotFound(w, r)
    }
}
