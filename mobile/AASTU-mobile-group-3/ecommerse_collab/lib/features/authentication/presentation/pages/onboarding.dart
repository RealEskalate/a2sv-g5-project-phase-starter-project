import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';
import '../../../product/presentation/pages/home_page.dart';
import '../../domain/usecases/login.dart';
import 'package:http/http.dart' as http;
import '../widgets/logo.dart';
import 'package:flutter/services.dart';
import '../../data/data_source/remote_data_source.dart';
import 'sign_in.dart';

class Onboarding extends StatefulWidget {
  const Onboarding({super.key});

  @override
  State<Onboarding> createState() => _OnboardingState();
}

class _OnboardingState extends State<Onboarding>
    with SingleTickerProviderStateMixin {
  @override
  void initState() {
    super.initState();
    Future.delayed(Duration(seconds: 10), () async {
      final cache = await SharedPreferences.getInstance();
      // final isLoggedIn = cache.getBool('isLoggedIn') ?? false;

      
          Navigator.push(
            context,
            MaterialPageRoute(builder: (context) { 
              return const SignIn();}),
          );
      
    });
  }
    @override
    Widget build(BuildContext context) {
      return Scaffold(
      body: Stack(
        children: [
          Container(
            decoration: const BoxDecoration(
              image: DecorationImage(
                image: AssetImage('assets/images/onboarding.png'),
                fit: BoxFit.cover,
              ),
            ),
          ),
          Container(
            decoration: BoxDecoration(
              gradient: LinearGradient(
                begin: Alignment.topCenter,
                end: Alignment.bottomCenter,
                colors: [
                  const Color(0XFF3F51F3).withOpacity(0.5),
                  const Color(0XFF3F51F3),
                ],
              ),
            ),
          ),
          const Center(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Logo(),
                  SizedBox(
                    width: 312,
                    height: 38,
                    child: Align(
                      alignment: Alignment.center,
                child: Text(
                  'ECOMMERCE APP',
                  style: TextStyle(
                    color: Colors.white,
                    fontSize: 35,
                    fontWeight: FontWeight.w500,
                  ),
                ),
                    ),
                  )
              ],
            ),
          ),
        ],
      ),
    );
    }


}
