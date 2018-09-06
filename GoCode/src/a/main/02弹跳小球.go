package main

import (
	"Clib"
	"fmt"
	"time"
)

type position struct{
	x int
	y int
}
type ball struct{
	position
}
var dxx int=1
var dyy int=1
func drawball(p position){
	Clib.GotoPostion(p.x,p.y)
	fmt.Printf("O")
}
func (b *ball)initball(){
	b.x=10
	b.y=10
}
func (b *ball)playball(){
	for{
		Clib.Cls()
		drawball(b.position)
		time.Sleep(time.Second/10)
		Clib.Cls()
		b.x+=dxx
		//b.y=-b.x*b.x+8*b.x+1
		b.y+=dyy
		if b.x<=0||b.x>=55{
			dxx=-dxx
		}
		if b.y<=0||b.y>=32{
			dyy=-dyy
		}
	}
}
func main(){
	var bal ball
bal.initball()
bal.playball()
}
