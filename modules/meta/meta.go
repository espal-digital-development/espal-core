package meta

// Config Meta configuration object.
type Config struct {
	UniqueIdentifier             string
	Version                      string
	MinimumCompatibleCoreVersion string
	MaximumCompatibleCoreVersion string
	Name                         string
	Author                       string
	Contact                      string
}

// Meta definition object.
type Meta struct {
	uniqueIdentifier             string
	version                      string
	minimumCompatibleCoreVersion string
	maximumCompatibleCoreVersion string
	name                         string
	author                       string
	contact                      string
}

// UniqueIdentifier gets the unique identifier.
func (d *Meta) UniqueIdentifier() string {
	return d.uniqueIdentifier
}

// Version gets the version.
func (d *Meta) Version() string {
	return d.version
}

// MinimumCompatibleCoreVersion gets the minimum compatible core version.
func (d *Meta) MinimumCompatibleCoreVersion() string {
	return d.minimumCompatibleCoreVersion
}

// MaximumCompatibleCoreVersion gets the maximum compatible core version.
func (d *Meta) MaximumCompatibleCoreVersion() string {
	return d.maximumCompatibleCoreVersion
}

// Name gets the name.
func (d *Meta) Name() string {
	return d.name
}

// Author gets the author.
func (d *Meta) Author() string {
	return d.author
}

// Contact gets the contact.
func (d *Meta) Contact() string {
	return d.contact
}

// New returns a new instance of Meta.
func New(config *Config) (*Meta, error) {
	m := &Meta{
		uniqueIdentifier:             config.UniqueIdentifier,
		version:                      config.Version,
		minimumCompatibleCoreVersion: config.MinimumCompatibleCoreVersion,
		maximumCompatibleCoreVersion: config.MaximumCompatibleCoreVersion,
		name:                         config.Name,
		author:                       config.Author,
		contact:                      config.Contact,
	}
	return m, nil
}
