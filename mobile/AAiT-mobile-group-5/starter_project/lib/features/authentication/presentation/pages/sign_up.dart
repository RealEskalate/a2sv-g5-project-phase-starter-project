import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:starter_project/features/authentication/presentation/pages/sign_in.dart';
import 'package:starter_project/features/authentication/presentation/pages/terms_and_policy.dart';
import 'package:starter_project/features/authentication/presentation/widgets/authentication_text_field.dart';
import 'package:starter_project/features/authentication/presentation/widgets/ecom.dart';
import 'package:starter_project/features/authentication/presentation/widgets/redirect.dart';

class SignUpPage extends StatelessWidget {
  const SignUpPage({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
            leading: IconButton(
              onPressed: () { Navigator.push(
                context,
                MaterialPageRoute(builder: (context) => const SignInPage()),
              );},
              icon: const Icon(
                Icons.arrow_back_ios_new_rounded,
                size: 20,
                color: Color(0xFF3F51F3),
              ),
            ),
            actions: const [ECOM(fontSize: 15), SizedBox(width: 30)]),
        body: Padding(
          padding: const EdgeInsets.all(25.0),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              Text("Sign into your account",
                  style: GoogleFonts.poppins(
                      textStyle: const TextStyle(
                    color: Colors.black,
                    fontWeight: FontWeight.bold,
                    fontSize: 25,
                  ))),
              Center(
                child: Padding(
                  padding: const EdgeInsets.all(16.0),
                  child: Column(
                    children: [
                      AuthenticationTextField(
                          labelText: 'Name',
                          hintText: 'ex: jon smith',
                          controller: TextEditingController()),
                      AuthenticationTextField(
                          labelText: 'Email',
                          hintText: 'ex: jonsmith@gmail.com',
                          controller: TextEditingController()),
                      AuthenticationTextField(
                          labelText: 'Password',
                          hintText: '***********',
                          controller: TextEditingController(),
                          isPassword: true),
                      AuthenticationTextField(
                          labelText: 'confirm Password',
                          hintText: '***********',
                          controller: TextEditingController(),
                          isPassword: true),
                      const Row(
                        children: [
                          TermsCheckbox(),
                          Redirect(
                              text: "I understood the terms ",
                              buttonText: "terms & policy.",
                              navigateTo: TermsAndPolicy()),
                        ],
                      ),
                      SizedBox(
                        height: 40,
                        width: double.infinity,
                        child: FilledButton(
                          onPressed: () {},
                          style: ButtonStyle(
                              backgroundColor: WidgetStateProperty.all(
                                  const Color(0xFF3F51F3)),
                              shape: WidgetStateProperty.all<
                                      RoundedRectangleBorder>(
                                  RoundedRectangleBorder(
                                      borderRadius:
                                          BorderRadius.circular(12.0)))),
                          child: const Text("SIGN IN"),
                        ),
                      ),
                      const Redirect(
                          text: " Have an account? ",
                          buttonText: "SIGN IN",
                          navigateTo: SignInPage())
                    ],
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}

class TermsCheckbox extends StatefulWidget {
  const TermsCheckbox({super.key});

  @override
  TermsCheckboxState createState() => TermsCheckboxState();
}

class TermsCheckboxState extends State<TermsCheckbox> {
  bool isChecked = false;

  @override
  Widget build(BuildContext context) {
    return Checkbox(
      fillColor:
          WidgetStateProperty.resolveWith<Color>((Set<WidgetState> states) {
        if (states.contains(WidgetState.selected)) {
          return Colors.blue;
        }
        return Colors.white;
      }),
      checkColor: Colors.white,
      value: isChecked,
      onChanged: (bool? value) {
        setState(() {
          isChecked = value!;
        });
      },
    );
  }
}
