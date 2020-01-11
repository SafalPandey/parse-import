package types

// ImportInfo type
type ImportInfo struct {
	Path      string
	IsDir     bool
	Importers []ImportedIn
}
