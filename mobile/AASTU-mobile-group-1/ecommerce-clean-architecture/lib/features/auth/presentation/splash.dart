  import 'package:ecommerce/features/auth/presentation/login.dart';
import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:google_fonts/google_fonts.dart';

class splash extends StatelessWidget {
  const splash({super.key});
  
  @override
  Widget build(BuildContext context) {
    Future.delayed(Duration(seconds: 5), () {
      Navigator.pushReplacement(
        context,
        MaterialPageRoute(builder: (context) => const login()),
      );
    });
    return Scaffold(
      body: Stack(
        children:[ 
          Container(
          decoration: BoxDecoration(
            image: DecorationImage(
              image: AssetImage('assets/images/splash.png'),
              fit: BoxFit.cover,
            ),
          ),
        ),
        Container(
          width: double.infinity,
          height: double.infinity,
          color : Color(0xFF3F51F3).withOpacity(0.6),
        ),
        Center(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Container(
                width: 250,
                height: 150,
                decoration: BoxDecoration(
                  color : Colors.white,
                  borderRadius: BorderRadius.circular(20),
                ),
                child: Center(
                  child: Text('ECOM',
                  style: GoogleFonts.caveatBrush(
                    textStyle: TextStyle(
                      fontSize: 100,
                      color: Color(0xFF3F51F3),
                    )
                  ),),
                ),
              ),
              const SizedBox(height: 20),
              Text('ECOMMERCE APP',
              style: GoogleFonts.poppins(
                textStyle: TextStyle(
                  fontSize: 30,
                  color: Colors.white,
                )
              ),),
              ],
            ),
          ),
        
        ],
      ),
    );
  }
}