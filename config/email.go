package config

// Email config section.
type Email interface {
	EmailHost() string
	EmailPort() int
	EmailUsername() string
	EmailPassword() string
	EmailNoReplyAddress() string
}

type email struct {
	Host           string
	Port           int
	Username       string
	Password       string
	NoReplyAddress string `yaml:"noReplyAddress"`
}

// EmailHost returns the email server host path.
func (c *Configuration) EmailHost() string {
	return c.email.Host
}

// EmailPort returns the email server port number.
func (c *Configuration) EmailPort() int {
	return c.email.Port
}

// EmailUsername returns the email server username.
func (c *Configuration) EmailUsername() string {
	return c.email.Username
}

// EmailPassword returns the email server password.
func (c *Configuration) EmailPassword() string {
	return c.email.Password
}

// EmailNoReplyAddress returns the email server no-reply address.
func (c *Configuration) EmailNoReplyAddress() string {
	return c.email.NoReplyAddress
}
