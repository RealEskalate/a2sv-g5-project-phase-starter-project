import 'package:flutter/material.dart';
import '../../features/auth/presentation/pages/sign_in_page.dart';
import '../../features/auth/presentation/pages/sign_up_page.dart';
import '../../features/auth/presentation/pages/splash_page.dart';
import '../../features/chat/presentation/pages/chat_home_screen.dart';
import '../../features/chat/presentation/pages/individiual_chat_screen.dart';
import '../../features/product/presentation/pages/add_update_page.dart';
import '../../features/product/presentation/pages/detail_page.dart';
import '../../features/product/presentation/pages/home_page.dart';
import '../../features/product/presentation/pages/search_page.dart';

const String homePage = 'home';
const String detailPage = 'detail';
const String searchPage = 'search';
const String addUpdatePage = 'add';
const String splashPage = 'splash';
const String signInPage = 'signIn';
const String signUpPage = 'signUp';
const String chatHomeScreen = 'chattHomeScreen';
const String individiualChatScreen= 'individiualChatScreen';

Route<dynamic> controller(RouteSettings settings) {
  switch (settings.name) {
    case addUpdatePage:
      final args = settings.arguments as Map<String, dynamic>?;
      return PageRouteBuilder(
        pageBuilder: (context, animation, secondaryAnimation) => AddUpdatePage(
          isUpdate: args?['isUpdate'] ?? false,
          product: args?['product'],
        ),
        transitionsBuilder: (context, animation, secondaryAnimation, child) {
          var begin = const Offset(1.0, 0.0);
          var end = Offset.zero;
          var curve = Curves.ease;

          var tween =
              Tween(begin: begin, end: end).chain(CurveTween(curve: curve));
          var offsetAnimation = animation.drive(tween);

          return SlideTransition(
            position: offsetAnimation,
            child: child,
          );
        },
      );
    case chatHomeScreen:
          return PageRouteBuilder(
        pageBuilder: (context, animation, secondaryAnimation) =>
            const ChatHomeScreen(),
        transitionsBuilder: (context, animation, secondaryAnimation, child) {
          var begin = const Offset(1.0, 0.0);
          var end = Offset.zero;
          var curve = Curves.ease;

          var tween =
              Tween(begin: begin, end: end).chain(CurveTween(curve: curve));
          var offsetAnimation = animation.drive(tween);

          return SlideTransition(
            position: offsetAnimation,
            child: child,
          );
        },
      );
    case detailPage:
      return PageRouteBuilder(
        pageBuilder: (context, animation, secondaryAnimation) =>
            const DetailPage(),
        transitionsBuilder: (context, animation, secondaryAnimation, child) {
          var begin = const Offset(1.0, 0.0);
          var end = Offset.zero;
          var curve = Curves.ease;

          var tween =
              Tween(begin: begin, end: end).chain(CurveTween(curve: curve));
          var offsetAnimation = animation.drive(tween);

          return SlideTransition(
            position: offsetAnimation,
            child: child,
          );
        },
      );
    case searchPage:
      return PageRouteBuilder(
        pageBuilder: (context, animation, secondaryAnimation) =>
            const SearchPage(),
        transitionsBuilder: (context, animation, secondaryAnimation, child) {
          var begin = const Offset(1.0, 0.0);
          var end = Offset.zero;
          var curve = Curves.ease;

          var tween =
              Tween(begin: begin, end: end).chain(CurveTween(curve: curve));
          var offsetAnimation = animation.drive(tween);

          return SlideTransition(
            position: offsetAnimation,
            child: child,
          );
        },
      );

      case individiualChatScreen:
      return PageRouteBuilder(
        pageBuilder: (context, animation, secondaryAnimation) =>
             IndividiualChatScreen(),
        transitionsBuilder: (context, animation, secondaryAnimation, child) {
          var begin = const Offset(1.0, 0.0);
          var end = Offset.zero;
          var curve = Curves.ease;

          var tween =
              Tween(begin: begin, end: end).chain(CurveTween(curve: curve));
          var offsetAnimation = animation.drive(tween);

          return SlideTransition(
            position: offsetAnimation,
            child: child,
          );
        },
      );
    case splashPage:
      return PageRouteBuilder(
          pageBuilder: (context, animation, secondaryAnimation) =>
              const SplashPage());

    case signUpPage:
      return PageRouteBuilder(
          pageBuilder: (context, animation, secondaryAnimation) =>
              const SignUpPage());
    case signInPage:
      return PageRouteBuilder(
          pageBuilder: (context, animation, secondaryAnimation) =>
              const SignInPage());
    case homePage:
    default:
      return PageRouteBuilder(
        pageBuilder: (context, animation, secondaryAnimation) =>
            const HomePage(),
        transitionsBuilder: (context, animation, secondaryAnimation, child) {
          var begin = const Offset(0.0, 1.0);
          var end = Offset.zero;
          var curve = Curves.ease;

          var tween =
              Tween(begin: begin, end: end).chain(CurveTween(curve: curve));
          var offsetAnimation = animation.drive(tween);

          return SlideTransition(
            position: offsetAnimation,
            child: child,
          );
        },
      );
  }
}
