package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type LocationRepositoryImpl struct {
}

func (repo *LocationRepositoryImpl) UpdateLocation() error {
	return nil
}

type LocationService struct {
	repo *LocationRepositoryImpl
}

func NewLocationService(repo *LocationRepositoryImpl) *LocationService {
	return &LocationService{repo: repo}
}

func main() {
	if err := godotenv.Load("configs/.env.dev"); err != nil {
		log.Fatal("Ошибка загрузки файла .env")
	}

	locationRepo := &LocationRepositoryImpl{}
	locationService := NewLocationService(locationRepo)

	r := mux.NewRouter()
	r.HandleFunc("/location", locationHandler(locationService)).Methods("PUT")

	log.Println("Запуск службы местоположения на порту:", os.Getenv("LOCATION_PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("LOCATION_PORT"), r); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}

func locationHandler(locationService *LocationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var locationUpdate map[string]string
		if err := json.NewDecoder(r.Body).Decode(&locationUpdate); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		driverID, ok := locationUpdate["IDВодителя"]
		if !ok {
			http.Error(w, "IDВодителя требуется", http.StatusBadRequest)
			return
		}

		newLocation, ok := locationUpdate["НовоеМестоположение"]
		if !ok {
			http.Error(w, "НовоеМестоположение требуется", http.StatusBadRequest)
			return
		}

		if err := locationService.repo.UpdateLocation(); err != nil {
			http.Error(w, "Не удалось обновить местоположение", http.StatusInternalServerError)
			return
		}

		fmt.Printf("Обновление местоположения для водителя %s на %s\n", driverID, newLocation)

		response := map[string]string{"message": "Местоположение успешно обновлено"}
		jsonResponse, _ := json.Marshal(response)
		_, err := w.Write(jsonResponse)
		if err != nil {
			return
		}
	}
}