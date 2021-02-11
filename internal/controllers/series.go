package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/luchacomics/ccdata-server/internal/models"
)

func (h *BaseHandler) getSeries(w http.ResponseWriter, r *http.Request) {
    // Extract the URL parameters.
    ctx := r.Context()
    pageTokenVal := ctx.Value("pageTokenParm").(uint64)
    pageSizeVal := ctx.Value("pageSizeParam").(uint64)

    results, totalCount := h.SeriesRepo.List(ctx, pageTokenVal, pageSizeVal)

    responseData := models.SeriesListRequest{
        Count: totalCount,
        Results: results,
    }
    if err := json.NewEncoder(w).Encode(&responseData); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func (h *BaseHandler) getSingleSeries(w http.ResponseWriter, r *http.Request, seriesIdStr string) {
    seriesId, err := strconv.ParseUint(seriesIdStr, 10, 64)
    if err != nil {
        http.Error(w, "`" + seriesIdStr + "` url argument is not valid integer for record id", http.StatusBadRequest)
        return
    }

    ctx := r.Context()
    series, err := h.SeriesRepo.GetById(ctx, seriesId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if series == nil {
        http.Error(w, "record does not exist for id "+seriesIdStr, http.StatusNotFound)
        return
    }
    if err := json.NewEncoder(w).Encode(&series); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
