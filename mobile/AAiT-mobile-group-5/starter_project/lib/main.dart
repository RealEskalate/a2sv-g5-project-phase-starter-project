import 'package:flutter/material.dart';
import 'package:starter_project/features/authentication/presentation/pages/sign_in.dart';
import 'package:starter_project/features/authentication/presentation/pages/sign_up.dart';
import 'package:starter_project/features/authentication/presentation/pages/splash_screen.dart';
import 'package:starter_project/features/chat/presentation/pages/chat_page.dart';

void main() {
  runApp( const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
          title: 'Flutter Demo',
          debugShowCheckedModeBanner: false,
          theme: ThemeData(
            colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
            useMaterial3: true,
          ),
          initialRoute: '/splash',
          routes: {
            '/signin': (context) => const SignInPage(),
            '/signup':(context)=> const SignUpPage(),
            '/splash': (context) => const SplashScreen(),
            '/chatpage': (context)=> const ChatPage()
            
          }
    );
  }
}
