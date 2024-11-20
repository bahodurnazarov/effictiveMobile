package models

type Settings struct {
	App struct {
		ServerName string `mapstructure:"APP_SERVERNAME"`
		PortRun    string `mapstructure:"APP_PORTRUN"`
		DebugFile  string `mapstructure:"APP_DEBUGFILE"`
		InfoFile   string `mapstructure:"APP_INFOFILE"`
	} `mapstructure:",squash"`

	Postgres struct {
		PGHost     string `mapstructure:"POSTGRES_PG_HOST"`
		PGPort     string `mapstructure:"POSTGRES_PG_PORT"`
		PGUser     string `mapstructure:"POSTGRES_PG_USER"`
		PGPassword string `mapstructure:"POSTGRES_PG_PASSWORD"`
		PGName     string `mapstructure:"POSTGRES_PG_NAME"`
	} `mapstructure:",squash"`
}
