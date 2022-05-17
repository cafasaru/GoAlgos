package reverse

import (
	"errors"
)

// reverse("hello") => "olleh"
func reverse(word string) (string, error) {
	var tmp string
	if len(word) == 0 {
		return "", errors.New("word needs to be grater than 0 chars")
	}
	for _, r := range word {
		tmp = string(r) + tmp
	}
	return tmp, nil
}
