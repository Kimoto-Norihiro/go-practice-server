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

func (mr *MemberRepository) CreateMember(m model.Member) error {
	return mr.db.Create(&m).Error
}

func (mr *MemberRepository) ShowMember(id string) (model.Member, error) {
	var m model.Member
	err := mr.db.Where("id = ?", id).First(&m).Error
	return m, err
}

func (mr *MemberRepository) DeleteMember(id string) error {
	return mr.db.Delete(&model.Member{}, id).Error
}

func (mr *MemberRepository) UpdateMember(id string, m model.Member) error {
	return mr.db.Model(&model.Member{}).Where("id = ?", id).Updates(m).Error
}


