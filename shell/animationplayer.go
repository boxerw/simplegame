package shell

import (
	"simplegame/core"
)

type AnimationPlayer interface {
	core.Object
	GetPosi() Posi2D
	SetPosi(posi Posi2D)
	GetFrames() int32
	SetFrames(v int32)
	GetHide() bool
	SetHide(v bool)
	GetPause() bool
	SetPause(v bool)
}

func NewAnimationPlayer(env core.Environment, screen Screen, layer int, posi Posi2D, animations []Animation, components ...core.ComponentBundle) AnimationPlayer {
	if screen == nil {
		panic("nil screen")
	}
	aniPlayer := &_AnimationPlayer{
		screen:     screen,
		layer:      layer,
		posi:       posi,
		animations: animations,
	}
	aniPlayer.ObjectModule.Init(env, aniPlayer, nil, components...)
	return aniPlayer
}

type _AnimationPlayer struct {
	core.ObjectModule
	screen     Screen
	layer      int
	posi       Posi2D
	animations []Animation
	frames     int32
	pause      bool
	hide       bool
}

func (aniPlayer *_AnimationPlayer) Update(frameCtx core.FrameContext) {
	if aniPlayer.GetDisabled() {
		return
	}

	aniPlayer.ObjectModule.Update(frameCtx)

	if aniPlayer.pause {
		return
	}

	if !aniPlayer.hide {
		for _, ani := range aniPlayer.animations {
			if maps, ok := ani.GetFrameMaps(aniPlayer.frames); ok {
				aniPlayer.screen.DrawMaps(aniPlayer.layer, aniPlayer.posi, maps)
			}
		}
	}

	aniPlayer.frames++
}

func (aniPlayer *_AnimationPlayer) GetPosi() Posi2D {
	return aniPlayer.posi
}

func (aniPlayer *_AnimationPlayer) SetPosi(posi Posi2D) {
	aniPlayer.posi = posi
}

func (aniPlayer *_AnimationPlayer) GetFrames() int32 {
	return aniPlayer.frames
}

func (aniPlayer *_AnimationPlayer) SetFrames(v int32) {
	aniPlayer.frames = v
}

func (aniPlayer *_AnimationPlayer) GetHide() bool {
	return aniPlayer.hide
}

func (aniPlayer *_AnimationPlayer) SetHide(v bool) {
	aniPlayer.hide = v
}

func (aniPlayer *_AnimationPlayer) GetPause() bool {
	return aniPlayer.pause
}

func (aniPlayer *_AnimationPlayer) SetPause(v bool) {
	aniPlayer.pause = v
}
