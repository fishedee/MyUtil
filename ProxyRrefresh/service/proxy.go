package service

import (
	"sync"
	"time"
    "github.com/franela/goreq"
	"github.com/PuerkitoBio/goquery"
)

type Proxy struct{
	proxyUrls map[string] bool
	proxyRwMutex sync.RWMutex
}

func init(){

}

func (self *Proxy)Get(number int)([]string){
	self.proxyRwMutex.RLock()
	result := []string{}
	for key,_ := range self.proxyUrls{
		result = append( result ,key )
		if len(result) == number{
			break
		}
	}
	self.proxyRwMutex.RUnlock()
	return result
}

func (self *Proxy)GetOnce()(string,error){
	resp, err := goreq.Request{
	    Method: "GET",
	    Uri: "http://vxer.daili666.com/ip/?tid=556861362403782&num=1&operator=1&delay=1&category=2&sortby=time&foreign=none&filter=on",
		Timeout: 5 * time.Second,
	}.Do()

	if err != nil {
		return "",err
	}

	return resp.Body.ToString()
}

func (self *Proxy)Add(url string){
	self.proxyRwMutex.Lock()
	self.proxyUrls[url] = true
	self.proxyRwMutex.Unlock()
}

func (self *Proxy)Remove(url string){
	self.proxyRwMutex.Lock()
	delete(self.proxyUrls,url)
	self.proxyRwMutex.Unlock()
}

func (self *Proxy)Clear(){
	self.proxyRwMutex.Lock()
	self.proxyUrls = make(map[string] bool)
	self.proxyRwMutex.Unlock()
}

func (self *Proxy)findSingle()error{
	result := []string{}

	doc, err := goquery.NewDocument("http://127.0.0.1:10003/http://pachong.org/")
	if err != nil {
    	return err
    }

    doc.Find("table.tb tbody tr").Each(func(i int, s *goquery.Selection) {
        ip := s.Find("td:nth-child(2)").Text()
        s.Find("td:nth-child(3) script").Remove()
        port := s.Find("td:nth-child(3)").Text()

        if ip == "" || port == ""{
        	return
        }
        result = append( result , ip + ":" +port )
	})

	if len(result) == 0 {
		return nil
	}

   	self.Clear()
   	for _,url := range result{
   		self.Add( url )
   	}

   	return nil
}

func (self *Proxy)Run(){
	for{
		self.findSingle()
		time.Sleep(10 * time.Second) 
	}
}