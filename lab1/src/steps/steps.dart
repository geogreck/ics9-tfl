// Workflow steps

import 'dart:io';

import '../arctic_semiring/arctic_expression.dart';
import '../srs/srs.dart';
import '../util/util.dart';
import '../util/validators.dart';

List<String> ReadInput() {
  List<String> lines = [];
  for (String? line = stdin.readLineSync();
      line != null && line != '';
      line = stdin.readLineSync()) {
    lines.add(line);
  }
  return lines;
}

/// Valid input consists of string in format:
/// [a-zA-Z]+->[a-zA-Z]+
///
/// For example:
/// fg -> g
bool ValidateInput(List<String> input) {
  RegExp exp = RegExp("[a-zA-Z]+->[a-zA-Z]+");

  for (var line in input) {
    bool valid = exp.hasMatch(line);
    if (!valid) {
      return false;
    }
  }
  return true;
}

File CreateFile(String fileName) {
  File file = File(fileName);
  return file;
}

/// Dump initial z3 program part to specified file
void WriteInitialPart(File f) {
  WriteFile(f, "(set-logic QF_NIA)\n\n");

  AppendFile(f, "(define-fun max ((x Int) (y Int)) Int (ite (> x y) x y))");

  AppendFile(f, "(define-fun arc_add ((x Int) (y Int)) Int (max x y))");
  AppendFile(f,
      "(define-fun arc_mul ((a Int) (b Int)) Int (ite (or (= a -1) (= b -1)) -1 (+ a b)))");

  AppendFile(f,
      "(define-fun arc_gt ((x Int) (y Int)) Bool (ite (or (> x y) (and (= x y) (= x -1))) true false))");
  AppendFile(f,
      "(define-fun arc_ge ((x Int) (y Int)) Bool (ite (or (>= x y) (and (= x y) (= x -1))) true false))\n\n");
}

void WriteFinalPart(File f) {
  AppendFile(f, "(check-sat)");
  AppendFile(f, "(get-model)");
  AppendFile(f, "(exit)");
}

SRS ParseSRS(List<String> input) {
  SRS srs = SRS();

  for (var line in input) {
    int ind = line.indexOf('->');
    srs.expressions.add(ArcticExpression.fromOperands(
        line.substring(0, ind), line.substring(ind + 2)));

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
