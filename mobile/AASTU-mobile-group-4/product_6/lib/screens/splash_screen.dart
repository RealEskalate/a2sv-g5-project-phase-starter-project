import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

import '../features/auth/presentation/pages/login_page.dart';

class SplashScreen extends StatefulWidget {
  const SplashScreen({super.key});

  @override
  State<SplashScreen> createState() => _SplashScreenState();
}

class _SplashScreenState extends State<SplashScreen> {
  @override
  void initState() {
    super.initState();
    // Timer to navigate to the next screen after 3 seconds
    Future.delayed(const Duration(seconds: 5), () {
      Navigator.of(context).pushReplacement(
        MaterialPageRoute(
          builder: (_) => LoginPage(
            text: '',
          ),
        ), // Replace with your home screen
      );
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        fit: StackFit.expand,
        children: [
          Image.asset(
            'images/splash/splash_background.png',
            fit: BoxFit.cover,
          ),
          Container(
            width: 200, // Adjust the width as needed
            height: 200, // Adjust the height as needed
            decoration: BoxDecoration(
                color: const Color.fromRGBO(63, 81, 243, 1).withOpacity(0.8)),
          ),
          Center(
            child: Column(
              mainAxisSize: MainAxisSize.min,
              children: [
                Container(
                  padding: const EdgeInsets.only(
                    left: 5,
                    right: 5,
                  ),
                  decoration: BoxDecoration(
                    color: Colors.white,
                    borderRadius: BorderRadius.circular(38),
                  ),
                  child: Text(
                    'ECOM',
                    style: GoogleFonts.caveatBrush(
                      textStyle: const TextStyle(
                        fontWeight: FontWeight.w400,
                        fontSize: 115,
                        color: Color.fromRGBO(63, 81, 243, 1),
                      ),
                    ),
                  ),
                ),
                const SizedBox(height: 10),
                Text(
                  'Ecommerce APP',
                  style: GoogleFonts.poppins(
                    textStyle: const TextStyle(
                      fontWeight: FontWeight.w500,
                      fontSize: 35,
                      color: Colors.white,
                    ),
                  ),
                ),
              ],
            ),
          )
        ],
      ),
    );
  }
}
