package color

// 粗体
func Bold(text string) string {
	// return value
	return "\x1b[1m" + text + "\x1b[0m"
}

// 变淡
func Faint(text string) string {
	// return value
	return "\x1b[2m" + text + "\x1b[0m"
}

// 青色
func Cyan(text string) string {
	// return value
	return "\x1b[36m" + text + "\x1b[0m"
}

// 绿色
func Green(text string) string {
	// return value
	return "\x1b[32m" + text + "\x1b[0m"
}

// 红色
func Red(text string) string {
	// return value
	return "\x1b[31m" + text + "\x1b[0m"
}

// 黄色
func Yellow(text string) string {
	// return value
	return "\x1b[33m" + text + "\x1b[0m"
}
