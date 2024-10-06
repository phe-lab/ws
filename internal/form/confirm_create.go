package form

import (
	"github.com/charmbracelet/huh"
)

func ConfirmCreateDirectory(path string) (bool, error) {
	var allowCreate bool

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("The workspace directory does not exist. Would you like to create it now?").
				Description(path).
				Value(&allowCreate),
		),
	).WithTheme(t).Run()

	return allowCreate, err
}
