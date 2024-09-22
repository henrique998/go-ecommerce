package main

//go:generate mockgen -source=internal/app/providers/auth-tokens-provider.go -destination=tests/mocks/auth-tokens-provider-mock.go -package=mocks
//go:generate mockgen -source=internal/app/providers/payments-provider.go -destination=tests/mocks/payments-provider-mock.go -package=mocks
//go:generate mockgen -source=internal/app/repositories/account-roles-repository.go -destination=tests/mocks/account-roles-repository-mock.go -package=mocks
//go:generate mockgen -source=internal/app/repositories/accounts-repository.go -destination=tests/mocks/accounts-repository-mock.go -package=mocks
//go:generate mockgen -source=internal/app/repositories/refresh-tokens-repository.go -destination=tests/mocks/refresh-tokens-repository-mock.go -package=mocks
//go:generate mockgen -source=internal/app/services/session-services-interfaces/login-service.go -destination=tests/services-mocks/login-service-mock.go -package=servicesmocks
//go:generate mockgen -source=internal/app/services/session-services-interfaces/refresh-token-service.go -destination=tests/services-mocks/refresh-token-service-mock.go -package=servicesmocks
