import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:url_launcher/url_launcher.dart';

import '../../../product/presentation/widgets/custom_outlined_button.dart';
import '../../domain/entities/register_entity.dart';
import '../bloc/auth_bloc.dart';
import '../widgets/auth_text_field.dart';

class SignUpPage extends StatefulWidget {
  const SignUpPage({super.key});

  @override
  State<SignUpPage> createState() => _SignUpPageState();
}

class _SignUpPageState extends State<SignUpPage> {
  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();
  final TextEditingController nameController = TextEditingController();
  final TextEditingController emailController = TextEditingController();
  final TextEditingController passwordController = TextEditingController();
  final TextEditingController confirmPasswordController =
      TextEditingController();
  bool isChecked = false;

  void unfocusTextFields() {
    FocusScope.of(context).unfocus();
  }

  @override
  Widget build(BuildContext context) {
    return BlocListener<AuthBloc, AuthState>(
      listener: (context, state) {
        if (state is AuthRegisterSuccess) {
          if (!context.mounted) return;
          ScaffoldMessenger.of(context).showSnackBar(SnackBar(
            content: const Text('Account created successfully'),
            backgroundColor: Theme.of(context).primaryColor,
          ));
          Navigator.pushNamed(context, '/login');
        } else if (state is AuthError) {
          if (!context.mounted) return;
          ScaffoldMessenger.of(context).showSnackBar(SnackBar(
            content: Text(state.message),
            backgroundColor: Colors.red,
          ));
        }
      },
      child: SafeArea(
        child: Scaffold(
          backgroundColor: Colors.white,
          appBar: AppBar(
            actions: [
              Container(
                  decoration: BoxDecoration(
                      borderRadius: BorderRadius.circular(10),
                      color: Colors.white,
                      border:
                          Border.all(color: Theme.of(context).primaryColor)),
                  child: Padding(
                    padding: const EdgeInsets.symmetric(horizontal: 5),
                    child: Text(
                      'ECOM',
                      style: GoogleFonts.caveatBrush(
                        fontWeight: FontWeight.w600,
                        fontSize: 25,
                        color: Theme.of(context).primaryColor,
                      ),
                    ),
                  )),
              const SizedBox(
                width: 15,
              )
            ],
          ),
          body: GestureDetector(
            onTap: unfocusTextFields,
            child: SingleChildScrollView(
              child: Padding(
                padding:
                    const EdgeInsets.symmetric(vertical: 70, horizontal: 40),
                child: Center(
                  child: Form(
                    key: _formKey,
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.center,
                      children: [
                        const Text(
                          'Create your account',
                          style: TextStyle(
                              fontWeight: FontWeight.w600, fontSize: 27),
                        ),
                        const SizedBox(height: 20),
                        AuthTextField(
                          controller: nameController,
                          hintText: 'ex: Jon Smith',
                          validator: (text) {
                            if (text == null || text.isEmpty) {
                              return 'Name cannot be empty';
                            }
                            return null;
                          },
                          label: 'Name',
                        ),
                        const SizedBox(
                          height: 10,
                        ),
                        AuthTextField(
                          controller: emailController,
                          hintText: 'ex: jon.smith@email.com',
                          validator: (text) {
                            if (text == null || text.isEmpty) {
                              return 'Email cannot be empty';
                            } else if (!RegExp(
                                    r'^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$')
                                .hasMatch(text)) {
                              return 'Please enter a valid email';
                            }
                            return null;
                          },
                          label: 'Email',
                        ),
                        const SizedBox(
                          height: 10,
                        ),
                        AuthTextField(
                          controller: passwordController,
                          hintText: '********',
                          isObscure: true,
                          validator: (text) {
                            if (text == null || text.isEmpty) {
                              return 'Password cannot be empty';
                            } else if (text.length < 6) {
                              return 'Password must be at least 6 characters long';
                            }
                            return null;
                          },
                          label: 'Password',
                        ),
                        const SizedBox(
                          height: 10,
                        ),
                        AuthTextField(
                          controller: confirmPasswordController,
                          hintText: '********',
                          isObscure: true,
                          validator: (text) {
                            if (text == null || text.isEmpty) {
                              return 'Confirm Password cannot be empty';
                            } else if (text != passwordController.text) {
                              return 'Passwords do not match';
                            }
                            return null;
                          },
                          label: 'Confirm Password',
                        ),
                        const SizedBox(
                          height: 10,
                        ),
                        Row(
                          children: [
                            Transform.scale(
                              scale: 0.8,
                              child: Checkbox(
                                activeColor: Theme.of(context).primaryColor,
                                side: BorderSide(
                                    color: Theme.of(context).primaryColor,
                                    width: 1),
                                value: isChecked,
                                onChanged: (bool? value) {
                                  setState(() {
                                    isChecked = value!;
                                  });
                                },
                              ),
                            ),
                            RichText(
                              text: TextSpan(
                                text: 'I understood the ',
                                style: GoogleFonts.poppins(color: Colors.black),
                                children: [
                                  TextSpan(
                                    text: 'terms and conditions',
                                    style: GoogleFonts.poppins(
                                        color: Theme.of(context).primaryColor),
                                    recognizer: TapGestureRecognizer()
                                      ..onTap = () {
                                        launchUrl(
                                          Uri.parse('https://a2sv.org/')
                                        );
                                      },
                                  ),
                                ],
                              ),
                            ),
                          ],
                        ),
                        const SizedBox(
                          height: 20,
                        ),
                        BlocBuilder<AuthBloc, AuthState>(
                          builder: (context, state) {
                            return CustomOutlinedButton(
                              backgroundColor: Theme.of(context).primaryColor,
                              foregroundColor: Colors.white,
                              borderColor: Theme.of(context).primaryColor,
                              onPressed: () {
                                
                                if (_formKey.currentState!.validate()) {
                                  if (isChecked) {
                                    context.read<AuthBloc>().add(RegisterEvent(
                                      registrationEntity:
                                          RegistrationEntity(
                                              name: nameController.text,
                                              email: emailController.text,
                                              password:
                                                  passwordController.text),
                                    ));
                                  } else {
                                    ScaffoldMessenger.of(context)
                                        .showSnackBar(const SnackBar(
                                      content: Text(
                                          'Please accept the terms and conditions'),
                                      backgroundColor: Colors.red,
                                    ));
                                  }
                                }
                              },
                              buttonWidth: double.maxFinite,
                              buttonHeight: 45,
                              child: (state is AuthLoading)
                                  ? const CircularProgressIndicator(
                                      color: Colors.white,
                                    )
                                  : const Text(
                                      'SIGN UP',
                                      style: TextStyle(
                                          fontWeight: FontWeight.w600,
                                          fontSize: 16),
                                    ),
                            );
                          },
                        ),
                        const SizedBox(
                          height: 120,
                        ),
                        RichText(
                          text: TextSpan(
                            text: 'Already have an account? ',
                            style: GoogleFonts.poppins(
                                color: Colors.black38, fontSize: 18),
                            children: [
                              TextSpan(
                                text: 'SIGN IN',
                                style: GoogleFonts.poppins(
                                    color: Theme.of(context).primaryColor,
                                    fontSize: 17),
                                recognizer: TapGestureRecognizer()
                                  ..onTap = () {
                                    Navigator.pushNamed(context, '/login');
                                  },
                              ),
                            ],
                          ),
                        )
                      ],
                    ),
                  ),
                ),
              ),
            ),
          ),
        ),
      ),
    );
  }
}
