package ansiparse

// Decorators is a mapping of ansi tags
var Decorators = map[string]string{
	"black":   "foregroundColorOpen",
	"red":     "foregroundColorOpen",
	"green":   "foregroundColorOpen",
	"yellow":  "foregroundColorOpen",
	"blue":    "foregroundColorOpen",
	"magenta": "foregroundColorOpen",
	"cyan":    "foregroundColorOpen",
	"white":   "foregroundColorOpen",

	"gray":          "foregroundColorOpen",
	"redBright":     "foregroundColorOpen",
	"greenBright":   "foregroundColorOpen",
	"yellowBright":  "foregroundColorOpen",
	"blueBright":    "foregroundColorOpen",
	"magentaBright": "foregroundColorOpen",
	"cyanBright":    "foregroundColorOpen",
	"whiteBright":   "foregroundColorOpen",

	"bgBlack":   "backgroundColorOpen",
	"bgRed":     "backgroundColorOpen",
	"bgGreen":   "backgroundColorOpen",
	"bgYellow":  "backgroundColorOpen",
	"bgBlue":    "backgroundColorOpen",
	"bgMagenta": "backgroundColorOpen",
	"bgCyan":    "backgroundColorOpen",
	"bgWhite":   "backgroundColorOpen",

	"bgGray":          "backgroundColorOpen",
	"bgRedBright":     "backgroundColorOpen",
	"bgGreenBright":   "backgroundColorOpen",
	"bgYellowBright":  "backgroundColorOpen",
	"bgBlueBright":    "backgroundColorOpen",
	"bgMagentaBright": "backgroundColorOpen",
	"bgCyanBright":    "backgroundColorOpen",
	"bgWhiteBright":   "backgroundColorOpen",

	"foregroundColorClose": "foregroundColorClose",
	"backgroundColorClose": "backgroundColorClose",

	"boldOpen":          "boldOpen",
	"dimOpen":           "dimOpen",
	"italicOpen":        "italicOpen",
	"underlineOpen":     "underlineOpen",
	"inverseOpen":       "inverseOpen",
	"hiddenOpen":        "hiddenOpen",
	"strikethroughOpen": "strikethroughOpen",

	"boldDimClose":       "boldDimClose",
	"italicClose":        "italicClose",
	"underlineClose":     "underlineClose",
	"inverseClose":       "inverseClose",
	"hiddenClose":        "hiddenClose",
	"strikethroughClose": "strikethroughClose",

	"reset": "reset",
}
