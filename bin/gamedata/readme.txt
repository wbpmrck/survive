本目录存放一些游戏逻辑相关配置数据：
1、基于 CSV 格式
2、用于管理游戏配置数据。
3、在 LeafServer 中使用 recordfile 非常简单：

例子：【在 gamedata 模块中调用函数 readRf 读取 CSV 文件】


// 确保 bin/gamedata 目录中存在 Test.txt 文件
// 文件名必须和此结构体名称相同（大小写敏感）
// 结构体的一个实例映射 recordfile 中的一行
type Test struct {
    // 将第一列按 int 类型解析
    // "index" 表明在此列上建立唯一索引
    Id  int "index"
    // 将第二列解析为长度为 4 的整型数组
    Arr [4]int
    // 将第三列解析为字符串
    Str string
}

// 读取 recordfile Test.txt 到内存中
// RfTest 即为 Test.txt 的内存镜像
var RfTest = readRf(Test{})

func init() {
    // 按索引查找
    // 获取 Test.txt 中 Id 为 1 的那一行
    r := RfTest.Index(1)

    if r != nil {
        row := r.(*Test)

        // 输出此行的所有列的数据
        log.Debug("%v %v %v", row.Id, row.Arr, row.Str)
    }
}