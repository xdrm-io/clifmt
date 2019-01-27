package markdown

func Transform(input string) (string, error) {

	// 1. bold
	bold, err := boldTransform(input)
	if err != nil {
		return "", err
	}

	// 2. italic
	italic, err := italicTransform(bold)
	if err != nil {
		return "", err
	}

	// 3. underline
	underline, err := underlineTransform(italic)
	if err != nil {
		return "", err
	}

	// 4. hyperlink
	hyperlinked, err := hyperlinkTransform(underline)
	if err != nil {
		return "", err
	}

	return hyperlinked, nil
}
