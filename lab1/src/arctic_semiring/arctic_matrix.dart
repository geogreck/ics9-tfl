import 'arctic_member.dart';

/// a11 a12 * (b11 b12 * x + b1 ) + a1
/// a21 a22   (b21 b22       b2 )   a2
///
/// a11 a12 * b11 b12 * x + a11 a12 * b1 + a1
/// a21 a22   b21 b22       a21 a22   b2   a2
///
/// a11 a12 * b11 b12   ====   (arc_sum (arc_mul a11 b11) (arc_mul a12 b21)) ...
/// a21 a22   b21 b22                      ....                              ...

class ArcticMatrix {
  List<List<ArcticMember>> data = [];

  ArcticMatrix();

  ArcticMatrix.fromSize(int m, int n) {
    for (var i = 0; i < m; i++) {
      data.add([]);
      for (var j = 0; j < n; j++) {
        data[i].add(ArcticMember());
      }
    }
  }

  ArcticMatrix.fromData(List<List<ArcticMember>> data) {
    this.data = data;
  }

  ArcticMatrix operator +(ArcticMatrix rhs) {
    int m1 = this.data.length;
    int m2 = rhs.data.length;
    if (!(m1 > 0 && m1 == m2)) {
      throw 'Arctic Matrix Addition: bad input';
    }

    int n1 = this.data[0].length;
    int n2 = rhs.data[0].length;
    if (!(n1 > 0 && n1 == n2)) {
      throw 'Arctic Matrix Addition: bad input';
    }

    ArcticMatrix newMatrix = ArcticMatrix.fromSize(m1, n1);

    for (var i = 0; i < m1; i++) {
      for (var j = 0; j < n1; j++) {
        newMatrix.data[i][j] = data[i][j] + rhs.data[i][j];
      }
    }

    return newMatrix;
  }

  ArcticMatrix operator *(ArcticMatrix rhs) {
    int m1 = this.data.length;
    int m2 = rhs.data.length;
    if (!(m1 > 0 && m2 > 0)) {
      throw 'Arctic Matrix Multiplication: bad input';
    }

    int n1 = this.data[0].length;
    int n2 = rhs.data[0].length;
    if (!(n1 > 0 && n2 > 0)) {
      throw 'Arctic Matrix Multiplication: bad input';
    }

    if (n1 != m2) {
      throw 'Arctic Matrix Multiplication: bad input (${n1} ${m2})';
    }

    ArcticMatrix newMatrix = ArcticMatrix.fromSize(m1, n2);

    for (var i = 0; i < m1; i++) {
      for (var j = 0; j < n2; j++) {
        ArcticMember buf = this.data[i][0] * rhs.data[0][j];
        for (var k = 1; k < n1; k++) {
          buf += this.data[i][k] * rhs.data[k][j];
        }
        newMatrix.data[i][j] = buf;
      }
    }

    return newMatrix;
  }

  @override
  String toString() {
    String res = "";
    for (var line in data) {
      res += line.join(" ") + "\n";
    }
    return res;
  }
}

ArcticMatrix arcticMulMatrixForChar(String char) {
  return ArcticMatrix.fromData([
    [
      ArcticMember.fromString('${char}_11'),
      ArcticMember.fromString('${char}_12')
    ],
    [
      ArcticMember.fromString('${char}_21'),
      ArcticMember.fromString('${char}_22')
    ],
  ]);
}

ArcticMatrix arcticAddMatrixForChar(String char) {
  return ArcticMatrix.fromData([
    [ArcticMember.fromString('${char}_1')],
    [ArcticMember.fromString('${char}_2')],
  ]);
}
