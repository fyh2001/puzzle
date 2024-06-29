export const scientificToDecimal = (num: number) => {
  // 判断输入是否为科学计数法表示的数字
  var n = String(num);
  if (/\d+\.?\d*e[\+\-]*\d+/i.test(n)) {
    var numStr = parseFloat(n).toString();
    var x = numStr.indexOf("e");
    var base = numStr.slice(0, x);
    var exponent = numStr.slice(x + 1);
    var result = "";

    if (parseFloat(base) === 0) {
      result = "0";
    } else {
      result = base.replace(".", "");
      var decimalPos = result.length + parseInt(exponent);
      if (decimalPos <= 0) {
        result = "0." + "0".repeat(Math.abs(decimalPos)) + result;
      } else {
        if (decimalPos >= result.length) {
          result = result + "0".repeat(decimalPos - result.length);
        } else {
          result = result.slice(0, decimalPos) + "." + result.slice(decimalPos);
        }
      }
    }

    return result;
  } else {
    return num;
  }
};
