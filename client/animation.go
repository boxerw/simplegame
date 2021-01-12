package client

type Animation interface {
	GetFrameMaps(frame int32) (Maps, bool)
	GetTotalFrames() int32
	GetFixedFPS() int32
}
