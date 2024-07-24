package etree

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// 定义结构体映射xml结构
type SConfig2 struct {
	XMLName      xml.Name   `xml:"config"`     // 指定最外层的标签为config
	SmtpServer   string     `xml:"smtpServer"` // 读取smtpServer配置项，并将结果保存到SmtpServer变量中
	SmtpPort     int        `xml:"smtpPort"`
	Sender       string     `xml:"sender"`
	SenderPasswd string     `xml:"senderPasswd"`
	Receivers    SReceivers `xml:"receivers"` // 读取receivers标签下的内容，以结构方式获取
}

type SReceivers2 struct {
	Age    int      `xml:"age"`
	Flag   string   `xml:"flag,attr"` // 读取flag属性
	User   []string `xml:"user"`      // 读取user数组
	Script string   `xml:"script"`    // 读取 <![CDATA[ xxx ]]> 数据
}

func ParseXMLDemo2() {
	file, err := os.Open("testdata/servers.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()

	// 创建带缓存的 Reader
	reader := bufio.NewReader(file)

	decoder := xml.NewDecoder(reader)

	for t, err := decoder.Token(); err == nil || err == io.EOF; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			fmt.Println(name)
			if name == "config" {
				// 解析 config
				var sConfig = SConfig{}
				configErr := decoder.DecodeElement(&sConfig, &token)
				if configErr != nil {
					fmt.Println("解析错误：")
					fmt.Println(configErr)
				} else {
					fmt.Println(sConfig)
				}
				return
			}
		}
	}
}
