package main

import (
	"math"

	"github.com/alivedise/tsuruki/point"
)

func triangleArea(p1, p2, p3 point.Point) float64 {
	return 0.5 * math.Abs((p1.X-p3.X)*(p2.Y-p1.Y)-(p1.X-p2.X)*(p3.Y-p1.Y))
}

func TriangleCollision(p1, p2, p3, collisionPoint point.Point) bool {
	// 計算三角形的面積
	area := triangleArea(p1, p2, p3)

	// 計算碰撞點到三個三角形頂點的面積和
	area1 := triangleArea(p1, p2, collisionPoint)
	area2 := triangleArea(p2, p3, collisionPoint)
	area3 := triangleArea(p3, p1, collisionPoint)

	// 如果碰撞點到三角形三個邊的面積和等於三角形的面積，表示碰撞點在三角形內部
	return math.Abs(area1+area2+area3-area) < 0.000001
}
