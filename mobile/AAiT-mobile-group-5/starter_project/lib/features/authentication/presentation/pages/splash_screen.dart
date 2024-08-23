import 'dart:async';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';


class SplashScreen extends StatefulWidget {
  const SplashScreen({super.key});

  @override
  State<SplashScreen> createState() => _SplashScreenState();
}


class _SplashScreenState extends State<SplashScreen> {

 @override
 void initState() {
    super.initState();

 
    Timer(const Duration(seconds: 5), () {
      Navigator.pushNamed(context, '/signin');
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: GestureDetector(
          onTap: (){Navigator.pushNamed(context, '/signin');},
          child: Stack(
                children: [
          Positioned.fill(
              child: Image.asset(
            'assets/ecommerce.jpg',
            fit: BoxFit.fill,
          )),
          Positioned.fill(
            child: Container(
              decoration: BoxDecoration(
                gradient: LinearGradient(
                  colors: [
                    Color(0xFF3F51F3).withOpacity(0.8), // 50% opacity blue
                    Color(0xFF3F51F3).withOpacity(0.8), // 50% opacity blue
                  ],
                  begin: Alignment.topCenter,
                  end: Alignment.bottomCenter,
                ),
              ),
            ),
          ),
          Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Center(
                  child: Container(
                width: 264,
                height: 125,
                decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(35), color: Colors.white),
                alignment: Alignment.center,
                child: Text(
                  'ECOM',
                  style: GoogleFonts.caveatBrush(
                      fontSize: 113,
                      fontWeight: FontWeight.w400,
                      color: const Color.fromARGB(255, 63, 81, 243),
                      height: 1.0
                      ),
                ),
              )),
              SizedBox(height: 20),
              Text(
                'ECOMMERCE APP',
                style: GoogleFonts.poppins(
                    fontSize: 36,
                    fontWeight: FontWeight.w500,
                    color: Colors.white),
              )
            ],
          )
                ],
              ),
        ));
  }
}
