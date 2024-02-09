package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// EncryptionParams 加密参数
type EncryptionParams struct {
	Dimension int    `json:"dimension"` // 阶数
	RandomIdx int64  `json:"ri"`        // 随机idx
	StepCount int    `json:"stepCount"` // 步数
	Scramble  string `json:"scramble"`  // 打乱
	Solution  string `json:"solution"`  // 解法
}

var salt = "defo1215_puzzle"

// GenerateRandomIdx 生成随机idx
func generateRandomIdx(timestamp int64) int64 {
	// 构造一个字符串，包含时间戳和(n-i)!至(n-1)!的相关信息
	inputString := fmt.Sprintf("%v%s", timestamp, salt)

	// 进行SHA-256加密
	sha256Result := sha256.Sum256([]byte(inputString))

	// 再次进行SHA-256加密
	sha256Result2 := sha256.Sum256(sha256Result[:])

	// 将字节数组转为16进制字符串
	truncatedHash := hex.EncodeToString(sha256Result2[:])[:13]

	// 将加密后的结果转为整数(16进制)
	encryptedInteger, err := strconv.ParseInt(truncatedHash, 16, 64)
	if err != nil {
		fmt.Println("Error parsing hex string:", err)
		return 0
	}

	return encryptedInteger
}

func isSolvable(n int, numsMap []int) bool {
	sum := 0

	// 若为偶数阶,需要考虑空格位置
	if n%2 == 0 {
		for i := 0; i < n*n; i++ {
			if numsMap[i] == 0 {
				sum += int(i/n) + ((i + 1) % n)
				continue
			}
			for j := 0; j < n*n-i; j++ {
				if numsMap[j+i] < numsMap[i] {
					sum++
				}
			}
		}
	} else {
		// 若为奇数阶
		for i := 0; i < n*n; i++ {
			for j := 0; j < n*n-i; j++ {
				if numsMap[j+i] < numsMap[i] && numsMap[j+i] != 0 && numsMap[i] != 0 {
					sum++
				}
			}
		}
	}

	return sum%2 == 0
}

// setNPerm 根据idx获取n阶排列
func setNPerm(idx, n int, ensureEven bool) []int {

	N := n * n

	arr := make([]int, N)

	even := 0
	if ensureEven {
		even = 1
	}

	parity := 0
	arr[N-1] = even

	if ensureEven {
		arr[N-2] = 0
	}

	for i := N - 2 - even; i >= 0; i-- {
		if idx%(N-i) == 0 {
			arr[i] = 0
		} else {
			arr[i] = idx % (N - i)
		}
		parity ^= arr[i]
		idx /= N - i

		for j := i + 1; j < N; j++ {
			if arr[j] >= arr[i] {
				arr[j]++
			}
		}
	}

	if ensureEven && (parity&1) != 0 {
		temp := arr[N-1]
		arr[N-1] = arr[N-2]
		arr[N-2] = temp
	}

	// 判断是否可解
	if !isSolvable(n, arr) {
		// 如果不可解, 交换最后两个元素, 保证可解.但若其中一个元素为0, 则交换倒数第三个元素
		if arr[N-1] == 0 {
			temp := arr[N-2]
			arr[N-2] = arr[N-3]
			arr[N-3] = temp
		}

		temp := arr[N-1]
		arr[N-1] = arr[N-2]
		arr[N-2] = temp
	}

	return arr
}

// Shuffle 打乱
func Shuffle(n, idx int) []int {
	N := n * n
	currentIndex := N
	randomIndex := 0

	array := make([]int, N)
	for i := 0; i < N; i++ {
		array[i] = i + 1
	}
	array[N-1] = 0

	seededRandom := func(min, max int) int {
		x := math.Sin(float64(idx)) * 10000
		idx++
		return int((x-math.Floor(x))*float64(max-min+1)) + min
	}

	for currentIndex != 0 {
		randomIndex = seededRandom(0, currentIndex-1)
		currentIndex--
		array[currentIndex], array[randomIndex] = array[randomIndex], array[currentIndex]
	}

	if !isSolvable(n, array) {
		if array[N-1] == 0 {
			temp := array[N-2]
			array[N-2] = array[N-3]
			array[N-3] = temp
		}
		temp := array[N-1]
		array[N-1] = array[N-2]
		array[N-2] = temp
	}

	return array
}

// CreateScramble 生成打乱
func CreateScramble() (int64, string) {
	// 生成随机idx
	randomIdx := generateRandomIdx(time.Now().UnixMilli())
	// 生成打乱字符串
	scramble := setNPerm(int(randomIdx), 16, true)
	scrambleStr := strings.Trim(strings.Replace(fmt.Sprint(scramble), " ", ",", -1), "[]")
	scrambleStr = strings.Replace(scrambleStr, ",0,", ",,", -1)
	re := regexp.MustCompile(`(?:^|,)0(?:,|$)`)
	scrambleStr = re.ReplaceAllString(scrambleStr, "${1},${2}")

	return randomIdx, scrambleStr
}

