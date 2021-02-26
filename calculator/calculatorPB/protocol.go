syntax = "proto3";
package calculator;
option go_package="endterm\\calculator\\calcPB;calcPB";

message PrimeNumberRequest{
int32 number = 1;
}

message PrimeNumberResponse{
int32 result = 1;
}

message AvgNumberRequest{
int32 num = 1;
}

message AvgNumberResponse{
float res = 1;
}

service CalculatorService{
rpc PrimeNumberDecomposition(PrimeNumberRequest) returns (stream PrimeNumberResponse);
rpc ComputeAverage(stream AvgNumberRequest) returns (AvgNumberResponse);
}