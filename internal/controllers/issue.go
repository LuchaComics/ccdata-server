package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/luchacomics/ccdata-server/internal/models"
)

func (h *BaseHandler) getIssues(w http.ResponseWriter, r *http.Request) {
    // Extract the URL parameters.
    ctx := r.Context()
    pageTokenVal := ctx.Value("pageTokenParm").(uint64)
    pageSizeVal := ctx.Value("pageSizeParam").(uint64)

    results, totalCount := h.IssueRepo.List(ctx, pageTokenVal, pageSizeVal)

    responseData := models.IssueListRequest{
        Count: totalCount,
        Results: results,
    }
    if err := json.NewEncoder(w).Encode(&responseData); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func (h *BaseHandler) getIssue(w http.ResponseWriter, r *http.Request, issueIdStr string) {
    issueId, err := strconv.ParseUint(issueIdStr, 10, 64)
    if err != nil {
        http.Error(w, "`" + issueIdStr + "` url argument is not valid integer for record id", http.StatusBadRequest)
        return
    }

    ctx := r.Context()
    issue, err := h.IssueRepo.GetById(ctx, issueId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if issue == nil {
        http.Error(w, "record does not exist for id "+issueIdStr, http.StatusNotFound)
        return
    }
    if err := json.NewEncoder(w).Encode(&issue); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
