package dependency_consumer

import (
	"fmt"
	"github.com/my-org-org/go-repo-two/internal"
	"github.com/my-org-org/go-repo-one/dependency"
)

func PrintDependencyMessage() {
	fmt.Println("go-repo-two. Release 6 NEW NAME: Use 1st repo NEW NAME ")
	dependency.PrintDependencyMessage()
	fmt.Println("go-repo-two. Release 6 !!!!!!!!!!!! : Dependency Consumer")
    internal.InternalFunction()
}
