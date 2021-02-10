package controllers

import (
    "encoding/json"
    "net/http"

    "github.com/luchacomics/ccdata-server/internal/models"
)

func (h *BaseHandler) getCountries(w http.ResponseWriter, r *http.Request) {
    // Extract the URL parameters.
    ctx := r.Context()
    pageTokenVal := ctx.Value("pageTokenParm").(uint64)
    pageSizeVal := ctx.Value("pageSizeParam").(uint64)

    results, totalCount := h.CountryRepo.List(ctx, pageTokenVal, pageSizeVal)

    responseData := models.CountryListRequest{
        Count: totalCount,
        Results: results,
    }
    if err := json.NewEncoder(w).Encode(&responseData); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
