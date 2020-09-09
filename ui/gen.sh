PROTOC_GEN_TS_PATH="./node_modules/.bin/protoc-gen-ts"
 
# Directory to write generated code to (.js and .d.ts files) 
OUT_DIR="./src/pb"
 
protoc \
    -I=./proto \
    -I=/usr/local/include/google/protobuf \
    --plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
    --js_out="import_style=commonjs,binary:${OUT_DIR}" \
    --ts_out="${OUT_DIR}" \
    proto/*.proto