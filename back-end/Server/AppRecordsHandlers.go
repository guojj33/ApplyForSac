package Server

import (
	"net/http"
	"strconv"

	"../Models"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func AppRecordsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		appRecords := Models.GetAllAppRecords()
		info := struct {
			Count      int                `json:"count"`
			AppRecords []Models.AppRecord `json:"appRecords"`
		}{
			Count:      len(appRecords),
			AppRecords: appRecords,
		}
		formatter.JSON(w, http.StatusOK, info)
	}
}

func GetAppRecordByAppRecordIdHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		var err error
		appRecordId, err := strconv.Atoi(vars["appRecordId"])
		if err == nil {
			appRecord, err := Models.GetAppRecord(appRecordId)
			if err == nil {
				formatter.JSON(w, http.StatusOK, appRecord)
			} else {
				w.WriteHeader(http.StatusNotFound)
				return
			}
		}
	}
}
