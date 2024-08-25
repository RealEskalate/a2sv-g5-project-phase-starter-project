import 'package:flutter/material.dart';
import 'package:hexcolor/hexcolor.dart';

class MyTheme {
  /// The following are the colors that i will be using
  ///
  static const Color ecRed = Color.fromARGB(255, 255, 19, 19);
  static const Color ecInputGrey = Color.fromARGB(255, 243, 243, 243);
  static const Color skeletonColor1 = Color.fromARGB(255, 236, 236, 236);
  static const Color skeletonColor2 = Color.fromARGB(255, 231, 231, 231);

  static const Color ecTextGrey = Color.fromARGB(255, 170, 170, 170);
  static const Color ecGrey = Color.fromARGB(255, 170, 170, 170);
  static const Color ecBlue = Color.fromARGB(255, 63, 81, 243);
  static const Color ecBlack = Color.fromARGB(255, 62, 62, 62);
  static const Color ecWhite = Color.fromARGB(255, 255, 255, 255);
  static const Color shadowColor = Color.fromARGB(255, 234, 234, 234);
  static const double smallFont = 12;
  static const double midSmallFont = 14;
  static const double mediumFont = 20;
  static const double largeFont = 30;
  static ThemeData lightTheme = ThemeData(
      floatingActionButtonTheme:
          const FloatingActionButtonThemeData(backgroundColor: ecBlue),
      scaffoldBackgroundColor: ecWhite);
}
