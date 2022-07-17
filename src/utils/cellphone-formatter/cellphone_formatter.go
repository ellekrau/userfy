package cellphoneformatter

import (
	"github.com/ellekrau/mercafacil/config"
	"strings"
)

func FormatCellphone(inputCellphone string) string {
	// If the CellphonePattern variable is empty just returns the cellphone input value
	if config.UserData.CellphonePattern == "" {
		return inputCellphone
	}

	outputCellphone := config.UserData.CellphonePattern
	cellphoneDigits := strings.Split(inputCellphone, "")

	for i := 0; i < len(cellphoneDigits); i++ {
		outputCellphone = strings.Replace(outputCellphone, "0", cellphoneDigits[i], 1)
	}

	return outputCellphone
}
