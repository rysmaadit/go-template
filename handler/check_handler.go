package handler

import (
	"github.com/rysmaadit/go-template/common/responder"
	"github.com/rysmaadit/go-template/service"
	"net/http"
)

func CheckRedis(service service.CheckService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := service.CheckRedis()
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, err)
		}
		responder.NewHttpResponse(r, w, http.StatusOK, resp, nil)
	}
}
