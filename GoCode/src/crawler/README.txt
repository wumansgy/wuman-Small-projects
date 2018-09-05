豆瓣电影爬取：

	1. 明确 url。 
	
		https://movie.douban.com/top250?start=0&filter=		1

		https://movie.douban.com/top250?start=25&filter=		2
	
		https://movie.douban.com/top250?start=50&filter=		3

		https://movie.douban.com/top250?start=75&filter=		4

	2. 待提取字符特性：

		1）电影名称： <img width="100" alt="（电影名称）" src="	—— `<img width="100" alt="(?s:(.*?))" src="`  	或 (.*?)

		2） 评分：<span class="rating_num" property="v:average">（评分）</span>

		3）评价人数： <span>（评价人数）人评价</span>		—— `<span>(\d*?)人评价</span>`

	3. 提示用户指定爬取起始、终止页

	4. 封装 doWork 函数， 按起始、终止页面循环爬取网页数据

	5. 组织每个网页的 url。 下一页 = +25

	6. 封装函数 HttpGetDB（url）result，err { http.Get(url), resp.Body.Read(buf), n==0 break, result+= string(buf[:n}) }

		爬取网页的所有数据 通过result 返回给调用者。

	7. 解析、编译正则表达式 —— 提取 “电影名称”fileNames 传出的是[][]string ， 下标为【1】是不带匹配参考项。

	8. 解析、编译正则表达式 —— 提取 “评分”传出的是[][]string ， 下标为【1】是不带匹配参考项。

	9. 解析、编译正则表达式 —— 提取 “评价人数”传出的是[][]string ， 下标为【1】是不带匹配参考项。

	10. 封装函数，将上述内容写入文件。save2File（ [][]string）

		1) Create()

		2) n = len(fileNames)
		
		3）for 循环 一次写入一条电影信息。 f.writeString()

	11. 创建并发go程 提取所有网页数据。
	
		for {
			go SpiderPageDB()
		}

	12. 创建阻止主go程提取退出的 channel ， SpiderPageDB() 末尾处，写channel

	13. doWork 中，添加新 for ，读channel
