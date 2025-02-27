## Run all tests
```sh
go  test ./...
```

## Run a specific test
```sh
go test -v -run TestCalculateValues
```

## Rerun all test test even the cached one
```sh
go test ./... -v -count=1
```
