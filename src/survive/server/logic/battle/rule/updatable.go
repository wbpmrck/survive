package rule

/*
	Updatable 表示一个可被定期更新状态的实体
 */

type Updatable interface {
	OnBeforeUpdate (int) bool
}