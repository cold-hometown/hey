package config

type Config struct {
	writer ConfigWriter
	reader ConfigReader
}

func (c Config) Load() error {
	meta := &Meta{}

	err := c.writer.config(*meta)

	if err != nil {
		return err
	}

	err = c.reader.config(*meta)

	if err != nil {
		return err
	}

	return nil
}
