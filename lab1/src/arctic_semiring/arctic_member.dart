/// a + b = (arc_sum a b) === (max a b)
/// a * b = (arc_mul a b) === (+ a b)

class ArcticMember {
  String data = "";

  ArcticMember.fromString(String str) {
    data = str;
  }

  ArcticMember();

  ArcticMember operator +(ArcticMember rhs) {
    return ArcticMember.fromString(
        "(arc_add ${ValidateArcMember(data)} ${ValidateArcMember(rhs.data)})");
  }

  ArcticMember operator *(ArcticMember rhs) {
    return ArcticMember.fromString(
        "(arc_mul ${ValidateArcMember(data)} ${ValidateArcMember(rhs.data)})");
  }

  @override
  String toString() {
    return data;
  }
}

/// TODO: Get decision of validation need
String ValidateArcMember(String str) {
  if (str.length == 1 || str.startsWith("(")) {
    return str;
  }
  return "${str}";
}
