package common

import "golang.org/x/text/language"
import "golang.org/x/text/message"

func FormatCurrency(localeID string, amount int64, useCurrencySymbol bool) (formattedCurrency string) {
	switch localeID {
	case "id_ID":
		p := message.NewPrinter(language.Indonesian)

		if useCurrencySymbol {
			formattedCurrency = p.Sprintf("Rp%d", amount)
		} else {
			formattedCurrency = p.Sprintf("%d", amount)
		}
	}

	return
}
