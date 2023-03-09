# Warning
Do not scroll if you have Trypophobia (fear of holes) as the images used in the first versions may be irritating to you. After official release I will fix that problem.

# simple4x
This project aims to create a simple 4x game using Golang. I want to try and build a fun simple 4x game and put it on App store and Google play for people to enjoy.

## First Step
For now the only functionality is creating a hexagonal grid with some debugging options. Just follow the main file at `cmd/main.go` to see what's going on up there.

## Requirements
For this to run you need a Clang compiler. You can checkout the [game engine repository](github.com/hajimehoshi/ebiten) for requirements

## Run the game
Just run `go run cmd/main.go`, and you should be good to go.

## Usage
This repository is starting as proof of concept for making hexagonal maps and working with them using ebiten game engine. As the project gets bigger it may be harder to track so check out all the tags and chose the one that is most relavent to your work.

## Want to see before you run?
Running this project will show you a window like this:

<img width="737" alt="image" src="https://user-images.githubusercontent.com/62159226/223891453-fa6a335a-9583-4641-9ceb-b1c41cb474d6.png">

As of [v0.0.1](https://github.com/MehrdadMiri/simple4x/blob/v0.0.1) the only feature is showing some debugging info including mouse position.

![image](https://user-images.githubusercontent.com/62159226/224164625-265eb558-d0ca-455e-a9eb-8bb36a7e2bc0.png)

As of [v0.0.2](https://github.com/MehrdadMiri/simple4x/blob/v0.0.2) you can move the mouse and you will see the hexagons around your mouse pointer.

## Comming soon
* Watching a slice of the main grid and moving by dragging your mouth or using arrow keys
