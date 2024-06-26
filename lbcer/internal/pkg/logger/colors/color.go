package colors

// Define constants for different colors using iota
const (
	Magenta Color = iota
	Blue
	Yellow
	Red
	Cyan
	White
)

// Color тип для цвета
type Color int

// Colorize функция для окрашивания строки в нужный цвет
func Colorize(s string, color Color) string {
	switch color {
	case Magenta:
		return "\033[35m" + s + "\033[0m"
	case Blue:
		return "\033[34m" + s + "\033[0m"
	case Yellow:
		return "\033[33m" + s + "\033[0m"
	case Red:
		return "\033[31m" + s + "\033[0m"
	case Cyan:
		return "\033[36m" + s + "\033[0m"
	case White:
		return "\033[37m" + s + "\033[0m"
	default:
		return s
	}
}
