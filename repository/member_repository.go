package repository

import (
	"gorm.io/gorm"

	"github.com/Kimoto-Norihiro/go-practice-server/model"
)

type MemberRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) *MemberRepository {
	return &MemberRepository{db}
}

func (r *MemberRepository) CreateMember(m model.Member) error {
	return r.db.Create(&m).Error
}


