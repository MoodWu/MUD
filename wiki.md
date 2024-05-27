一个地图分成若干格子,每个格子是一个场景
一个场景里面也分若干格子,但是场景里面不再套子场景
一个场景走到了边界处继续往外走,如果场景没有规定场景的出入点,或者在规定的出入点继续移动,则需要向系统请求是否有路,这时玩家调用move时,直接就去到新的场景,如果在边界查看外部,需要调用上级地图的查看方法,上级地图来判断是否有路
场景里面包含物品,每个场景有描述,描述中包含了物品列表,任务列表
每个场景都用一个markdown描述,比如

<scene>
  <id>村子</id>
  <desc>
  这是一个不大的村子,村口有一个<tree>大槐树</tree>,树下有两个<bottle>水壶</bottle>
  </desc>
  <item>
  <tree>  <action>climb up(爬上树)</action><scene ref= "树上"/></tree>
  </item>
</scene>
<scene>
  <id>树上</id>
  <desc>
  这是一颗高大的槐树,可以看到远处的大山和天边的云彩,不远处有一个<nest>鸟窝</nest>.
  </desc>
  <item>
  <nest>  <action>pick(掏鸟窝)</action></tree>
  </item>
  <actions>
  <action>climb down(溜下树)</action><scene ref = "村子">
  </actions>
</scene>


(0,2) | (1,2) | (2,2)
----------------------
(0,1) | (1,1) | (2,1)
----------------------
(0,0) | (1,0) | (2,0)
<map>
<scene x=1 y=1>村子</scene>
</map>


scene:新手村
scale: 3,5
initpoint: 1,1
entrypoint:
  东:2,3
  西:0,1
  南:1,0
  北:1,2
scenes:
  - id:村子
    desc:这是村子的描述
    position: 1,1
    items:
      - id:hulu
        qty: 2
        desc: 这是一个葫芦,看起来可以喝水
      - id: cloth
        name: 麻布衣服
        qty: 1
        desc: 一件普通的麻布衣服
    dirs: [east,west,north,south]
  - id:村东头
    position:0,1
    desc:这是村东头,没啥好看,后面是村子
    dirs:[west]





系统启动时,首先读取map,map中地一个scene就是默认场景,scene中地一个scene就是默认场景,将map,sence 读取到 列比哦啊中,读取是
登录,读取自己所在的map,scene,cordinate,进入scene

所有的场景 IScene,有保存功能,用于动态展开的地图下次再次进入会自动展开,
物体,有动作,比如 fillwater,drink water等,
动作分为几类,
1,观察类,直接对某个物体进行观察,可以是物体,也可以是人,IObserve 接口,被观察的物品实现此接口返回描述,可以传一个观察次数,以实现观察多次后物品描述发生变化,同时还可以接受一个context,以实现不同状态下可以观察到不同的东西,比如不同的位置看到的地图不一样
2.运动类,直接以方位为命令,是人的动作, IMove 接口,移动的物品调用此接口,返回IScene新状态,比如切换了一个新的Scene,或者更新当前Scene中的坐标
3.拾取,丢弃物品,人的动作,物品需要实现IGet 和IDrop 接口以实现拾取和丢弃,IGet 返回true/false 表示是否成功拾取,调用palyer.get,以物品为参数
3.装备,卸下装备,人对可穿戴物品(包括坐骑)的动作 IWear ITakeOff
4.服用,比如喝水,吃东西,人对可食用物品的动作, IEat 接口,plyer.eat,以物品为参数,调用物品的beEatten方法
5.格斗类,player.fight 以被打击物品为参数,如果 物品实现了IFightable接口,调用被打击物品的beaten,被打击物品,可以改变场景中物品,比如爆装备,这个在markdown中标注出来
6.超管类,比如召唤,切换场景
7.IM类

地图之间的切换，要有一个更高一级的物体来处理，比如world，或者 system啥的，它能读取所有的地图，知道地图之间的连通性，或者游戏在初始化地图时，将地图间的连通性写入地图的数据结构中，这样，地图自己就可以处理切换的问题了（这种更简单）。

所有的物品，和 地图都应该实现一个探查接口，表明，自己可以支持的动作

