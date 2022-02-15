package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	resp, err := client.Get("http://127.0.0.1:9933/hello")
	if err != nil {
		log.Fatal("请求出错了...")
		return
	}
  content, err := ioutil.ReadAll(resp.Body) //ReadAll返回[]byte 数组，后面需要string()强制转换
  /*
  func ReadAll(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}

  */
	if err != nil {
		log.Fatal("响应数据解析错误...")
		return
	}
	fmt.Println(string(content))

}
