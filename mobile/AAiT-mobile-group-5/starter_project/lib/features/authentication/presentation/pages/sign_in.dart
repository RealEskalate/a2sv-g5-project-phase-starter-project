import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:starter_project/features/authentication/presentation/pages/sign_up.dart';
import 'package:starter_project/features/authentication/presentation/widgets/authentication_text_field.dart';
import 'package:starter_project/features/authentication/presentation/widgets/ecom.dart';
import 'package:starter_project/features/authentication/presentation/widgets/redirect.dart';

class SignInPage extends StatelessWidget {
  const SignInPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: Padding(
          padding: const EdgeInsets.all(25.0),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              const ECOM(),
              const SizedBox(height: 60),
              Text(
                "Sign into your account",
                style: GoogleFonts.poppins(
                  textStyle: const TextStyle(
                    color: Colors.black,
                    fontWeight: FontWeight.bold,
                    fontSize: 25,
                  ),
                ),
              ),
              Center(
                child: Padding(
                  padding: const EdgeInsets.all(16.0),
                  child: Column(
                    children: [
                      AuthenticationTextField(
                        labelText: 'Name',
                        hintText: 'ex: jon smith',
                        controller: TextEditingController(),
                      ),
                      AuthenticationTextField(
                        labelText: 'Password',
                        hintText: '***********',
                        controller: TextEditingController(),
                        isPassword: true,
                      ),
                      SizedBox(
                        height: 40,
                        width: double.infinity,
                        child: FilledButton(
                          onPressed: () {},
                          style: ButtonStyle(
                            backgroundColor:
                                WidgetStateProperty.all(const Color(0xFF3F51F3)),
                            shape:
                                WidgetStateProperty.all<RoundedRectangleBorder>(
                              RoundedRectangleBorder(
                                borderRadius: BorderRadius.circular(12.0),
                              ),
                            ),
                          ),
                          child: const Text("SIGN IN"),
                        ),
                      ),
                      const Redirect(
                        text: " Don’t have an account?",
                        buttonText: "SIGN UP",
                        navigateTo: SignUpPage(),
                      )
                    ],
                  ),
                ),
              ),
            ],
          ),
        ),
      );
    
  }
}
