package domainmock

//go:generate moq -pkg domainmock -out store_synthesized.go .. Store
//go:generate moq -pkg domainmock -out domain_synthesized.go .. DomainEntity
