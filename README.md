# stdlist

helper package that lists all golang stdlib packages.

# release

To release this package, update the travis file to add the related go version.

RUn git push to trigger a travis build, it will generate the new data files.

Pull the updated repository.

run `go run cmd/merge/main.go > raw.go`, then commit, push and tag.
