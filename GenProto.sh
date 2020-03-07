protoc -I whinvitesvr/proto  --go_out=plugins=grpc:whinvitesvr/proto whinvitesvr.proto
echo "GenPb whinvitesvr Ok"
