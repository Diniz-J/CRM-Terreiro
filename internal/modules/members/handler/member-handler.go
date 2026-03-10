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
	if errors.Is(err, service.ErrMemberNotFound) {
		response.Error(w, http.StatusNotFound, "NOT_FOUND", err.Error())
		return
	}

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
	if id == "" {
		response.Error(w, http.StatusBadRequest, "BAD_REQUEST", "missing id")
		return
	}

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

	response.JSON(w, http.StatusOK, member)
}

func (h *MemberHandler) DeleteMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		response.Error(w, http.StatusBadRequest, "BAD_REQUEST", "missing id")
		return
	}

	err := h.service.DeleteMember(r.Context(), id)
	if err != nil {
		h.handleServiceError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
