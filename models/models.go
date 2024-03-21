package models

type Profile struct {
	Id       int64
	Username string
}

type ReportedProfile struct {
	Username        string
	BlockedUsername string
}

type Config struct {
	TgBotApi      string `env:"TELEGRAM_TOKEN_BOT"`
	PostgresUser  string `env:"POSTGRES_USER"`
	PostgresPass  string `env:"POSTGRES_PASSWORD"`
	PostgresDB    string `env:"POSTGRES_DB"`
	PostgresHost  string `env:"POSTGRES_HOST"`
	PostgresPorts string `env:"POSTGRES_PORTS"`
}
