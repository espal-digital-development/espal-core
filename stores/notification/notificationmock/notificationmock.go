package notificationmock

//go:generate moq -pkg notificationmock -out notification_synthesized.go .. NotificationEntity
//go:generate moq -pkg notificationmock -out notification_synthesized.go .. Store
