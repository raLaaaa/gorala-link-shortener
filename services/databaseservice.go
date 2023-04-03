package services

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/raLaaaa/gorala-link-shortener/models.go"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseService struct{}

func (d DatabaseService) FindAllLinks() ([]models.Link, error) {
	db, err := gorm.Open(sqlite.Open("database/data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Link{})

	links := []models.Link{}
	db.Order("times_visited desc, id, link, shortend_link, created_at, updated_at, deleted_at, shortend_full_link").Find(&links)

	return links, err
}

func (d DatabaseService) CreateShortLink(link *(models.Link), host string) error {

	db, err := gorm.Open(sqlite.Open("database/data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Link{})

	if err = db.Create(&link).Error; err != nil {
		return err
	}

	link.ShortendLink = d.randomString(5) + strconv.FormatUint(uint64(link.ID), 10)

	if !strings.HasPrefix(link.Link, "https://") {
		link.Link = "https://" + link.Link
	}

	link.ShortendFullLink = host + "/" + link.ShortendLink

	db.Save(&link)

	return err
}

func (d DatabaseService) FindLinkByShortLink(shortLinkID string) (*models.Link, error) {
	db, err := gorm.Open(sqlite.Open("database/data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Link{})

	var link = models.Link{}
	if result := db.First(&link, "shortend_link = ?", shortLinkID); result.Error != nil {
		return nil, result.Error
	}

	return &link, err
}

func (d DatabaseService) DeleteLinkByShortLink(shortLinkID string) error {
	db, err := gorm.Open(sqlite.Open("database/data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Link{})

	var link = models.Link{}
	if result := db.Where("shortend_link = ?", shortLinkID).Delete(&link); result.Error != nil {
		return result.Error
	}

	return err
}

func (d DatabaseService) IncreaseNumbersOfClicked(shortLinkID string) (*models.Link, error) {
	db, err := gorm.Open(sqlite.Open("database/data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Link{})

	var link = models.Link{}
	if result := db.First(&link, "shortend_link = ?", shortLinkID); result.Error != nil {
		return nil, result.Error
	}

	link.TimesVisited = link.TimesVisited + 1
	db.Save(&link)

	return &link, err
}

func (d DatabaseService) randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length+2)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[2 : length+2]
}
