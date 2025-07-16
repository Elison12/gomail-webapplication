package cmd

import(
	"github.com/Elison12/gomail-webapplication/internal/config"
	"github.com/Elison12/gomail-webapplication/internal/httpHandler"
	"github.com/Elison12/gomail-webapplication/internal/mailer"
	"github.com/gin-gonic/gin")

func Execute() {
	cfg := config.Load()

	app := gin.Default()

	h := httpHandler.Handler{
		Mailer: mailer.New(cfg.MailGun),
	}

	app.POST("/mail/send", h.SendMail)

	if err := app.Run(cfg.Server); err != nil {
		panic(err)
	}
}