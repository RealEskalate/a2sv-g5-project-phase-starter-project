import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:google_fonts/google_fonts.dart';
class Splash extends StatelessWidget {
  const Splash({super.key});

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      
      child: Scaffold(
        
        backgroundColor: Colors.white,
        body: Stack(
          children: [
            SizedBox(
              width: double.maxFinite,
              height: double.maxFinite,
              child: Image.asset('assets/smile.png', fit: BoxFit.cover,),)
      
              ,
              Opacity(
                opacity: 0.9,
                child: Container(
                  decoration: BoxDecoration(
                    gradient: LinearGradient(
                      begin: Alignment.topCenter,
                      end: Alignment.bottomCenter,
                      colors: [
                       Color.fromARGB(125, 63, 81, 243), // rgba(63, 81, 243, 0)
                        Color.fromARGB(255, 63, 81, 243,)
                      ],
                    ),
                  ),
                  height: double.maxFinite,
                  width: double.maxFinite,
                ),
              )
              ,
              
              Center(
                child: Positioned(
                  
                  
                    child: Column(
                      mainAxisAlignment: MainAxisAlignment.center,
                      children: [
                
                        GestureDetector(
                          onTap: () {
                            Navigator.pushReplacementNamed(context, '/signup');
                          },
                          child: Container(
                            
                            height: 121,
                            width: 264,
                            
                            decoration: BoxDecoration(
                              color: Colors.white,
                              borderRadius: BorderRadius.circular(20),
                            ),
                            
                              child: FittedBox(
                                
                                fit: BoxFit.cover,

                                child: Padding(
                                  padding: EdgeInsets.all(20),
                                  child: Center(
                                    
                                    child: Text('ECOM', 
                                    style: GoogleFonts.caveatBrush(
                                      color: Color.fromARGB(255,63, 81, 243),
                                      fontWeight: FontWeight.w400,
                                      fontSize: 112
                                      
                                      
                                    )
                                    // TextStyle(color: Colors.blue, fontSize: 45, fontWeight: FontWeight.bold)
                                    ,),
                                  ),
                                ),
                              ),
                            )
                          ),
                          SizedBox(height: 15,),
                        
                        Container(
                          child: Text('ECOMMERCE APP', style: GoogleFonts.poppins(
                            color: Colors.white,
                            fontWeight: FontWeight.w500,
                            fontSize: 36,
                          
                          )
                          // TextStyle(color: Colors.white, fontSize: 30, fontWeight: FontWeight.bold)
                          ,),
                        ),
                      ],
                    ),
                  ),
              )
          ],
        ),
      ),
    );
  }
}