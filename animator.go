package gozmo

import (
	"fmt"
	"math"
)

// The Animator component applies animations to gameObjects.
type Animator struct {
	currentAnimation string
	isPlaying        bool
	deltaT           float32
	currentFrame     int
	frameApplied     bool
}

func NewAnimator() *Animator {
	animator := Animator{isPlaying: false}
	return &animator
}

func (animator *Animator) Start(gameObject *GameObject) {
}

func (animator *Animator) Update(gameObject *GameObject) {
	if !animator.isPlaying {
		return
	}

	if animator.currentAnimation == "" {
		return
	}

	animation, ok := gameObject.Scene.animations[animator.currentAnimation]
	if !ok {
		return
	}

	if animator.deltaT > 0 {
		animator.deltaT -= gameObject.DeltaTime
	}

	if animator.deltaT <= 0 {
		// force drawing of the frame
		animator.frameApplied = false

		// Is it a new animation?
		if animator.currentFrame == -1 {
			if animation.Fps >= 0 {
				animator.currentFrame = 0
			} else {
				animator.currentFrame = len(animation.Frames) - 1
			}
		} else {
			// Switch frame.
			if animation.Fps > 0 {
				animator.currentFrame++
			} else if animation.Fps < 0 {
				animator.currentFrame++
			}
		}

		// Positive FPS.
		if animator.currentFrame >= len(animation.Frames) {
			if animation.Loop == false {
				animator.currentFrame--
				return
			}
			animator.currentFrame = 0
		}

		// Negative FPS.
		if animator.currentFrame < 0 {
			if animation.Loop == false {
				animator.currentFrame = 0
				return
			}
			animator.currentFrame = len(animation.Frames) - 1
		}

		animator.deltaT = float32(math.Abs(1.0 / float64(animation.Fps)))

	}

	if animator.frameApplied == false {
		frame := animation.Frames[animator.currentFrame]
		for _, action := range frame.actions {
			if action == nil {
				continue
			}
			err := gameObject.SetAttr(action.ComponentName, action.Attr, action.Value)
			if err != nil {
				fmt.Println(err)
			}
		}
		animator.frameApplied = true
	} else {
		// Check for interpolation.
		frame := animation.Frames[animator.currentFrame]
		for aid, action := range frame.actions {
			if action == nil {
				continue
			}
			if action.Interpolate {
				// This is only required if we want to interpolate linearly.
				/*
				   currentValue, err := gameObject.GetAttr(action.ComponentName, action.Attr)
				   if err != nil {
				       fmt.Println(err)
				       continue
				   }
				   value, ok := currentValue.(float32)
				   if !ok {
				       fmt.Println("error while interpolating %v: not a float32", action.Attr)
				       continue
				   }
				*/
				// Get the next frame.
				var nextFrame int
				if animation.Fps >= 0 {
					nextFrame = animator.currentFrame + 1
				} else {
					nextFrame = animator.currentFrame - 1
				}

				if nextFrame < 0 {
					if !animation.Loop {
						continue
					}
					nextFrame = len(animation.Frames) - 1
				}

				if nextFrame >= len(animation.Frames) {
					if !animation.Loop {
						continue
					}
					nextFrame = 0
				}

				// Got the next frame, check if an action is available
				// in the same position.
				if aid >= len(animation.Frames[nextFrame].actions) {
					continue
				}
				nextAction := animation.Frames[nextFrame].actions[aid]
				if nextAction == nil {
					continue
				}

				if nextAction.ComponentName != action.ComponentName {
					continue
				}

				if nextAction.Attr != action.Attr {
					continue
				}

				// Get the next action value.
				nextValue, ok := nextAction.Value.(float32)
				if !ok {
					fmt.Println("error while interpolating %v: not a float32", nextAction.Attr)
					continue
				}

				// Compute the gradient based on deltaT.
				var gradient, interpolatedValue float32
				frameTime := float32(math.Abs(1.0 / float64(animation.Fps)))
				gradient = (1.0 / frameTime) * (frameTime - animator.deltaT)
				interpolatedValue = action.Value.(float32) + ((nextValue - action.Value.(float32)) * gradient)

				err := gameObject.SetAttr(action.ComponentName, action.Attr, interpolatedValue)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}

}

func (animator *Animator) Play() {
	animator.isPlaying = true
}

func (animator *Animator) Stop() {
	animator.isPlaying = false
}

func (animator *Animator) SetAnimation(name string) {
	animator.currentAnimation = name
	animator.currentFrame = -1
	animator.deltaT = 0
}

func (animator *Animator) GetAnimation() string {
	return animator.currentAnimation
}

func (animator *Animator) SetAttr(attr string, value interface{}) error {
	switch attr {
	case "play":
		state, ok := value.(bool)
		if ok {
			if state {
				animator.Play()
			} else {
				animator.Stop()
			}
			return nil
		}
		return fmt.Errorf("%v attribute of %T expects a bool", attr, animator)
	case "animation":
		animation, ok := value.(string)
		if ok {
			animator.SetAnimation(animation)
			return nil
		}
		return fmt.Errorf("%v attribute of %T expects a string", attr, animator)
	}
	return nil
}

func (animator *Animator) GetAttr(attr string) (interface{}, error) {
	switch attr {
	case "animation":
		return animator.GetAnimation(), nil
	case "play":
		return animator.isPlaying, nil
	}
	return nil, fmt.Errorf("%v attribute of %T not found", attr, animator)

}

func (animator *Animator) GetType() string {
	return "Animator"
}

func initAnimator(args []interface{}) Component {
	return NewAnimator()
}

func init() {
	RegisterComponent("Animator", initAnimator)
}
