package handler

import (
	"net/http"
	"synapsis-test-backend/server/request"
	"synapsis-test-backend/usecase"

	validator "gopkg.in/go-playground/validator.v9"
)

// UserHandler ...
type UserHandler struct {
	Handler
}

// LoginHandler ...
func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	req := request.UserLoginRequest{}
	if err := h.Handler.Bind(r, &req); err != nil {
		SendBadRequest(w, err.Error())
		return
	}
	if err := h.Handler.Validate.Struct(req); err != nil {
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	userUC := usecase.UserUC{ContractUC: h.ContractUC}
	res, err := userUC.Login(req)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
	return
}

// CreateHandler ...
func (h *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	req := request.UserRequest{}
	if err := h.Handler.Bind(r, &req); err != nil {
		SendBadRequest(w, err.Error())
		return
	}
	if err := h.Handler.Validate.Struct(req); err != nil {
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	userUC := usecase.UserUC{ContractUC: h.ContractUC}
	res, err := userUC.Register(&req)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
	return
}
