#!/bin/sh

dart src/main.dart && z3 -smt2 res.smt2
