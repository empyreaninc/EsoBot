package utils

import (
	"github.com/gookit/color"
	"time"
)

func EndNeutralPrefix(tasknum string) string {
	time := time.Now()
	timefmt := time.Format("15:04:05.00")
	magenta := color.Magenta.Render
	ltgrey := color.FgDarkGray.Render

	return ltgrey("[esotera | "+timefmt+"] ") + magenta("END - Task "+tasknum+": ")

}

func EndSuccessPrefix(tasknum string) string {
	time := time.Now()
	timefmt := time.Format("15:04:05.00")
	magenta := color.Magenta.Render
	ltgrey := color.FgDarkGray.Render
	green := color.Green.Render

	return ltgrey("[esotera | "+timefmt+"] ") + magenta("END - Task "+tasknum+": ") + green("✔ ")

}

func EndFailurePrefix(tasknum string) string {
	time := time.Now()
	timefmt := time.Format("15:04:05.00")
	magenta := color.Magenta.Render
	ltgrey := color.FgDarkGray.Render
	red := color.Red.Render

	return ltgrey("[esotera | "+timefmt+"] ") + magenta("END - Task "+tasknum+": ") + red("! ")

}
