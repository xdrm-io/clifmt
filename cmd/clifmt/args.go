package main

import (
	"git.xdrm.io/go/clifmt"
	"git.xdrm.io/go/clifmt/internal/color"
)

func help() {
	clifmt.Printf("**NAME**\n")
	clifmt.Printf("\tclifmt - **c***ommand-***l***ine ***i***nterface ***f***or***m***a***t***ting*\n")

	clifmt.Printf("\n**SYNOPSIS**\n")
	clifmt.Printf("\tclifmt \"input text\"\n")
	clifmt.Printf("\tclifmt help\n")
	clifmt.Printf("\n")

	clifmt.Printf("\n**DESCRIPTION**\n")
	clifmt.Printf("\tclifmt allows you to upgrade your command-line experience with ${colo}(#000:#f00)${red}(#f00), **bold**, *italic*, \n")
	clifmt.Printf("\t_underlined_ text and [hyperlinks](https://git.xdrm.io/go/clifmt) with an intuitive syntax.\n")
	clifmt.Printf("\tSee the **COLORS** and **FORMATTING** sections for more details.\n")
	clifmt.Printf("\n")

	clifmt.Printf("\n**COLORS**\n")
	clifmt.Printf("\tThe syntax for colorizing text allows you to set the <foreground> color,\n")
	clifmt.Printf("\tthe <background> color, or both at the same time. Colors can be expressed in 3\n")
	clifmt.Printf("\tways : \n")
	clifmt.Printf("\t  (1) its name            (see **THEME** section for the list of available names)\n")
	clifmt.Printf("\t  (2) its hex value       (*e.g. '#ffaa88'*)\n")
	clifmt.Printf("\t  (3) its short hex value (*e.g. '#fa8'*)\n")

	clifmt.Printf("\n\t_foreground & background_\n")
	clifmt.Printf("\t  (1) \\${text to colorize}(red:green)          =>    ${text to colorize}(red:green)\n")
	clifmt.Printf("\t  (2) \\${text to colorize}(#ff0000:#00ff00)    =>    ${text to colorize}(#ff0000:#00ff00)\n")
	clifmt.Printf("\t  (3) \\${text to colorize}(#f00:#0f0)          =>    ${text to colorize}(#f00:#0f0)\n")
	clifmt.Printf("\n\t_foreground only_\n")
	clifmt.Printf("\t  (1) \\${text to colorize}(red)                =>    ${text to colorize}(red)\n")
	clifmt.Printf("\t  (2) \\${text to colorize}(#ff0000)            =>    ${text to colorize}(#ff0000)\n")
	clifmt.Printf("\t  (3) \\${text to colorize}(#f00)               =>    ${text to colorize}(#f00)\n")
	clifmt.Printf("\n\t_background only_\n")
	clifmt.Printf("\t  (1) \\${text to colorize}(:blue)               =>    ${text to colorize}(:blue)\n")
	clifmt.Printf("\t  (2) \\${text to colorize}(:#0000ff)            =>    ${text to colorize}(:#0000ff)\n")
	clifmt.Printf("\t  (3) \\${text to colorize}(:#00f)               =>    ${text to colorize}(:#00f)\n")

	clifmt.Printf("\n**FORMATTING**\n")
	clifmt.Printf("\tThe syntax for bold/italic/underline/hyperlink text is inspired by the markdown syntax :\n")
	clifmt.Printf("\t   +---------------------+----------------------+\n")
	clifmt.Printf("\t   |       **input**         |       **output**         |\n")
	clifmt.Printf("\t   |---------------------+----------------------|\n")
	clifmt.Printf("\t   | \\*\\*bold\\*\\*            | **bold**                 |\n")
	clifmt.Printf("\t   | \\*italic\\*            | *italic*               |\n")
	clifmt.Printf("\t   | \\_underlined\\_        | _underlined_           |\n")
	clifmt.Printf("\t   | \\[label](url)        | [label](url)                |\n")
	clifmt.Printf("\t   +---------------------+----------------------+\n")

	clifmt.Printf("\n\tNote that reserved characters must be escaped : \n")
	clifmt.Printf("\t  - **\\*** [asterisk]\tcan be escaped with '\\\\*'\n")
	clifmt.Printf("\t  - **$** [dollar]\t\tcan be escaped with '\\\\$'\n")
	clifmt.Printf("\t  - **_** [underscore]\tcan be escaped with '\\\\_'\n")
	clifmt.Printf("\t  - **[** [opening square brackets]\tcan be escaped with '\\\\['\n\n")

	clifmt.Printf("\n**THEME**\n")
	clifmt.Printf("\tThe theme contains the following color names :\n")

	theme := color.DefaultTheme()
	count := uint16(0)
	for k, v := range theme {
		count++
		// hexVal := strconv.FormatUint(uint64(v), 16)
		clifmt.Printf("\t (%2d) ${  }(:#%.6x) %s\n", count, v, k)
	}
	clifmt.Printf("\n")

	clifmt.Printf("**AUTHORS**\n")
	clifmt.Printf("\tclifmt has been entirely designed and developed by xdrm-brackets\n")
	clifmt.Printf("\t<xdrm.brackets.dev@gmail.com>. Feedback is really appreciated, feel\n")
	clifmt.Printf("\tfree to mail me.\n")

}
