import 'package:flutter/material.dart';

ThemeData lightmode = ThemeData(
  brightness: Brightness.light,
  colorScheme: ColorScheme.light(
      primary: Colors.blue,
      surface: Colors.white,
      onPrimary: const Color.fromARGB(255, 0, 0, 0),
      onSecondary: const Color.fromARGB(255, 75, 69, 69),
      onSurface: Colors.black,
      onTertiary: Colors.grey.shade300,
      onPrimaryContainer: Colors.white,
      onTertiaryFixed: Colors.white,
      onPrimaryFixedVariant: Color.fromRGBO(243, 243, 243, 1)),

  scaffoldBackgroundColor: Colors.white,

  // ------appbar theme for light------
  appBarTheme: AppBarTheme(
    // demo 1
    // color: const Color.fromARGB(255, 8, 85, 149),
    // iconTheme: IconThemeData(color: Colors.white),
    // titleTextStyle: TextStyle(
    //   color: Colors.white,
    //   fontSize: 20,
    //   fontWeight: FontWeight.bold,
    // ),
    // elevation: 4,

    // demo2
    color: Colors.white,
    iconTheme: IconThemeData(color: Colors.black),
    titleTextStyle: TextStyle(
      color: Colors.black,
      fontSize: 20,
      fontWeight: FontWeight.bold,
    ),
    // elevation: 4,
  ),
);

ThemeData darkmode = ThemeData(
  brightness: Brightness.dark,
  colorScheme: ColorScheme.dark(
      primary: const Color.fromARGB(255, 44, 42, 42),
      surface: Colors.black,
      onPrimary: Colors.white,
      onSecondary: Color.fromARGB(255, 170, 170, 170),
      onSurface: const Color.fromARGB(225, 255, 255, 255),
      onTertiary: Colors.grey.shade900,
      onPrimaryContainer: Colors.grey,
      onTertiaryFixed: const Color.fromARGB(255, 26, 25, 25),
      onPrimaryFixedVariant: Color.fromRGBO(52, 50, 50, 1)),

  scaffoldBackgroundColor: Colors.black,

  textTheme: TextTheme(
    bodyLarge: TextStyle(color: Colors.white),
    bodyMedium: TextStyle(color: Color.fromARGB(255, 255, 255, 255)),
  ),

  // ------appbar theme for dark------
  appBarTheme: AppBarTheme(
    color: const Color.fromARGB(255, 42, 40, 40),
    iconTheme: IconThemeData(color: const Color.fromARGB(255, 255, 255, 255)),
    titleTextStyle: TextStyle(
      color: const Color.fromARGB(229, 255, 255, 255),
      fontSize: 20,
      fontWeight: FontWeight.bold,
    ),
    // elevation: 4, // Shadow depth below the AppBar
  ),
);
