package repository

import "github.com/Kimoto-Norihiro/go-practice-server/model"

type Repository interface {
	CreateMember(m model.Member) error
}