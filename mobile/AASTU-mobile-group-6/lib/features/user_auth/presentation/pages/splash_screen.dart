import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:shared_preferences/shared_preferences.dart';

class SplashScreen extends StatefulWidget {
  @override
  State<SplashScreen> createState() => _SplashScreenState();
}

class _SplashScreenState extends State<SplashScreen> with SingleTickerProviderStateMixin {

  @override
  void initState() {
        super.initState();
        Future.delayed(
          Duration(seconds: 5),
          () async {

            var ch = await SharedPreferences.getInstance();

            var t = ch.getString('access_token');

            if (t==null){

              Navigator.pushReplacementNamed(context, '/login');
            }else if (t!=null){
              Navigator.pushNamed(context, '/home');
            }
          },
        );
      }
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Color.fromRGBO(63, 81, 243, 1),
      body:SafeArea(
        child: Stack(
          children: [
            Container(
                    height: MediaQuery.of(context).size.height,
                    // width: MediaQuery.of(context).size.width,
                    child: Image.asset('assets/girls.png',fit: BoxFit.cover,),

                  ),
                  
            Container(
              color: Color.fromRGBO(63, 81, 243, 1).withOpacity(0.6),
            ),

            Container(
              width: MediaQuery.of(context).size.width,
              height: MediaQuery.of(context).size.height,
              decoration: BoxDecoration(
                gradient: LinearGradient (
                  colors: const [
                    Colors.transparent,
                    Color.fromRGBO(63, 81, 243, 1),
                  ],
                  begin: Alignment.topLeft,
                  end: Alignment.bottomRight,
                ),
              ),
            ),
           Center(
              child: Column(
                mainAxisSize: MainAxisSize.min,   //the column take up only space of its children
                children: [
                  ClipRRect(
                    borderRadius: BorderRadius.circular(30),
                    child: Container(
                      padding: EdgeInsets.all(5),
                      color: Colors.white,
                      child: Text(
                        'ECOM',
                        style: GoogleFonts.caveatBrush(fontSize: 100,color:Color.fromRGBO(63, 81, 243, 1))
                      ),
                    ),
                  ),
                  SizedBox(height: 20), // Space between the two texts
                  Text(
                    'ECOMMERCE APP', // Replace with your desired text
                    style: GoogleFonts.poppins(fontSize: 40,color: Colors.white)
                  ),
                ],
            )
           )
          ]
        ),
      )
    );
  }
}