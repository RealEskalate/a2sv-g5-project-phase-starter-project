import 'dart:ui';

import 'package:flutter/material.dart';

import 'signin_page.dart';

class SplashScreen extends StatelessWidget {
  const SplashScreen({Key? key}) : super(key: key);

  Future<void> loadData() async {
    await Future.delayed(const Duration(seconds: 3));
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: loadData(),
      builder: (context, snapshot) {
        if (snapshot.connectionState == ConnectionState.waiting) {
          return Scaffold(
            body: Stack(children: [
              Container(
                decoration: const BoxDecoration(
                  image: DecorationImage(
                    image: AssetImage('assets/images/splash.jpg'),
                    fit: BoxFit.cover,
                  ),
                ),
              ),
              BackdropFilter(
                filter: ImageFilter.blur(sigmaX: 5, sigmaY: 5),
                child: Container(
                  color:
                      const Color.fromARGB(120, 28, 51, 231).withOpacity(0.5),
                ),
              ),
              Center(
                child: Container(
                  margin: const EdgeInsets.only(bottom: 100),
                  padding: const EdgeInsets.all(10),
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(20),
                    color: Colors.white,
                  ),
                  child: const Text(
                    'ECOM',
                    style: TextStyle(
                        color: Color.fromARGB(255, 9, 70, 225),
                        fontSize: 80,
                        fontFamily: 'Caveat Brush',
                        fontWeight: FontWeight.bold),
                  ),
                ),
              ),
              Center(
                  child: Container(
                      padding: const EdgeInsets.only(top: 100),
                      child: const Text(
                        'ECOMMERCE APP',
                        style: TextStyle(fontSize: 32, color: Colors.white),
                      )))
            ]),
          );
        } else {
          return SigninPage();
        }
      },
    );
  }
}
