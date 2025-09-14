package devices

import (
	"fmt"
	"testing"
)

//
//import (
//	"encoding/binary"
//	"fmt"
//	"github.com/quant1x/gox/api"
//	"testing"
//)
//
//func Test_asmCpuid(t *testing.T) {
//	cpuId, err := CpuPhysicalID()
//	fmt.Println(api.String2Bytes(cpuId), err)
//	a, b, c, d := asmCpuid(3)
//	fmt.Printf("%08x %08x %08x %08x\n", a, b, c, d)
//	a, b, c, d = asmCpuidex(1, 1)
//	fmt.Printf("%08x %08x %08x %08x\n", a, b, c, d)
//	data := []byte{}
//	data = binary.LittleEndian.AppendUint32(data, d)
//	data = binary.LittleEndian.AppendUint32(data, a)
//	fmt.Printf("%s\n", api.Bytes2String(data))
//}

func TestCpuPhysicalID(t *testing.T) {
	cpuId, err := CpuPhysicalID()
	fmt.Println(cpuId, err)
}
