package lua

import (
    goz "github.com/20tab/gozmo"
    "github.com/yuin/gopher-lua"
    "fmt"
)

type Lua struct {
    state *lua.LState
}

func NewLua(fileName string) *Lua {
    l := Lua{}
    ls := lua.NewState()
    l.state = ls
    err := ls.DoFile(fileName)
    if err != nil {
        panic(err)
    }

    mt := ls.NewTypeMetatable("gameobject")
    ls.SetField(mt, "__index", ls.SetFuncs(ls.NewTable(), gameobjectMethods))

    return &l
}

func gameobjectCheck(L *lua.LState) *goz.GameObject {
    ud := L.CheckUserData(1)
    if v, ok := ud.Value.(*goz.GameObject); ok {
        return v
    }
    L.ArgError(1, "gameobject expected")
    return nil
}

func gameobjectGetAttr(L *lua.LState) int {
    if L.GetTop() != 3 {
        L.ArgError(1, "invalid args")
        return 0
    }

    g := gameobjectCheck(L)

    v, err := g.GetAttr(L.CheckString(2), L.CheckString(3))
    if err != nil {
        fmt.Println(err)
        return 0
    }
    switch v.(type) {
        case string:
            L.Push(lua.LString(v.(string)))
            return 1
        case float32:
            L.Push(lua.LNumber(v.(float32)))
            return 1
        case bool:
            L.Push(lua.LBool(v.(bool)))
            return 1
    }
    return 0
}

func gameobjectSetAttr(L *lua.LState) int {
    if L.GetTop() != 4 {
        L.ArgError(1, "invalid args")
        return 0
    }

    g := gameobjectCheck(L)

    v := L.CheckAny(4)

    var err error

    switch v.(type) {
        case lua.LNumber:
            err = g.SetAttr(L.CheckString(2), L.CheckString(3), float32(lua.LVAsNumber(v)))
        case lua.LString:
            err = g.SetAttr(L.CheckString(2), L.CheckString(3), lua.LVAsString(v))
    }

    if err != nil {
        fmt.Println(err)
    }

    return 0
}

var gameobjectMethods = map[string]lua.LGFunction{
    "getattr": gameobjectGetAttr,
    "setattr": gameobjectSetAttr,
}

func (l *Lua) Start(g *goz.GameObject) {
    L := l.state

    ud := L.NewUserData()
    ud.Value = g 
    L.SetMetatable(ud, L.GetTypeMetatable("gameobject"))

    p := lua.P{Fn: L.GetGlobal("start"), NRet: 0, Protect: true}

    err := L.CallByParam(p, ud)
    if err != nil {
        fmt.Println(err)
    }
}

func (l *Lua) Update(g *goz.GameObject) {
    L := l.state

    ud := L.NewUserData()
    ud.Value = g 
    L.SetMetatable(ud, L.GetTypeMetatable("gameobject"))

    p := lua.P{Fn: L.GetGlobal("update"), NRet: 0, Protect: true}

    err := L.CallByParam(p, ud)
    if err != nil {
        fmt.Println(err)
    }
}

