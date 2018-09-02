package main

import (
	"Clib"
	"fmt"
	"os"
	"math/rand"
	"time"
)

const wide int = 20
const high int = 20

var key int64 = 1 //关卡
var food1 food    //定义一个全局食物结构体
//var size int = 2  //定义一个全局蛇的长度
var score int = 0 //定义一个全局分数
var dx int = 0
var dy int = 0 //蛇的偏移量
var barr1 barrier //障碍物结构体
var c cake //定义一个蛋糕
var FLAG bool=true
type postion struct {
	x int
	y int //父类坐标
}
type cake struct{
	  ca [5]postion
}                    //定义一个蛋糕
type snake struct {
	p    [wide * high]postion
	size int
	dir  byte
}
type barrier struct {
	barr [6]postion
} //障碍物结构体
func (c *cake)setcake(){
	x:=rand.Intn(wide-6)+3
	y:=rand.Intn(high-6)+3
	c.ca[0].x,c.ca[0].y=x,y
		c.ca[1].x,c.ca[1].y=x-1,y
	c.ca[2].x,c.ca[2].y= x-2,y
		c.ca[3].x,c.ca[3].y=x-1,y-1
	c.ca[4].x,c.ca[4].y=x-1,y+1
}
func (b *barrier)setbarrier(){   //定义一些随机障碍物
	b.barr[0].x,b.barr[0].y=rand.Intn(wide-1)+1,rand.Intn(high-3)+1
	b.barr[1].x,b.barr[1].y=rand.Intn(wide-1)+1,rand.Intn(high-3)+1
	b.barr[2].x,b.barr[2].y=rand.Intn(wide-1)+1,rand.Intn(high-3)+1
	//b.barr[3].x,b.barr[3].y=rand.Intn(wide-1)+1,rand.Intn(high-3)+1
	//b.barr[4].x,b.barr[4].y=rand.Intn(wide-1)+1,rand.Intn(high-3)+1
	//b.barr[5].x,b.barr[5].y=rand.Intn(wide-1)+1,rand.Intn(high-3)+1
}
type food struct {
	postion
} //食物
func drawui(p postion, ch byte) {
	Clib.GotoPostion(p.x*2+4, p.y+2+2)
	fmt.Fprintf(os.Stderr, "%c", ch)
}
func (s *snake) initsnake() { //蛇初始化
	s.p[0].x = wide / 2
	s.p[0].y = high / 2
	s.p[1].x = wide/2 - 1
	s.p[1].y = high / 2 //蛇头和第一个蛇结点初始化
	s.dir = 'R'
	s.size=2
	fmt.Fprintln(os.Stderr,
		`
  #-----------------------------------------#
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  #-----------------------------------------#
`)
	food1 = food{postion{rand.Intn(wide), rand.Intn(high) - 2}} //食物初始化
	drawui(food1.postion, 'o')                                  //画出食物

	//go func(){
	 Clib.GotoPostion(46,19)
	 fmt.Printf("正在进行第%d关，小心障碍物",key)
	//}()
	go func(){
		for{
			time.Sleep(time.Second)
			num:=rand.Intn(10)
			if num==6{
				c.setcake()
				break
			}
		}
		for i:=0;i<len(c.ca);i++{
			drawui(c.ca[i],'#')
		}
	}()                  //吃蛋糕的作用
	go func(){
		for i:=0;i<len(barr1.barr);i++{
			Clib.GotoPostion(barr1.barr[i].x,barr1.barr[i].y)
			drawui(barr1.barr[i],'!')
		}
	}()             //打印出障碍物
	go func() {
		for {
			switch Clib.Direction() {
			case 72, 87, 119:
				if s.dir == 'D' {
					break
				}
				s.dir = 'U'
			case 65, 97, 75:
				if s.dir == 'R' {
					break
				}
				s.dir = 'L'
			case 100, 68, 77:
				if s.dir == 'L' {
					break
				}
				s.dir = 'R'
			case 83, 115, 80:
				if s.dir == 'U' {
					break
				}
				s.dir = 'D'
			case 32:
				s.dir = 'P'
			}
		}
	}()   //获取蛇跑的方向
}

