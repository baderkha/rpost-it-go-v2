package repo

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// MYSQLPostRepo : Mysql gorm implementation of the repo
type MYSQLPostRepo struct {
	db *gorm.DB
}

func NewMYSQLPostRepo(db *gorm.DB) *MYSQLPostRepo {
	return &MYSQLPostRepo{
		db: db,
	}
}

func (mpr *MYSQLPostRepo) FindById(id string) *Post {
	var post Post
	mpr.db.First(&post, "id=?", id)
	return &post
}

func (mpr *MYSQLPostRepo) FindPostsByAccountOwner(posterId string) *[]Post {
	var posts []Post
	mpr.db.Where("poster_id=?", posterId).Find(&posts)
	return &posts
}

func (mpr *MYSQLPostRepo) FindPostsByCommunityId(communityId string) *[]Post {
	var posts []Post
	mpr.db.Where("community_id=?", communityId).Find(&posts)
	return &posts
}

func (mpr *MYSQLPostRepo) FindPostsByCommunityIdAndAccountOwner(communityId string, posterId string) *[]Post {
	var posts []Post
	mpr.db.Where("community_id=?", communityId).Where("poster_id=?", posterId).Find(&posts)
	return &posts
}

func (mpr *MYSQLPostRepo) Create(post *Post) (*Post, bool) {
	post.ID = uuid.NewV4().String()
	err := mpr.db.Create(post).Error
	return post, err == nil
}

func (mpr *MYSQLPostRepo) Update(post *Post) (*Post, bool) {
	err := mpr.db.Model(post).Updates(post).Error
	return post, err == nil
}

func (mpr *MYSQLPostRepo) Delete(id string) bool {
	return mpr.db.Unscoped().Delete(&Post{}, "id=?", id).Error != nil
}
