package shell

type SequenceAnimationFrame struct {
	Frame int32
	Maps  Maps
}

type SequenceAnimation struct {
	FixedFPS    int32
	TotalFrames int32
	FrameList   []SequenceAnimationFrame
}

func (ani *SequenceAnimation) GetFrameMaps(frame int32) (Maps, bool) {
	if frame < 0 || frame >= ani.TotalFrames {
		return nil, false
	}

	for i := len(ani.FrameList) - 1; i >= 0; i-- {
		if frame < ani.FrameList[i].Frame {
			continue
		}

		cur := i
		if cur < len(ani.FrameList)-1 {
			cur++
		}

		return ani.FrameList[cur].Maps, true
	}

	return nil, false
}

func (ani *SequenceAnimation) GetTotalFrames() int32 {
	return ani.TotalFrames
}

func (ani *SequenceAnimation) GetFixedFPS() int32 {
	return ani.FixedFPS
}

func (ani *SequenceAnimation) Marshal() ([]byte, error) {
	return nil, nil
}

func (ani *SequenceAnimation) Unmarshal(data []byte) error {
	return nil
}
