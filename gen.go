package main

//go:generate mockgen -source=internal/app/contracts/auth-tokens-provider.go -destination=tests/mocks/auth-tokens-provider-mock.go -package=mocks
//go:generate mockgen -source=internal/app/contracts/payments-provider.go -destination=tests/mocks/payments-provider-mock.go -package=mocks
//go:generate mockgen -source=internal/app/contracts/account-roles-repository.go -destination=tests/mocks/account-roles-repository-mock.go -package=mocks
//go:generate mockgen -source=internal/app/contracts/accounts-repository.go -destination=tests/mocks/accounts-repository-mock.go -package=mocks
//go:generate mockgen -source=internal/app/contracts/refresh-tokens-repository.go -destination=tests/mocks/refresh-tokens-repository-mock.go -package=mocks
//go:generate mockgen -source=internal/app/contracts/login-service.go -destination=tests/services-mocks/login-service-mock.go -package=servicesmocks
//go:generate mockgen -source=internal/app/contracts/refresh-token-service.go -destination=tests/services-mocks/refresh-token-service-mock.go -package=servicesmocks
