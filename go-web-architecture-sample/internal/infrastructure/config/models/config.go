package models

type Config struct {
	App         AppConfig         `mapstructure:",squash"`
	Logger      LoggerConfig      `mapstructure:",squash"`
	Persistence PersistenceConfig `mapstructure:",squash"`
}

func NewConfig() *Config {
	return &Config{
		App:         *NewAppConfig(),
		Logger:      *NewLoggerConfig(),
		Persistence: *NewPersistenceConfig(),
	}
}

func (c *Config) Validate() error {
	if err := c.App.Validate(); err != nil {
		return err
	}

	if err := c.Logger.Validate(); err != nil {
		return err
	}

	if err := c.Persistence.Validate(); err != nil {
		return err
	}

	return nil
}
