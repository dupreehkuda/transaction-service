package handlers

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

// FundsHandler gets request and passes it in the queue
func (h handlers) FundsHandler(w http.ResponseWriter, r *http.Request) {
	var data request

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		h.logger.Error("Unable to decode JSON", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if data.Account == "" || data.Amount.IsNegative() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if data.Operation == "add" && data.Operation == "withdraw" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	err = h.processor.WriteToQueue(data.Account, data.Operation, data.Amount)
	if err != nil {
		h.logger.Error("Error when writing to queue", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
