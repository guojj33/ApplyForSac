package Server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"../Models"
	"../Services"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func AdminsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		admins := Models.GetAllAdmins()
		info := struct {
			Count  int            `json:"count"`
			Admins []Models.Admin `json:"admins"`
		}{
			Count:  len(admins),
			Admins: admins,
		}
		formatter.JSON(w, http.StatusOK, info)
	}
}

func GetAdminByAdminIdHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		adminId := vars["adminId"]
		admin, err := Models.GetAdmin(adminId)
		if err == nil {
			formatter.JSON(w, http.StatusOK, *admin)
		} else {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
}

func GetAllAppRecordsAsAdminHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		adminId := vars["adminId"]
		appRecords, err := Services.GetAllAppRecordsAsAdmin(adminId)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, err.Error()+"\n")
			return
		} else {
			info := struct {
				Count      int `json:"Count"`
				AppRecords []Models.AppRecord
			}{
				Count:      len(appRecords),
				AppRecords: appRecords,
			}
			formatter.JSON(w, http.StatusOK, info)
		}
	}
}

func UpdateAppRecordAsAdminHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		adminId := vars["adminId"]
		appRecordIdStr := vars["appRecordId"]
		appRecordId, err := strconv.Atoi(appRecordIdStr)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to parse request.\n")
			return
		}
		updateAppRecordRequest := struct {
			UpdateField string
			NewValue    string
			ValueType   string
		}{}
		payload, _ := ioutil.ReadAll(req.Body)
		err = json.Unmarshal(payload, &updateAppRecordRequest)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to parse request.\n"+err.Error()+"\n")
			return
		}
		switch updateAppRecordRequest.ValueType {
		case "string":
			_, err = Services.UpdateAppRecordAsAdmin(adminId, appRecordId, updateAppRecordRequest.UpdateField, updateAppRecordRequest.NewValue)
		case "int":
			var newValueInt int
			newValueInt, err = strconv.Atoi(updateAppRecordRequest.NewValue)
			if err != nil {
				formatter.Text(w, http.StatusBadRequest, "Failed to parse request.\n"+err.Error()+"\n")
				return
			}
			_, err = Services.UpdateAppRecordAsAdmin(adminId, appRecordId, updateAppRecordRequest.UpdateField, newValueInt)
		}
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to update appRecord.\n"+err.Error()+"\n")
			return
		} else {
			formatter.Text(w, http.StatusOK, "Update succeed.\n")
			appRecord, _ := Models.GetAppRecord(appRecordId)
			formatter.JSON(w, http.StatusOK, *appRecord)
		}
	}
}
