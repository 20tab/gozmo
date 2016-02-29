package gozmo

type AnimationAction struct {
    ComponentName string
    Attr string
    Value interface{}
    Interpolate bool
}

type AnimationFrame struct {
    actions []*AnimationAction
}

type Animation struct {
    Name string
    Fps int
    Frames []*AnimationFrame
    Loop bool
}

func (animation *Animation) AddFrame(actions []*AnimationAction) *AnimationFrame {
    frame := AnimationFrame{}
    for _, action := range actions {
        frame.actions = append(frame.actions, action)
    }
    animation.Frames = append(animation.Frames, &frame)
    return &frame
}

func (animation *Animation) AddSimpleFrame(componentName string, attr string, value interface{}, interpolate bool) *AnimationFrame {
    frame := AnimationFrame{}
    frame.actions = append(frame.actions, &AnimationAction{componentName, attr, value, interpolate})
    animation.Frames = append(animation.Frames, &frame)
    return &frame
}

func (scene *Scene) AddAnimation(name string, fps int, loop bool) *Animation {
    animation := Animation{Name: name, Fps: fps, Loop: loop}
    scene.animations[name] = &animation
    return &animation
}
