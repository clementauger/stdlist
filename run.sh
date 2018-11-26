
mkdir -p data/

file="data/${TRAVIS_GO_VERSION}.txt"
if [ ! -f ${file} ]; then
  go run cmd/list/main.go > data/${TRAVIS_GO_VERSION}.txt
fi

# go run cmd/merge/main.go > data.go
