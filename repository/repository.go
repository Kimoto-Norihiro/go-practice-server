package repository

import "github.com/Kimoto-Norihiro/go-practice-server/model"

type Repository interface {
	CreateMember(m model.Member) error
	ShowMember(id uint) (model.Member, error)
	DeleteMember(id uint) error
	UpdateMember(id uint, m model.Member) error
}
