package spiders

import (
	"strconv"
	"strings"
	"regexp"
	"log"
	"encoding/json"

	"github.com/PuerkitoBio/goquery"

)

var (
	compileNumber = regexp.MustCompile("[0-9]")
	areaJson = `
        [{ "Id": 340000, "Name": "安徽", "FirstCharacter": "A", "Pinyin": "anhui", "City": [{ "Id": 340100, "Name": "合肥", "FirstCharacter": "H", "Pinyin": "hefei", "OldCityId": 1 }, { "Id": 340200, "Name": "芜湖", "FirstCharacter": "W", "Pinyin": "wuhu", "OldCityId": 9 }, { "Id": 340800, "Name": "安庆", "FirstCharacter": "A", "Pinyin": "anqing", "OldCityId": 11 }, { "Id": 341200, "Name": "阜阳", "FirstCharacter": "F", "Pinyin": "fu_yang", "OldCityId": 4 }] }, { "Id": 110000, "Name": "北京", "FirstCharacter": "B", "Pinyin": "beijing", "City": [{ "Id": 110100, "Name": "北京", "FirstCharacter": "B", "Pinyin": "beijing", "OldCityId": 646 }] }, { "Id": 500000, "Name": "重庆", "FirstCharacter": "C", "Pinyin": "chongqing", "City": [{ "Id": 500100, "Name": "重庆", "FirstCharacter": "C", "Pinyin": "chongqing", "OldCityId": 648 }] }, { "Id": 350000, "Name": "福建", "FirstCharacter": "F", "Pinyin": "fujian", "City": [{ "Id": 350100, "Name": "福州", "FirstCharacter": "F", "Pinyin": "fuzhou", "OldCityId": 24 }, { "Id": 350200, "Name": "厦门", "FirstCharacter": "X", "Pinyin": "xiamen", "OldCityId": 23 }, { "Id": 350500, "Name": "泉州", "FirstCharacter": "Q", "Pinyin": "quanzhou", "OldCityId": 28 }] }, { "Id": 620000, "Name": "甘肃", "FirstCharacter": "G", "Pinyin": "gansu", "City": [{ "Id": 620100, "Name": "兰州", "FirstCharacter": "L", "Pinyin": "lanzhou", "OldCityId": 46 }] }, { "Id": 440000, "Name": "广东", "FirstCharacter": "G", "Pinyin": "guangdong", "City": [{ "Id": 440100, "Name": "广州", "FirstCharacter": "G", "Pinyin": "guangzhou", "OldCityId": 62 }, { "Id": 440200, "Name": "韶关", "FirstCharacter": "S", "Pinyin": "shaoguan", "OldCityId": 65 }, { "Id": 440300, "Name": "深圳", "FirstCharacter": "S", "Pinyin": "shenzhen", "OldCityId": 670 }, { "Id": 440400, "Name": "珠海", "FirstCharacter": "Z", "Pinyin": "zhuhai", "OldCityId": 74 }, { "Id": 440500, "Name": "汕头", "FirstCharacter": "S", "Pinyin": "shantou", "OldCityId": 69 }, { "Id": 440600, "Name": "佛山", "FirstCharacter": "F", "Pinyin": "foshan", "OldCityId": 77 }, { "Id": 440700, "Name": "江门", "FirstCharacter": "J", "Pinyin": "jiangmen", "OldCityId": 76 }, { "Id": 441300, "Name": "惠州", "FirstCharacter": "H", "Pinyin": "huizhou", "OldCityId": 72 }, { "Id": 441900, "Name": "东莞", "FirstCharacter": "D", "Pinyin": "dongguan", "OldCityId": 73 }, { "Id": 442000, "Name": "中山", "FirstCharacter": "Z", "Pinyin": "zhongshan", "OldCityId": 75 }] }, { "Id": 450000, "Name": "广西", "FirstCharacter": "G", "Pinyin": "guangxi", "City": [{ "Id": 450100, "Name": "南宁", "FirstCharacter": "N", "Pinyin": "nanning", "OldCityId": 574 }, { "Id": 450200, "Name": "柳州", "FirstCharacter": "L", "Pinyin": "liuzhou", "OldCityId": 576 }, { "Id": 450300, "Name": "桂林", "FirstCharacter": "G", "Pinyin": "guilin", "OldCityId": 575 }] }, { "Id": 520000, "Name": "贵州", "FirstCharacter": "G", "Pinyin": "guizhou", "City": [{ "Id": 520100, "Name": "贵阳", "FirstCharacter": "G", "Pinyin": "guiyang", "OldCityId": 106 }, { "Id": 520300, "Name": "遵义", "FirstCharacter": "Z", "Pinyin": "zunyi", "OldCityId": 108 }] }, { "Id": 460000, "Name": "海南", "FirstCharacter": "H", "Pinyin": "hainan", "City": [{ "Id": 460100, "Name": "海口", "FirstCharacter": "H", "Pinyin": "haikou", "OldCityId": 655 }] }, { "Id": 130000, "Name": "河北", "FirstCharacter": "H", "Pinyin": "hebei", "City": [{ "Id": 130100, "Name": "石家庄", "FirstCharacter": "S", "Pinyin": "shijiazhuang", "OldCityId": 119 }, { "Id": 130200, "Name": "唐山", "FirstCharacter": "T", "Pinyin": "tangshan", "OldCityId": 121 }, { "Id": 130300, "Name": "秦皇岛", "FirstCharacter": "Q", "Pinyin": "qinhuangdao", "OldCityId": 123 }, { "Id": 130400, "Name": "邯郸", "FirstCharacter": "H", "Pinyin": "handan", "OldCityId": 120 }, { "Id": 130500, "Name": "邢台", "FirstCharacter": "X", "Pinyin": "xingtai", "OldCityId": 124 }, { "Id": 130600, "Name": "保定", "FirstCharacter": "B", "Pinyin": "baoding", "OldCityId": 122 }, { "Id": 130700, "Name": "张家口", "FirstCharacter": "Z", "Pinyin": "zhangjiakou", "OldCityId": 125 }, { "Id": 130800, "Name": "承德", "FirstCharacter": "C", "Pinyin": "chengde", "OldCityId": 126 }, { "Id": 130900, "Name": "沧州", "FirstCharacter": "C", "Pinyin": "cangzhou", "OldCityId": 127 }, { "Id": 131000, "Name": "廊坊", "FirstCharacter": "L", "Pinyin": "langfang", "OldCityId": 128 }, { "Id": 131100, "Name": "衡水", "FirstCharacter": "H", "Pinyin": "hengshui", "OldCityId": 129 }] }, { "Id": 230000, "Name": "黑龙江", "FirstCharacter": "H", "Pinyin": "heilongjiang", "City": [{ "Id": 230100, "Name": "哈尔滨", "FirstCharacter": "H", "Pinyin": "haerbin", "OldCityId": 153 }, { "Id": 230200, "Name": "齐齐哈尔", "FirstCharacter": "Q", "Pinyin": "qiqihaer", "OldCityId": 154 }, { "Id": 230600, "Name": "大庆", "FirstCharacter": "D", "Pinyin": "daqing", "OldCityId": 156 }] }, { "Id": 410000, "Name": "河南", "FirstCharacter": "H", "Pinyin": "henan", "City": [{ "Id": 410100, "Name": "郑州", "FirstCharacter": "Z", "Pinyin": "zhengzhou", "OldCityId": 183 }, { "Id": 410200, "Name": "开封", "FirstCharacter": "K", "Pinyin": "kaifeng", "OldCityId": 184 }, { "Id": 410300, "Name": "洛阳", "FirstCharacter": "L", "Pinyin": "luoyang", "OldCityId": 185 }, { "Id": 410400, "Name": "平顶山", "FirstCharacter": "P", "Pinyin": "pingdingshan", "OldCityId": 186 }, { "Id": 410500, "Name": "安阳", "FirstCharacter": "A", "Pinyin": "anyang", "OldCityId": 187 }, { "Id": 410700, "Name": "新乡", "FirstCharacter": "X", "Pinyin": "xinxiang", "OldCityId": 189 }, { "Id": 410800, "Name": "焦作", "FirstCharacter": "J", "Pinyin": "jiaozuo", "OldCityId": 190 }, { "Id": 410900, "Name": "濮阳", "FirstCharacter": "P", "Pinyin": "puyang", "OldCityId": 191 }, { "Id": 411000, "Name": "许昌", "FirstCharacter": "X", "Pinyin": "xuchang", "OldCityId": 192 }, { "Id": 411300, "Name": "南阳", "FirstCharacter": "N", "Pinyin": "nanyang", "OldCityId": 195 }, { "Id": 411400, "Name": "商丘", "FirstCharacter": "S", "Pinyin": "shangqiu", "OldCityId": 196 }, { "Id": 411500, "Name": "信阳", "FirstCharacter": "X", "Pinyin": "xinyang", "OldCityId": 199 }, { "Id": 411600, "Name": "周口", "FirstCharacter": "Z", "Pinyin": "zhoukou", "OldCityId": 197 }, { "Id": 411700, "Name": "驻马店", "FirstCharacter": "Z", "Pinyin": "zhumadian", "OldCityId": 198 }] }, { "Id": 420000, "Name": "湖北", "FirstCharacter": "H", "Pinyin": "hubei", "City": [{ "Id": 420100, "Name": "武汉", "FirstCharacter": "W", "Pinyin": "wuhan", "OldCityId": 221 }, { "Id": 420300, "Name": "十堰", "FirstCharacter": "S", "Pinyin": "shiyan", "OldCityId": 222 }, { "Id": 420500, "Name": "宜昌", "FirstCharacter": "Y", "Pinyin": "yichang", "OldCityId": 231 }, { "Id": 420600, "Name": "襄阳", "FirstCharacter": "X", "Pinyin": "xiangyang", "OldCityId": 223 }] }, { "Id": 430000, "Name": "湖南", "FirstCharacter": "H", "Pinyin": "hunan", "City": [{ "Id": 430100, "Name": "长沙", "FirstCharacter": "C", "Pinyin": "changsha", "OldCityId": 257 }, { "Id": 430200, "Name": "株洲", "FirstCharacter": "Z", "Pinyin": "zhuzhou", "OldCityId": 262 }, { "Id": 430300, "Name": "湘潭", "FirstCharacter": "X", "Pinyin": "xiangtan", "OldCityId": 263 }, { "Id": 430400, "Name": "衡阳", "FirstCharacter": "H", "Pinyin": "hengyang", "OldCityId": 264 }] }, { "Id": 220000, "Name": "吉林", "FirstCharacter": "J", "Pinyin": "jilin", "City": [{ "Id": 220100, "Name": "长春", "FirstCharacter": "C", "Pinyin": "changchun", "OldCityId": 286 }, { "Id": 220200, "Name": "吉林", "FirstCharacter": "J", "Pinyin": "jilinshi", "OldCityId": 287 }] }, { "Id": 320000, "Name": "江苏", "FirstCharacter": "J", "Pinyin": "jiangsu", "City": [{ "Id": 320100, "Name": "南京", "FirstCharacter": "N", "Pinyin": "nanjing", "OldCityId": 335 }, { "Id": 320200, "Name": "无锡", "FirstCharacter": "W", "Pinyin": "wuxi", "OldCityId": 346 }, { "Id": 320300, "Name": "徐州", "FirstCharacter": "X", "Pinyin": "xuzhou", "OldCityId": 336 }, { "Id": 320400, "Name": "常州", "FirstCharacter": "C", "Pinyin": "changzhou", "OldCityId": 345 }, { "Id": 320500, "Name": "苏州", "FirstCharacter": "S", "Pinyin": "suzhou", "OldCityId": 347 }, { "Id": 320600, "Name": "南通", "FirstCharacter": "N", "Pinyin": "nantong", "OldCityId": 343 }, { "Id": 320700, "Name": "连云港", "FirstCharacter": "L", "Pinyin": "lianyungang", "OldCityId": 337 }, { "Id": 320800, "Name": "淮安", "FirstCharacter": "H", "Pinyin": "huaian", "OldCityId": 339 }, { "Id": 320900, "Name": "盐城", "FirstCharacter": "Y", "Pinyin": "yancheng", "OldCityId": 340 }, { "Id": 321000, "Name": "扬州", "FirstCharacter": "Y", "Pinyin": "yangzhou", "OldCityId": 341 }, { "Id": 321100, "Name": "镇江", "FirstCharacter": "Z", "Pinyin": "zhenjiang", "OldCityId": 344 }, { "Id": 321200, "Name": "泰州", "FirstCharacter": "T", "Pinyin": "tai_zhou", "OldCityId": 342 }, { "Id": 321300, "Name": "宿迁", "FirstCharacter": "S", "Pinyin": "suqian", "OldCityId": 338 }] }, { "Id": 360000, "Name": "江西", "FirstCharacter": "J", "Pinyin": "jiangxi", "City": [{ "Id": 360100, "Name": "南昌", "FirstCharacter": "N", "Pinyin": "nanchang", "OldCityId": 314 }, { "Id": 360400, "Name": "九江", "FirstCharacter": "J", "Pinyin": "jiujiang", "OldCityId": 315 }, { "Id": 360700, "Name": "赣州", "FirstCharacter": "G", "Pinyin": "ganzhou", "OldCityId": 320 }, { "Id": 361100, "Name": "上饶", "FirstCharacter": "S", "Pinyin": "shangrao", "OldCityId": 321 }] }, { "Id": 210000, "Name": "辽宁", "FirstCharacter": "L", "Pinyin": "liaoning", "City": [{ "Id": 210100, "Name": "沈阳", "FirstCharacter": "S", "Pinyin": "shenyang", "OldCityId": 375 }, { "Id": 210200, "Name": "大连", "FirstCharacter": "D", "Pinyin": "dalian", "OldCityId": 376 }, { "Id": 210300, "Name": "鞍山", "FirstCharacter": "A", "Pinyin": "anshan", "OldCityId": 383 }, { "Id": 210800, "Name": "营口", "FirstCharacter": "Y", "Pinyin": "yingkou", "OldCityId": 385 }] }, { "Id": 150000, "Name": "内蒙古", "FirstCharacter": "N", "Pinyin": "namenggu", "City": [{ "Id": 150100, "Name": "呼和浩特", "FirstCharacter": "H", "Pinyin": "huhehaote", "OldCityId": 595 }, { "Id": 150200, "Name": "包头", "FirstCharacter": "B", "Pinyin": "baotou", "OldCityId": 596 }, { "Id": 150400, "Name": "赤峰", "FirstCharacter": "C", "Pinyin": "chifeng", "OldCityId": 598 }, { "Id": 150500, "Name": "通辽", "FirstCharacter": "T", "Pinyin": "tongliao", "OldCityId": 600 }, { "Id": 150600, "Name": "鄂尔多斯", "FirstCharacter": "E", "Pinyin": "eerduosi", "OldCityId": 602 }] }, { "Id": 640000, "Name": "宁夏", "FirstCharacter": "N", "Pinyin": "ningxia", "City": [{ "Id": 640100, "Name": "银川", "FirstCharacter": "Y", "Pinyin": "yinchuan", "OldCityId": 615 }] }, { "Id": 630000, "Name": "青海", "FirstCharacter": "Q", "Pinyin": "qinghai", "City": [{ "Id": 630100, "Name": "西宁", "FirstCharacter": "X", "Pinyin": "xining", "OldCityId": 571 }] }, { "Id": 610000, "Name": "陕西", "FirstCharacter": "S", "Pinyin": "shan_xi", "City": [{ "Id": 610100, "Name": "西安", "FirstCharacter": "X", "Pinyin": "xian", "OldCityId": 454 }, { "Id": 610800, "Name": "榆林", "FirstCharacter": "Y", "Pinyin": "yulin", "OldCityId": 461 }] }, { "Id": 370000, "Name": "山东", "FirstCharacter": "S", "Pinyin": "shandong", "City": [{ "Id": 370100, "Name": "济南", "FirstCharacter": "J", "Pinyin": "jinan", "OldCityId": 406 }, { "Id": 370200, "Name": "青岛", "FirstCharacter": "Q", "Pinyin": "qingdao", "OldCityId": 407 }, { "Id": 370300, "Name": "淄博", "FirstCharacter": "Z", "Pinyin": "zibo", "OldCityId": 411 }, { "Id": 370400, "Name": "枣庄", "FirstCharacter": "Z", "Pinyin": "zaozhuang", "OldCityId": 417 }, { "Id": 370500, "Name": "东营", "FirstCharacter": "D", "Pinyin": "dongying", "OldCityId": 410 }, { "Id": 370600, "Name": "烟台", "FirstCharacter": "Y", "Pinyin": "yantai", "OldCityId": 413 }, { "Id": 370700, "Name": "潍坊", "FirstCharacter": "W", "Pinyin": "weifang", "OldCityId": 412 }, { "Id": 370800, "Name": "济宁", "FirstCharacter": "J", "Pinyin": "jining", "OldCityId": 418 }, { "Id": 370900, "Name": "泰安", "FirstCharacter": "T", "Pinyin": "taian", "OldCityId": 419 }, { "Id": 371000, "Name": "威海", "FirstCharacter": "W", "Pinyin": "weihai", "OldCityId": 414 }, { "Id": 371100, "Name": "日照", "FirstCharacter": "R", "Pinyin": "rizhao", "OldCityId": 415 }, { "Id": 371200, "Name": "莱芜", "FirstCharacter": "L", "Pinyin": "laiwu", "OldCityId": 420 }, { "Id": 371300, "Name": "临沂", "FirstCharacter": "L", "Pinyin": "linyi", "OldCityId": 416 }, { "Id": 371400, "Name": "德州", "FirstCharacter": "D", "Pinyin": "dezhou", "OldCityId": 409 }, { "Id": 371500, "Name": "聊城", "FirstCharacter": "L", "Pinyin": "liaocheng", "OldCityId": 408 }, { "Id": 371600, "Name": "滨州", "FirstCharacter": "B", "Pinyin": "binzhou", "OldCityId": 421 }, { "Id": 371700, "Name": "菏泽", "FirstCharacter": "H", "Pinyin": "heze", "OldCityId": 422 }] }, { "Id": 310000, "Name": "上海", "FirstCharacter": "S", "Pinyin": "shanghai", "City": [{ "Id": 310100, "Name": "上海", "FirstCharacter": "S", "Pinyin": "shanghai", "OldCityId": 649 }] }, { "Id": 140000, "Name": "山西", "FirstCharacter": "S", "Pinyin": "shanxi", "City": [{ "Id": 140100, "Name": "太原", "FirstCharacter": "T", "Pinyin": "taiyuan", "OldCityId": 467 }, { "Id": 140200, "Name": "大同", "FirstCharacter": "D", "Pinyin": "datong", "OldCityId": 468 }, { "Id": 140400, "Name": "长治", "FirstCharacter": "C", "Pinyin": "changzhi", "OldCityId": 471 }, { "Id": 140500, "Name": "晋城", "FirstCharacter": "J", "Pinyin": "jincheng", "OldCityId": 472 }, { "Id": 140700, "Name": "晋中", "FirstCharacter": "J", "Pinyin": "jinzhong", "OldCityId": 475 }, { "Id": 140800, "Name": "运城", "FirstCharacter": "Y", "Pinyin": "yuncheng", "OldCityId": 477 }, { "Id": 141000, "Name": "临汾", "FirstCharacter": "L", "Pinyin": "linfen", "OldCityId": 476 }] }, { "Id": 510000, "Name": "四川", "FirstCharacter": "S", "Pinyin": "sichuan", "City": [{ "Id": 510100, "Name": "成都", "FirstCharacter": "C", "Pinyin": "chengdu", "OldCityId": 489 }, { "Id": 510600, "Name": "德阳", "FirstCharacter": "D", "Pinyin": "deyang", "OldCityId": 492 }, { "Id": 510700, "Name": "绵阳", "FirstCharacter": "M", "Pinyin": "mianyang", "OldCityId": 491 }, { "Id": 511300, "Name": "南充", "FirstCharacter": "N", "Pinyin": "nanchong", "OldCityId": 493 }] }, { "Id": 120000, "Name": "天津", "FirstCharacter": "T", "Pinyin": "tianjin", "City": [{ "Id": 120100, "Name": "天津", "FirstCharacter": "T", "Pinyin": "tianjin", "OldCityId": 647 }] }, { "Id": 650000, "Name": "新疆", "FirstCharacter": "X", "Pinyin": "xinjiang", "City": [{ "Id": 650100, "Name": "乌鲁木齐", "FirstCharacter": "W", "Pinyin": "wulumuqi", "OldCityId": 624 }] }, { "Id": 530000, "Name": "云南", "FirstCharacter": "Y", "Pinyin": "yunnan", "City": [{ "Id": 530100, "Name": "昆明", "FirstCharacter": "K", "Pinyin": "kunming", "OldCityId": 521 }, { "Id": 530300, "Name": "曲靖", "FirstCharacter": "Q", "Pinyin": "qujing", "OldCityId": 522 }] }, { "Id": 330000, "Name": "浙江", "FirstCharacter": "Z", "Pinyin": "zhejiang", "City": [{ "Id": 330100, "Name": "杭州", "FirstCharacter": "H", "Pinyin": "hangzhou", "OldCityId": 538 }, { "Id": 330200, "Name": "宁波", "FirstCharacter": "N", "Pinyin": "ningbo", "OldCityId": 539 }, { "Id": 330300, "Name": "温州", "FirstCharacter": "W", "Pinyin": "wenzhou", "OldCityId": 547 }, { "Id": 330400, "Name": "嘉兴", "FirstCharacter": "J", "Pinyin": "jiaxing", "OldCityId": 541 }, { "Id": 330500, "Name": "湖州", "FirstCharacter": "H", "Pinyin": "huzhou", "OldCityId": 540 }, { "Id": 330600, "Name": "绍兴", "FirstCharacter": "S", "Pinyin": "shaoxing", "OldCityId": 543 }, { "Id": 330700, "Name": "金华", "FirstCharacter": "J", "Pinyin": "jinhua", "OldCityId": 545 }, { "Id": 330800, "Name": "衢州", "FirstCharacter": "Q", "Pinyin": "quzhou", "OldCityId": 544 }, { "Id": 331000, "Name": "台州", "FirstCharacter": "T", "Pinyin": "taizhou", "OldCityId": 546 }] }]
    `
)

