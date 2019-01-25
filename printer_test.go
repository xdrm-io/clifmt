package clifmt

import (
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
			"start \033[48;2;0;255;0m\033[38;2;255;0;0msome text input\033[0m\033[0m end\n",
		}, {
			"start ${some text input}(#f00:#0f0) end\n",
			"start \033[48;2;0;255;0m\033[38;2;255;0;0msome text input\033[0m\033[0m end\n",
		}, {
			"start ${some text input}(red:green) end\n",
			"start \033[48;2;0;255;0m\033[38;2;255;0;0msome text input\033[0m\033[0m end\n",
		},

		// mixed notations
		{
			"start ${some text input}(red:#00ff00) end\n",
			"start \033[48;2;0;255;0m\033[38;2;255;0;0msome text input\033[0m\033[0m end\n",
		}, {
			"start ${some text input}(red:#0f0) end\n",
			"start \033[48;2;0;255;0m\033[38;2;255;0;0msome text input\033[0m\033[0m end\n",
		}, {
			"start ${some text input}(#ff0000:green) end\n",
			"start \033[48;2;0;255;0m\033[38;2;255;0;0msome text input\033[0m\033[0m end\n",
		}, {
			"start ${some text input}(#f00:green) end\n",
			"start \033[48;2;0;255;0m\033[38;2;255;0;0msome text input\033[0m\033[0m end\n",
		},

		// foreground only
		{
			"start ${some text input}(red) end\n",
			"start \033[38;2;255;0;0msome text input\033[0m end\n",
		}, {
			"start ${some text input}(#ff0000) end\n",
			"start \033[38;2;255;0;0msome text input\033[0m end\n",
		}, {
			"start ${some text input}(#f00) end\n",
			"start \033[38;2;255;0;0msome text input\033[0m end\n",
		},

		// background only
		{
			"start ${some text input}(:blue) end\n",
			"start \033[48;2;0;0;255msome text input\033[0m end\n",
		}, {
			"start ${some text input}(:#0000ff) end\n",
			"start \033[48;2;0;0;255msome text input\033[0m end\n",
		}, {
			"start ${some text input}(:#00f) end\n",
			"start \033[48;2;0;0;255msome text input\033[0m end\n",
		},

		// multi matches
		{
			"start ${text1}(red) separation ${text2}(#0f0) end\n",
			"start \033[38;2;255;0;0mtext1\033[0m separation \033[38;2;0;255;0mtext2\033[0m end\n",
		}, {
			"start ${text1}(:red) separation ${text2}(:#0f0) end\n",
			"start \033[48;2;255;0;0mtext1\033[0m separation \033[48;2;0;255;0mtext2\033[0m end\n",
		}}

	for i, test := range tests {
		output, err := Sprintf(test.Input)
		if err != nil {
			t.Errorf("[%d] unexpected error <%v>", i, err)
			break
		}

		if output != test.Output {
			t.Errorf("[%d] expected '%s', got '%s'", i, test.Output, output)
		}

	}
}
