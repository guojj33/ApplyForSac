package Server

import (
	"net/http"

	"github.com/unrolled/render"
)

func ApiHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			Users      string `json:"users"`
			User       string `json:"user"`
			Admins     string `json:"admins"`
			Admin      string `json:"admin"`
			Rooms      string `json:"rooms"`
			Room       string `json:"room"`
			AppRecords string `json:"appRecords"`
			AppRecord  string `json:"appRecord"`
		}{
			Users:      "https://localhost:8080/api/users",
			User:       "https://localhost:8080/api/users/{userId}",
			Admins:     "https://localhost:8080/api/admins",
			Admin:      "https://localhost:8080/api/admins/{adminId}",
			Rooms:      "https://localhost:8080/api/rooms",
			Room:       "https://localhost:8080/api/rooms/{roomName}",
			AppRecords: "https://localhost:8080/api/appRecords",
			AppRecord:  "https://localhost:8080/api/appRecords/{appRecordId}",
		})
	}
}
