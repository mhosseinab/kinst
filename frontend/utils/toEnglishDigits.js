const toEnglishDigits = function(valau) {
  const charCodeZero = '۰'.charCodeAt(0);
  return parseInt(
    valau.replace(/[۰-۹]/g, function(w) {
      return w.charCodeAt(0) - charCodeZero;
    }),
    10,
  );
};

export default toEnglishDigits;
