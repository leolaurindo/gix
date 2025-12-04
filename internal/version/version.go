package version

// Version is set at build time via -ldflags "-X github.com/leolaurindo/gix/internal/version.Version=vX.Y.Z".
// Defaults to "dev" for untagged/local builds.
var Version = "dev"
