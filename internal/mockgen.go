package internal

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozonmp/ise-apartment-api/internal/app/repo EventRepo
//go:generate mockgen -destination=./mocks/sender_mock.go -package=mocks github.com/ozonmp/ise-apartment-api/internal/app/sender EventSender
//go:generate mockgen -destination=./mocks/apartment_repo_mock.go -package=mocks github.com/ozonmp/ise-apartment-api/internal/repo Repo
