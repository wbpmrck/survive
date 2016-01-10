##what is this

this is a part-time project,it can be a funny game.

####DOING
* 开发
    * skillBase.go
        * 定义技能基类的结构，里面应该分别包含技能项、技能类定义
        * 实现skill的接口
            * 注意技能接口的实现应该也是依赖effect接口的实现
    * effect.go
        * 添加效果字典集合
        * 在builtIn里添加几个内置的效果
    * battle.go
        * 定义粗略的战斗流程
        * 定义战斗判定阶段，考虑与技能、单位动作的衔接问题

####TODO
* 设计
    > 完成首批需求收集，确定0.1版本目标

* 沙箱和脚本

    > 建立基本的沙箱机制
    
    > 使用javascript来描述业务逻辑
  
