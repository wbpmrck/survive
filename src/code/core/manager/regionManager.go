package manager
import (
	"code/core/timeRule"
	"code/core/maps"
	"code/core/agent"
)

/*负责一个区域的整体管理包括：
*	1:timeRule
	2:region区域的变化
	3:管理region里所有agent之间的消息通信和中转，确保任何2个agent可以互相通信【agent之间也可以直接互相通信，但也许有时候
	需要regionManager来转达，比如广播】
	4:管理区域agent的行动顺序，并与timeRule交互，隔离timeRule与agent
 */

type RegionManager struct {
	timeRuler *timeRule.TimeRuler //负责区域内的时间管理
	region *maps.Region	//区域的地图数据
	godAgent *agent.Agenter //当前负责管理region的agent
}

func BuildRegionManager()