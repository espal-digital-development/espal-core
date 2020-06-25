package config

// Assets config section.
type Assets interface {
	AssetsGZip() bool
	AssetsBrotli() bool
	AssetsGZipFiles() bool
	AssetsBrotliFiles() bool
	AssetsCacheMaxAge() string
}

type assets struct {
	Gzip        bool
	Brotli      bool
	GzipFiles   bool   `yaml:"gzipFiles"`
	BrotliFiles bool   `yaml:"brotliFiles"`
	CacheMaxAge string `yaml:"cacheMaxAge"` // String to not have to convert on runtime every time
}

// AssetsGZip returns an indicator if assets calls like CSS, JS
// and images should be served GZipped.
// This feature works dynamically and if an agent doesn't support it
// it will still serve the raw data variation.
// This option can be disabled if you know for sure that your visitors
// don't support (or want) GZip and you want to save some memory on
// the stored Gzip variations of the assets.
func (c *Configuration) AssetsGZip() bool {
	return c.assets.Gzip
}

// AssetsBrotli returns an indicator if assets calls like CSS, JS
// and images should be served as Brotli compressed.
// This function has the same behavior as GzipAssets
func (c *Configuration) AssetsBrotli() bool {
	return c.assets.Brotli
}

// AssetsGZipFiles returns an indicator if files calls should be served
// as GZip compressed.
// This function has the same behavior as GzipAssets
func (c *Configuration) AssetsGZipFiles() bool {
	return c.assets.GzipFiles
}

// AssetsBrotliFiles returns an indicator if files calls should be served
// as Brotli compressed.
// This function has the same behavior as GzipAssets
func (c *Configuration) AssetsBrotliFiles() bool {
	return c.assets.BrotliFiles
}

// AssetsCacheMaxAge returns a string variation for HTTP MaxCache.
// It is a string so empty literally so it doesn't need int conversion
// every time when it's 0.
func (c *Configuration) AssetsCacheMaxAge() string {
	return c.assets.CacheMaxAge
}
