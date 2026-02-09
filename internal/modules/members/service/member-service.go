package service

import (
	"github.com/Diniz-J/teiunecc-admin/internal/modules/members/repository"
)

type MemberService struct {
	repo *repository.MemberRepository
}

func NewMemberService(repo *repository.MemberRepository) *MemberService {
	return &MemberService{repo: repo}
}
