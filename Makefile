.PHONY: proto
proto:
	rm -f pb/*.pb.go && protoc -Iproto \
      --go_out=pb \
      --go_opt=module=github.com/shibafu528/utissue/pb \
      --go-grpc_out=pb \
      --go-grpc_opt=module=github.com/shibafu528/utissue/pb \
      proto/shibafu528/utissue/*.proto
