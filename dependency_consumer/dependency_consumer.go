package dependency_consumer

import (
	"fmt"
	"github.com/my-org-for-test/go-repo-two/internal"
)

func PrintDependencyMessage() {
	fmt.Println("go-repo-two. Release 2 !!!!!!!!!!!! : Dependency Consumer")
    internal.InternalFunction()
}
