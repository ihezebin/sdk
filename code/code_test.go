package code

import (
	"fmt"
	"testing"
)

func TestCode(t *testing.T) {
	fmt.Println(Default())
	fmt.Println(Blend(4))
	fmt.Println(Verify("rfbd67", "RFbD67"))
	fmt.Println(Verify("rfbd67", "RFbD67", true))
	fmt.Println(Verify("RFbD67", "RFbD67", true))
	fmt.Println(ZnCn())
	fmt.Println(NewZnCn(4))
	fmt.Println(Verify("生却块正步成", "生却块正步成"))
	fmt.Println(Verify("生却块正步啊", "生却块正步成"))
	fmt.Println(Custom("123abcDEF却块正步"))
	fmt.Println(NewCustom(4, "123abcDEF却块正步"))
	fmt.Println(Verify("块d步a1d", "块d步A1D"))
	fmt.Println(Verify("块d步a1d", "块d步A1D", true))
	fmt.Println(Verify("块d步A1D", "块d步A1D", true))
}
