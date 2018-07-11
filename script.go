package bitcointx

import (
	"github.com/golang-collections/collections/stack"
)

const (
	// 入栈操作
	op0     = byte(0x00) // 一个字节空串被压入堆栈中
	opFalse = op0
	// 1-75 	0x01-0x4b 把接下来的N个字节压入堆栈中，N 的取值在 1 到 75 之间
	opPushData1 = byte(0x4c) // 下一个脚本字节包括N，会将接下来的N个字节压入堆栈
	opPushData2 = byte(0x4d) // 下两个脚本字节包括N，会将接下来的N个字节压入堆栈
	opPushData4 = byte(0x4e) // 下四个脚本字节包括N，会将接下来的N个字节压入堆栈
	op1Negate   = byte(0x4f) // 将脚本-1压入堆栈
	opReserve   = byte(0x50) // 终止 - 交易无效（除非在未执行的 OP_IF 语句中）
	op1         = byte(0x51) // 将脚本1压入堆栈
	opTrue      = op1
	op2         = byte(0x52) // 将脚本N压入堆栈，例如 OP_2 压入脚本“2”

	// 有条件的流量控制操作
	opNop      = byte(0x61) // 无操作
	opVer      = byte(0x62) // 终止 - 交易无效（除非在未执行的 OP_IF 语句中）
	opIf       = byte(0x63) // 如果栈项元素值为0，语句将被执行
	opNotif    = byte(0x64) // 如果栈项元素值不为0，语句将被执行
	opVerif    = byte(0x65) // 终止 - 交易无效
	opVernotif = byte(0x66) // 终止 - 交易无效
	opElse     = byte(0x67) // 如果前述的 OP_IF 或 OP_NOTIF 未被执行，这些语句就会被执行
	opEndif    = byte(0x68) // 终止 OP_IF, OP_NOTIF, OP_ELSE 区块
	opVerify   = byte(0x69) // 如果栈项元素值非真，则标记交易无效，OP_VERIFY
	opReturn   = byte(0x6a) // 标记交易无效

	// 堆栈操作
	opToaltStack   = byte(0x6b) // 从主堆栈中取出元素，推入辅堆栈。
	opFromaltStack = byte(0x6c) // 从辅堆栈中取出元素，推入主堆栈
	op2drop        = byte(0x6d) // 删除栈顶两个元素
	op2dup         = byte(0x6e) // 复制栈顶两个元素
	op3dup         = byte(0x6f) // 复制栈顶三个元素
	op2over        = byte(0x70) // 把栈底的第三、第四个元素拷贝到栈顶
	op2rot         = byte(0x71) // 移动第五、第六元素到栈顶
	op2swap        = byte(0x72) // 交换栈顶两个元素
	opIfdup        = byte(0x73) // 如果栈项元素值不为0，复制该元素值
	opDepth        = byte(0x74) // 把堆栈元素的个数压入堆栈
	opDrop         = byte(0x75) // 删除栈顶元素
	opDup          = byte(0x76) // 复制栈顶元素
	opNip          = byte(0x77) // 删除栈顶的下一个元素
	opOver         = byte(0x78) // 复制栈顶的下一个元素到栈顶
	opPick         = byte(0x79) // 把堆栈的第n个元素拷贝到栈顶
	opRoll         = byte(0x7a) // 把堆栈的第n个元素移动到栈顶
	opRot          = byte(0x7b) // 翻转栈顶的三个元素
	opSwap         = byte(0x7c) // 栈顶的三个元素交换
	opTuck         = byte(0x7d) // 拷贝栈顶元素并插入到栈顶第二个元素之后

	// 字符串接操作
	opCat    = byte(0x7e) // 已禁用, 连接两个字符串
	opSubstr = byte(0x7f) // 已禁用, 返回字符串的一部分
	opLeft   = byte(0x80) // 已禁用, 在一个字符串中保留左边指定长度的子串
	opRight  = byte(0x81) // 已禁用, 在一个字符串中保留右边指定长度的子串
	opSize   = byte(0x82) // 把栈顶元素的字符串长度压入堆栈

	// 二进制算术和条件
	opInvert      = byte(0x83) // 已禁用, 所有输入的位取反，
	opAnd         = byte(0x84) // 已禁用, 对输入的所有位进行布尔与运算
	opOr          = byte(0x85) // 已禁用, 对输入的每一位进行布尔或运算
	opXor         = byte(0x86) // 已禁用, 对输入的每一位进行布尔异或运算
	opEqual       = byte(0x87) // 如果输入的两个数相等，返回1，否则返回0
	opEqualVerify = byte(0x88) // 与 OP_EQUAL 一样，如结果为0，之后运行 OP_VERIFY
	opReserved1   = byte(0x89) // 终止 - 无效交易（除非在未执行的OP_IF语句中）
	opReserved2   = byte(0x8a) // 终止-无效交易（除非在未执行的OP_IF语句中）

	// 数值操作
	op1add               = byte(0x8b) // 输入值加1
	op1sub               = byte(0x8c) // 输入值减1
	op2mul               = byte(0x8d) // 无效（输入值乘2）
	op2div               = byte(0x8e) // 无效（输入值除2）
	opNegate             = byte(0x8f) // 输入值符号取反
	opAbs                = byte(0x90) // 输入值符号取正
	opNot                = byte(0x91) // 如果输入值为0或1，则输出1或0；否则输出0
	op0NotEqual          = byte(0x92) // 输入值为0输出0；否则输出1
	opADD                = byte(0x93) // 输出输入两项之和
	opSub                = byte(0x94) // 输出输入（第二项减去第一项）之差
	opMul                = byte(0x95) // 禁用, （输出输入两项的积）
	opDiv                = byte(0x96) // 禁用, （输出用第二项除以第一项的倍数）
	opMod                = byte(0x97) // 禁用, （输出用第二项除以第一项得到的余数）
	opLshift             = byte(0x98) // 禁用, （左移第二项，移动位数为第一项的字节数）
	opRshift             = byte(0x99) // 禁用, （右移第二项，移动位数为第一项的字节数）
	opBoolAnd            = byte(0x9a) // 两项都不会为0，输出1，否则输出0
	opBoolOr             = byte(0x9b) // 两项有一个不为0，输出1，否则输出0
	opNumEqual           = byte(0x9c) // 两项相等则输出1，否则输出为0
	opNumEqualVerify     = byte(0x9d) // 和 NUMEQUAL 相同，如结果为0运行 OP_VERIFY
	opNumNotEqual        = byte(0x9e) // 如果栈顶两项不是相等数的话，则输出1
	opLessThan           = byte(0x9f) // 如果第二项小于栈顶项，则输出1
	opGreaterThan        = byte(0xa0) // 如果第二项大于栈顶项，则输出1
	opLessThanOrEqual    = byte(0xa1) // 如果第二项小于或等于第一项，则输出1
	opGreaterThanOrEqual = byte(0xa2) // 如果第二项大于或等于第一项，则输出1
	opMin                = byte(0xa3) // 输出栈顶两项中较小的一项
	opMax                = byte(0xa4) // 输出栈顶两项中较大的一项
	opWithin             = byte(0xa5) // 如果第三项的数值介于前两项之间，则输出1

	// 加密和散列操作
	opRipemd160           = byte(0xa6) // 返回栈顶项的 RIPEMD160 哈希值
	opSha1                = byte(0xa7) // 返回栈顶项 SHA1 哈希值
	opSha256              = byte(0xa8) // 返回栈顶项 SHA256 哈希值
	opHash160             = byte(0xa9) // 栈顶项进行两次HASH，先用SHA-256，再用RIPEMD-160
	opHash256             = byte(0xaa) // 栈顶项用SHA-256算法HASH两次
	opCodeSeparator       = byte(0xab) // 标记已进行签名验证的数据
	opCheckSig            = byte(0xac) // 交易所用的签名必须是哈希值和公钥的有效签名，如果为真，则返回1
	opCheckSigVerify      = byte(0xad) // 与 CHECKSIG 一样，但之后运行 OP_VERIFY
	opCheckMultiSig       = byte(0xae) // 对于每对签名和公钥运行 CHECKSIG。所有的签名要与公钥匹配。因为存在BUG，一个未使用的外部值会从堆栈中删除。
	opCheckMultiSigVerify = byte(0xaf) // 与 CHECKMULTISIG 一样，但之后运行 OP_VERIFY

	// 非操作
	opNop1  = byte(0xb0) // 无操作 忽略
	opNop2  = byte(0xb1) // 无操作 忽略
	opNop3  = byte(0xb2) // 无操作 忽略
	opNop4  = byte(0xb3) // 无操作 忽略
	opNop5  = byte(0xb4) // 无操作 忽略
	opNop6  = byte(0xb5) // 无操作 忽略
	opNop7  = byte(0xb6) // 无操作 忽略
	opNop8  = byte(0xb7) // 无操作 忽略
	opNop9  = byte(0xb8) // 无操作 忽略
	opNop10 = byte(0xb9) // 无操作 忽略

	// 仅供内部使用的保留关键字
	opSmallData     = byte(0xf9) // 代表小数据域
	opSmallInteger  = byte(0xfa) // 代表小整数数据域
	opPubKeys       = byte(0xfb) // 代表公钥域
	opPubKeyHash    = byte(0xfd) // 代表公钥哈希域
	opPubKey        = byte(0xfe) // 代表公钥域
	opInvalidOpCode = byte(0xff) // 代表当前未指定的操作码
)

// ScriptExecutor something that execute the script
type ScriptExecutor struct {
	stack *stack.Stack
}

func (scriptExe *ScriptExecutor) init() {
	scriptExe.stack = stack.New()
}

func (scriptExe *ScriptExecutor) execute(codes []interface{}) {
	for _, code := range codes {

		scriptExe.stack.Push(code)
	}
}
