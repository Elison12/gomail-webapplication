package config

func Default() Config {
	return Config{
		Server: ":5000",
		MailGun: mailer.Config{
			Domain: "",
			APIKEY: "",
		},
	}
}