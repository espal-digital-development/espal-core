package contexts

// RenderContext for renderer service parsing.
type RenderContext interface {
	CreatedBy(entity Entity) string
	UpdatedBy(entity Entity) string
	CountryName(countryID uint16) string
	LanguageName(languageID uint16) string
}

// CreatedBy returns the presentable name for the User that created this entity.
func (httpContext *HTTPContext) CreatedBy(entity Entity) string {
	return httpContext.rendererService.CreatedBy(entity, httpContext.language.ID())
}

// UpdatedBy returns the presentable name for the User that last updated this entity.
func (httpContext *HTTPContext) UpdatedBy(entity Entity) string {
	return httpContext.rendererService.UpdatedBy(entity, httpContext.language.ID())
}

// CountryName returns the localized name for the given countryID and languageID.
func (httpContext *HTTPContext) CountryName(countryID uint16) string {
	return httpContext.rendererService.CountryName(countryID, httpContext.language.ID())
}

// LanguageName returns the localized name for the given languageID and targetLanguageID.
func (httpContext *HTTPContext) LanguageName(languageID uint16) string {
	return httpContext.rendererService.LanguageName(languageID, httpContext.language.ID())
}
