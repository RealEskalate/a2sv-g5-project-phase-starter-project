import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../product/presentation/widgets/components/styles/text_style.dart';
import '../bloc/auth_bloc.dart';

class CoverPage extends StatefulWidget {
  const CoverPage({super.key});

  @override
  State<CoverPage> createState() => _CoverPageState();
}

class _CoverPageState extends State<CoverPage> {
  @override
  void initState() {
    super.initState();
    BlocProvider.of<AuthBloc>(context).add(GetCurrentUserEvent());
    Future.delayed(const Duration(seconds: 4), () {
      if (BlocProvider.of<AuthBloc>(context).state is AuthUserLoaded) {
        Navigator.pushReplacementNamed(context, '/home_page');
      } else {
        Navigator.pushReplacementNamed(context, '/sign_in_page');
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: const BoxDecoration(
        image: DecorationImage(
            image: AssetImage('assets/splash_screen.jpg'), fit: BoxFit.cover),
      ),
    child: Container(
      decoration: const BoxDecoration(
          gradient: LinearGradient(
            begin: Alignment.bottomCenter,
            end: Alignment.topCenter,
            colors: [Color.fromRGBO(63, 82, 243, 1),Color.fromRGBO(63, 81, 243, 0)]
      ),),
      child: Scaffold(
        backgroundColor: Colors.transparent,
        body: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              Container(
                decoration: BoxDecoration(
                  borderRadius: BorderRadius.circular(32),
                  color: Colors.white,
                ),
                child: const Padding(
                  padding: EdgeInsets.symmetric(horizontal: 24),
                  child: CustomTextStyle(
                    color: Color.fromARGB(255, 56, 15, 219),
                    name: 'ECOM',
                    weight: FontWeight.w400,
                    size: 113,
                    family: 'CaveatBrush',
                  ),
                ),
              ),
              const SizedBox(height: 10),
              const CustomTextStyle(
                name: 'ECOMMERCE APP',
                weight: FontWeight.w500,
                size: 36,
                color: Colors.white,
                family: 'Poppins',
              ),
            ],
          ),
        ),

      ),
    ),
    );
  }
}
