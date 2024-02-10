package num

import (
	"fmt"
	"gitee.com/quant1x/gox/num/internal/functions"
	"golang.org/x/sys/cpu"
)

// Misc

// SetAcceleration toggles simd acceleration. Not thread safe.
func SetAcceleration(enabled bool) error {
	if enabled && !(cpu.X86.HasAVX2 && cpu.X86.HasFMA) {
		functions.UseAVX2 = false
		return fmt.Errorf("acceleration not supported on this platform")

	}
	functions.UseAVX2 = enabled
	return nil
}

func init() {
	_ = SetAcceleration(true)
}
