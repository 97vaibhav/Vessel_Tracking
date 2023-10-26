.PHONY: codegen-proto
codegen-proto:
	buf generate

.PHONY: mysql
mysql:
	 docker run --name Vessel -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.7

.PHONY: mocks
mocks:
	mockery --dir=internal/infrastructure/datastore --name=VesselRepository --filename=repository.go --output=./mocks/ --outpkg=mocks
	mockery --dir=internal/domain/service --name=VesselUsecase --filename=usecase.go --output=./mocks/ --outpkg=mocks

.PHONY: test
test:
	go test ./...

.PHONY: up-dev
up-dev:
	docker compose up -d

.PHONY: down-dev
down-dev:
	docker compose down