import { generateRandomIdx } from "@/utils/encipherment.js";

// 单元格坐标
interface CellCoordinate {
  row: number;
  column: number;
}

// 表格分组
const redGroup = new Set([1, 2, 3, 4, 5, 9, 13]);
const blueGroup = new Set([6, 7, 8, 10, 14]);
const yellowGroup = new Set([11, 12, 15]);

// 获取单元格的类名
export const getCellClass = (cellValue: number) => {
  if (redGroup.has(cellValue)) return "bg-red-4 shadow";
  if (blueGroup.has(cellValue)) return "bg-blue-4 shadow";
  if (yellowGroup.has(cellValue)) return "bg-yellow-4 shadow";
  return "";
};

// 创建游戏地图哈希表
export const createHashMap = (map: number[][]) => {
  const hashMap = new Map<number, CellCoordinate>();
  for (let i = 0; i < map.length; i++) {
    for (let j = 0; j < map[i].length; j++) {
      hashMap.set(map[i][j], {
        row: i,
        column: j,
      });
    }
  }
  return hashMap;
};

// 根据打乱字符串获取游戏地图
export const getGameMapByScramble = (scramble: string) => {
  // 将字符串拆分成数组
  const scrambleArray = scramble.split(",");

  // 将一维数组转换为二维数组
  const rows = 4;
  const cols = 4;
  const gameMap = new Array(rows)
    .fill(null)
    .map((_, i) =>
      scrambleArray
        .slice(i * cols, (i + 1) * cols)
        .map((value) => (value !== "" ? parseInt(value) : null))
    );

  return gameMap;
};

// 点击单元格
export const handleClick = (
  map: number[][],
  hashMap: Map<number, CellCoordinate>,
  row?: number,
  column?: number,
  value?: number
) => {
  let newSolution = "";
  let newStepCount = 0;

  // 获取点击单元格的坐标
  if (!row && !column && value) {
    row = hashMap.get(value)!.row;
    column = hashMap.get(value!)!.column;
  }

  // 获取点击单元格的值
  const item = map[row!][column!];
  // 点击的是空白单元格
  if (item === 0)
    return {
      newSolution,
      newStepCount,
      map,
      hashMap,
    };

  // 获取空白单元格的坐标
  const nullRow = hashMap.get(0)!.row;
  const nullColumn = hashMap.get(0)!.column;

  // 同一行
  if (row == nullRow) {
    // 点击的单元格在空白单元格的右边
    if (column! > nullColumn) {
      // 保存移动的单元格
      newSolution = item.toString();
      // 步数加一
      newStepCount++;
      for (let i = nullColumn; i < column!; i++) {
        map[row][i] = map[row][i + 1];
        hashMap.set(map[row][i + 1], {
          row: row,
          column: i,
        });
      }
    }
    // 点击的单元格在空白单元格的左边
    else if (column! < nullColumn) {
      // 保存移动的单元格
      newSolution = item.toString();
      // 步数加一
      newStepCount++;
      for (let i = nullColumn; i > column!; i--) {
        map[row][i] = map[row][i - 1];
        hashMap.set(map[row][i - 1], {
          row: row,
          column: i,
        });
      }
    }
    map[row][column!] = 0;
    hashMap.set(0, { row: row, column: column! });
  }

  // 同一列
  if (column == nullColumn) {
    // 点击的单元格在空白单元格的下边
    if (row! > nullRow) {
      // 保存移动的单元格
      newSolution = item.toString();
      // 步数加一
      newStepCount++;
      for (let i = nullRow; i < row!; i++) {
        map[i][column] = map[i + 1][column];
        hashMap.set(map[i + 1][column], {
          row: i,
          column: column,
        });
      }
    }
    // 点击的单元格在空白单元格的上边
    else if (row! < nullRow) {
      // 保存移动的单元格
      newSolution = item.toString();
      // 步数加一
      newStepCount++;
      for (let i = nullRow; i > row!; i--) {
        map[i][column] = map[i - 1][column];
        hashMap.set(map[i - 1][column], {
          row: i,
          column: column,
        });
      }
    }
    map[row!][column] = 0;
    hashMap.set(0, { row: row!, column: column });
  }

  return {
    newSolution,
    newStepCount,
    map,
    hashMap,
  };
};

// 判断是否可解
export const isSolvable = (numsMap: number[]) => {
  let sum = 0;
  for (let i = 0; i < 16; i++) {
    if (numsMap[i] === 0) {
      sum += Math.floor(i / 4) + ((i + 1) % 4);
      continue;
    }
    for (let j = 0; j < 16 - i; j++) {
      if (numsMap[j + i] < numsMap[i]) {
        sum++;
      }
    }
  }
  return sum % 2 === 0;
};

// 根据idx获取n阶排列
export const setNPerm = (n: number, idx: number, ensureEven: boolean) => {
  const N = n * n;

  const arr = Array.from({ length: N }, (_, index) => index + 1);
  arr[N - 1] = 0;

  const even = ensureEven ? 1 : 0;
  let parity = 0;

  arr[N - 1] = even;

  if (ensureEven) {
    arr[N - 2] = 0;
  }

  for (let i = N - 2 - even; i >= 0; i--) {
    arr[i] = idx % (N - i);
    parity ^= arr[i];
    idx = Math.floor(idx / (N - i));

    for (let j = i + 1; j < N; j++) {
      if (arr[j] >= arr[i]) {
        arr[j]++;
      }
    }
  }

  if (ensureEven && (parity & 1) !== 0) {
    const temp = arr[N - 1];
    arr[N - 1] = arr[N - 2];
    arr[N - 2] = temp;
  }

  // 判断是否可解
  if (!isSolvable(arr)) {
    // 如果不可解, 交换最后两个元素, 保证可解.但若其中一个元素为null, 则交换倒数第三个元素
    if (arr[N - 1] === null) {
      const temp = arr[N - 2];
      arr[N - 2] = arr[N - 3];
      arr[N - 3] = temp;
    }

    const temp = arr[N - 1];
    arr[N - 1] = arr[N - 2];
    arr[N - 2] = temp;
  }

  return arr;
};

// 获取打乱数组
export const getScramble = (n: number, timestampe: number) => {
  const randomIdx = generateRandomIdx(timestampe);

  const scramble = setNPerm(n, randomIdx, true);

  return { scramble, randomIdx };
};
