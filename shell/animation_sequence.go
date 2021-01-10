package shell

type SequenceAnimationFrame struct {
	BeginFrame, EndFrame int32
	Maps                 Maps
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

	for i := 0; i < len(ani.FrameList); i++ {
		frameData := &ani.FrameList[i]

		if frame < frameData.BeginFrame || frame > frameData.EndFrame {
			continue
		}

		return ani.FrameList[i].Maps, true
	}

	return nil, false
}

func (ani *SequenceAnimation) GetTotalFrames() int32 {
	return ani.TotalFrames
}

func (ani *SequenceAnimation) GetFixedFPS() int32 {
	return ani.FixedFPS
}
