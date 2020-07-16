package config

// Assets config section.
type Assets interface {
	AssetsBrotli() bool
	AssetsGZip() bool
	AssetsBrotliFiles() bool
	AssetsGZipFiles() bool
	OptimizePngs() bool
	OptimizeJpegs() bool
	OptimizeGifs() bool
	OptimizeSvgs() bool
	AssetsCacheMaxAge() string
}

type assets struct {
	Brotli        bool
	Gzip          bool
	BrotliFiles   bool   `yaml:"brotliFiles"`
	GzipFiles     bool   `yaml:"gzipFiles"`
	OptimizePngs  bool   `yaml:"optimizePngs"`
	OptimizeJpegs bool   `yaml:"optimizeJpegs"`
	OptimizeGifs  bool   `yaml:"optimizeGifs"`
	OptimizeSvgs  bool   `yaml:"optimizeSvgs"`
	CacheMaxAge   string `yaml:"cacheMaxAge"` // String to not have to convert on runtime every time
}

// AssetsBrotli returns an indicator if assets calls like CSS, JS
// and images should be served as Brotli compressed.
// This function has the same behavior as GzipAssets.
func (c *Configuration) AssetsBrotli() bool {
	return c.assets.Brotli
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

// AssetsBrotliFiles returns an indicator if files calls should be served
// as Brotli compressed.
// This function has the same behavior as GzipAssets.
func (c *Configuration) AssetsBrotliFiles() bool {
	return c.assets.BrotliFiles
}

// AssetsGZipFiles returns an indicator if files calls should be served
// as GZip compressed.
// This function has the same behavior as GzipAssets.
func (c *Configuration) AssetsGZipFiles() bool {
	return c.assets.GzipFiles
}

// OptimizePngs returns an indicator if the system should attempt to optimize
// PNG assets that get loaded.
func (c *Configuration) OptimizePngs() bool {
	return c.assets.OptimizePngs
}

// OptimizeJpegs returns an indicator if the system should attempt to optimize
// JPEG assets that get loaded.
func (c *Configuration) OptimizeJpegs() bool {
	return c.assets.OptimizeJpegs
}

// OptimizeGifs returns an indicator if the system should attempt to optimize
// GIF assets that get loaded.
func (c *Configuration) OptimizeGifs() bool {
	return c.assets.OptimizeGifs
}

// OptimizeSvgs returns an indicator if the system should attempt to optimize
// SVG assets that get loaded.
func (c *Configuration) OptimizeSvgs() bool {
	return c.assets.OptimizeSvgs
}

// AssetsCacheMaxAge returns a string variation for HTTP MaxCache.
// It is a string so empty literally so it doesn't need int conversion
// every time when it's 0.
func (c *Configuration) AssetsCacheMaxAge() string {
	return c.assets.CacheMaxAge
}
