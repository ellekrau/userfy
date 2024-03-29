package cellphoneformatter

import (
	clientconfig "github.com/ellekrau/userfy/config/client-config"
	"strings"
)

func FormatCellphone(clientUserConfig clientconfig.User, inputCellphone string) string {
	// If the CellphonePattern variable is empty just returns the cellphone input value
	if clientUserConfig.Cellphone == "" {
		return inputCellphone
	}

	outputCellphone := strings.ToUpper(clientUserConfig.Cellphone)
	cellphoneDigits := strings.Split(inputCellphone, "")

	for i := 0; i < len(cellphoneDigits); i++ {
		outputCellphone = strings.Replace(outputCellphone, "X", cellphoneDigits[i], 1)
	}

	return outputCellphone
}
