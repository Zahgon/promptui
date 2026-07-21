//go:build !windows
// +build !windows

package promptui

var (
	IconInitial = Styler(FGBlue)("?")

	IconGood = Styler(FGGreen)("✔")

	IconWarn = Styler(FGYellow)("⚠")

	IconBad = Styler(FGRed)("✗")

	IconSelect = Styler(FGBold)("▸")
)
