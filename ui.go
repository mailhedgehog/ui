package ui

import (
	"embed"
)

var DirStaticName = ".go-static"

//go:embed .go-static/*
var DirStaticEmbed embed.FS
