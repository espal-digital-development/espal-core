package contextsmock

//go:generate moq -pkg contextsmock -out factory.go .. Factory
//go:generate moq -pkg contextsmock -out context.go .. Context
