package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	roomWidth  = 40
	roomHeight = 20
)

type Game struct {
	app     *tview.Application
	grid    *tview.TextView
	playerX int
	playerY int
	objects map[[2]int]rune
}

func NewGame() *Game {
	g := &Game{
		app:     tview.NewApplication(),
		grid:    tview.NewTextView().SetDynamicColors(true),
		playerX: roomWidth / 2,
		playerY: roomHeight / 2,
		objects: make(map[[2]int]rune),
	}
	g.grid.SetChangedFunc(func() {
		g.app.Draw()
	})
	return g
}

func (g *Game) addObjects() {
	// Add a few static objects
	g.objects[[2]int{5, 5}] = '#'
	g.objects[[2]int{10, 8}] = '#'
	g.objects[[2]int{15, 15}] = '#'
	g.objects[[2]int{30, 10}] = '#'
}

func (g *Game) render() {
	g.grid.Clear()
	for y := 0; y <= roomHeight; y++ {
		for x := 0; x <= roomWidth; x++ {
			if x == g.playerX && y == g.playerY {
				g.grid.Write([]byte("@"))
			} else if ch, ok := g.objects[[2]int{x, y}]; ok {
				g.grid.Write([]byte(string(ch)))
			} else if x == roomWidth || y == roomHeight {
				g.grid.Write([]byte("|"))
			} else {
				g.grid.Write([]byte(" "))
			}
		}
		g.grid.Write([]byte("\n"))
	}
}

func (g *Game) move(dx, dy int) {
	newX := g.playerX + dx
	newY := g.playerY + dy
	if newX >= 0 && newX < roomWidth && newY >= 0 && newY < roomHeight {
		g.playerX = newX
		g.playerY = newY
	}
}

func (g *Game) setupInput() {
	g.grid.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'w':
			g.move(0, -1)
		case 's':
			g.move(0, 1)
		case 'a':
			g.move(-1, 0)
		case 'd':
			g.move(1, 0)
		// support dvorak
		case ',':
			g.move(0, -1)
		case 'o':
			g.move(0, 1)
		case 'e':
			g.move(1, 0)
		}
		g.render()
		return nil
	})
}

func main() {
	game := NewGame()
	game.addObjects()
	game.setupInput()
	game.render()

	if err := game.app.SetRoot(game.grid, true).Run(); err != nil {
		panic(err)
	}
}
