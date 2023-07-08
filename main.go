package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SkillCast struct {
	precast  float64
	cast     float64
	backwing float64
}

type FloorSkill struct {
	skillcast *SkillCast
	width     float64
	height    float64
	areaType  string
}

type Castbar struct {
	Image *ebiten.Image
}

type Creature struct {
	Image         *ebiten.Image
	X             float64
	Y             float64
	Skills        *SkillRotation
	state         string
	Casting       time.Time
	CastStartTime time.Time
	CastEndTime   time.Time
	CastbarImage  *ebiten.Image
	casted        bool
	rushing       bool
	TargetX       float64
	TargetY       float64
	Speed         float64
	ElapsedTime   float64
}

type Game struct {
	Player1   *Creature
	Player2   *Creature
	Player3   *Creature
	Player4   *Creature
	Player5   *Creature
	Boss      *Creature
	input     *Input
	StartTime time.Time
}

type SkillRotation struct {
	current int64
	skills  []*FloorSkill
}

func NewCreature(path string, x float64, y float64) *Creature {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		fmt.Println(err)
	}
	bar := ebiten.NewImage(img.Bounds().Dx(), 15)
	return &Creature{
		Image:        img,
		X:            x,
		Y:            y,
		CastbarImage: bar,
	}
}

func calculateSlope(p1, p2 Point) (float64, error) {
	if p1.X == p2.X {
		return 0, fmt.Errorf("无法计算斜率，两点的 x 值相同")
	}

	slope := (p2.Y - p1.Y) / (p2.X - p1.X)
	return slope, nil
}

func CalculateRectanglePoints2(x, y, l, w, s float64) [5]Point {
	radians := math.Pi * s / 180.0
	halfLength := l / 2.0
	halfWidth := w / 2.0

	// 计算矩形的四个角点坐标
	// 第一个角点（左上角）
	x1 := x - halfLength*math.Cos(radians) + halfWidth*math.Sin(radians)
	y1 := y - halfLength*math.Sin(radians) - halfWidth*math.Cos(radians)

	// 第二个角点（右上角）
	x2 := x + halfLength*math.Cos(radians) + halfWidth*math.Sin(radians)
	y2 := y + halfLength*math.Sin(radians) - halfWidth*math.Cos(radians)

	// 第三个角点（右下角）
	x3 := x + halfLength*math.Cos(radians) - halfWidth*math.Sin(radians)
	y3 := y + halfLength*math.Sin(radians) + halfWidth*math.Cos(radians)

	// 第四个角点（左下角）
	x4 := x - halfLength*math.Cos(radians) - halfWidth*math.Sin(radians)
	y4 := y - halfLength*math.Sin(radians) + halfWidth*math.Cos(radians)

	return [5]Point{{X: x1, Y: y1}, {X: x2, Y: y2}, {X: x3, Y: y3}, {X: x4, Y: y4}, {X: x, Y: y}}
}

func calculateRectanglePoints(start, end Point, width float64) [5]Point {
	length := math.Sqrt(math.Pow(end.X-start.X, 2) + math.Pow(end.Y-start.Y, 2))
	centerX := (start.X + end.X) / 2
	centerY := (start.Y + end.Y) / 2

	slope, _ := calculateSlope(start, end)
	angle := math.Atan(slope) * 180 / math.Pi

	return CalculateRectanglePoints2(centerX, centerY, length, 25, angle)
}

var (
	bossSkillImage = ebiten.NewImage(3, 3)
)

