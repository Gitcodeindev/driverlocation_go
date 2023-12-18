package driver

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var driverService *DriverService

func (ds *DriverService) GetDriver(ctx context.Context, id int64) (*Driver, error) {
	return nil, nil
}

func RegisterDriver(w http.ResponseWriter, r *http.Request) {
	var driver Driver
	err := json.NewDecoder(r.Body).Decode(&driver)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = driverService.RegisterDriver(context.Background(), &driver)
	if err != nil {
		http.Error(w, "Failed to register driver", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Driver registered")
}

func GetDriverInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid driver ID", http.StatusBadRequest)
		return
	}

	driver, err := driverService.GetDriver(context.Background(), id)
	if err != nil {
		http.Error(w, "Failed to get driver", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(driver)
}