客户端传来的命令 到底由那个对象处理，客户端传来的命令，对应一个用户，用户属性中有物品 和 地图 ，如果 一个命令，只有一个物品有，比如drink ，当物品中只有一个水壶有 drink 动作时，可以直接让 水壶处理，但是 如果有水壶 和 酒壶 都有 drink动作时，就返回一个询问，你是想 喝水(drink water) 还是 喝酒(drink wine)，这个喝水(drink water) 喝酒（drink wine）是动作的描述性文本
如果传过来 的命令 没有任何物品可以实现，那么就返回 疑问，不识别的命令，或者可以设计一个显示当前所有命令的 命令来

地图类的作用主要是 生成大的 地图 ShowWorld 和 生成当前地图 ShowMap
用户是走在场景上，物品也是挂在场景上，场景实现 观察接口 返回 场景描述(desc字段)，物品 也实现 观察接口(desc字段) ，都实现 探查接口，返回可以的动作 (actions) 一个action 有 三个字段，action:[drink, drink water] 第一个action是struct实现的名称，第二个是别名，可以随便修改 actiondesc: 喝水 effect:player.hp +=30 effectdesc:你喝了一口水，感觉神清气爽
所有物品都有一个type字段，表明实现自己的struct的名称

有些动作需要满足一些条件，比如水壶装水，必须当前场景有水源，action 增加一个PreCondition:watersource,目前就只有一种前置条件：场景 或者 库存中有某种物品

有些动作必须是在人身上才能执行，比如，食物，首先要捡起来放入Inventory，然后才能eat，可以Item设计一个owner，默认为nil，被人拾取后设置为player

前置动作,为了简单尽量不要写为表达式
HasOwner  //有主
HasNoOwner //无主
ContainItem(watersource) //包含物品
NoInventory // *身上没有
HasSpace // * 有库存容量

所有的动作 都会改变玩家 或者 场景的状态，少数也许会改变系统的状态，因此Action应该获得玩家这个参数，还是说，Action的Excute是挂在player上

在场景的移动中，如果移动到边界，继续向外移动，允许切换到其他地图时，用Jmp(sceneCode)来表示，如果只有Jmp 没有sceneCode则表示如果地图旁边有相邻地图则跳转过去，否则不过去


加载地图时，可以解析出地图上所有的动作，这些动作 和 地图都应该缓存起来，这样有新的玩家进来可以很快读取到地图信息，目前系统共用一套地图，没有副本，缓存起来，刷新系统也方便。

玩家执行动作时，要先判断动作是否存在，因此玩家


所有动作直接用lua脚本，目前阶段传递player 和 item 两个参数

player实现PerformAction，接收一个实现了Action接口的Struct，为参数，如果Action有PreCondiiton，先进行PreCondition的判断，在执行Action，Action返回动作的效果，Object 可以是player 表明对玩家其效果，可以是 map，表示对地图起效果，比如可以产生新的通道，可以是system表示对系统起效果
[{
  Object:Player
  Effect:[{Field:hp,Value:10}]
}]

todo:场景中应该可以对场景做动作，比如挖洞？还是将土地作为一个item？

特别的接口
IPlayer ,
修改经验值,生命值,库存

所有动作返回一个大的结构体,用于指示是否需要对玩家的属性进行修改
{
  Type:"hp",
  Value:30
}


map模板
#新手村地图
map: 新手村
code: m00001
desc: 这是一个不大的村子,看起来安静祥和,你可以在这里学习如何探索世界
scale: 3,3
initpoint: 1,1
scenes:
  - name:
    code:
    desc:
    position:
    items:
      - qty:
        item:
          - name:
            desc:
            qty:
            actions:
              - cmd:
                desc:
                effect:
                effectDesc:
                condition:
    path:
      west: 

markdown中可以标识当前物品的ID,名字,每个Scene都是一个plughin,在系统启动时,会加载所有的物品 和 场景 ,以及将物品分布在场景中,场景有刷新方法,让物品定期生成.每个物品会读取markdown,获得一些属性,每个物品都可以调用系统功能,比如增加

不同的地图可能有新的物品,

玩家是一个巨大的json

程序启动流程
1.从命令行获取配置文件地址
2.从配置文件读取插件目录,加载所有的插件
3.读取地图配置文件,初始场景
4.启动WebSocket server

WebSocket功能
1.接收命令,登录/指令/登出
2.转发命令给不同的模块
  登录 给登录函数,返回一个用户ID
  指令 给命令中心,先根据用户ID构建用户实例,在调用具体方法,返回 json

场景的初始化
1.挨个读取子场景,放入场景列表,
2.读取地图中的物品,实例化
3.读取场景中的玩家列表