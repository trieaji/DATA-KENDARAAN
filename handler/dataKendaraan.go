package handlers

import (
	"encoding/json"
	"net/http"
	dataKendaraandto "prokdrn/dto/dataKendaraan"
	dto "prokdrn/dto/result"
	"prokdrn/models"
	"prokdrn/repositories"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerKendaraan struct {
	DataKendaraanRepository repositories.DataKendaraanRepository
}

func HandlerDataKendaraan(DataKendaraanRepository repositories.DataKendaraanRepository) *handlerKendaraan {
	return &handlerKendaraan{DataKendaraanRepository}
}

func (h *handlerKendaraan) FindDatas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dKendaraan, err := h.DataKendaraanRepository.FindDatas()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: dKendaraan}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerKendaraan) GetData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var kendaraan models.DataKendaraan
	kendaraan, err := h.DataKendaraanRepository.GetData(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProduct(kendaraan)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerKendaraan) CreateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(dataKendaraandto.DatakendaraanRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	kendaraan := models.DataKendaraan{
		NoRegistrasi:   request.NoRegistrasi,
		NamaPemilik:    request.NamaPemilik,
		MerkKendaraan:  request.MerkKendaraan,
		TahunPembuatan: request.TahunPembuatan,
		Kapasitas:      request.Kapasitas,
		Warna:          request.Warna,
		BahanBakar:     request.BahanBakar,
	}

	kendaraan, err = h.DataKendaraanRepository.CreateData(kendaraan)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	kendaraan, _ = h.DataKendaraanRepository.GetData(kendaraan.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: kendaraan}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerKendaraan) UpdateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	request := new(dataKendaraandto.DatakendaraanUpdate)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	kendaraan, _ := h.DataKendaraanRepository.GetData(int(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	if (request.NoRegistrasi) != "" {
		kendaraan.NoRegistrasi = request.NoRegistrasi
	}

	if request.NamaPemilik != "" {
		kendaraan.NamaPemilik = request.NamaPemilik
	}

	if request.MerkKendaraan != "" {
		kendaraan.MerkKendaraan = request.MerkKendaraan
	}

	if request.TahunPembuatan != 0 {
		kendaraan.TahunPembuatan = request.TahunPembuatan
	}

	if request.Kapasitas != 0 {
		kendaraan.Kapasitas = request.Kapasitas
	}

	if request.Warna != "" {
		kendaraan.Warna = request.Warna
	}

	if request.BahanBakar != "" {
		kendaraan.BahanBakar = request.BahanBakar
	}

	data, err := h.DataKendaraanRepository.UpdateData(kendaraan, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProduct(data)}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerKendaraan) DeleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	kendaraan, err := h.DataKendaraanRepository.GetData(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.DataKendaraanRepository.DeleteData(kendaraan, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProduct(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseProduct(p models.DataKendaraan) dataKendaraandto.DatakendaraanResponse {
	return dataKendaraandto.DatakendaraanResponse{
		ID:             p.ID,
		NoRegistrasi:   p.NoRegistrasi,
		NamaPemilik:    p.NamaPemilik,
		MerkKendaraan:  p.MerkKendaraan,
		TahunPembuatan: p.TahunPembuatan,
		Kapasitas:      p.Kapasitas,
		Warna:          p.Warna,
		BahanBakar:     p.BahanBakar,
	}
}
