package repo

import "gorm.io/gorm"

type MYSQLComment struct {
	db *gorm.DB
}

func NewMYSQLCommentRepo(db *gorm.DB) *MYSQLComment {
	return &MYSQLComment{
		db: db,
	}
}

func (mcom *MYSQLComment) FindById(id string) *Comment {
	var com Comment
	mcom.db.First(&com, "id=?", id)
	return &com
}

func (mcom *MYSQLComment) FindCommentsByAccountOwnerAndPost(accountId string, postedId string) *[]Comment {
	var coms []Comment
	mcom.db.Where("account_owner_id=?", accountId).Where("post_id", postedId)
	return &coms
}

func (mcom *MYSQLComment) FindCommentsByPostId(postId string) *[]Comment {
	var coms []Comment
	mcom.db.Where("post_id", postId)
	return &coms
}

func (mcom *MYSQLComment) Create(comment *Comment) (*Comment, bool) {
	err := mcom.db.Create(comment).Error
	return comment, err == nil
}

func (mcom *MYSQLComment) Update(comment *Comment) (*Comment, bool) {
	err := mcom.db.Model(comment).Updates(comment).Error
	return comment, err == nil
}

func (mcom *MYSQLComment) DeleteByIdAndAccountId(id string, accountId string) bool {
	err := mcom.db.Where("id=?", id).Where("account_owner_id=?", accountId).Error
	return err == nil
}
