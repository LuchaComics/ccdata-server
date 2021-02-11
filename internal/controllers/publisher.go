package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/luchacomics/ccdata-server/internal/models"
)

func (h *BaseHandler) getPublishers(w http.ResponseWriter, r *http.Request) {
    // Extract the URL parameters.
    ctx := r.Context()
    pageTokenVal := ctx.Value("pageTokenParm").(uint64)
    pageSizeVal := ctx.Value("pageSizeParam").(uint64)

    results, totalCount := h.PublisherRepo.List(ctx, pageTokenVal, pageSizeVal)

    responseData := models.PublisherListRequest{
        Count: totalCount,
        Results: results,
    }
    if err := json.NewEncoder(w).Encode(&responseData); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func (h *BaseHandler) getPublisher(w http.ResponseWriter, r *http.Request, publisherIdStr string) {
    publisherId, err := strconv.ParseUint(publisherIdStr, 10, 64)
    if err != nil {
        http.Error(w, "`" + publisherIdStr + "` url argument is not valid integer for record id", http.StatusBadRequest)
        return
    }

    ctx := r.Context()
    publisher, err := h.PublisherRepo.GetById(ctx, publisherId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if publisher == nil {
        http.Error(w, "record does not exist for id "+publisherIdStr, http.StatusNotFound)
        return
    }
    if err := json.NewEncoder(w).Encode(&publisher); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
