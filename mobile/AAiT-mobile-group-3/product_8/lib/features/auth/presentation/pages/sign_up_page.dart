import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';

import '../../../product/presentation/widgets/custom_buttom.dart';
import '../../domain/entities/sign_up_user_entitiy.dart';
import '../bloc/auth_bloc/auth_bloc.dart';
import '../widget/text_field.dart';

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
            content: const Text('Sign Up successful'),
            backgroundColor: Theme.of(context).primaryColor,
          ));
          Navigator.pushNamed(context, '/signin');
        } else if (state is AuthError) {
          if (!context.mounted) return;
          ScaffoldMessenger.of(context).showSnackBar(SnackBar(
            content: Text(state.message),
            backgroundColor: Colors.red,
          ));
        }
      },
      child: Scaffold(
        appBar: AppBar(
          leading: IconButton(
              onPressed: () => {Navigator.pop(context)},
              icon: Icon(
                Icons.arrow_back_ios_outlined,
                color: Colors.indigoAccent.shade400,
              )),
          actions: [
            Container(
                decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(10),
                    color: Colors.white,
                    border: Border.all(color: Theme.of(context).primaryColor)),
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
            child: Center(
              child: Padding(
                padding: const EdgeInsets.all(20.0),
                child: Form(
                  key: _formKey,
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.start,
                    crossAxisAlignment: CrossAxisAlignment.center,
                    children: [
                      const Text('Create your account',
                          style: TextStyle(
                            color: Colors.black,
                            fontSize: 30,
                            fontWeight: FontWeight.bold,
                          )),
                      TextFieldWidget(
                        controller: nameController,
                        hintText: 'ex: jon smith',
                        obscureText: 'Name',
                        validator: (text) {
                          if (text == null || text.isEmpty) {
                            return 'Name cannot be empty';
                          }
                          return null;
                        },
                      ),
                      TextFieldWidget(
                        controller: emailController,
                        hintText: 'ex: jon.smith@email.com',
                        obscureText: 'Email',
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
                      ),
                      TextFieldWidget(
                        isObscure: true,
                        controller: passwordController,
                        hintText: '**********',
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
                      TextFieldWidget(
                        isObscure: true,
                        controller: confirmPasswordController,
                        hintText: '**********',
                        obscureText: 'Confirm Password',
                        validator: (text) {
                          if (text == null || text.isEmpty) {
                            return 'Confirm Password cannot be empty';
                          } else if (text != passwordController.text) {
                            return 'Passwords do not match';
                          }
                          return null;
                        },
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
                                }),
                          ),
                          const Text('I understood the '),
                          const Text(
                            'term & policy.',
                            style: TextStyle(
                                color: Color.fromRGBO(63, 81, 243, 1),
                                fontWeight: FontWeight.w200),
                          ),
                        ],
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
                                    'SIGN UP',
                                    style:
                                        TextStyle(fontWeight: FontWeight.w600),
                                  ),
                            onPressed: () {
                              if (_formKey.currentState!.validate()) {
                                if (isChecked) {
                                  context.read<AuthBloc>().add(SignupEvent(
                                      signUpUserEntitiy: SignUpUserEntitiy(
                                          email: emailController.text,
                                          password: passwordController.text,
                                          name: nameController.text)));
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
                          );
                        },
                      ),
                      const SizedBox(
                        height: 40,
                      ),
                      Row(
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [
                          const Text(
                            'Have an account?',
                            style: TextStyle(
                                color: Color.fromRGBO(111, 111, 111, 15)),
                          ),
                          TextButton(
                            onPressed: () {
                              Navigator.pushNamed(context, '/signin');
                            },
                            child: const Text(
                              'SIGN IN',
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
