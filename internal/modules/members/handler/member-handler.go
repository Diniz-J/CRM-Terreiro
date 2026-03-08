package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Diniz-J/teiunecc-admin/internal/modules/members/service"
	"github.com/Diniz-J/teiunecc-admin/internal/shared/response"
)

type MemberHandler struct {
	service *service.MemberService
}

func NewMemberHandler(service *service.MemberService) *MemberHandler {
	return &MemberHandler{service: service}
}

func (h *MemberHandler) CreateMember(w http.ResponseWriter, r *http.Request) {
	var input service.MemberInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.Error(w, http.StatusBadRequest, "BAD_REQUEST", "invalid body")
		return
	}

	member, err := h.service.CreateMember(r.Context(), input)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCPF) ||
			errors.Is(err, service.ErrInvalidEmail) ||
			errors.Is(err, service.ErrInvalidPhone) {
			response.Error(w, http.StatusBadRequest, "BAD_REQUEST", err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, member)
}
