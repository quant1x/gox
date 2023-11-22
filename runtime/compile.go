package runtime

// Debug go编译tag debug状态
//
//	运行和编译增加 -tags="dev"
func Debug() bool {
	return tagDebug
}
