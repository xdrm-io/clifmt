package clifmt

import (
	"strings"
	"testing"
)

func TestColoring(t *testing.T) {

	tests := []struct {
		Input  string
		Output string
	}{
		// foreground + background
		{
			"start ${some text input}(#ff0000:#00ff00) end\n",
			"start \x1b[38;2;255;0;0;48;2;0;255;0msome text input\x1b[0m end\n",
		}, {
			"start ${some text input}(#f00:#0f0) end\n",
			"start \x1b[38;2;255;0;0;48;2;0;255;0msome text input\x1b[0m end\n",
		}, {
			"start ${some text input}(red:green) end\n",
			"start \x1b[38;2;255;0;0;48;2;0;255;0msome text input\x1b[0m end\n",
		},

		// mixed notations
		{
			"start ${some text input}(red:#00ff00) end\n",
			"start \x1b[38;2;255;0;0;48;2;0;255;0msome text input\x1b[0m end\n",
		}, {
			"start ${some text input}(red:#0f0) end\n",
			"start \x1b[38;2;255;0;0;48;2;0;255;0msome text input\x1b[0m end\n",
		}, {
			"start ${some text input}(#ff0000:green) end\n",
			"start \x1b[38;2;255;0;0;48;2;0;255;0msome text input\x1b[0m end\n",
		}, {
			"start ${some text input}(#f00:green) end\n",
			"start \x1b[38;2;255;0;0;48;2;0;255;0msome text input\x1b[0m end\n",
		},

		// foreground only
		{
			"start ${some text input}(red) end\n",
			"start \x1b[38;2;255;0;0msome text input\x1b[0m end\n",
		}, {
			"start ${some text input}(#ff0000) end\n",
			"start \x1b[38;2;255;0;0msome text input\x1b[0m end\n",
		}, {
			"start ${some text input}(#f00) end\n",
			"start \x1b[38;2;255;0;0msome text input\x1b[0m end\n",
		},

		// background only
		{
			"start ${some text input}(:blue) end\n",
			"start \x1b[48;2;0;0;255msome text input\x1b[0m end\n",
		}, {
			"start ${some text input}(:#0000ff) end\n",
			"start \x1b[48;2;0;0;255msome text input\x1b[0m end\n",
		}, {
			"start ${some text input}(:#00f) end\n",
			"start \x1b[48;2;0;0;255msome text input\x1b[0m end\n",
		},

		// multi matches
		{
			"start ${text1}(red) separation ${text2}(#0f0) end\n",
			"start \x1b[38;2;255;0;0mtext1\x1b[0m separation \x1b[38;2;0;255;0mtext2\x1b[0m end\n",
		}, {
			"start ${text1}(:red) separation ${text2}(:#0f0) end\n",
			"start \x1b[48;2;255;0;0mtext1\x1b[0m separation \x1b[48;2;0;255;0mtext2\x1b[0m end\n",
		}}

	for i, test := range tests {
		output := Sprintf(test.Input)
		if output != test.Output {
			t.Errorf("[%d] expected '%s', got '%s'", i, test.Output, output)
		}

	}
}

func TestMarkdown(t *testing.T) {

	tests := []struct {
		Input  string
		Output string
	}{
		// each notation
		{
			"start **bold text** end\n",
			"start \x1b[1mbold text\x1b[22m end\n",
		}, {
			"start *italic text* end\n",
			"start \x1b[3mitalic text\x1b[23m end\n",
		}, {
			"start _underlined text_ end\n",
			"start \x1b[4munderlined text\x1b[24m end\n",
		},

		// mixed notations
		{
			"start ***bold italic*** end\n",
			"start \x1b[3m\x1b[1mbold italic\x1b[23m\x1b[22m end\n",
		}, {
			"start **_bold underline_** end\n",
			"start \x1b[1m\x1b[4mbold underline\x1b[24m\x1b[22m end\n",
		}, {
			"start _**bold underline**_ end\n",
			"start \x1b[4m\x1b[1mbold underline\x1b[22m\x1b[24m end\n",
		}, {
			"start *_italic underline_* end\n",
			"start \x1b[3m\x1b[4mitalic underline\x1b[24m\x1b[23m end\n",
		}, {
			"start _*italic underline*_ end\n",
			"start \x1b[4m\x1b[3mitalic underline\x1b[23m\x1b[24m end\n",
		}, {
			"start _***bold italic underline***_ end\n",
			"start \x1b[4m\x1b[3m\x1b[1mbold italic underline\x1b[23m\x1b[22m\x1b[24m end\n",
		}, {
			"start **_*bold italic underline*_** end\n",
			"start \x1b[1m\x1b[4m\x1b[3mbold italic underline\x1b[23m\x1b[24m\x1b[22m end\n",
		}, {
			"start *_**bold italic underline**_* end\n",
			"start \x1b[3m\x1b[4m\x1b[1mbold italic underline\x1b[22m\x1b[24m\x1b[23m end\n",
		}, {
			"start _***bold italic underline***_ end\n",
			"start \x1b[4m\x1b[3m\x1b[1mbold italic underline\x1b[23m\x1b[22m\x1b[24m end\n",
		},

		// mixed notations not matching
		{
			"start ***bold** italic* end\n",
			"start \x1b[3m\x1b[1mbold\x1b[22m italic\x1b[23m end\n",
		}, {
			"start **_bold** underline_ end\n",
			"start \x1b[1m\x1b[4mbold\x1b[22m underline\x1b[24m end\n",
		}, {
			"start _**bold_ underline** end\n",
			"start \x1b[4m\x1b[1mbold\x1b[24m underline\x1b[22m end\n",
		},
	}

	for i, test := range tests {
		output := Sprintf(test.Input)

		if output != test.Output {
			t.Errorf("[%d] expected '%s'\n", i, strings.Replace(strings.Replace(test.Output, "\n", "\\n", -1), "\x1b", "\\e", -1))
			t.Errorf("[%d]      got '%s'\n", i, strings.Replace(strings.Replace(output, "\n", "\\n", -1), "\x1b", "\\e", -1))
		}

	}
}