func (c *Creature) DrawSkill(screen *ebiten.Image) {
	if c.state == "casting" {
		//s := ebiten.DeviceScaleFactor()
		//ebitenutil.DrawRect(screen, c.X , c.Y, 640*s, 480*s, color.RGBA{255, 0, 0, 1})
		op := &ebiten.DrawTrianglesOptions{}
		op.Address = ebiten.AddressUnsafe
		indices := []uint16{}
		for i := 0; i < 4; i++ {
			indices = append(indices, uint16(i), uint16(i+1)%uint16(4), uint16(4))
		}
		points := calculateRectanglePoints(Point{c.X, c.Y}, Point{c.TargetX, c.TargetY}, 20.0)
		scale := ebiten.DeviceScaleFactor()
		screen.DrawTriangles([]ebiten.Vertex{
			{DstX: float32(points[0].X * scale), DstY: float32(points[0].Y * scale), SrcX: 0, SrcY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
			{DstX: float32(points[1].X * scale), DstY: float32(points[1].Y * scale), SrcX: 0, SrcY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
			{DstX: float32(points[2].X * scale), DstY: float32(points[2].Y * scale), SrcX: 0, SrcY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
			{DstX: float32(points[3].X * scale), DstY: float32(points[3].Y * scale), SrcX: 0, SrcY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
			{DstX: float32(points[4].X * scale), DstY: float32(points[4].Y * scale), SrcX: 0, SrcY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
		}, indices, bossSkillImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image), op)
	}
}

type Point struct {
	X, Y float64
}

func (c *Creature) RandomMove() {
	randomInt := rand.Intn(4)
	move := [4][]int{{2, 0}, {0, 2}, {-2, 0}, {0, -2}}
	c.X = c.X + float64(move[randomInt][0])
	if c.X < 0 {
		c.X = 0
	}
	c.Y = c.Y + float64(move[randomInt][1])
	if c.Y < 0 {
		c.Y = 0
	}
}

func (c *Creature) GetGravityCenter() (float64, float64) {
	return c.X + float64(c.Image.Bounds().Size().X/2), c.Y + float64(c.Image.Bounds().Size().Y/2)
}

var TARGET_W = 40.0
var TARGET_H = 40.0

func (c *Creature) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	dw := c.Image.Bounds().Size().X
	dh := c.Image.Bounds().Size().Y
	scaleX := TARGET_W / float64(dw)
	scaleY := TARGET_H / float64(dh)
	op.GeoM.Translate((c.X-20)/scaleX, (c.Y-20)/scaleY)
	op.GeoM.Scale(scaleX, scaleY)
	scale := ebiten.DeviceScaleFactor()
	op.GeoM.Scale(scale, scale)
	op.Filter = ebiten.FilterLinear
	if !c.Casting.IsZero() {
		op2 := &ebiten.DrawImageOptions{}
		progress := float64(c.Casting.Sub(c.CastStartTime).Milliseconds()) / float64(c.CastEndTime.Sub(c.CastStartTime).Milliseconds())

		c.CastbarImage.Fill(color.Gray{Y: 128})
		c.Image.DrawImage(c.CastbarImage, op2) // 绘制进度条
		if progress != 0 {
			barWidth := int(float64(c.CastbarImage.Bounds().Dx()) * progress)
			bar := ebiten.NewImage(barWidth, 15)
			bar.Fill(color.RGBA{4, 59, 92, 1})
			c.Image.DrawImage(bar, op2)
		}
	}
	screen.DrawImage(c.Image, op)
}

func (c *Creature) Move(dir Dir) {
	vx, vy := dir.Vector()
	c.X = c.X + float64(vx)
	c.Y = c.Y + float64(vy)
	return
}

func (c *Creature) Update(input *Input) {
	if dir, ok := input.Dir(); ok {
		c.Move(dir)
	}
}

func (c *Creature) CastSkill() {
}

var PRECAST = 300
var CAST = 1500
var BACKWING = 500

func (g *Game) distance(x1, y1, x2, y2 float64) float64 {
	dx := float64(x2 - x1)
	dy := float64(y2 - y1)
	return dx*dx + dy*dy
}

func (c *Creature) CastMelee(g *Game) {
	c.casted = true
	c.Casting = time.Now()
	if c.state == "precast" {
		if c.CastEndTime.Before(c.Casting) {
			c.state = "casting"
			c.CastStartTime = time.Now()
			c.CastEndTime = c.CastStartTime.Add(time.Millisecond * 1500)
			players := [5]*Creature{}
			players[0] = g.Player1
			players[1] = g.Player2
			players[2] = g.Player3
			players[3] = g.Player4
			players[4] = g.Player5

			target := players[0]
			for _, char := range players {
				if char != g.Boss {
					if g.distance(g.Boss.X, g.Boss.Y, target.X, target.Y) > g.distance(g.Boss.X, g.Boss.Y, char.X, char.Y) {
						target = char
					}
				}
			}
			c.TargetX = target.X
			c.TargetY = target.Y
		}
	} else if c.state == "casting" {
		if c.CastEndTime.Before(c.Casting) {
			c.state = "backwing"
			c.CastStartTime = time.Now()
			c.CastEndTime = c.CastStartTime.Add(time.Millisecond * 500)
		}
	} else if c.state == "backwing" {
		if c.CastEndTime.Before(c.Casting) {
			c.state = "precast"
			c.CastStartTime = time.Now()
			c.CastEndTime = c.CastStartTime.Add(time.Millisecond * 2000)
		}
	} else {
		c.state = "precast"
		c.CastStartTime = time.Now()
		c.CastEndTime = c.CastStartTime.Add(time.Millisecond * 2000)
	}
}

func (c *Creature) CastRush(g *Game) {
	c.casted = true
	c.Casting = time.Now()
	if c.state == "precast" {
		if c.CastEndTime.Before(c.Casting) {
			c.state = "casting"
			c.CastStartTime = time.Now()
			c.CastEndTime = c.CastStartTime.Add(time.Millisecond * 1500)
			players := [5]*Creature{}
			players[0] = g.Player1
			players[1] = g.Player2
			players[2] = g.Player3
			players[3] = g.Player4
			players[4] = g.Player5

			target := players[0]
			for _, char := range players {
				if char != g.Boss {
					if g.distance(g.Boss.X, g.Boss.Y, target.X, target.Y) < g.distance(g.Boss.X, g.Boss.Y, char.X, char.Y) {
						target = char
					}
				}
			}
			c.TargetX = target.X
			c.TargetY = target.Y
		}
	} else if c.state == "casting" {
		if c.CastEndTime.Before(c.Casting) {
			c.state = "rushing"
			c.CastStartTime = time.Now()
			c.CastEndTime = c.CastStartTime.Add(time.Millisecond * 500)
		}
	} else if c.state == "rushing" {
		// find farest player

		c.Speed = 2.0
		c.state = "moving"
	} else if c.state == "moving" {
		dt := 1.0 / ebiten.ActualTPS()
		c.ElapsedTime += dt

		if c.ElapsedTime >= 3.0 {
			c.X = c.TargetX
			c.Y = c.TargetY
			c.state = "backwing"
			c.ElapsedTime = 0.0
			c.Speed = 0
		} else {
			c.X = c.X + (c.TargetX-c.X)*c.Speed*dt
			c.Y = c.Y + (c.TargetY-c.Y)*c.Speed*dt
		}
	} else if c.state == "backwing" {
		if c.CastEndTime.Before(c.Casting) {
			c.state = "precast"
			c.CastStartTime = time.Now()
			c.CastEndTime = c.CastStartTime.Add(time.Millisecond * 2000)
		}
	} else {
		c.state = "precast"
		c.CastStartTime = time.Now()
		c.CastEndTime = c.CastStartTime.Add(time.Millisecond * 2000)
	}
}

func NewGame() *Game {
	g := &Game{
		input:   NewInput(),
		Boss:    NewCreature("images/boss.png", 100, 100),
		Player1: NewCreature("images/luin.png", 150.0, 240.0),
		Player2: NewCreature("images/yuna.png", 240.0, 150.0),
		Player3: NewCreature("images/kaito.png", 40.0, 250.0),
		Player4: NewCreature("images/oruta.png", 250.0, 40.0),
		Player5: NewCreature("images/namalie.png", 250.0, 250.0),
	}
	return g
}

func (g *Game) Update() error {
	g.input.Update()
	g.Player1.Update(g.input)
	g.Player2.RandomMove()
	g.Player3.RandomMove()
	g.Player4.RandomMove()
	g.Player5.RandomMove()

	g.Boss.CastMelee(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Boss.DrawSkill(screen)
	g.Boss.Draw(screen)
	g.Player1.Draw(screen)
	g.Player2.Draw(screen)
	g.Player3.Draw(screen)
	g.Player4.Draw(screen)
	g.Player5.Draw(screen)
	ebitenutil.DebugPrint(screen, "Hello, World!\n"+g.Boss.Casting.String()+"\n"+g.Boss.state+"\n"+strconv.FormatFloat(g.Boss.X, 'f', -1, 64)+","+strconv.FormatFloat(g.Boss.Y, 'f', -1, 64))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// The unit of outsideWidth/Height is device-independent pixels.
	// By multiplying them by the device scale factor, we can get a hi-DPI screen size.
	s := ebiten.DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}

func main() {
	bossSkillImage.Fill(color.RGBA{189, 22, 64, 1})
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
