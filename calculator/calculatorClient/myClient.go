package main

import (
	"context"
	"endterm/calculator/calculatorPB"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error! %v", err)
	}
	defer conn.Close()
	c := calcSide.NewCalculatorServiceClient(conn)
	pnd(c)
	compAvg(c)
}
func pnd(c calcSide.CalculatorServiceClient) {
	ctx := context.Background()
	request := &calcSide.PrimeNumberRequest{Number: 120}

	stream, err := c.pnd(ctx, request)
	if err != nil {
		log.Fatalf("Error prime number decomposition! %v", err)
	}
	defer stream.CloseSend()

	//the loop for checking
checking:
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break checking
			}
			log.Fatalf("Error prime number decomposition! %v", err)
		}
		log.Printf("Response from Prime Number decomposition: %v \n", res.getResultat())
	}

}
func compAvg(c calcSide.CalculatorServiceClient) {
	ctx := context.Background()
	stream, err := c.compAvg(ctx)
	numbers := []int32{2, 5, 7, 8, 9}
	if err != nil {
		log.Fatalf("Error compute average! %v", err)
	}
	for _, number := range numbers {
		stream.Send(&calcSide.AvgNumberRequest{Num: number})
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.closeAll()
	if err != nil {
		log.Fatalf("Error during the receive! %v", err)
	}
	fmt.Printf("Response from Compute Average: %v\n", res.resultat())
}
