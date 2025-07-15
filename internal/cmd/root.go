package cmd

import("github.com/gin-gonic/gin")

func Execute() {
	cfg := config.Load()

	app := gin.Default()

	h := handler.Handler{
		Mailer: mailer.New(cfg.MailGun),
	}

	app.POST("/mail/send", h.SendMail)

	if err := app.Run(cfg.Server); err != nil {
		panic(err)
	}
}