package util

import (
	"encoding/binary"
	"errors"
	"log"
	"net"
	grpc_proto "test/bk_proto"
	"time"

	"github.com/gogo/protobuf/proto"
)

var (
	G_transfer *Transfer
)

func InitTransfer() {
	var (
		pTCPAddr *net.TCPAddr
		conn     net.Conn
		err      error
	)
	if pTCPAddr, err = net.ResolveTCPAddr("tcp", "192.168.3.123:3181"); err != nil {
		log.Fatalln("连接tcp失败", err)
		return
	}

	if conn, err = net.DialTCP("tcp", nil, pTCPAddr); err != nil {
		log.Fatalln("连接tcp失败", err)
		return
	}

	go func() {
		//1024 * 2
		for true {
			buf := make([]byte, 1024*2)
			_, err = conn.Read(buf[:1024])
			if len(buf) > 0 {
				continue
			}
			log.Println("receive:", len(buf))
			response := &grpc_proto.Response{}
			if err = proto.Unmarshal(buf[:], response); err != nil {
				return
			}
			log.Println("txnId:", response.Header.TxnId)
		}
	}()
	time.Sleep(1 * time.Second)

	// 定义 Transfer 指针变量
	G_transfer = &Transfer{
		Conn: conn,
	}
}

// 声明 Transfer 结构体
type Transfer struct {
	Conn net.Conn       // 连接
	Buf  [1024 * 2]byte // 传输时，使用的缓冲
}

// // 获取并解析服务器的消息
// func (transfer *Transfer) ReadResponse(response *grpc_proto.ProtocolMessage) (err error) {
// 	_, err = transfer.Conn.Read(transfer.Buf[:4])
// 	if err != nil {
// 		return
// 	}

// 	// 根据 buf[:4] 转成一个 uint32 类型
// 	var pkgLen uint32
// 	pkgLen = binary.BigEndian.Uint32(transfer.Buf[:4])
// 	//根据pkglen 读取消息内容
// 	n, err := transfer.Conn.Read(transfer.Buf[:pkgLen])
// 	if n != int(pkgLen) || err != nil {
// 		return
// 	}

// 	if err = proto.Unmarshal(transfer.Buf[:pkgLen], response); err != nil {
// 		return
// 	}
// 	returgrpc_proto

// 发送消息到服务器
func (transfer *Transfer) SendMsg(action *grpc_proto.Request) (err error) {
	var (
		sendBytes []byte
		readLen   int
	)
	//sendBytes, ints := action.Descriptor()
	if sendBytes, err = proto.Marshal(action); err != nil {
		return err
	}

	pkgLen := uint32(len(sendBytes))
	log.Println(pkgLen)
	var buf [32]byte
	binary.BigEndian.PutUint32(buf[:pkgLen], pkgLen)

	if readLen, err = transfer.Conn.Write(buf[:4]); readLen != 4 && err != nil {
		if readLen == 0 {
			return errors.New("发送数据长度发生异常，长度为0")
		}
		return err
	}
	// 发送消息
	if readLen, err = transfer.Conn.Write(sendBytes); err != nil {
		if readLen == 0 {
			return errors.New("检查到服务器关闭，客户端也关闭")
		}
		return err
	}
	return
}
