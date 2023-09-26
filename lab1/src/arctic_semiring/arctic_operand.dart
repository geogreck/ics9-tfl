import 'arctic_matrix.dart';

class ArcticOperand {
  ArcticMatrix mulOperand = ArcticMatrix();

  ArcticMatrix addOperand = ArcticMatrix();

  ArcticOperand.fromOperand(String operand) {
    // TODO: implement fromOperand
    String lastChar = operand.substring(operand.length - 1);

    mulOperand = arcticMulMatrixForChar(lastChar);
    addOperand = arcticAddMatrixForChar(lastChar);

    for (var i = operand.length - 2; i >= 0; i--) {
      String char = operand.substring(i, i + 1);
      ArcticMatrix curMulOperand = arcticMulMatrixForChar(char);
      ArcticMatrix curAddOperand = arcticAddMatrixForChar(char);

      mulOperand = mulOperand * curMulOperand;
      addOperand = mulOperand * curAddOperand + addOperand;
    }
  }

  ArcticOperand();

  @override
  String toString() {
    // TODO: implement toString
    return mulOperand.toString() + '\n' + addOperand.toString();
  }
}
