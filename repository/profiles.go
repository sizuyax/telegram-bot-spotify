package repository

import (
	"github.com/sirupsen/logrus"
	"telegram-bot-spotify/database"
	"telegram-bot-spotify/models"
)

func AddProfileToDB(id int64, username string) error {
	profile := models.Profile{Id: id, Username: username}

	if err := database.DB.Create(&profile).Error; err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func IsProfileExist(id int64) (bool, error) {
	var profile models.Profile

	if err := database.DB.Where("id = ?", id).First(&profile).Error; err != nil {
		logrus.Error(err)
		return false, err
	}

	return true, nil
}

func GetProfiles() ([]string, error) {
	var profiles []models.Profile

	if err := database.DB.Find(&profiles).Error; err != nil {
		logrus.Error(err)
		return nil, err
	}

	var usernames []string
	for _, profile := range profiles {
		usernames = append(usernames, profile.Username)
	}

	return usernames, nil
}

func BlockProfile(username, blockedUsername string) error {
	profile := models.ReportedProfile{Username: username, BlockedUsername: blockedUsername}

	if err := database.DB.Create(&profile).Error; err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func IsProfileBlocked(blockedUsername string) (bool, error) {
	var profile models.ReportedProfile

	if err := database.DB.Where("blocked_username = ?", blockedUsername).First(&profile).Error; err != nil {
		logrus.Error(err)
		return false, err
	}

	return true, nil
}

func UnBlockProfile(blockedUsername string) error {
	if err := database.DB.Where("blocked_username = ?", blockedUsername).Delete(&models.ReportedProfile{}).Error; err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func GetBlockedProfiles() ([]string, error) {
	var profiles []models.ReportedProfile

	if err := database.DB.Find(&profiles).Error; err != nil {
		logrus.Error(err)
		return nil, err
	}

	var usernames []string
	for _, profile := range profiles {
		usernames = append(usernames, profile.BlockedUsername)
	}

	return usernames, nil
}
