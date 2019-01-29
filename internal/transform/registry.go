package transform

// Registry is used to apply a stack of transformations
// over an input string
type Registry struct {
	// cursor is the current transformer
	cursor uint
	// Transformers represents the transformer stack
	// ; each one will be executed in ascending order
	Transformers []Transformer
}

// Transform executes each transformer of the stack in ascending order feeding
// each one with the output of its predecessor (@input for the first). Note that if one returns an error
// the process stops here and the error is directly returned.
func (r *Registry) Transform(input string) (string, error) {
	in := input

	// execute each transformer by order
	for _, t := range r.Transformers {

		// 1.  execute ; dispatch error on failure
		out, err := execute(t, in)
		if err != nil {
			return "", err
		}

		// 2. replace next input with current output
		in = out
	}

	return in, nil
}

// execute 1 given transformer @t with its @input string and returns the output,
// and the error if one.
func execute(t Transformer, input string) (string, error) {
	var (
		output string
		cursor int
	)

	// apply transformatione for each match
	for _, match := range t.Regex().FindAllStringSubmatchIndex(input, -1) {

		// (1) append gap between input start OR previous match
		output += input[cursor:match[0]]
		cursor = match[1]

		// (2) build transformation arguments
		args := make([]string, 0, len(match)/2+1)
		for i, l := 2, len(match); i < l; i += 2 {
			// match exists (not both -1, nor negative length)
			if match[i+1]-match[i] > 0 {
				args = append(args, input[match[i]:match[i+1]])
				continue
			}
			args = append(args, "")
		}

		// (3) execute transformation
		transformed, err := t.Transform(args...)
		if err != nil {
			return "", &TransformerError{t, err, input[match[0]:match[1]]}
		}

		// (4) apply transformation
		output += transformed

	}

	// Add end of input (if not covered by matches)
	if cursor < len(input) {
		output += input[cursor:]
	}

	// return final output
	return output, nil
}
