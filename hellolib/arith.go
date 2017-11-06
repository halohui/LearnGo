package hellolib

import (
	"fmt"
	"unsafe"
)

/****************************************************************************
* 功能描述: 二叉树节点
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-08-09 08:28:25          0.00           cxh                创建
*
*****************************************************************************/

type BinTreeNode struct {
	data       int
	leftChild  *BinTreeNode
	rightChild *BinTreeNode
}

const MAXSIZE = 1000

/****************************************************************************
* 功能描述: 二叉树的先序非递归算法
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-08-08 13:17:14          0.00           cxh                创建
*
*****************************************************************************/

func nonRecPreOrder(root * BinTreeNode) {

	var (
		stack [MAXSIZE]*BinTreeNode
		top   = -1
		cur   *BinTreeNode
	)

	if root == nil {
		panic("Parameter is wrong!")
	}

	cur = root
	top++
	stack[top] = cur

	for top != -1 {
		cur = stack[top]
		top--

		fmt.Print(cur.data, "\t")

		if cur.rightChild != nil {
			top++
			stack[top] = cur.rightChild
		}

		if cur.leftChild != nil {
			top++
			stack[top] = cur.leftChild
		}
	}
}

/****************************************************************************
* 功能描述: 二叉树的中序非递归算法
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-08-08 13:17:41          0.00           cxh                创建
*
*****************************************************************************/

func nonRecInorder(root *BinTreeNode) {
	var (
		stack [MAXSIZE]*BinTreeNode
		top                = -1
		cur   *BinTreeNode = root
	)

	for top != -1 || cur != nil {

		for cur != nil {
			top++
			stack[top] = cur
			cur = cur.leftChild
		}

		if top != -1 {
			cur = stack[top]
			top--
			fmt.Print(cur.data, "\t")

			cur = cur.rightChild
		}
	}
}

/****************************************************************************
* 功能描述: 二叉树的后序非递归算法
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-08-08 13:17:54          0.00           cxh                创建
*
*****************************************************************************/

func nonRecPostOrder(root *BinTreeNode) {
	var (
		stack [MAXSIZE]*BinTreeNode
		top                = -1
		cur   *BinTreeNode = root
		pre   *BinTreeNode
		flag  bool
	)

	for {

		for cur != nil {
			top++
			stack[top] = cur
			cur = cur.leftChild
		}

		flag = true
		pre = nil

		for top != -1 && true == flag {
			cur = stack[top]

			if cur.rightChild == pre {
				top--
				fmt.Print(cur.data, "\t")
				pre = cur
			} else {
				flag = false
				cur = cur.rightChild
			}
		}

		if top == -1 {
			break
		}
	}
}

/****************************************************************************
* 功能描述: 二叉树的层次遍历
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-08-08 13:18:06          0.00           cxh                创建
*
*****************************************************************************/

func nonLevelOrder(root *BinTreeNode) {
	var (
		qu    [MAXSIZE]*BinTreeNode
		front              = 0
		rear               = 0
		cur   *BinTreeNode = root
	)

	if root == nil {
		panic("Parameter is wrong!")
	}

	rear = (rear + 1) % MAXSIZE
	qu[rear] = cur

	for rear != front {
		front = (front + 1) % MAXSIZE
		cur = qu[front]
		fmt.Print(cur.data, "\t")

		if cur.leftChild != nil {
			rear = (rear + 1) % MAXSIZE
			qu[rear] = cur.leftChild
		}

		if cur.rightChild != nil {
			rear = (rear + 1) % MAXSIZE
			qu[rear] = cur.rightChild
		}
	}
}

/****************************************************************************
* 功能描述: 直接插入排序
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-08-09 08:29:27          0.00           cxh                创建
*
*****************************************************************************/

func insertSort(data []int) {
	var j int
	for i := 1; i < len(data); i++ {
		cur := data[i]

		for j = i - 1; j >= 0 && data[j] > cur; j-- {
			data[j+1] = data[j]
		}
		data[j+1] = cur
	}
}

