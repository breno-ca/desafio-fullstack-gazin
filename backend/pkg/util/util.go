package util

import (
	"fmt"
	"strings"
)

func Order(by string, mode string) string {
	if by == "" {
		return ""
	}

	order := `
		ORDER BY
			%s %s
	`

	mode = strings.ToUpper(mode)

	if mode == "ASC" {
		return fmt.Sprintf(order, by, "ASC")
	}
	if mode == "DESC" {
		return fmt.Sprintf(order, by, "DESC")
	}

	return fmt.Sprintf(order, by, "")
}
