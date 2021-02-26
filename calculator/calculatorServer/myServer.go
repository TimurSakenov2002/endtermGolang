package main

import (
	"context"
	"endterm/calculator/calculatorPB"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"time"
)

type Server struct {
	calcSide.UnimplementedCalculatorServiceServer
}

func (s *Server) PrimeNumberDecomposition(req *calcSide.PrimeNumberRequest, stream calcSide.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("Trying to call prime number decomposition \n")
	prime_number := req.sumNumber()
	prime_numbers_factors := takePrimeNum(prime_number)

	for i := 0; i < len(prime_numbers_factors); i++ {
		res := &calcSide.PrimeNumberResponse{Result: prime_numbers_factors[i]}
		if err := stream.Send(res); err != nil {
			log.Fatalf("Error! %v", err.Error())
		}
		time.Sleep(time.Second)
	}
	return nil
}
func takePrimeNum(num int32) []int32 {
	res_arr := []int32{}
	for {
		res_arr = append(res_arr, 2)
		num /= 2
		if num%2 != 0 {
			break
		}
	}
	var i int32 = 0
	for i = 3; i <= num*num; i += 2 {
		for {
			res_arr = append(res_arr, 3)
			num /= i
			if num%i != 0 {
				break
			}
		}
	}
	if num > 2 {
		res_arr = append(res_arr, num)
	}

	return res_arr
}
func (s *Server) ComputeAverage(stream calcSide.CalculatorService_ComputeAverageServer) error {
	fmt.Printf("Trying to call Compute Average\n")
	var sum int32 = 0
	var quantity int32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			var response = &calcSide.AvgNumberResponse{Res: float32(sum / quantity)}
			return stream.SendAndClose(response)
		}
		if err != nil {
			log.Fatalf("Error! %v", err)
		}
		sum += req.sumNum()
		quantity++
	}
}
func main() {
	l, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		log.Fatalf("Error! %v", err)
	}

	s := grpc.NewServer()
	calcSide.RegisterCalculatorServiceServer(s, &Server{})
	log.Println("Running: 8081")
	if err := s.Serve(l); err != nil {
		log.Fatalf("Error!%v", err)
	}
}
