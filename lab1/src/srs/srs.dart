import '../arctic_semiring/arctic_expression.dart';

/// I guess i need:
/// 1. Parse input for Unique Strings and Expressions
/// 2. Calculate expression composition for lhs and rhs
/// 3. Dump it all with by-component comparasion
/// 4. Write application for academic leave

class SRS {
  Set<UniqueString> uniqueStrings = {};

  List<ArcticExpression> expressions = [];

  @override
  String toString() {
    String uniqueStringsSmt = uniqueStrings.join('\n') + '\n';
    String expressionsSmt = expressions.join('\n') + '\n';
    return uniqueStringsSmt + expressionsSmt;
  }
}

class UniqueString {
  String string = "";

  UniqueString(String string) {
    this.string = string;
  }

  @override
  String toString() {
    List<String> funs = [];

    for (var i = 1; i <= 2; i++) {
      // Elements of string's vector sum
      funs.add("(declare-fun ${string}_${i} () Int)");
      for (var j = 1; j <= 2; j++) {
        // Elements of string's matrix by multiplication
        funs.add("(declare-fun ${string}_${i}${j} () Int)");
      }
    }

    for (var i = 1; i <= 2; i++) {
      if (i == 1) {
        funs.add("(assert (> ${string}_${i} -1))");
      } else {
        funs.add("(assert (>= ${string}_${i} -1))");
      }
      for (var j = 1; j <= 2; j++) {
        if (j == 1 && i == 1) {
          funs.add("(assert (> ${string}_${i}${j} -1))");
        } else {
          funs.add("(assert (>= ${string}_${i}${j} -1))");
        }
      }
    }

    return funs.join('\n') + "\n";
  }

  @override
  bool operator ==(covariant UniqueString rhs) {
    return (rhs.string == string);
  }

  @override
  int get hashCode => string.hashCode;
}
