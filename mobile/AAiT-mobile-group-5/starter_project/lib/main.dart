import 'package:flutter/material.dart';

import 'features/authentication/presentation/pages/sign_in.dart';
import 'features/authentication/presentation/pages/sign_up.dart';
import 'features/authentication/presentation/pages/terms_and_policy.dart';
import 'features/chat/presentation/pages/home_page.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: "Chat App",
      initialRoute: "/",
      debugShowCheckedModeBanner: false,
      routes: {
        "/": (context) => const HomePage(),
        "/signUp": (context) => const SignUpPage(),
        "/signIn": (context) => const SignInPage(),
        "/terms": (context) => const TermsAndPolicy(),
      },
      theme: ThemeData(
        primaryColor: const Color(0xFFFFFFFF), // White color
        appBarTheme: const AppBarTheme(
          backgroundColor: Color(0xFFFFFFFF), // AppBar background color
          titleTextStyle: TextStyle(
            color: Color(0xFF000000), // Black color for AppBar title text
          ),
          iconTheme: IconThemeData(
            color: Color(0xFF000000), // Black color for AppBar icons
          ),
        ),
        buttonTheme: const ButtonThemeData(
          buttonColor: Color(0xFF3F51F3), // Button background color
          textTheme: ButtonTextTheme.primary, // Makes the text color white
        ),
        textTheme: const TextTheme(
          labelLarge: TextStyle(
              color:
                  Color.fromARGB(255, 8, 8, 8)), // White color for button text
          bodyLarge: TextStyle(
              color:
                  Color.fromARGB(255, 9, 9, 9)), // White color for general text
          bodyMedium: TextStyle(
              color: Color.fromARGB(
                  255, 15, 15, 15)), // White color for general text
        ),
        elevatedButtonTheme: ElevatedButtonThemeData(
          style: ButtonStyle(
            backgroundColor:
                WidgetStateProperty.all<Color>(const Color(0xFF3F51F3)),
            foregroundColor:
                WidgetStateProperty.all<Color>(const Color(0xFFFFFFFF)),
          ),
        ),
      ),
    );
  }
}
