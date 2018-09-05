package main

import (
	"fmt"
	"strconv"
	"net/http"
	"io"
	"regexp"
	"os"
)

func HttpGetDB(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return
}

func SpiderPageDB(i int, page chan<- int)  {
	url := "https://movie.douban.com/top250?start=" + strconv.Itoa((i-1)*25) + "&filter="
	result, err := HttpGetDB(url)
	if err != nil {
		fmt.Println("HttpGetDB err:", err)
		return
	}
	// 编译、解析正则表达式 —— 电影名
	ret1 := regexp.MustCompile(`<img width="100" alt="(?s:(.*?))" src="` )
	// 提取有效信息
	fileNames := ret1.FindAllStringSubmatch(result, -1)

	// 编译、解析正则表达式 —— 分数
	pattern := `<span class="rating_num" property="v:average">(.*?)</span>`
	ret2 := regexp.MustCompile(pattern )
	// 提取有效信息
	fileScore := ret2.FindAllStringSubmatch(result, -1)
/*	for _, one := range fileScore {
		fmt.Println("fileName:", one[1])
	}*/
	// 编译、解析正则表达式 —— 评分人数
	ret3 := regexp.MustCompile(`<span>(\d*?)人评价</span>`)
	// 提取有效信息
	peopleNum := ret3.FindAllStringSubmatch(result, -1)

	// 写入到一个文件中
	save2file(i, fileNames, fileScore, peopleNum)

	page <- i	// 写入channel ，协调主go程与子go程调用顺序。
}

func save2file(idx int, fileNames, fileScore, peopleNum [][]string)  {
	// 组织保存文件路径及名程
	path := "C:/exec/第" + strconv.Itoa(idx) + "页.txt"

	f, err := os.Create(path)
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}
	defer f.Close()
	// 获取 一个网页中的条目数 —— 25
	n := len(fileNames)

	// 写一行标题
	f.WriteString("电影名称" + "\t" + "评分" + "\t" + "评价人数" + "\n")

	// 依次按序写入电影相关条目。
	for i:=0; i<n; i++ {
		f.WriteString(fileNames[i][1] + "\t" + fileScore[i][1] + "\t" + peopleNum[i][1] + "\n")
	}
}

func doWork(start, end int)  {
	page := make(chan int)

	// 循环创建多个goroutine，提高爬取效率
	for i:=start; i<=end; i++ {
		go SpiderPageDB(i, page)
	}
	// 循环读取 channel， 协调主、子go程调用顺序
	for i:=start; i<=end; i++ {
		fmt.Printf("第%d页爬取完成\n", <-page)
	}
}

func main()  {
	var start, end int
	fmt.Print("请输入爬取起始页面（>=1）:")
	fmt.Scan(&start)

	fmt.Print("请输入爬取终止页面（>=start）:")
	fmt.Scan(&end)

	doWork(start, end)
}
