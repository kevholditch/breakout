# breakout

Breakout clone written in Go using OpenGL.

## About
I wanted to learn how to write raw Open GL code, so what better project than to implement the classic game breakout!

The game engine logic uses an [Entity Component System](https://en.wikipedia.org/wiki/Entity_component_system).  This makes it easier to separate up the game logic into systems.  Thanks to [Liam Galvin](https://github.com/liamg) for the [ECS library](https://github.com/liamg/ecs).

![Gameplay](./gameplay.png)

## Controls
To play use spacebar to kick off and left and right arrows to move the bat.
