package database

import (
	"telegram-bot-spotify/backend/database/models"
)

func AddProfileToDB(id int64, username string) error {
	profile := models.Profile{Id: id, Username: username}
	err := db.Create(&profile).Error
	if err != nil {
		return err
	}
	return nil
}

func IsProfileExist(id int64) bool {
	var profile models.Profile
	err := db.Where("id = ?", id).First(&profile).Error
	if err != nil {
		return false
	}
	return true
}

func GetProfile() ([]string, error) {
	var profiles []models.Profile

	err := db.Find(&profiles).Error
	if err != nil {
		return nil, err
	}

	var usernames []string
	for _, profile := range profiles {
		usernames = append(usernames, profile.Username)
	}

	return usernames, nil
}

func GetUsers() []models.Profile {
	var profiles []models.Profile
	db.Find(&profiles)
	return profiles
}

func BlockProfile(username, blockedUsername string) error {
	profile := models.ReportedProfile{Username: username, BlockedUsername: blockedUsername}
	err := db.Create(&profile).Error
	if err != nil {
		return err
	}
	return nil
}

func IsProfileBlocked(blockedUsername string) bool {
	var profile models.ReportedProfile
	err := db.Where("blocked_username = ?", blockedUsername).First(&profile).Error
	if err != nil {
		return false
	}
	return true
}

func UnBlockProfile(blockedUsername string) error {
	err := db.Where("blocked_username = ?", blockedUsername).Delete(&models.ReportedProfile{}).Error
	if err != nil {
		return err
	}
	return nil
}

func GetBlockedProfiles() ([]string, error) {
	var profiles []models.ReportedProfile

	err := db.Find(&profiles).Error
	if err != nil {
		return nil, err
	}

	var usernames []string
	for _, profile := range profiles {
		usernames = append(usernames, profile.BlockedUsername)
	}

	return usernames, nil
}
