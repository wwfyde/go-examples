package simple_text_extractor

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Content is used for storage bufio
var Content string

// Extract extract
func Extract(file bufio.Reader) (res []string) {
	// 逐行读取并按regexp匹配行, 将内容元素存储到[]string中
	data, _ := os.ReadFile("/Users/wwfyde/Temp/17-58-42-04--{fbc89d2c-424a-456d-8f37-a16b158dec6b}.json")
	fmt.Println(data)
	f, _ := os.Open("/Users/wwfyde/Temp/17-58-42-04--{fbc89d2c-424a-456d-8f37-a16b158dec6b}.json")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// TODO add content to res
		if ok, _ := regexp.MatchString(`content:(.*?),`, strings.TrimSpace(line)); ok {
			res = append(res, line)
		}

	}
	return
}
