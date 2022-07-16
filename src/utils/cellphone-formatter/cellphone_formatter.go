package cellphoneformatter

import (
	"github.com/ellekrau/mercafacil/config"
	"strings"
)

func FormatCellphone(cellphone string) string {
	// If the CellphonePattern variable is empty just returns the cellphone input value
	if config.UserData.CellphonePattern == "" {
		return cellphone
	}

	result := config.UserData.CellphonePattern
	cellphoneDigits := strings.Split(cellphone, "")

	for i := 0; i < len(cellphoneDigits); i++ {
		result = strings.Replace(result, "0", cellphoneDigits[i], 1)
	}

	return result
}