/****************************************************************************
* 功能描述: 折半插入排序
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-08-09 08:29:41          0.00           cxh                创建
*
*****************************************************************************/

func binInsertSort(data [] int) {
	var start, end, cur, mid int

	for i := 1; i < len(data); i++ {
		start = 0
		end = i - 1
		cur = data[i]

		for start <= end {
			mid = (start + end) / 2
			if data[mid] > cur {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}

		for j := i - 1; j >= end+1; j-- {
			data[j+1] = data[j]
		}
		data[end+1] = cur
	}
}

/****************************************************************************
* 功能描述: shell排序
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-08-09 08:29:55          0.00           cxh                创建
*
*****************************************************************************/

func shellCore(data []int, inc int) {
	var cur, j int
	for i := inc; i < len(data); i++ {
		cur = data[i]
		for j = i - inc; j >= 0 && data[j] > cur; j -= inc {
			data[j+inc] = data[j]
		}
		data[j+inc] = cur
	}
}

func shellSort(data []int) {
	for i := len(data) / 2; i >= 1; i = i / 2 {
		shellCore(data, i)
	}
}

/****************************************************************************
* 功能描述: 冒泡排序
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-08-09 08:30:05          0.00           cxh                创建
*
*****************************************************************************/

func bubbleSort(data []int) {
	var length = len(data)
	var flag bool
	for i := 1; i < length; i++ {
		flag = false
		for j := 0; j < length-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				flag = true
			}
		}
		if false == flag {
			break
		}
	}
}

/****************************************************************************
* 功能描述: 快速排序
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-08-09 08:30:20          0.00           cxh                创建
*
*****************************************************************************/

func quickAdjust(data []int, start, end int) {
	var cur = data[start]
	var i, j int
	for i, j = start, end; i < j; {
		for i < j && data[j] >= cur {
			j--
		}
		if i < j {
			data[i] = data[j]
			i++
		}

		for i < j && data[i] <= cur {
			i++
		}
		if i < j {
			data[j] = data[i]
			j--
		}
	}
	data[i] = cur

	if i > start {
		quickAdjust(data, start, i-1)
	}

	if i < end {
		quickAdjust(data, i+1, end)
	}
}

