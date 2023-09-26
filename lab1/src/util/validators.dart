bool isAlpha(int code) {
  return (code <= 'z'.codeUnitAt(0) && code >= 'a'.codeUnitAt(0) || code <= 'Z'.codeUnitAt(0) && code >= 'A'.codeUnitAt(0));
}
