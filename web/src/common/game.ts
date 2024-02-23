// 单元格坐标
interface CellCoordinate {
  row: number;
  column: number;
}

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
export const getGameMapByScramble = (
  scramble: string,
  dimension: number
): number[][] => {
  // 将字符串拆分成数组
  const scrambleArray = scramble.split(",");

  // 将一维数组转换为二维数组
  const rows = dimension;
  const cols = dimension;
  const gameMap = new Array(rows)
    .fill(0)
    .map((_, i) =>
      scrambleArray
        .slice(i * cols, (i + 1) * cols)
        .map((value) => (value !== "" ? parseInt(value) : 0))
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
  const newMap = map.map((row) => row.slice());

  // 获取点击单元格的坐标
  if (!row && !column && value) {
    row = hashMap.get(value)!.row;
    column = hashMap.get(value!)!.column;
  }

  // 获取点击单元格的值
  const item = newMap[row!][column!];
  // 点击的是空白单元格
  if (item === 0)
    return {
      newSolution,
      newStepCount,
      newMap,
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
        newMap[row][i] = newMap[row][i + 1];
        hashMap.set(newMap[row][i + 1], {
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
        newMap[row][i] = newMap[row][i - 1];
        hashMap.set(newMap[row][i - 1], {
          row: row,
          column: i,
        });
      }
    }
    newMap[row][column!] = 0;
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
        newMap[i][column] = newMap[i + 1][column];
        hashMap.set(newMap[i + 1][column], {
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
        newMap[i][column] = newMap[i - 1][column];
        hashMap.set(newMap[i - 1][column], {
          row: i,
          column: column,
        });
      }
    }
    newMap[row!][column] = 0;
    hashMap.set(0, { row: row!, column: column });
  }

  return {
    newSolution,
    newStepCount,
    newMap,
    hashMap,
  };
};

// 判断是否可解
export const isSolvable = (n: number, numsMap: number[]) => {
  let sum = 0;

  // 若为偶数阶,需要考虑空格位置
  if (n % 2 === 0) {
    for (let i = 0; i < n * n; i++) {
      if (numsMap[i] === 0) {
        sum += Math.floor(i / n) + ((i + 1) % n);
        continue;
      }
      for (let j = 0; j < n * n - i; j++) {
        if (numsMap[j + i] < numsMap[i]) {
          sum++;
        }
      }
    }
  } else {
    // 若为奇数阶
    for (let i = 0; i < n * n; i++) {
      for (let j = 0; j < n * n - i; j++) {
        if (
          numsMap[j + i] < numsMap[i] &&
          numsMap[j + i] !== 0 &&
          numsMap[i] !== 0
        ) {
          sum++;
        }
      }
    }
  }

  return sum % 2 === 0;
};

//  Fisher-Yates洗牌算法
export const shuffle = (n: number, idx: number) => {
  const N = n * n;

  let currentIndex = N,
    randomIndex;

  const array = Array.from({ length: n * n }, (_, index) => index + 1);
  array[n * n - 1] = 0;

  // 设置随机数种子
  const seededRandom = (min: number, max: number) => {
    let x = Math.sin(idx++) * 10000;
    return Math.floor((x - Math.floor(x)) * (max - min + 1)) + min;
  };

  // 当还有未洗牌的元素时
  while (currentIndex !== 0) {
    // 随机选取一个未洗牌的元素
    randomIndex = seededRandom(0, currentIndex - 1);
    currentIndex--;

    // 交换它与当前元素
    [array[currentIndex], array[randomIndex]] = [
      array[randomIndex],
      array[currentIndex],
    ];
  }

  // 判断是否可解
  if (!isSolvable(n, array)) {
    // 如果不可解, 交换最后两个元素, 保证可解.
    if (array[N - 1] !== 0 && array[N - 2] !== 0) {
      [array[N - 1], array[N - 2]] = [array[N - 2], array[N - 1]];
    } else {
      [array[N - 4], array[N - 3]] = [array[N - 3], array[N - 4]];
    }
  }

  return array;
};

// 根据阶数将一维数组转换为二维数组
export const transformArray = (arr: number[], n: number) => {
  const res = [];
  for (let i = 0; i < n; i++) {
    res.push(arr.slice(i * n, (i + 1) * n));
  }
  return res;
};

// 根据idx获取n阶排列(已弃用)
// export const setNPerm = (n: number, idx: number, ensureEven: boolean) => {
//   const N = n * n;

//   const arr = Array.from({ length: N }, (_, index) => index + 1);
//   arr[N - 1] = 0;

//   const even = ensureEven ? 1 : 0;
//   let parity = 0;

//   arr[N - 1] = even;

//   if (ensureEven) {
//     arr[N - 2] = 0;
//   }

//   for (let i = N - 2 - even; i >= 0; i--) {
//     arr[i] = idx % (N - i);
//     parity ^= arr[i];
//     idx = Math.floor(idx / (N - i));

//     for (let j = i + 1; j < N; j++) {
//       if (arr[j] >= arr[i]) {
//         arr[j]++;
//       }
//     }
//   }

//   if (ensureEven && (parity & 1) !== 0) {
//     const temp = arr[N - 1];
//     arr[N - 1] = arr[N - 2];
//     arr[N - 2] = temp;
//   }

//   // 判断是否可解
//   if (!isSolvable(n, arr)) {
//     // 如果不可解, 交换最后两个元素, 保证可解.但若其中一个元素为0, 则交换倒数第三个元素
//     if (arr[N - 1] === 0) {
//       const temp = arr[N - 2];
//       arr[N - 2] = arr[N - 3];
//       arr[N - 3] = temp;
//     }

//     const temp = arr[N - 1];
//     arr[N - 1] = arr[N - 2];
//     arr[N - 2] = temp;
//   }

//   return arr;
// };
