package config

type Config struct {
	PORT     int
	filename string
}

func CreateConfig() *Config {
	return &Config{}
}

func (c *Config) SetPORT(port int) {
	c.PORT = port
}

func (c *Config) SetFilename(filename string) {
	c.filename = filename
}

func (c *Config) GetPORT() int {
	return c.PORT
}

func (c *Config) GetFilename() string {
	return c.filename
}
