package currency_formats
import (
	"github.com/leekchan/accounting"
)

var ac accounting.Accounting
func init(){
	ac = accounting.Accounting{Symbol: "$", Precision: 2}
}

func format(money int) string{
	return ac.FormatMoney(money);
}