import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';

import '../bloc/authentication/authentication_bloc.dart';

class SplashScreen extends StatefulWidget {
  const SplashScreen({super.key});

  @override
  // ignore: library_private_types_in_public_api
  _SplashScreenState createState() => _SplashScreenState();
}

class _SplashScreenState extends State<SplashScreen> with SingleTickerProviderStateMixin {
  @override
  void initState() {
    super.initState();

  
    Timer(const Duration(seconds: 3), () {
      // Navigator.pushReplacementNamed(context, '/home');

      if (mounted) {
        context.read<AuthenticationBloc>().add(CheckCurrentStatus());
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: [
          // Background Image
          Container(
            height: double.infinity,
            width: double.infinity,
            decoration: const BoxDecoration(
              image: DecorationImage(
                image: AssetImage('assets/images/back_ground_lady.png'),
                fit: BoxFit.cover,
              ),
            ),
          ),
          // Overlay gradient
          Container(
            decoration: BoxDecoration(
              gradient: LinearGradient(
                colors: [
                  const Color.fromARGB(255, 25, 78, 239).withOpacity(0.9),
                  const Color.fromARGB(255, 54, 104, 255).withOpacity(0.4),
                ],
                begin: Alignment.bottomCenter,
                end: Alignment.topCenter,
              ),
            ),
          ),
          // Logo and text
          Center(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Container(
                  padding: const EdgeInsets.symmetric(horizontal: 10.0),
                  decoration: BoxDecoration(
                    color: Colors.white.withOpacity(0.8),
                    borderRadius: BorderRadius.circular(40),
                  ),
                  child: Text(
                    'ECOM',
                    style: GoogleFonts.caveatBrush(
                      fontSize: 112.89,
                      fontWeight: FontWeight.w400,
                      color: const Color.fromARGB(255, 54, 104, 255),
                      height: 1.04,
                      letterSpacing: 2.25,
                    ),
                  ),
                ),
                const SizedBox(height: 20),
                Text(
                  'ECOMMERCE APP',
                  textAlign: TextAlign.center, // Center align the text
                  style: GoogleFonts.poppins(
                    fontSize: 35.98,
                    fontWeight: FontWeight.w500,
                    color: Colors.white,
                    height: 1.04, // Calculated line height
                    letterSpacing: 0.72, // Calculated letter spacing
                  ),
                ),
                BlocListener<AuthenticationBloc, AuthenticationState>(
                  listener: (context, state) {
                  if (state is LoggedInState) {
                    Navigator.pushReplacementNamed(context, '/home');
                  } else if (state is LoggedOutState) {
                    Navigator.pushReplacementNamed(context, '/signin');
                  }
           },
                 child: const SizedBox(),
                )
              ],
            ),
          ),
        ],
      ),
    );
  }
}
