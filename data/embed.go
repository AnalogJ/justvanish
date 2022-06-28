package data

import "embed"

//go:embed templates/*
var Templates embed.FS

//go:embed organizations/*
var Organizations embed.FS
