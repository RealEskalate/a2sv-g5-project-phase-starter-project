import 'package:flutter/material.dart';
import 'package:starter_project/features/authentication/presentation/pages/sign_in.dart';
import 'package:starter_project/features/authentication/presentation/pages/sign_up.dart';
import 'package:flutter/material.dart';
import 'package:starter_project/features/authentication/presentation/pages/sign_in.dart';
import 'package:starter_project/features/authentication/presentation/pages/sign_up.dart';
import 'package:starter_project/features/authentication/presentation/pages/splash_screen.dart';
import 'package:starter_project/features/chat/presentation/pages/chat_page.dart';
import 'package:starter_project/features/chat/presentation/pages/home_page.dart';
import 'features/authentication/presentation/pages/terms_and_policy.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Chat App',
      debugShowCheckedModeBanner: false,
      initialRoute: '/splash',
      routes: {
        '/signin': (context) => const SignInPage(),
        '/signup': (context) => const SignUpPage(),
        '/splash': (context) => const SplashScreen(),
        '/chatpage': (context) => const ChatPage(),
        "/terms": (context) => const TermsAndPolicy(),
        '/home': (context) => const HomePage()
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
