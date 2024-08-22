package config

import models "github.com/aait.backend.g5.main/backend/Domain/Models"

func NewEmailServer(env Env) models.EmailConfig {
	return models.EmailConfig{
		SMTPServer: env.SMTP_SERVER,
		Port:       env.SMTP_PORT,
		Username:   env.SMTP_USERNAME,
		Password:   env.SMTP_PASSWORD,
	}
}
