// Workflow steps

import 'dart:io';

import '../srs/srs.dart';
import '../util/util.dart';
import '../util/validators.dart';

List<String> ReadInput() {
  List<String> lines = [];
  for (String? line = stdin.readLineSync(); line != null && line != ''; line = stdin.readLineSync()) {
    lines.add(line);
  }
  return lines;
}

/// Valid input consists of string in format:
/// string->string
///
/// For example:
/// fg -> g
bool ValidateInput(List<String> input) {
  //TODO: Yep, it surely does nothing
  return true;
}

File CreateFile(String fileName) {
  File file = File(fileName);
  return file;
}

/// Dump initial z3 program part to specified file
void WriteInitialPart(File f) {
  WriteFile(f, "(set-logic QF_NIA)\n\n");

  AppendFile(f, "(define-fun max ((a Int) (b Int)) Int (ite (> x y) x y))");

  AppendFile(f, "(define-fun arc_add ((a Int) (b Int)) Int (max a b))");
  AppendFile(f, "(define-fun arc_mul ((a Int) (b Int)) Int (max a b))"); // TODO: invent me!!

  AppendFile(
      f, "(define-fun arc_gt ((a Int) (b Int)) Bool (ite (or (> x y) (and (= x y) (= x -1))) true false))");
  AppendFile(f,
      "(define-fun arc_ge ((a Int) (b Int)) Bool (ite (or (>= x y) (and (= x y) (= x -1))) true false))\n\n");
}

void WriteFinalPart(File f) {
  AppendFile(f, "(get-model)");
  AppendFile(f, "(check-sat)");
  AppendFile(f, "(exit)");
}

SRS ParseSRS(List<String> input) {
  SRS srs = SRS();

  for (var line in input) {
    int ind = line.indexOf('->');
    srs.expressions.add(Expression.fromOperands(line.substring(0, ind),line.substring(ind+2)));
    print(line.substring(0, ind));
    print(line.substring(ind + 2));
    for (var rune in line.runes) {
      if (isAlpha(rune)) {
        String char = new String.fromCharCode(rune);
        srs.uniqueStrings.add(UniqueString(char));
      }
    }
  }

  return srs;
}


void DumpSRSToFile(File f, SRS srs) {
  AppendFile(f, srs.toString());
}
