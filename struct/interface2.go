package main

import "fmt"

type alipayUser struct {
	phone   string
	qr_code string
	huabei  int
}

type wechatpayUser struct {
	phone         string
	qr_code       string
	weixin_change int
}

type cashUser struct {
	cash int
}

func (a *alipayUser) pay() {
	fmt.Println("支付宝支付成功...")
}

func (w *wechatpayUser) pay() {
	fmt.Println("微信支付成功...")
}
func (ca *cashUser) pay() {
	fmt.Println("现金支付成功...")
}

type consume interface {
	pay()
}

func main() {
	a := &alipayUser{
		phone:   "iphone/android",
		qr_code: "qr_code",
		huabei:  1000,
	}
	w := &wechatpayUser{
		phone:         "iphone/android",
		qr_code:       "qr_code",
		weixin_change: 1000,
	}

	ca := &cashUser{
		cash: 1000,
	}
	var c consume
	c = a
	c.pay()

	c = w
	c.pay()

	c = ca
	c.pay()

}

/*
支付宝支付成功...
微信支付成功...
现金支付成功...


**/
