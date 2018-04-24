mkdir client/src/protos && echo "creating proto directory"
mkdir server/protos && echo "creating proto directory"

echo "generating proto files"
protoc \
  --plugin=protoc-gen-ts=client/node_modules/.bin/protoc-gen-ts \
  --js_out=import_style=commonjs,binary:client/src/protos \
  --ts_out=service=true:client/src/protos \
  --go_out=plugins=grpc,import_path=protos:server/protos \
  -I protos/ protos/*.proto && echo "generated proto files"
