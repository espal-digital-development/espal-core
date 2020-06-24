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
func (configuration *Configuration) EmailHost() string {
	return configuration.email.Host
}

// EmailPort returns the email server port number.
func (configuration *Configuration) EmailPort() int {
	return configuration.email.Port
}

// EmailUsername returns the email server username.
func (configuration *Configuration) EmailUsername() string {
	return configuration.email.Username
}

// EmailPassword returns the email server password.
func (configuration *Configuration) EmailPassword() string {
	return configuration.email.Password
}

// EmailNoReplyAddress returns the email server no-reply address.
func (configuration *Configuration) EmailNoReplyAddress() string {
	return configuration.email.NoReplyAddress
}
