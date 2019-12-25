package types

// ImportInfo type
type ImportInfo struct {
	Line       int
	Path       string
	Name       string
	Module     string
	IsDir      bool
	ImportedIn string
}
