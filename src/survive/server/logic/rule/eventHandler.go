package rule
import "survive/server/logic/dataStructure"
/*
	基础的事件处理器

	这个处理器可以被外部初始化，并插入到各个实体的执行阶段里
 */
type EventHandler interface {
	//能够执行事件处理
	//返回是否取消实体店本次操作，以及操作结果
	//注意，只有在beforeXXX事件中返回 true的时候，才有可能取消后续操作
	Handle(time *dataStructure.Time) (isCancel bool,handleResult string)
}
