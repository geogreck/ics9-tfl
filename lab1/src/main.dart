import 'dart:io';

import 'srs/srs.dart';
import 'steps/steps.dart';

const fileName = "res.smt2";

void RunWorkflow() {
  print("input srs: ");
  List<String> input = ReadInput();
  print("input: " + input.toString());

  bool valid = ValidateInput(input);
  if (!valid) {
    throw ("input invalid, terminating");
  }
  print("input validated");

  File file = CreateFile(fileName);
  print("created file ${fileName}");
  WriteInitialPart(file);
  print("wrote initial part to file");

  SRS srs = ParseSRS(input);
  print("parsed srs");

  DumpSRSToFile(file, srs);
  print("dumped srs to file");

  WriteFinalPart(file);
  print("wrote final part to file");
  print("all done\n\n");
}

void main(List<String> args) {
  RunWorkflow();
}
