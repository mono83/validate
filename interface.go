package validate

// Interface describes entities, able to validate it's contents.
// This is core for this package
type Interface interface {
	Validate() error
}
