package main

import (
	"fmt"
	"github.com/my-org-org/go-repo-one/dependency"
	"github.com/my-org-org/go-repo-two/dependency_consumer"
)

func main() {
    fmt.Println("========> go-repo-tro: Release 6")

	fmt.Println("go-repo-two: Dependency on go-repo-one:")
	dependency.PrintDependencyMessage()

    fmt.Println("go-repo-two: Self dependency")
	dependency_consumer.PrintDependencyMessage()
}
