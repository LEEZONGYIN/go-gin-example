package article_service

import (
	"go-gin-example/models"
)

type Article struct {
	ID            int
	TagId         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreatedBy     string
	ModifiedBy    string

	PageNum  int
	PageSize int
}

func (a *Article) Add() error {
	article := map[string]interface{}{
		"tag_id":          a.TagId,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"cover_iamge_url": a.CoverImageUrl,
		"state":           a.State,
	}
	if err := models.AddArticle(article); err != nil {
		return err
	}
	return nil
}

func (a *Article) Edit() error {
	return models.EditArticle(a.ID, map[string]interface{}{
		"tag_id":          a.TagId,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"cover_image_url": a.CoverImageUrl,
		"modified_by":     a.ModifiedBy,
	})
}
