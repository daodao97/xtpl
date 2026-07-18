package web

import "embed"

// Dist contains the client assets and the Goja-compatible server bundle.
//go:embed all:dist
var Dist embed.FS
