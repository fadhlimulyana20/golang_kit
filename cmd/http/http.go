package http

import (
	"context"
	"template/config"
	"template/database"
	h "template/internal/server/http"
	mail "template/utils/mailer"
	"template/utils/minio"

	"github.com/spf13/cobra"
)

func StartServer(ctx context.Context, port int) {
	dbConfig := config.NewDbConfig().Load().Get()
	db := database.NewSqlDB(dbConfig.Driver, dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Database).ORM()

	smtpConfig := config.NewSMTPConfig().Load().Get()
	smtp := mail.NewMailer(smtpConfig.Host, smtpConfig.Port, smtpConfig.AuthEmail, smtpConfig.Password).GetMailer()

	secretKey := config.NewSecretCfg().Load()

	minioConfig := config.NewMinioCfg().Load()
	minio := minio.NewMinioStorage(minioConfig.Endpoint, minioConfig.AccessKeyID, minioConfig.SecretAccessKey, minioConfig.BucketName, minioConfig.UseSSL)

	ht := h.NewServer(config.Env(), db, *smtp, secretKey.Key, minio)
	defer ht.Done()
	ht.Run(ctx, port)

	// return
	// http.ListenAndServe(":3000", r)
}

func ServerCmd(ctx context.Context) *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "serve",
		Short: "Start HTTP server",
		Long:  "Start HTTP Server",
		Run: func(cmd *cobra.Command, args []string) {
			port, _ := cmd.Flags().GetInt("port")
			if port == 0 {
				port = 3000
			}
			StartServer(ctx, port)
		},
	}

	serverCmd.PersistentFlags().Int("port", 3000, "step for rolling back migration")

	return serverCmd
}
