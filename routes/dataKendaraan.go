package routes

import (
	handlers "prokdrn/handler"
	"prokdrn/pkg/mysql"
	"prokdrn/repositories"

	"github.com/gorilla/mux"
)

func KendaraanRoutes(r *mux.Router) {
	kendaraanRepository := repositories.RepositoryDataKendaraan(mysql.DB)
	h := handlers.HandlerDataKendaraan(kendaraanRepository)

	r.HandleFunc("/kendaraans", h.FindDatas).Methods("GET")
	r.HandleFunc("/kendaraan/{id}", h.GetData).Methods("GET")
	r.HandleFunc("/kendaraan", h.CreateData).Methods("POST")
	r.HandleFunc("/kendaraan/{id}", h.UpdateData).Methods("PATCH")
	r.HandleFunc("/kendaraan/{id}", h.DeleteData).Methods("DELETE")
}
