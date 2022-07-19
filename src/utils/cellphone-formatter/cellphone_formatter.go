package cellphoneformatter

import (
	"github.com/ellekrau/mercafacil/config"
	"strings"
)

func FormatCellphone(inputCellphone string) string {
	// If the CellphonePattern variable is empty just returns the cellphone input value
	if config.OldUserData.CellphonePattern == "" {
		return inputCellphone
	}

	outputCellphone := strings.ToUpper(config.OldUserData.CellphonePattern)
	cellphoneDigits := strings.Split(inputCellphone, "")

	for i := 0; i < len(cellphoneDigits); i++ {
		outputCellphone = strings.Replace(outputCellphone, "X", cellphoneDigits[i], 1)
	}

	return outputCellphone
}
