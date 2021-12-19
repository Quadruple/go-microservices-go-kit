package eigen

import (
	"context"
	"encoding/json"
	"net/http"
)

func DecodeEigenValueRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request eigenValueRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeEigenValueResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
