import CryptoJS from "crypto-js";

const privateKey = {
  "001": "defo1215_puzzle",
};

// 生成随机值
export const generateRandomIdx = (timestamp: number, n: number) => {
  const inputString = timestamp.toString() + privateKey["001"];

  // 进行SHA-256加密
  const sha256Result = CryptoJS.SHA256(inputString);

  // 进行SHA-256加密
  const sha256Result2 = CryptoJS.SHA256(sha256Result);

  // 截取前13个字符
  let truncatedHash = sha256Result2.toString(CryptoJS.enc.Hex);
  // truncatedHash = truncatedHash.substring(0, 13);
  if (n === 4) {
    truncatedHash = truncatedHash.substring(0, 13);
  } else if (n === 5) {
    truncatedHash = truncatedHash.substring(0, 20);
  } else if (n === 6 || n === 3) {
    truncatedHash = truncatedHash.substring(0, 33);
  }
  // 将加密后的结果转为整数(16进制)
  const encryptedInteger = parseInt(truncatedHash, 16);

  return encryptedInteger;
};
