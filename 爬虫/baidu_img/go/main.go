package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

type Spider_baidu_image struct {
	SearchKeyword string
	SearchNumber  int
	PageCount     int
	ReqUrl        string
	ReqMethod     string
	ReqHeaders    map[string]string
	CurrentPath   string
	ImgDirName    string
	ImgDir        string
	ImgUrls       []string
}

type RespJson struct {
	Data []struct {
		Url string `json:"thumbURL"`
	} `json:"data"`
}

func main() {
	newSpider().run()
	//300张刘亦菲的图片已下载完成！耗时: 1.269014675s
}

func (s *Spider_baidu_image) run() {
	defer func(start time.Time) {
		terminal := time.Since(start)
		fmt.Println("耗时:", terminal)
	}(time.Now())
	s.getAllUrls()
	s.getAllImgData()
}

func (s *Spider_baidu_image) getAllImgData() {
	saveNum := s.SearchNumber
	allCount := len(s.ImgUrls)
	if allCount == 0 {
		fmt.Println("没有爬取到图片")
		return
	}
	if allCount < s.SearchNumber {
		fmt.Printf("只爬取到%d张图片，没有达到%d张，是否继续？(y/n)", allCount, saveNum)
		var answer string
		fmt.Scanln(&answer)
		if answer == "n" {
			return
		} else if answer == "y" {
			saveNum = allCount
		} else {
			s.getAllImgData()
			return
		}
	}

	var wg sync.WaitGroup
	for k := 1; k <= saveNum; k++ {
		wg.Add(1)
		go func(index int, url string) {
			c := s.newColly(1)
			c.OnResponse(func(resp *colly.Response) {
				file, err := os.Create(fmt.Sprintf("%s/%d.jpg", s.ImgDir, index))
				if err != nil {
					fmt.Println("文件创建失败", err.Error())
					return
				}
				file.Write(resp.Body)
				defer file.Close()
			})
			c.Visit(url)
			c.Wait()
			defer wg.Done()
		}(k, s.ImgUrls[k])
	}
	wg.Wait()
	fmt.Printf("%d张%s的图片已下载完成！", s.SearchNumber, s.SearchKeyword)
}

func (s *Spider_baidu_image) getAllUrls() {
	var wg sync.WaitGroup
	for pageNum := 1; pageNum <= s.PageCount; pageNum++ {
		wg.Add(1)
		go func(pageNum int) {
			s.getSinglePageImgUrls(pageNum)
			defer wg.Done()
		}(pageNum)
	}
	wg.Wait()
}

func (s *Spider_baidu_image) getSinglePageImgUrls(pageNum int) {
	c := s.newColly(pageNum)
	c.OnResponse(func(resp *colly.Response) {
		var respinfo RespJson
		e := json.Unmarshal(resp.Body, &respinfo)
		if e != nil {
			fmt.Println("json.Unmarshal failed")
			return
		}
		for _, v := range respinfo.Data {
			s.ImgUrls = append(s.ImgUrls, v.Url)
		}
	})
	c.Visit(s.ReqUrl)
	c.Wait()
}

