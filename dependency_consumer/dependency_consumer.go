package dependency_consumer

import (
	"fmt"
	"github.com/my-org-org/go-repo-two/internal"
	"github.com/my-org-for-test/go-repo-one/dependency"
)

func PrintDependencyMessage() {
	fmt.Println("go-repo-two. Release 5 NEW NAME: Use 1st repo OLD NAME")
	dependency.PrintDependencyMessage()
	fmt.Println("go-repo-two. Release 4 !!!!!!!!!!!! : Dependency Consumer")
    internal.InternalFunction()
}
