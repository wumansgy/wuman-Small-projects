package main

import (
	"Clib"
	"fmt"
	"os"
	"runtime"
	"time"
)

const wide1 int = 40
const high1 int = 20
var bullet1 bullet
type position1 struct {
	x int
	y int
} //父类结构体
type bullet struct {
	position1 //子弹
}
type plane struct {
	dir  byte
	body [7]position1
} //飞机结构体
func drawui1(p position1, ch byte) {
	Clib.GotoPostion(p.x*2+4, p.y+2+2)
	fmt.Fprintf(os.Stderr, "%c", ch)
} //画图函数
func (p *plane) initplane() { //初始化飞机
	p.body[0].x = wide1 / 2
	p.body[0].y = high1 / 2
	p.body[1].x = wide1 / 2
	p.body[1].y = high1/2 + 1
	p.body[2].x = wide1 / 2
	p.body[2].y = high1/2 + 2
	p.body[3].x = wide1/2 - 1
	p.body[3].y = high1/2 + 1
	p.body[4].x = wide1/2 + 1
	p.body[4].y = high1/2 + 1
	p.body[5].x = wide1/2 - 1
	p.body[5].y = high1/2 + 2
	p.body[6].x = wide1/2 + 1
	p.body[6].y = high1/2 + 2
	for i := 0; i < len(p.body); i++ {
		drawui1(p.body[i], '#') //
	}
	go func() { //得到方向
		for {
			switch Clib.Direction() {
			case 72, 87, 119:
				if p.dir == 'D' {
					break
				}
				p.dir = 'U'
			case 65, 97, 75:
				if p.dir == 'R' {
					break
				}
				p.dir = 'L'
			case 100, 68, 77:
				if p.dir == 'L' {
					break
				}
				p.dir = 'R'
			case 83, 115, 80:
				if p.dir == 'U' {
					break
				}
				p.dir = 'D'
			case 32:
				p.dir = 'P'
			}
		}
	}()
}
func (p *plane) playplane() { //玩飞机
	bullet1.y=high1/2-1
	bullet1.x=wide1/2
	go func() {
		for{
			or:=bullet1.position1
			drawui1(bullet1.position1, 'O')
			time.Sleep(time.Second)
			bullet1.y--
			drawui1(or,' ')
			if bullet1.y<=-5{
				runtime.Goexit()
			}
		}
	}()
	for {
		if p.dir == 'P' {
			continue
		}
		if p.dir == 'L' {
			Clib.Cls()
			for i := 0; i < len(p.body); i++ {
				p.body[i].x--
			}
			for i := 0; i < len(p.body); i++ {
				drawui1(p.body[i], '#')
			}
			if p.body[0].x<=-1{
				fmt.Println("死亡")
				return
			}
			p.dir = 'P'
		}
		if p.dir == 'R' {
			Clib.Cls()
			for i := 0; i < len(p.body); i++ {
				p.body[i].x++
			}
			for i := 0; i < len(p.body); i++ {
				drawui1(p.body[i], '#')
			}
			p.dir = 'P'
		}
		if p.dir == 'D' {
			Clib.Cls()
			for i := 0; i < len(p.body); i++ {
				p.body[i].y++
			}
			for i := 0; i < len(p.body); i++ {
				drawui1(p.body[i], '#')
			}
			p.dir = 'P'
		}
		if p.dir == 'U' {
			Clib.Cls()
			for i := 0; i < len(p.body); i++ {
				p.body[i].y--
			}
			for i := 0; i < len(p.body); i++ {
				drawui1(p.body[i], '#')
			}
			if p.body[0].y<=-4{
				fmt.Println("死亡")
				return
			}
			p.dir = 'P'
		}
	}
}
func main() {
	var p1 plane
	Clib.HideCursor()
	p1.initplane()
	p1.playplane()
	time.Sleep(time.Second*3)

}
