package profiles

// ProfilePayload is the interface implemented by all profile payload types.
type ProfilePayload interface {
	PayloadType() string
}
