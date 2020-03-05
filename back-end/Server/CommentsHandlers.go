package Server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../Models"
	"../Services"
	"github.com/unrolled/render"
)

func CommentsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		comments := Models.GetAllComments()
		info := struct {
			Count    int              `json:"count"`
			Comments []Models.Comment `json:"comments"`
		}{
			Count:    len(comments),
			Comments: comments,
		}
		formatter.JSON(w, http.StatusOK, info)
	}
}

func AddCommentAsGuestHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := ioutil.ReadAll(req.Body)
		newCommentReq := struct {
			Name    string
			Content string
		}{}
		err := json.Unmarshal(payload, &newCommentReq)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to parse request.\n"+err.Error()+"\n")
			return
		}
		_, err = Services.AddCommentAsGuest(newCommentReq.Name, newCommentReq.Content)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to add comment.\n"+err.Error()+"\n")
			return
		} else {
			formatter.Text(w, http.StatusOK, "Add comment succeed.\n")
		}
	}
}
