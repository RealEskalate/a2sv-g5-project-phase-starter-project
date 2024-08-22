import 'package:flutter/material.dart';

class AppTheme  {
  final darkTheme = ThemeData(
    scaffoldBackgroundColor: Colors.black,
    appBarTheme: AppBarTheme(
      color: Colors.black,
      iconTheme: IconThemeData(color: Colors.white),

    ),
    textTheme: const TextTheme(
      bodyMedium: TextStyle(color: Colors.white),
    )
  );
  final lightTheme= ThemeData(
    scaffoldBackgroundColor: Colors.white,
    appBarTheme: AppBarTheme(
      color: Colors.black,
      iconTheme: IconThemeData(color: Colors.black),

    ),
    textTheme: const TextTheme(
      bodyMedium: TextStyle(color: Colors.black),
    )
  );
}