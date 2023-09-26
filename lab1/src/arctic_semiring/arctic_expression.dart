import 'arctic_operand.dart';

class ArcticExpression {
  ArcticOperand lhs = ArcticOperand();
  ArcticOperand rhs = ArcticOperand();

  ArcticExpression.fromOperands(String lhs, String rhs) {
    this.lhs = ArcticOperand.fromOperand(lhs);
    this.rhs = ArcticOperand.fromOperand(rhs);
  }

  ArcticExpression();

  @override
  String toString() {
    List<String> funs = [];

    for (var i = 0; i <= 1; i++) {
      for (var j = 0; j <= 1; j++) {
        funs.add(
            "(assert (arc_gt ${lhs.mulOperand.data[i][j]} ${rhs.mulOperand.data[i][j]}))");
      }
    }

    for (var i = 0; i <= 1; i++) {
      funs.add(
          "(assert (arc_gt ${lhs.addOperand.data[i][0]} ${rhs.addOperand.data[i][0]}))");
    }

    return funs.join('\n') + "\n";
  }
}
