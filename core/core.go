package core

//go:generate mockgen -package core -destination mocks_test.go github.com/petrkotek/go-ash/core EntityListener
//go:generate sed -i .bak "s/core\\.//g; s/core\\ .*//g" mocks_test.go
// above is an ugly hack to have mocks in the same package

/*

Mocks:
To generate mocks for (some) core interfaces, run `go generate ./...`

*/
