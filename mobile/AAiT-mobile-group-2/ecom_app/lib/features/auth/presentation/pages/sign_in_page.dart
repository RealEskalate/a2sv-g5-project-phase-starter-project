import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';

import '../../../product/presentation/widgets/custom_outlined_button.dart';
import '../../domain/entities/login_entity.dart';
import '../bloc/auth_bloc.dart';
import '../widgets/auth_text_field.dart';

class SignInPage extends StatefulWidget {
  const SignInPage({super.key});

  @override
  State<SignInPage> createState() => _SignInPageState();
}

class _SignInPageState extends State<SignInPage> {
  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();
  final TextEditingController emailController = TextEditingController();
  final TextEditingController passwordController = TextEditingController();

  void unfocusTextFields() {
    FocusScope.of(context).unfocus();
  }

  @override
  Widget build(BuildContext context) {
    return BlocListener<AuthBloc, AuthState>(
      listener: (context, state) {
        if (state is AuthSuccess) {
          if (!context.mounted) return;
          ScaffoldMessenger.of(context).showSnackBar(SnackBar(
            content: const Text('Login successful'),
            backgroundColor: Theme.of(context).primaryColor,
          ));
          Navigator.pushNamed(context, '/home');
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
          body: GestureDetector(
            onTap: unfocusTextFields,
            child: SingleChildScrollView(
              child: Padding(
                padding:
                    const EdgeInsets.symmetric(vertical: 100, horizontal: 40),
                child: Center(
                  child: Form(
                    key: _formKey,
                    child: Column(
                      // mainAxisAlignment: MainAxisAlignment.center,
                      crossAxisAlignment: CrossAxisAlignment.center,
                      children: [
                        Container(
                          decoration: BoxDecoration(
                              borderRadius: BorderRadius.circular(10),
                              color: Colors.white,
                              border: Border.all(
                                  color: Theme.of(context).primaryColor)),
                          child: Padding(
                            padding: const EdgeInsets.symmetric(horizontal: 20),
                            child: Text(
                              'ECOM',
                              style: GoogleFonts.caveatBrush(
                                  fontWeight: FontWeight.w600,
                                  fontSize: 50,
                                  color: Theme.of(context).primaryColor,
                                  decoration: TextDecoration.none),
                            ),
                          ),
                        ),
                        const SizedBox(
                          height: 60,
                        ),
                        const Text(
                          'Sign into your account',
                          style: TextStyle(
                              fontWeight: FontWeight.w600, fontSize: 27),
                        ),
                        const SizedBox(
                          height: 40,
                        ),
                        AuthTextField(
                            controller: emailController,
                            hintText: 'ex.jon.smith@email.com',
                            validator: (text) {
                              if (text!.isEmpty) {
                                return 'Email cannot be empty';
                              } else if (!RegExp(
                                      r'^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$')
                                  .hasMatch(text)) {
                                return 'Please enter a valid email';
                              }
                              return null;
                            },
                            label: 'Email'),
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
                            label: 'Password'),
                        const SizedBox(
                          height: 60,
                        ),
                        BlocBuilder<AuthBloc, AuthState>(
                          builder: (context, state) {
                            return CustomOutlinedButton(
                                backgroundColor: Theme.of(context).primaryColor,
                                foregroundColor: Colors.white,
                                borderColor: Theme.of(context).primaryColor,
                                onPressed: () {
                                  if (_formKey.currentState!.validate()) {
                                    unfocusTextFields();
                                    context.read<AuthBloc>().add(LoginEvent(
                                          loginEntity: LoginEntity(
                                              email: emailController.text,
                                              password:
                                                  passwordController.text),
                                        ));
                                  }
                                },
                                buttonWidth: double.maxFinite,
                                buttonHeight: 45,
                                child: (state is AuthLoading)
                                    ? const CircularProgressIndicator(
                                        color: Colors.white,
                                      )
                                    : const Text(
                                        'SIGN IN',
                                        style: TextStyle(
                                            fontWeight: FontWeight.w600,
                                            fontSize: 16),
                                      ));
                          },
                        ),
                        const SizedBox(
                          height: 120,
                        ),
                        RichText(
                          text: TextSpan(
                              text: "Don't have an account? ",
                              style: GoogleFonts.poppins(
                                  color: Colors.black38, fontSize: 18),
                              children: [
                                TextSpan(
                                    text: 'SIGN UP',
                                    style: GoogleFonts.poppins(
                                        color: Theme.of(context).primaryColor,
                                        fontSize: 17),
                                    recognizer: TapGestureRecognizer()
                                      ..onTap = () {
                                        Navigator.pushNamed(context, '/signup');
                                      })
                              ]),
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
