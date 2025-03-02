package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"
)

func (app *Application) GetProvincesByCountryId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	countryIdString := r.PathValue("countryId")
	countryIdNormal, err := strconv.Atoi(countryIdString)
	if err != nil {
		NewProblemDetailsErrorResponse(w, http.StatusBadRequest, "Invalid country id provided", "Invalid url path parameter provided.")
		return
	}
	countryId := int32(countryIdNormal)

	provinces, err := app.service.GetProvincesByCountryId(ctx, countryId)
	if err != nil {
		log.Printf("Error: %v, ", err)
		NewProblemDetailsErrorResponse(w, http.StatusInternalServerError, "Unhandled server error", err.Error())
		return
	}

	sort.Slice(provinces, func(i, j int) bool {
		return provinces[i].ProvinceName < provinces[j].ProvinceName
	})

	json.NewEncoder(w).Encode(provinces)
}

func (h *Application) GetProvincesByProvinceId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	provinceIdString := r.PathValue("provinceId")
	provinceIdConverted, err := strconv.Atoi(provinceIdString)
	if err != nil {
		NewProblemDetailsErrorResponse(w, http.StatusBadRequest, "Invalid country id provided", "Invalid url path parameter provided.")
		return
	}
	provinceId := int32(provinceIdConverted)
	province, err := h.service.GetProvincesById(ctx, provinceId)
	if err != nil {
		log.Printf("Error: %v, ", err)
		NewProblemDetailsErrorResponse(w, http.StatusInternalServerError, "Unhandled server error", err.Error())
		return
	}
	json.NewEncoder(w).Encode(province)
}
