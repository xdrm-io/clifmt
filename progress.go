package clifmt

import (
	"fmt"
	"reflect"
	"strings"
)

var ErrNoNewline = fmt.Errorf("no newline allowed in progress mode")

// Printpf prints a progress (dynamic) line that rewrites itself
// on arguments' update
func Printpf(format string, args ...interface{}) error {

	// 1. check format
	if strings.ContainsAny(format, "\n\r") {
		return ErrNoNewline
	}

	// 2. init
	fixed := make([]interface{}, len(args), len(args)) // actual values
	update := make([]chan interface{}, 0, len(args))   // channels that update values
	updateIndex := make([]int, 0, len(args))           // association [order -> index in @fixed]

	// 3. manage fixed values vs. updatable values (channels)
	for i, arg := range args {

		// channel -> keep Zero value + store channel
		if reflect.TypeOf(arg).Kind() == reflect.Chan {
			updateIndex = append(updateIndex, i)
			update = append(update, arg.(chan interface{}))
			continue
		}

		// raw -> set value
		fixed[i] = arg
	}

	// 4. launch dynamic select for each channel
	maxlen := 0
	nselect(update, func(i int, value interface{}, ok bool) {

		// channel is closed -> do nothing
		if !ok {
			return
		}

		// extract real index
		index := updateIndex[i]

		// update value
		fixed[index] = value

		// ignore on errors (updatable values still NIL)
		str, err := Sprintf(format, fixed...)
		if err != nil {
			return
		}
		reallen := displaySize(str)

		// print string
		fmt.Printf("\r%s", str)

		// pad right to end of max size
		if reallen < maxlen {
			pad := make([]byte, 0, maxlen-reallen)
			for i := reallen; i < maxlen; i++ {
				pad = append(pad, ' ')
			}
			fmt.Printf("%s", pad)
		} else {
			maxlen = reallen
		}

	})

	fmt.Printf("\n")

	return nil

}

// nselect selects on N channels
func nselect(channels []chan interface{}, handler func(int, interface{}, bool)) {

	// 1. build the case list containing each channel
	cases := make([]reflect.SelectCase, len(channels))
	for i, ch := range channels {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
	}

	// 2. wait for selections
	remaining := len(cases)
	for remaining > 0 {
		index, value, ok := reflect.Select(cases)

		// (1) Closed
		if !ok {
			cases[index].Chan = reflect.ValueOf(nil)
			remaining--
			continue
		}

		// (2) Received data
		handler(index, value, ok)
	}
}
