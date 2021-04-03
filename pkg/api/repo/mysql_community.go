package repo

import "gorm.io/gorm"

type MYSQLCommunityRepo struct {
	db *gorm.DB
}

func NewMYSQLCommunityRepo(db *gorm.DB) *MYSQLCommunityRepo {
	return &MYSQLCommunityRepo{
		db: db,
	}
}

func (mcomm *MYSQLCommunityRepo) FindById(id string) *Community {
	var community Community
	mcomm.db.First(&mcomm, "id=?", id)
	return &community
}

func (mcomm *MYSQLCommunityRepo) FindByLikeIdInput(input string) *[]Community {
	var communities []Community
	mcomm.db.Where("id like ?", input).Find(communities)
	return &communities
}
func (mcomm *MYSQLCommunityRepo) FindByAccountOwnerId(accountId string) *Community {
	var community Community
	mcomm.db.Where("account_owner_id=?", accountId).First(community)
	return &community
}

func (mcomm *MYSQLCommunityRepo) Create(com *Community) (*Community, error) {
	err := mcomm.db.Create(com).Error
	return com, err
}

func (mcomm *MYSQLCommunityRepo) Update(com *Community) (*Community, error) {
	err := mcomm.db.Model(com).Updates(com).Error
	return com, err
}

func (mcomm *MYSQLCommunityRepo) Delete(id string) error {
	return mcomm.db.Delete(&Community{}, "id=?", id).Error
}
