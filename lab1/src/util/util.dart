import 'dart:io';

/// I am surely noob and don't know if such thing exists
void AppendFile(File f, String str) {
  if (!str.endsWith("\n")) {
    str = str + "\n";
  }
  f.writeAsStringSync(str, mode: FileMode.writeOnlyAppend);
}

void WriteFile(File f, String str) {
  if (!str.endsWith("\n")) {
    str = str + "\n";
  }
  f.writeAsStringSync(str, mode: FileMode.writeOnly);
}
