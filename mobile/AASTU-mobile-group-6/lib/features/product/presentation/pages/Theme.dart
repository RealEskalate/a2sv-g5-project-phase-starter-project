import 'package:flutter/material.dart';

ThemeData lightmode = ThemeData(
  brightness: Brightness.light,
  colorScheme: ColorScheme.light(
      primary: Colors.purple,
      surface: Colors.white,
      onPrimary: const Color.fromARGB(255, 0, 0, 0),
      onSecondary: const Color.fromARGB(255, 75, 69, 69),
      onSurface: Colors.black,
      onTertiary: Colors.grey.shade300),
  scaffoldBackgroundColor: Colors.white,
  textTheme: TextTheme(
    bodyLarge: TextStyle(color: Colors.black),
    bodyMedium: TextStyle(color: Color.fromARGB(255, 54, 51, 51)),
  ),
);

ThemeData darkmode = ThemeData(
  brightness: Brightness.dark,
  colorScheme: ColorScheme.dark(
    primary: Colors.purple,
    surface: Colors.black,
    onPrimary: Colors.white,
    onSecondary: Color.fromARGB(255, 170, 170, 170),
    onSurface: const Color.fromARGB(255, 225, 224, 224),
    onTertiary: Colors.grey.shade900,
  ),
  scaffoldBackgroundColor: Colors.black,
  textTheme: TextTheme(
    bodyLarge: TextStyle(color: Colors.white),
    bodyMedium: TextStyle(color: Color.fromARGB(255, 255, 255, 255)),
  ),
);
