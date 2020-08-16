package types

// MapNode type
type MapNode struct {
	IsLocal      bool
	IsEntrypoint bool
	Path         string
	Info         ImportInfo
}