func (s *Spider_baidu_image) newColly(pageNum int) *colly.Collector {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		for k, v := range s.ReqHeaders {
			r.Headers.Add(k, v)
		}
		r.Headers.Add("pn", strconv.Itoa(pageNum*60))
	})
	return c
}
func newSpider() *Spider_baidu_image {
	url := "https://image.baidu.com/search/acjson"
	params := map[string]string{
		"tn":        "resultjson_com",
		"logid":     "11763858014708962078",
		"ipn":       "rj",
		"ct":        "201326592",
		"is":        "",
		"fp":        "result",
		"fr":        "",
		"cg":        "star",
		"cl":        "2",
		"lm":        "-1",
		"ie":        "utf-8",
		"oe":        "utf-8",
		"adpicid":   "",
		"st":        "-1",
		"z":         "",
		"ic":        "0",
		"hd":        "",
		"latest":    "",
		"copyright": "",
		"s":         "",
		"se":        "",
		"tab":       "",
		"width":     "",
		"height":    "",
		"face":      "0",
		"istype":    "2",
		"qc":        "",
		"nc":        "1",
		"expermode": "",
		"nojc":      "",
		"isAsync":   "",
		"rn":        "60",
		"gsm":       "1e",
	}
	var keyword string
	fmt.Print("请输入搜索图片关键字:")
	fmt.Scanln(&keyword)
	params["word"] = keyword
	params["queryWord"] = keyword

	var number int
	fmt.Print("请输入搜索数量:")
	fmt.Scanln(&number)
	pageCount := (number + 60 - 1) / 60
	// params["pn"] = strconv.Itoa(pageCount * 60)
	var now = time.Now().Local()
	var nowtime = strconv.FormatInt(now.UnixMilli(), 10)
	params[nowtime] = ""
	url = fmt.Sprintf("%s?%s=", url, nowtime)
	for k, v := range params {
		url = fmt.Sprintf("%s&%s=%s", url, k, v)
	}
	headers := map[string]string{
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"Accept-Encoding":           "gzip, deflate, br",
		"Accept-Language":           "zh-CN,zh;q=0.9",
		"Cache-Control":             "max-age=0",
		"Connection":                "keep-alive",
		"Cookie":                    "BDqhfp=%E5%88%98%E4%BA%A6%E8%8F%B2%26%260-10-1undefined%26%260%26%261; BIDUPSID=7947D31494D7968F801FC22F739B3F61; PSTM=" + strconv.FormatInt(now.Unix(), 10) + "; BAIDUID=7947D31494D7968F0A75DA00A84C5631':F'G=1; H_PS_PSSID=35410_35104_31253_35736_35489_34584_35491_34812_35684_26350_35751_35746; BAIDUID_BFESS=7947D31494D7968F0A75DA00A84C5631':F'G=1; BDRCVFR[X_XKQks0S63]=mk3SLVN4HKm; __yjs_duid=1_dd2b650a1268e02c97f9ca0ca635d11a1642646666312; firstShowTip=1; indexPageSugList=%5B%22%E5%88%98%E4%BA%A6%E8%8F%B2%22%5D; cleanHistoryStatus=0; BDRCVFR[dG2JNJb_ajR]=mk3SLVN4HKm; BDRCVFR[-pGxjrCMryR]=mk3SLVN4HKm",
		"DNT":                       "1",
		"Host":                      "image.baidu.com",
		"sec-ch-ua":                 "'Not A;Brand';v='99', 'Chromium';v='96', 'Google Chrome';v='96'",
		"sec-ch-ua-mobile":          "?0",
		"sec-ch-ua-platform":        "macOS",
		"Sec-Fetch-Dest":            "document",
		"Sec-Fetch-Mode":            "navigate",
		"Sec-Fetch-Site":            "none",
		"Sec-Fetch-User":            "?1",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36",
	}
	dir, _ := os.Getwd()
	s := &Spider_baidu_image{
		SearchKeyword: keyword,
		SearchNumber:  number,
		PageCount:     pageCount,
		ReqUrl:        url,
		ReqMethod:     "GET",
		ReqHeaders:    headers,
		CurrentPath:   dir,
		ImgDirName:    keyword,
		ImgDir:        fmt.Sprintf("%s/%s", dir, keyword),
	}
	var dirExist = func() bool {
		dir, err := os.Stat(s.ImgDir)
		if err != nil {
			return false
		}
		return dir.IsDir()
	}
	var verifyDir func()
	verifyDir = func() {
		if dirExist() {
			fmt.Printf("文件夹'%s'已经存在，是否使用该文件夹保存图片？(y/n):", s.ImgDirName)
			var answer string
			fmt.Scan(&answer)
			answer = strings.Trim(answer, " ")
			if answer == "y" || answer == "Y" || answer == "yes" {
				return
			} else if answer == "n" || answer == "N" || answer == "no" {
				fmt.Print("请设置新的文件夹名字:")
				fmt.Scan(&s.ImgDirName)
				s.ImgDir = fmt.Sprintf("%s/%s", s.CurrentPath, s.ImgDirName)
			}
			verifyDir()
			return
		}
		err := os.Mkdir(s.ImgDir, os.ModePerm)
		if err != nil {
			log.Fatalf("创建文件出错![%v]\n", err)
		}
	}
	verifyDir()
	return s
}
