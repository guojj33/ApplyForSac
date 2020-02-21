package Server

import (
	"os"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		Extensions: []string{".html"},
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()
	mx.StrictSlash(false)
	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
		}
	}
	//mx.HandleFunc("/api/", ApiHandler(formatter)).Methods("GET")

	mx.Use(ValidateTokenMiddleware)

	mx.HandleFunc("/api/login", LoginHandler(formatter)).Methods("POST")
	mx.HandleFunc("/api/register", RegisterHandler(formatter)).Methods("POST")
	mx.HandleFunc("/api/logout", LogoutHandler(formatter)).Methods("POST")

	//mx.HandleFunc("/api/users/", UsersHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/users/{userId}", GetUserByUserIdHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/users/{userId}/appRecords", GetUserAppRecordsAsUserHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/users/{userId}/appRecords", CreateAppRecordAsUserHandler(formatter)).Methods("POST")
	mx.HandleFunc("/api/users/{userId}/appRecords/{appRecordId}", UpdateAppRecordAsUserHandler(formatter)).Methods("PUT")
	mx.HandleFunc("/api/users/{userId}/rooms/{roomName}/appRecords", GetAppRecordByRoomNameAsUserHandler(formatter)).Methods("GET")

	//mx.HandleFunc("/api/admins/", AdminsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/admins/{adminId}", GetAdminByAdminIdHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/admins/{adminId}/appRecords", GetAllAppRecordsAsAdminHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/admins/{adminId}/appRecords", CreateAppRecordAsAdminHandler(formatter)).Methods("POST")
	mx.HandleFunc("/api/admins/{adminId}/appRecords/{appRecordId}", UpdateAppRecordAsAdminHandler(formatter)).Methods("PUT")
	mx.HandleFunc("/api/admins/{adminId}/rooms/{roomName}/appRecords", GetAppRecordByRoomNameAsAdminHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/admins/{adminId}/register", RegisterNewAdminAsAdminHandler(formatter)).Methods("POST")

	mx.HandleFunc("/api/rooms", RoomsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/rooms", CreateNewRoomAsAdminHandler(formatter)).Methods("POST")
	//mx.HandleFunc("/api/rooms/{roomName}", GetRoomByRoomNameHandler(formatter)).Methods("GET")

	//mx.HandleFunc("/api/appRecords", AppRecordsHandler(formatter)).Methods("GET")
	//mx.HandleFunc("/api/appRecords/{appRecordId}", GetAppRecordByAppRecordIdHandler(formatter)).Methods("GET")
}