/****************************************************************************
* 功能描述: 选择排序
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-08-09 08:30:46          0.00           cxh                创建
*
*****************************************************************************/
func selectSort(data []int) {
	var (
		minIndex = 0
		length   = len(data)
	)

	for i := 0; i < length-1; i++ {
		minIndex = i
		for j := i + 1; j < length; j++ {
			if data[j] < data[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			data[i], data[minIndex] = data[minIndex], data[i]
		}
	}
}

/****************************************************************************
* 功能描述: 调整大顶堆，索引从0开始计数
* 输入参数: start表示开始调整的索引，end则表示最后的索引
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容 
* ---------------------------------------------------------------------------
*  2017-08-08 14:29:22          0.00           cxh                创建    
* 
*****************************************************************************/

func heapAdjust(data [] int, start, end int) {
	for parent, child := start, 2*start+1; child <= end; child = 2*parent + 1 {

		if child < end && data[child+1] > data[child] {
			child++
		}

		if data[parent] >= data[child] {
			break
		}

		data[parent], data[child] = data[child], data[parent]
		parent = child
	}
}

/****************************************************************************
* 功能描述: 堆排序（大顶堆）
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-08-08 14:36:17          0.00           cxh                创建
*
*****************************************************************************/
func heapSort(data [] int) {
	var length = len(data)

	for i := (length - 1) / 2; i >= 0; i-- {
		heapAdjust(data, i, length-1)
	}

	for i := length - 1; i > 0; i-- {
		data[0], data[i] = data[i], data[0]
		heapAdjust(data, 0, i-1)
	}
}

/****************************************************************************
* 功能描述: 计数排序
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-08-08 15:03:58          0.00           cxh                创建
*
*****************************************************************************/
func countSort(data []int, k int) []int {

	var count [MAXSIZE] int
	var sorted [MAXSIZE]int
	var length = len(data)

	for i := 0; i < length; i++ {
		count[data[i]]++
	}

	for i := 1; i < MAXSIZE; i++ {
		count[i] += count[i-1]
	}

	for i := length - 1; i >= 0; i-- {
		count[data[i]]--
		sorted[count[data[i]]] = data[i]
	}
	return sorted[0:length:length]
}

/****************************************************************************
* 功能描述: 归并排序
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-08-08 15:22:04          0.00           cxh                创建
*
*****************************************************************************/

func mergeCore(data []int, start, mid, end int) {
	var (
		i, j   int
		sorted []int //切片，保存结果
	)
	for i, j = start, mid+1; i <= mid && j <= end; {
		if data[i] <= data[j] {
			sorted = append(sorted, data[i])
			i++
		} else {
			sorted = append(sorted, data[j])
			j++
		}
	}
	for i <= mid {
		sorted = append(sorted, data[i])
		i++
	}
	for j <= mid {
		sorted = append(sorted, data[j])
		j++
	}

	for i, j = 0, start; i < len(sorted); i++ {
		data[j] = sorted[i]
		j++
	} //此处不能使用copy函数

}
func mergeSort(data []int, start, end int) {
	if start < end {
		mid := (start + end) / 2
		mergeSort(data, start, mid)
		mergeSort(data, mid+1, end)
		mergeCore(data, start, mid, end)
	}
}

/****************************************************************************
* 功能描述: 求数组中除了本元素之外其他元素的乘积
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-08-09 14:54:59          0.00           cxh                创建
*
*****************************************************************************/

func multiElement(input []int) []int {
	var (
		result []int
		length = len(input)
	)

	result = append(result, 1)
	for i := 1; i < length; i++ {
		result = append(result, input[i-1]*result[i-1])
	}

	result[0] = input[length-1]
	for i := length - 2; i > 0; i-- {
		result[i] *= result[0]
		result[0] *= input[i]
	}

	return result
}

/****************************************************************************
* 功能描述: 找出所有满足其左边的数都小于它，右边的数都大于它的这样的数
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-08-09 15:24:34          0.00           cxh                创建
*
*****************************************************************************/

func findElemts(input []int) []int {
	var (
		result, assist [] int
		length         = len(input)
		curMax         = input[0]
		curMin         = input[length-1]
	)

	assist = append(assist, curMax)
	for i := 1; i < length; i++ {
		if input[i] > curMax {
			curMax = input[i]
		}
		assist = append(assist, curMax)
	}

	fmt.Println("assist:", assist)

	for i := length - 1; i >= 0; i-- {
		if curMin > input[i] {
			curMin = input[i]
		}

		if curMin >= assist[i] {
			result = append(result, input[i])
		}
	}
	return result
}

func findMoreHalfElment(input []int) int {
	var (
		count = 1
		cur   = input[0]
	)

	for i := 1; i < len(input); i++ {
		if input[i] == cur {
			count++
		} else {
			count--

			if 0 == count {
				count = 1
				cur = input[i]
			}
		}
	}
	return cur //因为题目说了存在这样一个数，所以没有检查
}

/****************************************************************************
* 功能描述: 大数相乘
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-08-09 16:08:29          0.00           cxh                创建
*
*****************************************************************************/
func bigNumMulti(multiplicand string, multiplier string) string {
	var (
		lenLier = len(multiplier)
		lenCand = len(multiplicand)
		assist  = make([]byte, lenLier+lenCand) //[]int{lenLier+lenCand-1:0}
		result  []byte
	)

	for i := 0; i < lenCand; i++ {
		for j := 0; j < lenLier; j++ {
			cur := (multiplicand[i] - '0') * (multiplier[j] - '0')
			assist[i+j+1] += cur
		}
	}

	for i := lenLier + lenCand - 1; i > 0; i-- {
		assist[i-1] += assist[i] / 10
		assist[i] = assist[i] % 10
	}

	for i, value := range assist {
		if value != 0 {
			result = append(result, assist[i]+'0')
		}
	}
	//fmt.Println(string(result))
	return string(result)

}

/****************************************************************************
* 功能描述: 找出字符串第一个只出现一次的字符
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-08-09 17:35:59          0.00           cxh                创建
*
*****************************************************************************/

func findOnceChar(input string) (result byte) {
	var hash [256]byte

	for i := 0; i < len(input); i++ {
		hash[input[i]]++
	}

	for i := 0; i < len(input); i++ {
		if hash[input[i]] == 1 {
			result = input[i]
			break
		}
	}
	return
}

/****************************************************************************
* 功能描述: 替换字符串中单个字符为某个字符串
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-08-09 17:49:30          0.00           cxh                创建
*
*****************************************************************************/

func replaceCharToString(intput string, key byte, replace string) (string) {
	count := 0
	length := len(intput)

	for i := 0; i < length; i++ {
		if intput[i] == key {
			count++
		}
	}

	newLength := length + (len(replace)-1)*count
	result := make([]byte, newLength) //因为Go中string是不可修改的，所以新建切片

	for i, j := length-1, newLength-1; i >= 0; i-- {
		if intput[i] == key {
			for k := len(replace) - 1; k >= 0; k-- {
				result[j] = replace[k]
				j--
			}
		} else {
			result[j] = intput[i]
			j--
		}
	}
	return string(result)
}

func removeRepeatChars(input string) string {
	var
	(
		result []byte
		length = len(input)
		hash   [256]bool
	)

	result = append(result, input[0])
	hash[input[0]] = true
	for i := 1; i < length; i++ {
		if false == hash[input[i]] {
			result = append(result, input[i])
			hash[input[i]] = true
		}
	}
	return string(result)
}

/****************************************************************************
* 功能描述: 寻找只出现一次的的数字
* 输入参数:
* 输出参数:
* 返 回 值:
* 其他说明:

* 修改日期                      版本号          修改人            修改内容
* ---------------------------------------------------------------------------
*  2017-08-10 09:58:12          0.00           cxh                创建
*
*****************************************************************************/
func appearOnce(input []int) (int, int) {
	var length = len(input)
	var result1, result2 int

	for i := 0; i < length; i++ {
		result1 ^= input[i]
	}

	position := findFirst1Bit(result1)

	result1 = 0
	for i := 0; i < length; i++ {
		if true == isBit1(input[i], position) {
			result1 ^= input[i]
		} else {
			result2 ^= input[i]
		}
	}
	return  result1,result2
}

func findFirst1Bit(key int) uint8 {
	bits := int(unsafe.Sizeof(key))
	var position uint8

	for i := 0; i < bits && (key&1 == 0); i++ {
		position++
		key = key >> 1
	}
	return position
}

func isBit1(key int, index uint8) bool {
	key = key >> index
	return (key & 1) == 1
}

func TestSort() {
	data := []int{0, 1, 2, 1, 3,3}
	fmt.Println(data)

	//selectSort(data)
	//bubbleSort(data)
	//quickAdjust(data, 0, len(data)-1)
	//heapSort(data)
	//mergeSort(data, 0, len(data)-1)
	//result := multiElement(data)
	//result := findElemts(data)
	//result := findMoreHalfElment(data)
	//fmt.Println(bigNumMulti("12", "12"))
	//fmt.Println(findOnceChar("133di229fm"))
	//fmt.Println(replaceCharToString("chenccxiaocc", 'c', "mei"))
	//fmt.Println(removeRepeatChars("cdhencchdxiaoccxi"))
	//fmt.Println(result)
	fmt.Println(appearOnce(data))
	fmt.Println(data)

	var xx =45

	fmt.Println(xx)
}
