package routes

import (
	"github.com/Diniz-J/teiunecc-admin/internal/modules/members/handler"
	"github.com/gorilla/mux"
)

func MemberRoutes(r *mux.Router, h *handler.MemberHandler) {
	r.HandleFunc("/members", h.ListMembers).Methods("GET")
	r.HandleFunc("/members", h.CreateMember).Methods("POST")
	r.HandleFunc("/members/{id}", h.GetMember).Methods("GET")
	r.HandleFunc("/members/{id}", h.UpdateMember).Methods("PUT")
	r.HandleFunc("/members/{id}", h.DeleteMember).Methods("DELETE")
}
