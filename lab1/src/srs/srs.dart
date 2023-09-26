/// I guess i need:
/// 1. Parse input for Unique Strings and Expressions
/// 2. Calculate expression composition for lhs and rhs
/// 3. Dump it all with by-component comparasion
/// 4. Write application for academic leave


class SRS {
  Set<UniqueString> uniqueStrings = {};

  Expression finalExpression = Expression();

  @override
  String toString() {
    // TODO: implement toString
    return uniqueStrings.join('\n') + '\n';
  }
}

class UniqueString {
  String string = "";
  
  UniqueString(String string) {
    this.string  = string;
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

    return funs.join('\n') + "\n";
  }

  @override
  bool operator ==(covariant UniqueString rhs) {
    return (rhs.string == string);
  }

  @override
  int get hashCode => string.hashCode;
}

class ArcticOperand {
  /// How to limit it to 2x2 size?
  List<List<String>> mulOperand = [];

  /// How to limit it to 2x1 size?
  List<List<String>> addOperand = [];

  @override
  String toString() {
    // TODO: implement toString
    return super.toString();
  }
}

class Expression {
  ArcticOperand lhs = ArcticOperand();
  ArcticOperand rhs = ArcticOperand();
}