func (s *snake) playgame() {
	//barr:=barrier{postion{rand.Intn(wide-5)+5,rand.Intn(high-5)+3}
		//drawui(barr.postion,'p')
	for {
		switch key {
		case 1: time.Sleep(time.Second / 3)
		case 2:time.Sleep(time.Second / 5)
		case 3:time.Sleep(time.Second / 6)
		case 4:time.Sleep(time.Second / 7)
		case 5:time.Sleep(time.Second / 8)
		case 6:time.Sleep(time.Second / 9)    //用来每增加一关蛇的速度加快
		}


		if s.dir == 'P' {
			continue
		}
		if s.p[0].x < 0 || s.p[0].x >= wide || s.p[0].y+2 < 0 || s.p[0].y >= high-2 {
			Clib.GotoPostion(wide*3, high-3)
			FLAG=false
			return //如果蛇头碰墙就死亡
		}
		//if s.p[0].x==barr.postion.x&&s.p[0].y==barr.postion.y{
		//	Clib.GotoPostion(wide*3, high-3)
		//	return //如果蛇头碰障碍物就死亡
		//}

		for i := 1; i <s.size; i++ {
			if s.p[0].x == s.p[i].x && s.p[0].y == s.p[i].y {
				Clib.GotoPostion(wide*3, high-3)
				FLAG=false
				return
			}
		}
		for j:=0;j<len(barr1.barr);j++{
			if s.p[0].x==barr1.barr[j].x&&s.p[0].y==barr1.barr[j].y{
				Clib.GotoPostion(wide*3, high-3)
				FLAG=false
				return
			}                //碰到障碍物死亡
		}
		for m:=0;m<len(c.ca);m++{
			if s.p[0].x==c.ca[m].x&&s.p[0].y==c.ca[m].y{
				s.size++
				score++
			}
			if score >= int(6+key*2) {
				key++
				return
			}
		}
		if s.p[0].x == food1.x && s.p[0].y == food1.y {
			s.size++
			score++
			if score >= int(6+key*2) {
				key++
				return
			}
			//画蛇
			//food1 = food{postion{rand.Intn(wide), rand.Intn(high) - 2}}
			for {
				flag := true
				temp := food{postion{rand.Intn(wide), rand.Intn(high) - 2}}
				for i := 1; i < s.size; i++ {
					if (temp.postion.x == s.p[i].x && temp.postion.y == s.p[i].y)  {
						flag = false
						break
					}
				}
				for i:=0;i<len(barr1.barr);i++{
					if temp.postion.x==barr1.barr[i].x&&temp.postion.y==barr1.barr[i].y{
						flag=false
						break
					}
				}
				if flag == true {
					food1 = temp
					break
				}

			}
			drawui(food1.postion, 'o')
		}

		switch s.dir {
		case 'U':
			dx = 0
			dy = -1
		case 'D':
			dx = 0
			dy = 1
		case 'L':
			dx = -1
			dy = 0
		case 'R':
			dx = 1
			dy = 0
		}
		lp := s.p[s.size-1] //蛇尾位置
		for i := s.size - 1; i > 0; i-- {
			s.p[i] = s.p[i-1]
			drawui(s.p[i], '*')
		}
		drawui(lp, ' ') //蛇尾画空格

		s.p[0].x += dx
		s.p[0].y += dy      //更新蛇头
		drawui(s.p[0], 'O') //画蛇头
	}

}
func main() {
	rand.Seed(time.Now().UnixNano())
	var s snake

	for k:=1;k<=6;k++{    //用来循环6次代表6个关卡，这里可以自己设置多少关卡
		s.initsnake()        //初始化
		barr1.setbarrier()  //障碍物
		s.playgame()       //玩游戏开始
		if FLAG==false{        //这个代表蛇死亡返回的，所以这样就退出了
			Clib.GotoPostion(46,21)
			fmt.Printf("你已死亡，第%d关总分：%d分",k, score)
			break
		}
		Clib.GotoPostion(46,21)
		fmt.Printf("第%d关总分：%d分,稍等进入下一关",k, score)
		//key++
		time.Sleep(time.Second * 5)  //延时5秒
		Clib.Cls()                  //每一关清屏一下
		//size=2
		score=0                   //每一关分数置为0
	}

	time.Sleep(time.Second * 5)   //延时5秒
}
