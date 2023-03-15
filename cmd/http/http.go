package http

import (
	"context"
	"template/config"
	"template/database"
	h "template/internal/server/http"
	mail "template/utils/mailer"
)

func StartServer(ctx context.Context) {
	dbConfig := config.NewDbConfig().Load().Get()
	db := database.NewSqlDB(dbConfig.Driver, dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Database).ORM()

	smtpConfig := config.NewSMTPConfig().Load().Get()
	smtp := mail.NewMailer(smtpConfig.Host, smtpConfig.Port, smtpConfig.AuthEmail, smtpConfig.Password).GetMailer()

	secretKey := config.NewSecretCfg().Load()

	ht := h.NewServer(&h.HttpServerCfg{
		DB:        db,
		SMTP:      *smtp,
		Secret:    secretKey.Key,
		AesSecret: secretKey.AesKey,
	})
	defer ht.Done()
	ht.Run(ctx)

	// return
	// http.ListenAndServe(":3000", r)
}
