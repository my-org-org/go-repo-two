package dependency_consumer

import (
	"fmt"
	"github.com/my-org-for-test/go-repo-two/internal"
	"github.com/my-org-for-test/go-repo-one/dependency"
)

func PrintDependencyMessage() {
	fmt.Println("go-repo-two. Release 3 $$$$$ : Use 1st repo")
	dependency.PrintDependencyMessage()
	fmt.Println("go-repo-two. Release 3 !!!!!!!!!!!! : Dependency Consumer")
    internal.InternalFunction()
}
