package main

import (
	"os"
	"log"
	"io"
	"Btos"
)
/*func main(){
	s:="11101001 01101001 10110110 00101000"
	bs:=Btos.BinaryStringToBytes(s)
	fmt.Println(bs)

}*/
func main() {
/*	bs := []byte{1, 'a', 3}
	s := Btos.BytesToBinaryString(bs)
   fmt.Println(s)*/
	fp,err:=os.Open("d:/exec/222.jpg")
	fp1,err1:=os.Create("d:/exec/222加密.jpg")
	if err!=nil||err1!=nil{
		log.Fatal(err)
	}
	defer fp.Close()
	defer fp1.Close()
	buf:=make([]byte,1)
	for{
		n,err:=fp.Read(buf)
		buf1:=Btos.BinaryStringToBytes(string(buf[:n]))
		fp1.Write([]byte(buf1))
		if err==io.EOF{
			break
		}
	}
}