// getGameMapByScrambleHandler 根据打乱字符串获取游戏地图
func getGameMapByScrambleHandler(n int, scramble string) [][]int {
	// 将字符串拆分成数组
	scrambleArray := strings.Split(scramble, ",")

	// 定义行数和列数
	rows := n
	cols := n

	// 创建二维数组
	gameMap := make([][]int, rows)
	for i := range gameMap {
		gameMap[i] = make([]int, cols)
	}

	// 将一维数组转换为二维数组
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			index := i*cols + j
			if index < len(scrambleArray) {
				// 将字符串转换为整数，如果为空则为零值
				value, err := strconv.Atoi(scrambleArray[index])
				if err == nil {
					gameMap[i][j] = value
				}
			}
		}
	}

	return gameMap
}

// getSolutionHandler 根据解法字符串获取解法
func getSolutionHandler(solutionStr string) []int {
	// 使用 strings.Split 将字符串分割成字符串切片
	strValues := strings.Split(solutionStr, ",")

	// 初始化整数切片
	intValues := make([]int, len(strValues))

	// 将字符串切片中的每个元素转换为整数
	for i, strValue := range strValues {
		// 使用 strconv.Atoi 进行字符串到整数的转换
		intValue, err := strconv.Atoi(strValue)
		if err != nil {
			return nil
		}
		intValues[i] = intValue
	}

	return intValues
}

// createMap 创建游戏地图哈希表
func createHashMap(arr [][]int) map[int]map[string]int {
	hashMap := make(map[int]map[string]int)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			element := arr[i][j]
			if _, ok := hashMap[element]; !ok {
				hashMap[element] = make(map[string]int)
			}
			hashMap[element]["row"] = i
			hashMap[element]["column"] = j
		}
	}

	return hashMap
}

// ClickRules 点击规则
func ClickRules(gameMap [][]int, value int) {

	newGameHashMap := createHashMap(gameMap)

	row := newGameHashMap[value]["row"]
	column := newGameHashMap[value]["column"]

	// 获取点击单元格的值
	item := gameMap[row][column]
	// 点击的是空白单元格
	if item == 0 {
		return
	}

	// 获取空白单元格的坐标
	nullRow := newGameHashMap[0]["row"]
	nullColumn := newGameHashMap[0]["column"]

	// 同一行
	if row == nullRow {
		// 点击的单元格在空白单元格的右边
		if column > nullColumn {
			for i := nullColumn; i < column; i++ {
				gameMap[row][i] = gameMap[row][i+1]
				newGameHashMap[gameMap[row][i+1]] = map[string]int{
					"row":    row,
					"column": i,
				}
			}
		} else if column < nullColumn { // 点击的单元格在空白单元格的左边
			for i := nullColumn; i > column; i-- {
				gameMap[row][i] = gameMap[row][i-1]
				newGameHashMap[gameMap[row][i-1]] = map[string]int{
					"row":    row,
					"column": i,
				}
			}
		}

		gameMap[row][column] = 0
		newGameHashMap[0] = map[string]int{"row": row, "column": column}
	}

	// 同一列
	if column == nullColumn {
		// 点击的单元格在空白单元格的下边
		if row > nullRow {
			for i := nullRow; i < row; i++ {
				gameMap[i][column] = gameMap[i+1][column]
				newGameHashMap[gameMap[i+1][column]] = map[string]int{
					"row":    i,
					"column": column,
				}
			}
		} else if row < nullRow { // 点击的单元格在空白单元格的上边
			for i := nullRow; i > row; i-- {
				gameMap[i][column] = gameMap[i-1][column]
				newGameHashMap[gameMap[i-1][column]] = map[string]int{
					"row":    i,
					"column": column,
				}
			}
		}

		gameMap[row][column] = 0
		newGameHashMap[0] = map[string]int{"row": row, "column": column}
	}
}

// checkCondition 判断是否完成
func checkCondition(gameMapValue [][]int) bool {

	flatMap := make([]int, 0)
	for _, row := range gameMapValue {
		flatMap = append(flatMap, row...)
	}

	for index, item := range flatMap {
		if index >= len(flatMap)-1 {
			break
		}
		if item != index+1 {
			return false
		}
	}

	return true
}

// VerifyScramble 校验函数
func (ep EncryptionParams) VerifyScramble() bool {
	n := ep.Dimension

	// 生成打乱字符串
	scramble := Shuffle(n, int(ep.RandomIdx))
	scrambleStr := strings.Trim(strings.Replace(fmt.Sprint(scramble), " ", ",", -1), "[]")

	if scrambleStr != ep.Scramble {
		fmt.Println("scramble错误!")
		fmt.Println("scramble:   ", scrambleStr)
		fmt.Println("ep.Scramble:   ", ep.Scramble)
		return false
	}

	scrambleMap := getGameMapByScrambleHandler(n, ep.Scramble)
	solutionList := getSolutionHandler(ep.Solution)

	if len(solutionList) != ep.StepCount {
		fmt.Println("stepCount错误!")
		fmt.Println("stepCount:   ", len(solutionList))
		fmt.Println("ep.StepCount:   ", ep.StepCount)
		return false
	}

	// 根据解法字符串获取解法
	for _, value := range solutionList {
		ClickRules(scrambleMap, value)
	}

	// 判断是否完成
	isFinish := checkCondition(scrambleMap)
	if !isFinish {
		fmt.Println("未完成!")
		return false
	}

	return true
}
