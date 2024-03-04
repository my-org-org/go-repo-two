package main

import (
	"fmt"
	"github.com/my-org-for-test/go-repo-one/dependency"
	"github.com/my-org-for-test/go-repo-two/dependency_consumer"
)

func main() {
    fmt.Println("========> go-repo-tro: Release 4")

	fmt.Println("go-repo-two: Dependency on go-repo-one:")
	dependency.PrintDependencyMessage()

    fmt.Println("go-repo-two: Self dependency")
	dependency_consumer.PrintDependencyMessage()
}
