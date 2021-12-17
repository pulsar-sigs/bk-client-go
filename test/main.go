package main

import (
	"log"
	"net/http"
	"os"
	proto "test/bk_proto"
	"test/util"
	"time"

	"github.com/gin-gonic/gin"
)

// func init() {
// 	util.InitTransfer()
// 	send()
// }

func addEntry(ledgerId, entryId int64, masterKey, body []byte) {

	flag := proto.AddRequest_RECOVERY_ADD

	addRequest := &proto.AddRequest{
		LedgerId:  &ledgerId,
		EntryId:   &entryId,
		MasterKey: masterKey,
		Body:      body,
		Flag:      &flag,
	}
	log.Println(addRequest)
}

//获取ledger metadata
// 接口实现
func getLedgerMetadata() {

}

func getBookieInfo() {

}

func send() {

	var txnId uint64 = 1
	version := proto.ProtocolVersion_VERSION_THREE
	operation := proto.OperationType_ADD_ENTRY
	header := &proto.BKPacketHeader{
		Version:   &version,
		Operation: &operation,
		TxnId:     &txnId,
	}

	var ledgerId int64 = 10
	var entryId int64 = 2

	flag := proto.AddRequest_RECOVERY_ADD

	addRequest := &proto.AddRequest{
		LedgerId:  &ledgerId,
		EntryId:   &entryId,
		MasterKey: []byte("hello"),
		Body:      []byte("world"),
		Flag:      &flag,
	}
	

	message := &proto.Request{
		Header:     header,
		AddRequest: addRequest,
	}

	log.Println("begin send message")
	if err := util.G_transfer.SendMsg(message); err != nil {
		log.Println("send.failed!")
		return
	}

	// if err := util.G_transfer.ReadResponse(message); err != nil {
	// 	c.JSON(500, gin.H{
	// 		"err": err.Error(),
	// 	})
	// 	return
	// }
}

func main() {

	util.InitTransfer()
	send()
	if 2 > 1 {
		time.Sleep(5 * time.Second)
		os.Exit(-1)
	}

	router := gin.Default()
	// search 测试
	router.GET("/search", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "world",
		})
	})

	ReadTimeout := time.Duration(60) * time.Second
	WriteTimeout := time.Duration(60) * time.Second

	s := &http.Server{
		Addr:           ":8090",
		Handler:        router,
		ReadTimeout:    ReadTimeout,
		WriteTimeout:   WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