type QcArea struct {
	Name string `json:"name"`
	City []QcCity
}

type QcCity struct {
	Name string `json:"name"`
	Pinyin string `json:"pinyin"`
}

type QcCar struct {
	CityName string
	Title string
	Price float64
	Kilometer float64
	Year int
}

func GetCitys() []QcCity {
	var areas []QcArea
	var citys []QcCity

	json.Unmarshal([]byte(areaJson), &areas)

	for _, area := range areas {
		for _, city := range area.City {
			citys = append(citys, city)
		}
	}

	return citys
}

func GetCityName(doc *goquery.Document) string {
	return doc.Find(".citycont .fn-left").Text()
}

func GetNextPageUrl(doc *goquery.Document) (val string, exists bool) {
	return doc.Find(".page .page-item-next").Attr("href")
}

func GetCurrentPage(doc *goquery.Document) (page int) {
	pageS := doc.Find(".page .current").Text()

	if pageS != "" {
		var err error
		page, err = strconv.Atoi(pageS)
		if err != nil {
			log.Printf("spiders.GetCurrentPage err: %v", err)
		}
	}

	return page
}

func GetCars(doc *goquery.Document) (cars []QcCar) {
	cityName := GetCityName(doc)
	doc.Find(".piclist ul li:not(.line)").Each(func(i int, selection *goquery.Selection) {
		title := selection.Find(".title a").Text()
		price := selection.Find(".detail .detail-r").Find(".colf8").Text()
		kilometer := selection.Find(".detail .detail-l").Find("p").Eq(0).Text()
		year := selection.Find(".detail .detail-l").Find("p").Eq(1).Text()

		// 数据处理
		kilometer = strings.Join(compileNumber.FindAllString(kilometer, -1), "")
		year = strings.Join(compileNumber.FindAllString(strings.TrimSpace(year), -1), "")
		priceS, _ := strconv.ParseFloat(price, 64)
		kilometerS, _ := strconv.ParseFloat(kilometer, 64)
		yearS, _ := strconv.Atoi(year)

		cars = append(cars, QcCar{
			CityName: cityName,
			Title: title,
			Price: priceS,
			Kilometer: kilometerS,
			Year: yearS,
		})
	})

	return cars
}