package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrResponse struct {
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}

func ResponsedJSON(ctx context.Context, w http.ResponseWriter, body any, status int) {
	w.Header().Set("Content-Type", "application/json")
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		rsp := ErrResponse{
			Message: "failed to marshal response body",
		}
		if err := json.NewEncoder(w).Encode(rsp); err != nil {
			fmt.Printf("failed to encode response body: %v", err)
		}
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(bodyBytes); err != nil {
		fmt.Printf("failed to write response body: %v", err)
	}
}
