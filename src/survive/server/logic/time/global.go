package time
import "survive/server/logic/dataStructure"

var now *dataStructure.Time

func GetNow() *dataStructure.Time{
	return now
}
