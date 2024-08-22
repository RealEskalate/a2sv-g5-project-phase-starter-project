import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

class SplashPage extends StatelessWidget {
  const SplashPage({super.key});

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () {
        Navigator.pushNamed(context, '/signin');
      },
      child: Stack(
        children: [
          Image.asset(
            'assets/splash.png',
            height: MediaQuery.of(context).size.height,
            width: MediaQuery.of(context).size.width,
            fit: BoxFit.cover,
          ),
          Scaffold(
            backgroundColor: const Color.fromRGBO(29, 29, 236, 0.6),
            body: Padding(
              padding: const EdgeInsets.only(top: 260, right: 50, left: 50),
              child: Column(
                mainAxisAlignment: MainAxisAlignment.start,
                children: [
                  Card(
                    color: Colors.white,
                    margin: const EdgeInsets.only(top: 0),
                    shape: const RoundedRectangleBorder(
                      borderRadius: BorderRadius.all(Radius.circular(25)),
                    ),
                    child: Padding(
                      padding: const EdgeInsets.only(
                          top: 0, bottom: 0,
                         right: 10, left: 10),
                      child: Text(
                        'ECOM',
                        style: GoogleFonts.caveatBrush(
                          
                          color: const Color.fromRGBO(63, 81, 243, 1),
                          fontSize: 112.89,
                          
                          fontWeight: FontWeight.w400,
                        ),
                      ),
                    ),
                  ),
                  const SizedBox(
                    height: 20,
                  ),
                  const Text('ECOMMERCE APP',
                      style: TextStyle(
                        color: Colors.white,
                        fontSize: 32,
                        fontWeight: FontWeight.bold,
                      )),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}
