package main

import (
	"fmt"
	cregex "github.com/mingrammer/commonregex"
)

func main() {
	text := `John, please get that article on www.linkedin.com to me by 5:00PM on Jan 9th 2012. 4:00 would be ideal, actually. If you have any questions, You can reach me at (519)-236-2723x341 or get in touch with my associate at harold.smith@gmail.com`
	dateList := cregex.Date(text)
	timeList := cregex.Time(text)
	linkList := cregex.Links(text)
	phoneList := cregex.Phones(text)
	emailList := cregex.Emails(text)

	fmt.Println("date list:", dateList)
	fmt.Println("time list:", timeList)
	fmt.Println("link list:", linkList)
	fmt.Println("phone list:", phoneList)
	fmt.Println("email list:", emailList)

	text = `mac address: ac:de:48:00:11:22, ip: 192.168.3.20, md5: fdbf72fdabb67ea6ef7ff5155a44def4`
	macList := cregex.MACAddresses(text)
	ipList := cregex.IPs(text)
	md5List := cregex.MD5Hexes(text)

	fmt.Println("mac list:", macList)
	fmt.Println("ip list:", ipList)
	fmt.Println("md5 list:", md5List)
}
