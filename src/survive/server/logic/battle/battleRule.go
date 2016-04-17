package battle
import (
	"survive/server/logic/consts/nature"
	"survive/server/logic/math"
	sysmath "math"
)


//判断本次攻击是否致命攻击
func IsCriticalAttack(attacker,defender *Warrior,attackNature nature.Nature) bool{
	if attackNature == nature.Physical{
		//如果是物理攻击，则看【物理暴击率】
		return math.IsHitByFloatProbability(attacker.CriticalRatePhysical.GetValue().Get())
	}else if attackNature == nature.Magical{
		//如果是魔法攻击，则看【魔法暴击率】

		return math.IsHitByFloatProbability(attacker.CriticalRateMagical.GetValue().Get())
	}else{
		return false
	}
}

//获取造成的伤害
func GetDamage(attacker,defender *Warrior,attackNature nature.Nature,critical bool) float64{
	damage := 0.0
	if attackNature == nature.Physical{
		atk := attacker.AttackPhysical.GetValue().Get()
		var def float64 = 0
		//如果是非致命攻击，则防御力起作用
		if !critical{
			def = defender.DefencePhysical.GetValue().Get()
		}

		damage = sysmath.Max(atk - def,0)
	}else if attackNature == nature.Magical{

		atk := attacker.AttackMagical.GetValue().Get()
		def := 0.0
		//如果是非致命攻击，则防御力起作用
		if !critical{
			def = defender.DefenceMagical.GetValue().Get()
		}

		damage = sysmath.Max(atk - def,0)
	}
	return damage
}

