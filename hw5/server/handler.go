package server

import (
	"net/http"
	"strconv"

	"github.com/freakkid/Service-Computing/hw5/entities"
	"github.com/unrolled/render"
)

func postUserInfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if req.FormValue("username") == "" {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input!"})
		} else {
			userInfo := entities.UserInfo{UserName: req.FormValue("username"),
				DepartName: req.FormValue("departname")}
			entities.UserInfoService.Save(&userInfo)
			formatter.JSON(w, http.StatusOK, userInfo)
		}
	}
}

func getUserInfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if req.FormValue("userid") != "" {
			id, err := strconv.ParseInt(req.FormValue("userid"), 10, 32)
			if err != nil {
				formatter.JSON(w, http.StatusBadRequest, entities.UserInfo{})
			} else {
				formatter.JSON(w, http.StatusOK, entities.UserInfoService.FindByID(int(id)))
			}
		} else {
			userList := entities.UserInfoService.FindAll()
			formatter.JSON(w, http.StatusOK, userList)
		}
	}
}

func getUserCountHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, entities.UserInfoService.Count())
	}
}
