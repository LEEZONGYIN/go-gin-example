package tag_service

import "go-gin-example/models"

type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int
}

func (t *Tag) Add() error {
	if err := models.AddTag(t.Name, t.State, t.ModifiedBy); err != nil {
		return err
	}
	return nil
}

func (t *Tag) Edit() error {
	tag := &Tag{
		Name:       t.Name,
		ModifiedBy: t.ModifiedBy,
		State:      t.State,
	}
	if err := models.EditTag(t.ID, tag); err != nil {
		return err
	}
	return nil
}
