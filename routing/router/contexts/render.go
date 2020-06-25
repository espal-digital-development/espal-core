package contexts

// RenderContext for renderer service parsing.
type RenderContext interface {
	CreatedBy(entity Entity) string
	UpdatedBy(entity Entity) string
	CountryName(countryID uint16) string
	LanguageName(languageID uint16) string
}

// CreatedBy returns the presentable name for the User that created this entity.
func (c *HTTPContext) CreatedBy(entity Entity) string {
	return c.rendererService.CreatedBy(entity, c.language.ID())
}

// UpdatedBy returns the presentable name for the User that last updated this entity.
func (c *HTTPContext) UpdatedBy(entity Entity) string {
	return c.rendererService.UpdatedBy(entity, c.language.ID())
}

// CountryName returns the localized name for the given countryID and languageID.
func (c *HTTPContext) CountryName(countryID uint16) string {
	return c.rendererService.CountryName(countryID, c.language.ID())
}

// LanguageName returns the localized name for the given languageID and targetLanguageID.
func (c *HTTPContext) LanguageName(languageID uint16) string {
	return c.rendererService.LanguageName(languageID, c.language.ID())
}
