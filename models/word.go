package models

import (
	"RecordWords/dao"
)

// Word Model
type Word struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreateWord(word *Word) (err error) {
	err = dao.DB.Create(&word).Error
	return
}

func GetAllWord() (wordList []*Word, err error) {
	if err = dao.DB.Find(&wordList).Error; err != nil {
		return nil, err
	}
	return
}

func GetWord(id string) (word *Word, err error) {
	word = new(Word)
	if err = dao.DB.Debug().Where("id=?", id).First(word).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateWord(word *Word) (err error) {
	err = dao.DB.Save(word).Error
	return
}

func DeleteWord(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Word{}).Error
	return
}
