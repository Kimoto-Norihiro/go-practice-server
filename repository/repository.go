package repository

import "github.com/Kimoto-Norihiro/go-practice-server/model"

type Repository interface {
	CreateMember(m model.Member) error
	ShowMember(id string) (model.Member, error)
	DeleteMember(id string) error
	UpdateMember(id string, m model.Member) error
}