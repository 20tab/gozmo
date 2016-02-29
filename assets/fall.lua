function start(g)
    g:setattr("{}", "customFooBar", "C-u-stoM vAlue")
end

function update(g)
    print(g:getattr("", "name"))
    print(g:getattr("", "positionY"))
    y = g:getattr("", "positionY")
    y = y - 2 * g:getattr("", "deltaTime")
    g:setattr("", "positionY", y)
    print(g:getattr("{}", "customFooBar"))

    if g:getattr("", "getKeyRight") then
        g:setattr("", "positionAddX", 3 * g:getattr("", "deltaTime"))
    end
end

print("Hello, i am Lua")
