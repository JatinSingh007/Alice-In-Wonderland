package movement

import (
	"image/color"

	"github.com/hajemoshi/ebiten/v2"
	"github.com/hajemoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	PlayerX float64
	PlayerY float64
}

func (g *Game) Update() error {
	speed := 4.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.PlayerX -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.PlayerX += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.PlayerX -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.PlayerX += speed
	}

	// For screen Boundaries
	if g.PlayerX < 0 {
		g.PlayerX = 0
	}
	if g.PlayerY < 0 {
		g.PlayerY = 0
	}
	if g.PlayerX > 304 {
		g.PlayerX = 304
	}
	if g.playerY > 224 {
		g.playerY = 224
	}

	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	ebitenutil.DrawRect(
		screen,
		g.PlayerX,
		g.PlayerY,
		16,
		16,
		color.White,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}
