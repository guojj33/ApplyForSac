package Server

import (
	"net/http"

	"../Models"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func RoomsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		rooms := Models.GetAllRooms()
		info := struct {
			Count int           `json:"count"`
			Rooms []Models.Room `json:"rooms"`
		}{
			Count: len(rooms),
			Rooms: rooms,
		}
		formatter.JSON(w, http.StatusOK, info)
	}
}

func GetRoomByRoomNameHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		roomName := vars["roomName"]
		room, err := Models.GetRoom(roomName)
		if err == nil {
			formatter.JSON(w, http.StatusOK, *room)
		} else {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
}
