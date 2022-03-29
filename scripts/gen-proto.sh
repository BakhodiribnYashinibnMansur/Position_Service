CURRENT_DIR=$1
for x in $(find ${CURRENT_DIR}/protos/* -type d); do
  sudo protoc -I=${x} -I=${CURRENT_DIR}/protos -I /usr/local/include --go_out=plugins=grpc:${CURRENT_DIR} ${x}/*.proto
done
# CURRENT_DIR=$1
# for x in $(find ${CURRENT_DIR}/protos/* -type d); do
#   protoc -I=${x} -I=${CURRENT_DIR}/protos -I /usr/local/include --go_out=plugins=grpc:${CURRENT_DIR} ${x}/*.proto
# done