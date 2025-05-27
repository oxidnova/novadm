package sql

import "embed"

//go:embed migrations/*
var MigrationsEmbedFS embed.FS
