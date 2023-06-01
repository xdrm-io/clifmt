package clifmt

import (
	"fmt"
	"reflect"
	"strings"
)

// Printpf prints with possible updatable values.
// Arguments can be interface{} (standard fmt.Printf)
// or can be channels (chan interface{}) that will make the
// output update on channel reception.
//
// You SHOULD launch it in a goroutine i.e. go clifmt.Printpf()
// in order for the select{} to work
//
// It can work with multiple lines thanks to ANSI escape sequences
// that allows to rewrite previously written lines
func Printpf(format string, args ...interface{}) error {
	values := make([]interface{}, len(args), len(args)) // actual values
	channels := make([]chan interface{}, 0, len(args))  // channels that update values
	indexes := make([]int, 0, len(args))                // association [channel order -> index in @fixed]

	// manage values vs. channels
	for i, arg := range args {
		// channel -> keep Zero value + store channel
		if reflect.TypeOf(arg).Kind() == reflect.Chan {
			indexes = append(indexes, i)
			channels = append(channels, arg.(chan interface{}))
			continue
		}

		// raw -> set value
		values[i] = arg
	}

	// launch dynamic select for each channel
	var rows int = -1
	nselect(channels, func(i int, value interface{}, ok bool) {
		// channel is closed -> do nothing
		if !ok {
			return
		}

		// update value
		index := indexes[i]
		values[index] = value

		// don't print on error (values still NIL)
		str, err := Sprintf(format, values...)
		if err != nil {
			return
		}

		// rewind N lines (written previous time)
		if rows > 0 {
			fmt.Printf("\x1b[%dF\x1b[K", rows)
		} else {
			fmt.Printf("\r\x1b[K")
		}
		rows = strings.Count(str, "\n")

		// make each line rewrite the previous line
		str = strings.Replace(str, "\n", "\n\x1b[K", 11)

		fmt.Printf("%s", str)

	})

	return nil

}

// nselect selects on N channels
func nselect(channels []chan interface{}, handler func(int, interface{}, bool)) {
	// build the case list containing each channel
	cases := make([]reflect.SelectCase, len(channels))
	for i, ch := range channels {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
	}

	// wait for selections
	remaining := len(cases)
	for remaining > 0 {
		index, value, ok := reflect.Select(cases)
		// closed
		if !ok {
			cases[index].Chan = reflect.ValueOf(nil)
			remaining--
			continue
		}

		// received data
		handler(index, value, ok)
	}
}
