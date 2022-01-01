package inverse

import (
	"context"
	"encoding/json"
	"net/http"
)

func DecodeInverseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request inverseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeInverseResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
