package service

import (
	"time"
	"errors"
	"math/rand"
	"github.com/franela/goreq"
	log "github.com/cihub/seelog"
)

type Refresh struct{
}

func (self *Refresh) catch(proxy* Proxy)error{
	proxyUrl,err := proxy.GetOnce();
	if err != nil{
		return err
	}

	resp, err := goreq.Request{
	    Method: "GET",
	    Proxy: "http://"+proxyUrl,
	    Uri: "http://www.baidu.com",
	    Timeout: 5 * time.Second,
	}.Do()
	if( err != nil ){
		return err
	}
	if( resp.StatusCode != 200 ){
		return errors.New("no 200!")
	}
	return nil	
}

func (self *Refresh)Run(proxy* Proxy){
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		//wait
		timeout := random.Intn(10*1000)+5000
		time.Sleep( time.Duration(timeout) * time.Millisecond) 

		//go
		err := self.catch(proxy)
		log.Info("timeout:",timeout,",error:",err)
	}
}