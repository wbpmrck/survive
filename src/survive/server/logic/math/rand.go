package math
import (
	"math/rand"
	"time"
	"github.com/wbpmrck/golibs/random"
)
var rand_seed *rand.Rand
/*
	提供随机事件组功能，可以在多个事件中得到随机命中的结果
 */
func NewRandomEventGroup(probability ...uint32)*random.RandomEventGroup{
	return random.CreateEventGroup(rand_seed,probability...)
}

//返回一个 [0,max)的整数
func NextRandomInt(max int) int{
	return rand.Intn(max)
}

func init(){
	rand_seed = rand.New(rand.NewSource(time.Now().Unix()))
	//修改默认的seed
	rand.Seed(time.Now().UnixNano())
}

