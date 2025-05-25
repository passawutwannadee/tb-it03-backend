package app

import "github.com/passawutwannadee/tb-it03/pkg/postgres"

type AppConfig struct {
	Database *postgres.Postgres
}
