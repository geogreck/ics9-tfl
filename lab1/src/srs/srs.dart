/// I guess i need:
/// 1. Parse input for Unique Strings and Expressions
/// 2. Calculate expression composition for lhs and rhs
/// 3. Dump it all with by-component comparasion
/// 4. Write application for academic leave

class SRS {
  Set<UniqueString> uniqueStrings = {};

  List<Expression> expressions = [];

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

  ArcticOperand.fromOperand(String operand) {
    // TODO: implement fromOperand
    mulOperand = [
      ['a1', 'a2'],
      ['a3', 'a4'],
    ];
    addOperand = [
      ['b1'],
      ['b2']
    ];
  }

  ArcticOperand();

  @override
  String toString() {
    // TODO: implement toString
    return super.toString();
  }
}

class Expression {
  ArcticOperand lhs = ArcticOperand();
  ArcticOperand rhs = ArcticOperand();

  Expression.fromOperands(String lhs, String rhs) {
    this.lhs = ArcticOperand.fromOperand(lhs);
    this.rhs = ArcticOperand.fromOperand(rhs);
  }

  Expression();

  @override
  String toString() {
    List<String> funs = [];

    for (var i = 0; i <= 1; i++) {
      for (var j = 0; j <= 1; j++) {
        funs.add("(assert (arc_gt (${lhs.mulOperand[i][j]}) (${rhs.mulOperand[i][j]})))");
      }
    }

    for (var i = 0; i <= 1; i++) {
      funs.add("(assert (arc_gt (${lhs.addOperand[i][0]}) (${rhs.addOperand[i][0]})))");
    }

    return funs.join('\n') + "\n";
  }
}
