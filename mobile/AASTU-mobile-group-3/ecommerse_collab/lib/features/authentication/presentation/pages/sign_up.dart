import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../bloc/blocs.dart';
import '../bloc/events.dart';
import '../bloc/states.dart';
import '../widgets/custom_button.dart';
import '../widgets/linked_text.dart';
import '../widgets/logo.dart';
import '../widgets/textField.dart';
import '../widgets/title.dart';
import 'sign_in.dart';

class SignUp extends StatefulWidget {
  const SignUp({super.key});

  @override
  State<SignUp> createState() => _SignUpState();
}

class _SignUpState extends State<SignUp> {
  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();
  final TextEditingController nameController = TextEditingController();
  final TextEditingController emailController = TextEditingController();
  final TextEditingController passwordController = TextEditingController();
  final TextEditingController confirmPasswordController = TextEditingController();
  bool _isChecked = false;

  @override
  Widget build(BuildContext context) {
    return BlocConsumer<UserBloc, UserState>(
      listener: (context, state) {
        if (state is RegisterLoadingState) {
          ScaffoldMessenger.of(context).showSnackBar(
            const SnackBar(
              content: Text("Registering User"),
            ),
          );
        } else if (state is RegisterErrorState) {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(
              content: Text(state.message),
              backgroundColor: Colors.red,
            ),
          );
        } else if (state is UserRegisteredState) {
          ScaffoldMessenger.of(context).showSnackBar(
            const SnackBar(
              content: Text("User Registered"),
              backgroundColor: Colors.green,
            ),
          );
          Navigator.push(
            context,
            MaterialPageRoute(builder: (context) => const SignIn()),
          );
        }
      },
      builder: (context, state) {
        return Scaffold(
          appBar: AppBar(
            actions: const [Logo(width: 37, height: 15, fontSize: 24)],
          ),
          body: Padding(
            padding: const EdgeInsets.all(16.0),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                Expanded(
                  child: SingleChildScrollView(
                    child: Form(
                      key: _formKey,
                      child: Column(
                        children: [
                        const BigTitle(text: "Create an account"),
                          CustomTextField(
                            name: 'Name',
                            controller: nameController,
                            hintText: 'ex: jon smith',
                            isPassword: false,
                            validator: (value) {
                              if (value == null || value.isEmpty) {
                                return 'Please enter your name';
                              }
                              return null;
                            },
                          ),
                          CustomTextField(
                            name: 'Email',
                            controller: emailController,
                            hintText: 'ex: jon.smith@email.com',
                            isPassword: false,
                            validator: (value) {
                              if (value == null || value.isEmpty) {
                                return 'Please enter your email';
                              } else if (!RegExp(r'^[^@]+@[^@]+\.[^@]+').hasMatch(value)) {
                                return 'Enter a valid email';
                              }
                              return null;
                            },
                          ),
                          CustomTextField(
                            name: 'Password',
                            controller: passwordController,
                            hintText: '*********',
                            isPassword: true,
                            validator: (value) {
                              if (value == null || value.isEmpty) {
                                return 'Please enter your password';
                              } else if (value.length < 6) {
                                return 'Password must be at least 6 characters';
                              }
                              return null;
                            },
                          ),
                          CustomTextField(
                            name: 'Confirm Password',
                            controller: confirmPasswordController,
                            hintText: '*********',
                            isPassword: true,
                            validator: (value) {
                              if (value == null || value.isEmpty) {
                                return 'Please confirm your password';
                              } else if (value != passwordController.text) {
                                return 'Passwords do not match';
                              }
                              return null;
                            },
                          ),
                          Row(
                            children: [
                              Checkbox(
                                value: _isChecked,
                                onChanged: (bool? value) {
                                  setState(() {
                                    _isChecked = value ?? false;
                                  });
                                },
                              ),
                              const Text(
                                'I understood the ',
                                style: TextStyle(
                                  fontSize: 14,
                                  fontFamily: 'Poppins',
                                ),
                              ),
                              const Text(
                                'terms & policy.',
                                style: TextStyle(
                                  fontSize: 14,
                                  fontFamily: 'Poppins',
                                  color: Color(0xFF3E50F3),
                                ),
                              ),
                            ],
                          ),
                          CustomButton(
                            name: "SIGN UP",
                            login: false,
                            formKey: _formKey,
                            isCheckboxChecked: _isChecked,
                            controllers: [
                              emailController,
                              passwordController,
                              nameController,
                              confirmPasswordController,
                            ],
                            onPressed: () {},
                          ),
                          const SizedBox(height: 40),
                          LinkedText(
                            description: "Have an account?",
                            linkedText: " SIGN IN",
                            navigateTo: const SignIn(),
                          ),
                        ],
                      ),
                    ),
                  ),
                ),
              ],
            ),
          ),
        );
      },
    );
  }
}
