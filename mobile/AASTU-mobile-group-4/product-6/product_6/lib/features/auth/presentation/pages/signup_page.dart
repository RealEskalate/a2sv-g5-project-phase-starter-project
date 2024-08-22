import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';

import '../../domain/entity/auth_entity.dart';
import '../bloc/auth_bloc.dart';
import '../bloc/auth_event.dart';
import '../bloc/auth_state.dart';
import '../widgets/logo.dart';
import '../widgets/text_fields.dart';
import 'login_page.dart';

class SignupPage extends StatelessWidget {
  final String text;
  SignupPage({super.key, required this.text});

  final usernameController = TextEditingController();
  final passwordController = TextEditingController();
  final emailController = TextEditingController();
  final confirmPasswordController = TextEditingController();

  // Key for the form
  final _formKey = GlobalKey<FormState>();

  // Validation functions
  String? validateUsername(String? value) {
    if (value == null || value.isEmpty) {
      return 'Please enter a username';
    }
    return null;
  }

  String? validateEmail(String? value) {
    if (value == null || value.isEmpty) {
      return 'Please enter an email';
    }
    // Regular expression for email validation
    final RegExp emailRegExp = RegExp(
        r'^[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z]{2,}$');
    if (!emailRegExp.hasMatch(value)) {
      return 'Please enter a valid email';
    }
    return null;
  }

  String? validatePassword(String? value) {
    if (value == null || value.isEmpty) {
      return 'Please enter a password';
    }
    if (value.length < 6) {
      return 'Password must be at least 6 characters';
    }
    return null;
  }

  String? validateConfirmPassword(String? value) {
    if (value == null || value.isEmpty) {
      return 'Please confirm your password';
    }
    if (value != passwordController.text) {
      return 'Passwords do not match';
    }
    return null;
  }

