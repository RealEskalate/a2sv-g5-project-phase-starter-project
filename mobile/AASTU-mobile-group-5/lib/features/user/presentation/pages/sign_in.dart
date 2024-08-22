import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';

import '../../../../service_locator.dart';
import '../../domain/repositories/user_repository.dart';
import '../bloc/sign_in_page/sign_in_page_bloc.dart';

class SignInPage extends StatelessWidget {
  const SignInPage({super.key});

  @override
  Widget build(BuildContext context) {
    final TextEditingController emailController = TextEditingController();
    final TextEditingController passwordController = TextEditingController();
    
    return Scaffold(
      backgroundColor: Colors.white,
      body: SingleChildScrollView(
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 16.0),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: [
              const SizedBox(height: 100),
              Center(
                child: Container(
                  height: 50,
                  width: 143,
                  decoration: BoxDecoration(
                    color: Colors.white,
                    borderRadius: BorderRadius.circular(8.0),
                    border: Border.all(
                        color: const Color.fromARGB(255, 54, 104, 255), width: 2),
                    boxShadow: [
                      BoxShadow(
                        color: Colors.black
                            .withOpacity(0.25), // Shadow color with opacity
                        offset: const Offset(0, 4), // Horizontal and vertical offset
                        blurRadius: 4.0, // Blur radius
                        spreadRadius: 0.0, // Spread radius
                      ),
                    ],
                  ),
                  child: Center(
                    child: Text(
                      'ECOM',
                      style: GoogleFonts.caveatBrush(
                        textStyle: const TextStyle(
                          fontSize: 48.0,
                          fontWeight: FontWeight.bold,
                          color: Color.fromARGB(255, 54, 104, 255),
                          height: 24.26 / 48.0,
                          letterSpacing: 2 / 100 * 48.0,
                        ),
                      ),
                    ),
                  ),
                ),
              ),
              const SizedBox(height: 50.0),
              // Sign into your account text
              const Text(
                'Sign into your account',
                textAlign: TextAlign.left,
                style: TextStyle(
                  fontSize: 24.0,
                  fontWeight: FontWeight.bold,
                  color: Colors.black,
                ),

              ),
              const SizedBox(height: 32.0),
              Text(
                'Email',
                style: GoogleFonts.poppins(
                  textStyle: const TextStyle(
                    fontSize: 16.0,
                    fontWeight: FontWeight.w400,
                    color: Color.fromRGBO(111, 111, 111, 1),
                    height: 49.85 / 16.0,
                    letterSpacing: 2 / 100 * 16.0,
                  ),
                ),
              ),

              TextField(
                controller: emailController,
                decoration: InputDecoration(
                  hintText: 'ex: jon.smith@email.com',
                  hintStyle: GoogleFonts.poppins(
                    textStyle: const TextStyle(
                      fontWeight: FontWeight.w400,
                      color: Color.fromRGBO(111, 111, 111, 1),
                    ),
                  ),
                  border: OutlineInputBorder(
                    borderRadius: BorderRadius.circular(8.0),
                    borderSide: BorderSide.none,
                  ),
                  filled: true,
                  fillColor: Colors.grey[200],
                ),
              ),
              const SizedBox(height: 16.0),

              Text(
                'Password',
                style: GoogleFonts.poppins(
                  textStyle: const TextStyle(
                    fontSize: 16.0,
                    fontWeight: FontWeight.w400,
                    // color: Theme.of(context).brightness == Brightness.dark ? Color.fromRGBO(111, 111, 111, 1) : Color.fromRGBO(111, 111, 111, 1),
                    color: Color.fromRGBO(111, 111, 111, 1),
                    height: 49.85 / 16.0,
                    letterSpacing: 2 / 100 * 16.0,
                  ),
                ),
              ),

              TextField(
                controller: passwordController,
                
                obscureText: true,
                decoration: InputDecoration(
                  hintText: '********',
                  hintStyle: GoogleFonts.poppins(
                    textStyle: const TextStyle(
                      fontWeight: FontWeight.w400,
                      color: Color.fromRGBO(111, 111, 111, 1),
                    ),
                  ),
                  alignLabelWithHint: true,
                  border: OutlineInputBorder(
                      borderRadius: BorderRadius.circular(8.0),
                      borderSide: BorderSide.none),
                  filled: true,
                  fillColor: Colors.grey[200],
                ),
              ),
              const SizedBox(height: 32.0),

              BlocProvider(
                create: (context) => SignInPageBloc(userRepository: getIt<UserRepository>()),
                child: BlocListener<SignInPageBloc, SignInPageState>(
                  listener: (context, state) {
                    // if (state is SignInPageFailure) {
                    //   ScaffoldMessenger.of(context).showSnackBar(
                    //     SnackBar(content: Text(state.error)),
                    //   );
                    // } else if (state is SignInPageSuccess) {
                    //   Navigator.pushReplacementNamed(context, '/home');
                    // }
                    if (state is SignInPageFailure) {
                      String errorMessage;
                      // Check for password mismatch error
                      if (state.error == 'Password does not match') {
                        errorMessage =
                            'Password does not match. Please try again.';
                      } else {
                        errorMessage =
                            'Failed to sign in. Please check your credentials and try again.';
                      }
                      ScaffoldMessenger.of(context).showSnackBar(
                        SnackBar(
                          content: Text(errorMessage),
                          backgroundColor: const Color.fromARGB(255, 84, 80, 79),
                        ),
                      );
                    } else if (state is SignInPageSuccess) {
                      Navigator.pushReplacementNamed(context, '/home');
                    }
                  },
                  child: BlocBuilder<SignInPageBloc, SignInPageState>(
                    builder: (context, state) {
                      return Column(
                        children: [
                          state is SignInPageLoading
                              ? const CircularProgressIndicator()
                              : SizedBox(
                                width: double.infinity,
                                child: ElevatedButton(
                                    onPressed: () {
                                      context.read<SignInPageBloc>().add(
                                            SignInPageButtonPressed(
                                              email: emailController.text,
                                              password: passwordController.text,
                                            ),
                                          );
                                    },
                                    style: ElevatedButton.styleFrom(
                                      padding: const EdgeInsets.symmetric(vertical: 16.0),
                                      backgroundColor: const Color.fromARGB(255, 54, 104, 255),
                                      shape: RoundedRectangleBorder(
                                        borderRadius: BorderRadius.circular(8.0),
                                      ),
                                    ),
                                    child: const Text(
                                      'SIGN IN',
                                      style: TextStyle(
                                        fontSize: 16.0,
                                        fontWeight: FontWeight.w600,
                                        color: Colors.white,
                                      ),
                                    ),
                                  ),
                              ),
                        ],
                      );
                    },
                  ),
                ),
              ),
              const SizedBox(height: 100.0),
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  const Text(
                    "Don't have an account?",
                    style: TextStyle(
                      color: Colors.grey,
                    ),
                  ),
                  TextButton(
                    onPressed: () {
                      Navigator.pushNamed(context, '/signup');
                    },
                    child: const Text(
                      'SIGN UP',
                      style: TextStyle(
                        color: Color.fromARGB(255, 54, 104, 255),
                        fontWeight: FontWeight.w600,
                      ),
                    ),
                  ),
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }
}