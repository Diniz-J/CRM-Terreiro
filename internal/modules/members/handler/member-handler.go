package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Diniz-J/teiunecc-admin/internal/modules/members/service"
	"github.com/Diniz-J/teiunecc-admin/internal/shared/response"
	"github.com/gorilla/mux"
)

type MemberHandler struct {
	service *service.MemberService
}

func NewMemberHandler(service *service.MemberService) *MemberHandler {
	return &MemberHandler{service: service}
}

func (h *MemberHandler) handleServiceError(w http.ResponseWriter, err error) {
	if errors.Is(err, service.ErrInvalidCPF) ||
		errors.Is(err, service.ErrInvalidEmail) ||
		errors.Is(err, service.ErrInvalidPhone) {
		response.Error(w, http.StatusBadRequest, "BAD_REQUEST", err.Error())
		return
	}
	response.Error(w, http.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
}

func (h *MemberHandler) CreateMember(w http.ResponseWriter, r *http.Request) {
	var input service.MemberInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "BAD_REQUEST", "invalid body")
		return
	}

	member, err := h.service.CreateMember(r.Context(), input)
	if err != nil {
		h.handleServiceError(w, err)
		return
	}

	response.JSON(w, http.StatusCreated, member)
}

func (h *MemberHandler) UpdateMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var input service.MemberInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "BAD_REQUEST", "invalid body")
		return
	}

	member, err := h.service.UpdateMember(r.Context(), id, input)
	if err != nil {
		h.handleServiceError(w, err)
		return
	}

	response.JSON(w, http.StatusCreated, member)
}
