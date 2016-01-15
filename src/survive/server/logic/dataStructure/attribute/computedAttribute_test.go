package attribute
import (
	"testing"
	"fmt"
)


func TestCreate(t *testing.T){

	attrMap := make(map[string]*Attribute)

	str:=NewAttribute("力量","",10)
	agi:=NewAttribute("敏捷","",12)
	attrMap["str"] = str
	attrMap["agi"] = agi

	//创建一个计算属性cp1,它的值由另外2个属性决定
	cp1 := NewComputedAttribute("战斗力","依赖力量,敏捷",100,attrMap,func(dependencies map[string]*Attribute) float64{
		return dependencies["str"].GetValue().Get()*10+dependencies["agi"].GetValue().Get()*2+2
	})

	//没有任何加成情况下的值
	fmt.Printf("没有加成时,cp1= %v \n",cp1.GetValue().Get())
	if cp1.GetValue().GetRaw()!=126{
		t.Errorf("原始值应该为126")
	}
	if cp1.GetValue().Get()!=126{
		t.Errorf("未加成情况下，最终值也应该为126")
	}

	//给力量和敏捷进行加成
	str.GetValue().Add(50) //力量+50
	agi.GetValue().AddByPercent(0.5) //敏捷+50%
	fmt.Printf("力量和敏捷修正时,cp1= %v \n",cp1.GetValue().Get())
	if cp1.GetValue().GetRaw()!=638{
		t.Errorf("原始值应该为638")
	}
	if cp1.GetValue().Get()!=638{
		t.Errorf("未加成情况下，最终值也应该为638")
	}

	//给计算属性本身进行加成
	cp1.GetValue().Add(200)
	fmt.Printf("给计算属性本身进行加成时,cp1= %v \n",cp1.GetValue().Get())
	if cp1.GetValue().GetRaw()!=638{
		t.Errorf("原始值应该为638")
	}
	if cp1.GetValue().Get()!=838{
		t.Errorf("加成情况下，最终值应该为838")
	}
}
