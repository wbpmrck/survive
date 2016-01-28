##what is this

this is a part-time project,it can be a funny game.

####DOING
* 开发
    * 完整流程编写：
        * 开启一个time.Source,直接把时间片交给Arena
        * Arena 创建1个battle,然后battle开始处理战斗逻辑
            * 订阅warrior的ComputedAttribute(行动顺序)的变化，重新排序

    * ComputedAttribute.go (OK)
        * 现在的模式是每次读取的时候都去计算
        * 修改为：
            * 订阅依赖属性的变化事件
            * 内部记录脏数据标记
                * 只要数据不脏，就直接返回
                * 如果数据脏了，则重新计算
    * attributeCarrier.go(OK)
        * 添加事件机制，包括
            * 属性值修改

    * targetChooser的实现、丰富(ing)
        * 测试各种chooser,此时可以使用简单的battle对象来测试
    * arena.go
        * 作为整个探索游戏的一种游戏模式，“竞技场”对象是不受地图限制而存在的
        * 竞技场对象同样持有时间片
    * battle.go
        * 定义粗略的战斗流程
            * 战斗过程就是时间流逝的过程
                * 一个战斗对象，持有一个时间管道，从外部接受时间输入
                    * 在设想里，战斗可能发生在野外探索，也可能发生在竞技场(0.1版本),那么，在不同环境下，时间片的输入来源不同
                        * 如果是野外，那么战斗的时间输入可能就是当前地图的god协程
                        * 如果是竞技场，那么战斗的时间输入可能就是arena对象
                * 每个时间片里，战斗对象来驱动战斗过程的进行
            * 每个时间tick到之后，要执行：
                * 回复单位的行动力
                * 回复单位状态，重新计算状态
                * 根据最新状态判断行动顺序
                * 按照顺序先后，执行角色行动逻辑
        * 定义战斗判定阶段，考虑与技能、单位动作的衔接问题
    * character.go / warrior.go
        * 实现更多的战斗过程代码
        * 编写模拟战斗测试代码


####TODO
* 设计
    > 完成首批需求收集，确定0.1版本目标
* 装备系统
    > 装备的效果、和技能效果，如何区分和设计
* 效果和技能系统完善
    * 效果的排斥
        * 目前效果直接并没有直接关系，效果可以叠加。
        * 如果实现同一效果的叠加次数控制，以及不同效果的叠加互斥判断等，需要再考虑一下
* 沙箱和脚本

    > 建立基本的沙箱机制
    
    > 使用javascript来描述业务逻辑

####DONE
* 开发
    * skillBase.go
        * 定义技能基类的结构，里面应该分别包含技能项、技能类定义
        * 实现skill的接口
            * 注意技能接口的实现应该也是依赖effect接口的实现
    * effect.go
        * 添加效果字典集合
        * 在builtIn里添加几个内置的效果
    * eventEmitter.go
        * 支持通用的事件订阅、发布模型 (OK)
        * 在订阅事件的时候，对EventHandler进行构造和调用  (OK)
        * 让character内嵌Eventbase来提供事件发射、取消后续事件执行等功能  (OK)
        * 测试一下事件机制能否正常工作 (OK)
    * 规则的默认实现，以及实体的组合
        * 各种规则，如果有默认通用实现的，实现通用的版本
            * effectCarrier (OK)
            * attributeCarrier  (OK)
        * character对这些默认实现进行组合，从而实现自己的功能 (OK)
            * 定义一个角色 (OK)
            * 添加属性、效果 (OK)
            * 测试用例验证效果是否生效 (OK)
    * skill.go
        * 技能包、技能项的设计(OK)
        * 技能项的实现
            * 实现技能包 (ing)
            * 实现技能项 (ing)
                * PluginStep:(OK)
                    * 技能是主动还是被动触发，要看SkillItem里配置的Step
                        * step里如果配置的是角色上线，且目标选择器是用户自己，那么基本上就是被动技能了
                        * step里如果配置的是发起攻击后，且目选择器选择的是攻击对象，那么基本上就是主动技能
                * 技能项里的目标选择器+效果列表，到后面要手动选取成对的配置好，再随机抽取
                * 目标选择器，要能分为3大类，选择敌人的，选择友军的，和敌我不分的。 (ing)
                    * 每个choose通过配置项来决定选择对象。比如配置：只选择对象自身(适用于一些被动加成类效果)
                    * chooser可以从技能发动者出发，根据不同的阶段，选取不同的目标，比如：
                        * UserOnLine阶段：只能选择from(自己)
                        * UserAfterAttack阶段：可以选择攻击对象，也可以以攻击对象为中心，选择其周围200单位长度的敌友(灵活可控)

                * 技能项调用目标选择器得到目标，再对每1个目标putOn效果
            * 定义skillCarrier接口、提供默认实现类 (OK)
            * character内嵌默认实现类，具备技能使用能力 (OK)
            * 给效果添加SetLevel方法，并对已有效果实现根据level修正效果属性的功能 (ing)
    * time.pipe.go
        * 新增时间管道的概念：
            * 时间管道可以定义自己的时间流逝速度
            * 时间管道可以级联，可以把自己获得的时间片传递给下一阶段(获得的时间片，会以自己的倍率进行重设后交给后端)