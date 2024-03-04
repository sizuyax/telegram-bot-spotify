package models

type Profile struct {
	Id       int64
	Username string
}

type ReportedProfile struct {
	Username        string
	BlockedUsername string
}
