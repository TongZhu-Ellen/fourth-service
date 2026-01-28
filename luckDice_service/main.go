package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	pb "luckDice_service/pb" // proto 生成的 Go 包名，不是硬盘地址！！！
	"sync"

	"google.golang.org/grpc"
)

// LuckDiceServer 是 proto 的 gRPC 服务实现
// 下面的字段和方法名必须对应 proto 或 gRPC 生成的接口


var (

	luckMap map[string]bool = make(map[string]bool)              
	mu      sync.Mutex  
)

type LuckDiceServer struct {
	pb.UnimplementedLuckDiceServiceServer // proto 生成的默认结构体，必须保留
	                   
}



// gRPC 方法实现，完全对应 proto
func (sp *LuckDiceServer) IndemPay(ctx context.Context, rp *pb.RequestMsg) (*pb.ResponseMsg, error) {
	mu.Lock()               
	defer mu.Unlock()       

	lucky, exists := luckMap[rp.OrderID] 
	if !exists {
		lucky = rand.Intn(2) == 0          
		luckMap[rp.OrderID] = lucky      
	}


	return &pb.ResponseMsg{
		Lucky: lucky,
	}, nil 
}

func main() {
	listen, err := net.Listen("tcp", ":50051") // listen GPT 自己起名
	if err != nil { panic(err) }

	s := grpc.NewServer() 
	pb.RegisterLuckDiceServiceServer(s, &LuckDiceServer{}) // proto 注册方法必须用 proto 名



	fmt.Println("LuckDice gRPC server running on :50051")
	err = s.Serve(listen)
	if err != nil { panic(err) }

	
}
