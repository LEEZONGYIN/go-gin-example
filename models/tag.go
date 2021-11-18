package models

import "github.com/jinzhu/gorm"

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

//func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("CreatedOn", time.Now().Unix())
//	return nil
//}
//
//func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
//	scope.SetColumn("ModifiedOn", time.Now().Unix())
//	return nil
//}

func GetTags(pageNum int, pageSize int, maps interface{}) ([]Tag, error) {
	var tags []Tag
	if pageSize > 0 && pageNum > 0 {
		db = db.Offset(pageNum).Limit(pageSize)
	}
	err := db.Where(maps).Find(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tags, nil
}

func GetTagTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Tag{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ? AND deleted_on = ?", name, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

func AddTag(name string, state int, createBy string) error {
	tag := Tag{
		Name:      name,
		State:     state,
		CreatedBy: createBy}
	if err := db.Create(&tag).Error; err != nil {
		return err
	}
	return nil
}

func ExistTagById(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ?", id).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

func DeleteTag(id int) error {
	if err := db.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}
	return nil
}

func EditTag(id int, data interface{}) error {
	if err := db.Model(&Tag{}).Where("id = ?", id).Update(data).Error; err != nil {
		return err
	}
	return nil
}

func CleanAllTag() (bool, error) {
	if err := db.Unscoped().Where("deleted_on != ?", 0).Delete(&Tag{}).Error; err != nil {
		return false, err
	}
	return true, nil
}
