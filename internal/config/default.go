package config
import (
	"github.com/Elison12/gomail-webapplication/internal/mailer"
)

func Default() Config {
	return Config{
		Server: ":5000",
		MailGun: mailer.Config{
			Domain: "elison.monteiro.pge@gmail.com",
			APIKEY: "PGe@2024!",
		},
	}
}