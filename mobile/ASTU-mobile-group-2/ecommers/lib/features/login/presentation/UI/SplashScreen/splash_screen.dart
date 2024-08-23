import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

import '../../../../../core/text/text.dart';
import '../../../../../core/user_validation/user_validation.dart';


class SplashScreen extends StatefulWidget {
  const SplashScreen({super.key});

  @override
  State<SplashScreen> createState() => _SplashScreenState();
}

class _SplashScreenState extends State<SplashScreen> with SingleTickerProviderStateMixin {
  
  @override
  void initState() {
    super.initState();
    _initializeApp();
  }
  Future<void> _initializeApp() async {
    SystemChrome.setEnabledSystemUIMode(SystemUiMode.immersive);
    final checker = await userValidation();
    await Future.delayed(const Duration(seconds: 3));

    if (checker) {
      // ignore: use_build_context_synchronously
      Navigator.pushReplacementNamed(context, '/home',
      );
    } else {
      // ignore: use_build_context_synchronously
      Navigator.pushReplacementNamed(context, '/login');
    }
  }
  @override
  void dispose(){
    SystemChrome.setEnabledSystemUIMode(SystemUiMode.manual,
    overlays: SystemUiOverlay.values);
    super.dispose();
  }
  @override
  Widget build(BuildContext context) {
    final width = MediaQuery.of(context).size.width;
    return  SafeArea(
      child: Scaffold(
        body: SizedBox(
          width: double.infinity,
          height: double.infinity,
          child: Container(
              decoration: const BoxDecoration(
                image: DecorationImage(
                    image: AssetImage(
                        'assets/image/splashScreen.png'),
                    fit: BoxFit.fill),
              ),
      
              child: Container(
                width: double.infinity,
                height: double.infinity,
                decoration: const BoxDecoration(
                  gradient: LinearGradient(
                    colors: [
                      Color.fromARGB(180, 91, 30, 233),
                      Color.fromARGB(200, 91, 30, 233),
                      Color.fromARGB(240, 91, 30, 233),
                    ],
                    begin: Alignment.topCenter,
                  )
                ),
      
                child:  Column(
                  
                  mainAxisAlignment: MainAxisAlignment.center,
                  crossAxisAlignment: CrossAxisAlignment.center,
                  children: [
                    Container(
                      width : width*0.64,
                      height: 120,
                      decoration: BoxDecoration(
                        color: Colors.white,
                        borderRadius: BorderRadius.circular(30),
      
                      ),
                      child: const Center(
                        child:  ConStTexts(
                          text: 'ECOM',
                          fontSize: 100,
                          color: Color.fromARGB(255, 91, 30, 233),
                          fontWeight: FontWeight.bold,
                          fontFamily: 'CaveatBrush',
                        )),
                    ),
                    const SizedBox(
                      height: 20,
                    ),
                    const ConStTexts(
                      text: 'ECOMMERCE APP', 
                      fontSize: 35, 
                      color: Colors.white, 
                      fontWeight: FontWeight.w500,
                      fontFamily: 'Poppins',
                      )
                  ],
                ),
              ),
        ),
        )
      ),
    );
  }
}