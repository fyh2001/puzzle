import CryptoJS from "crypto-js";

const privateKey = {
  "001": "defo1215_puzzle",
};

// 生成随机值
export const generateRandomIdx = (timestamp: number) => {
  const inputString = timestamp.toString() + privateKey["001"];

  // 进行SHA-256加密
  const sha256Result = CryptoJS.SHA256(inputString);

  // 进行SHA-256加密
  const sha256Result2 = CryptoJS.SHA256(sha256Result);

  // 截取前13个字符
  const truncatedHash = sha256Result2
    .toString(CryptoJS.enc.Hex)
    .substring(0, 13);

  // 将加密后的结果转为整数(16进制)
  const encryptedInteger = parseInt(truncatedHash, 16);

  return encryptedInteger;
};
