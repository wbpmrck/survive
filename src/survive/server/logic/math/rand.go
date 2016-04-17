package math
import (
	"math/rand"
	"time"
	"github.com/wbpmrck/golibs/random"
)
var rand_seed *rand.Rand
const MaxProbability int = 1000 + 1 //表示最大可能性的概率值(1000/1000)
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

//简便的判断一个事件是否发生的方法，参数为百分比的概率 [0,1]
func IsHitByFloatProbability(prob float64) bool{
	r := float64(NextRandomInt(MaxProbability))
	return r<prob*1000
}

func init(){
	rand_seed = rand.New(rand.NewSource(time.Now().Unix()))
	//修改默认的seed
	rand.Seed(time.Now().UnixNano())
}

