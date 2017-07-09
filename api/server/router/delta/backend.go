package delta

// Backend is the methods that need to be implemented to provide
// delta specific functionality.
type Backend interface {
	DeltaCreate(deltaSrc, deltaDest string) (imageID string, err error)
}
