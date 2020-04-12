CONFIG_PATH=${HOME}/dev/.chronicle/

.PHONY: init
init:
	mkdir -p ${CONFIG_PATH}

.PHONY: test
test:
	go test -race ./...

.PHONY: compile
compile:
	protoc api/v1/*.proto \
	--gogo_out=\
	Mgogoproto/gogo.proto=github.com/gogo/protobuf/proto,plugins=grpc:. \
	--proto_path=\
	$$(go list -f '{{ .Dir }}' -m github.com/gogo/protobuf) \
	--proto_path=.

.PHONY: gencert
gencert:
    # CA
	cfssl gencert \
		-initca conf/ca-csr.json | cfssljson -bare ca

    # Server
	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=conf/ca-config.json \
		-profile=server \
		conf/server-csr.json | cfssljson -bare server

    # Client
	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=conf/ca-config.json \
		-profile=client \
		conf/client-csr.json | cfssljson -bare client

	# Root Client
	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=conf/ca-config.json \
		-profile=client \
		-cn="root" \
		conf/client-csr.json | cfssljson -bare root-client

	mv *.pem *.csr ${CONFIG_PATH}