  @override
  Widget build(BuildContext context) {
    return BlocListener<AuthBloc, AuthState>(
      listener: (context, state) {
        print('from sign up page');
        print(state);
        if (state is UserSuccessRegister) {
          ScaffoldMessenger.of(context).showSnackBar(
              SnackBar(content: Text('Sign Up successfully')));
          Navigator.push(
            context,
            MaterialPageRoute(
              builder: (context) => LoginPage(
                text: 'SIGN IN',
              ),
            ),
          );
        } else if (state is AuthError) {
          ScaffoldMessenger.of(context).showSnackBar(
              SnackBar(content: Text('Sign Up failed')));
        }
      },
      child: Scaffold(
        backgroundColor: Colors.white,
        appBar: AppBar(
          backgroundColor: Colors.white,
          leading: Padding(
            padding: const EdgeInsets.symmetric(horizontal: 25),
            child: IconButton(
              onPressed: () {
                Navigator.pop(context);
              },
              icon: const Icon(
                Icons.arrow_back_ios,
                size: 18,
                color: Color.fromRGBO(63, 81, 243, 1),
              ),
            ),
          ),
          actions: [
            Padding(
              padding: EdgeInsets.symmetric(horizontal: 25),
              child: Logo(
                height: 30,
                textheight: 20,
                paddingBottom: 0,
                paddingHorizontal: 10,
              ),
            )
          ],
        ),
        body: SafeArea(
          child: Center(
            child: SingleChildScrollView(
              child: Form(
                key: _formKey,
                child: Column(
                  children: [
                    SizedBox(height: 40),
                    Text(
                      'Create an account',
                      style: GoogleFonts.poppins(
                        textStyle: const TextStyle(
                          fontWeight: FontWeight.w600,
                          fontSize: 25,
                          color: Color.fromRGBO(0, 0, 0, 1),
                        ),
                      ),
                    ),
                    SizedBox(height: 20),
                    Padding(
                      padding: const EdgeInsets.symmetric(horizontal: 25),
                      child: Row(
                        children: [
                          Text(
                            'Username',
                            style: GoogleFonts.poppins(
                              textStyle: const TextStyle(
                                fontWeight: FontWeight.w400,
                                fontSize: 16,
                                color: Color.fromRGBO(111, 111, 111, 1),
                              ),
                            ),
                          ),
                        ],
                      ),
                    ),
                    SizedBox(height: 5),
                    Fields(
                      controller: usernameController,
                      hintText: 'ex: jon smith',
                      obsecureText: false,
                      validator: validateUsername,
                    ),
                    SizedBox(height: 5),
                    Padding(
                      padding: const EdgeInsets.symmetric(horizontal: 25),
                      child: Row(
                        children: [
                          Text(
                            'Email',
                            style: GoogleFonts.poppins(
                              textStyle: const TextStyle(
                                fontWeight: FontWeight.w400,
                                fontSize: 16,
                                color: Color.fromRGBO(111, 111, 111, 1),
                              ),
                            ),
                          ),
                        ],
                      ),
                    ),
                    SizedBox(height: 5),
                    Fields(
                      controller: emailController,
                      hintText: 'ex: jon.smith@email.com',
                      obsecureText: false,
                      validator: validateEmail,
                    ),
                    SizedBox(height: 5),
                    Padding(
                      padding: const EdgeInsets.symmetric(horizontal: 25),
                      child: Row(
                        children: [
                          Text(
                            'Password',
                            style: GoogleFonts.poppins(
                              textStyle: const TextStyle(
                                fontWeight: FontWeight.w400,
                                fontSize: 16,
                                color: Color.fromRGBO(111, 111, 111, 1),
                              ),
                            ),
                          ),
                        ],
                      ),
                    ),
                    SizedBox(height: 5),
                    Fields(
                      controller: passwordController,
                      hintText: '*********',
                      obsecureText: true,
                      validator: validatePassword,
                    ),
                    SizedBox(height: 5),
                    Padding(
                      padding: const EdgeInsets.symmetric(horizontal: 25),
                      child: Row(
                        children: [
                          Text(
                            'Confirm Password',
                            style: GoogleFonts.poppins(
                              textStyle: const TextStyle(
                                fontWeight: FontWeight.w400,
                                fontSize: 16,
                                color: Color.fromRGBO(111, 111, 111, 1),
                              ),
                            ),
                          ),
                        ],
                      ),
                    ),
                    SizedBox(height: 5),
                    Fields(
                      controller: confirmPasswordController,
                      hintText: '*********',
                      obsecureText: true,
                      validator: validateConfirmPassword,
                    ),
                    Padding(
                      padding: const EdgeInsets.symmetric(horizontal: 20.0),
                      child: Row(
                        children: [
                          Checkbox(value: false, onChanged: (bool? value) {}),
                          Text(
                            'I understood ',
                            style: GoogleFonts.poppins(
                              textStyle: const TextStyle(
                                fontWeight: FontWeight.w400,
                                fontSize: 12,
                                color: Color.fromRGBO(0, 0, 0, 1),
                              ),
                            ),
                          ),
                          Text(
                            'terms & policy.',
                            style: GoogleFonts.poppins(
                              textStyle: const TextStyle(
                                fontWeight: FontWeight.w400,
                                fontSize: 12,
                                color: Color.fromRGBO(63, 81, 243, 1),
                              ),
                            ),
                          ),
                        ],
                      ),
                    ),
                    ElevatedButton(
                      style: ElevatedButton.styleFrom(
                        minimumSize: Size(308, 45),
                        padding: const EdgeInsets.symmetric(horizontal: 25),
                        backgroundColor: Color.fromRGBO(63, 81, 243, 1),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(8),
                        ),
                      ),
                      onPressed: () {
                        if (_formKey.currentState!.validate()) {
                          var sign = SentUserEntity(
                            name: usernameController.text,
                            email: emailController.text,
                            password: passwordController.text,
                          );
                          context
                              .read<AuthBloc>()
                              .add(RegisterEvent(senduserentity: sign));
                        }
                      },
                      child: Text(
                        'Sign Up',
                        style: GoogleFonts.poppins(
                          textStyle: const TextStyle(
                            fontWeight: FontWeight.w600,
                            fontSize: 15,
                            color: Color.fromRGBO(255, 255, 255, 1),
                          ),
                        ),
                      ),
                    ),
                    SizedBox(height: 40),
                    Row(
                      mainAxisAlignment: MainAxisAlignment.center,
                      children: [
                        Text(
                          'Have an account? ',
                          style: GoogleFonts.poppins(
                            textStyle: const TextStyle(
                              fontWeight: FontWeight.w400,
                              fontSize: 16,
                              color: Color.fromRGBO(136, 136, 136, 1),
                            ),
                          ),
                        ),
                        GestureDetector(
                          onTap: () {
                            Navigator.push(
                              context,
                              MaterialPageRoute(
                                builder: (context) => LoginPage(text: 'SIGN IN'),
                              ),
                            );
                          },
                          child: Text(
                            'Sign In.',
                            style: GoogleFonts.poppins(
                              textStyle: const TextStyle(
                                fontWeight: FontWeight.w600,
                                fontSize: 16,
                                color: Color.fromRGBO(63, 81, 243, 1),
                              ),
                            ),
                          ),
                        ),
                      ],
                    ),
                  ],
                ),
              ),
            ),
          ),
        ),
      ),
    );
  }
}
