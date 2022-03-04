package randoms

import (
	"fmt"
	"testing"
)

func TestGetRandomNumString(t *testing.T)  {

	fmt.Println(GetRandomNumString(3,"ABC"))
	fmt.Println(GetRandomNumString(3,"EFG"))
}
