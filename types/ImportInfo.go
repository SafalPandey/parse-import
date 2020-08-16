package types

// ImportInfo type
type ImportInfo struct {
	Path      string
	IsDir     bool
	Imports   []string
	Importers []ImportedIn
}
