import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';

import '../../../product/presentation/widgets/custom_buttom.dart';
import '../../domain/entities/sign_in_user_entitiy.dart';
import '../bloc/auth_bloc/auth_bloc.dart';
import '../widget/text_field.dart';

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
          // if (!context.mounted) return;
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
      child: Scaffold(
        body: GestureDetector(
          onTap: unfocusTextFields,
          child: SingleChildScrollView(
            child: Center(
              child: Padding(
                padding: const EdgeInsets.all(20.0),
                child: Form(
                  key: _formKey,
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.start,
                    crossAxisAlignment: CrossAxisAlignment.center,
                    children: [
                      const SizedBox(
                        height: 100,
                      ),
                      Card(
                        shadowColor: Colors.black,
                        color: Colors.white,
                        shape: const RoundedRectangleBorder(
                          side: BorderSide(
                            color: Color.fromRGBO(63, 81, 243, 1),
                          ),
                          borderRadius: BorderRadius.all(Radius.circular(10)),
                        ),
                        child: Padding(
                          padding: const EdgeInsets.only(
                              top: 0, bottom: 0, right: 13, left: 13),
                          child: Text(
                            'ECOM',
                            style: GoogleFonts.caveatBrush(
                              color: const Color.fromRGBO(63, 81, 243, 1),
                              fontSize: 50,
                              fontWeight: FontWeight.bold,
                            ),
                          ),
                        ),
                      ),
                      const SizedBox(
                        height: 45,
                      ),
                      const Text('Sign into your account',
                          style: TextStyle(
                            color: Colors.black,
                            fontSize: 30,
                            fontWeight: FontWeight.bold,
                          )),
                      const SizedBox(
                        height: 20,
                      ),
                      TextFieldWidget(
                        controller: emailController,
                        hintText: 'ex: jon.smith@email.com',
                        obscureText: 'Email',
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
                      ),
                      TextFieldWidget(
                        isObscure: true,
                        controller: passwordController,
                        hintText: '*********',
                        obscureText: 'Password',
                        validator: (text) {
                          if (text == null || text.isEmpty) {
                            return 'Password cannot be empty';
                          } else if (text.length < 6) {
                            return 'Password must be at least 6 characters long';
                          }
                          return null;
                        },
                      ),
                      const SizedBox(
                        height: 30,
                      ),
                      BlocBuilder<AuthBloc, AuthState>(
                        builder: (context, state) {
                          return CustomButton(
                            backgroundColor: Theme.of(context).primaryColor,
                            foregroundColor: Colors.white,
                            borderColor: Theme.of(context).primaryColor,
                            buttonWidth: double.maxFinite,
                            buttonHeight: 45,
                            child: (state is AuthLoading)
                                ? const CircularProgressIndicator(
                                    color: Colors.white,
                                  )
                                : const Text(
                                    'SIGN IN',
                                    style:
                                        TextStyle(fontWeight: FontWeight.w600),
                                  ),
                            onPressed: () {
                              if (_formKey.currentState!.validate()) {
                                unfocusTextFields();
                                context.read<AuthBloc>().add(SigninEvent(
                                    signInUserEntitiy: SignInUserEntitiy(
                                        email: emailController.text,
                                        password: passwordController.text)));
                              }
                            },
                          );
                        },
                      ),
                      const SizedBox(
                        height: 50.0,
                      ),
                      Row(
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [
                          const Text(
                            'Don\'t have an account?',
                            style: TextStyle(
                                color: Color.fromRGBO(111, 111, 111, 15)),
                          ),
                          TextButton(
                            onPressed: () {
                              Navigator.pushNamed(context, '/signup');
                            },
                            child: const Text(
                              'SIGN UP',
                              style: TextStyle(
                                  color: Color.fromRGBO(63, 81, 243, 1),
                                  fontWeight: FontWeight.w200),
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
      ),
    );
  }
}
