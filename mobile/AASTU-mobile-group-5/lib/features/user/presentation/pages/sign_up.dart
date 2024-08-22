import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';

import '../../../../service_locator.dart';
import '../../../product/presentation/widgets/custom_back_button.dart';
import '../../domain/repositories/user_repository.dart';
import '../bloc/sign_up_page/sign_up_page_bloc.dart';

class SignUpPage extends StatefulWidget {
  const SignUpPage({super.key});

  @override
  // ignore: library_private_types_in_public_api
  _SignUpPageState createState() => _SignUpPageState();
}

class _SignUpPageState extends State<SignUpPage> {
  final TextEditingController nameController = TextEditingController();
  final TextEditingController emailController = TextEditingController();
  final TextEditingController passwordController = TextEditingController();
  final TextEditingController confirmPasswordController =
      TextEditingController();
  bool isChecked = false;
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        backgroundColor: Colors.white,
        elevation: 0,
        leading: const CustomBackButton(),
        actions: [
          Padding(
            padding: const EdgeInsets.all(10.0),
            child: Container(
              width: 60,
              height: 40,
              padding: const EdgeInsets.only(left: 5, right: 5),
              decoration: BoxDecoration(
                color: Colors.white,
                borderRadius: BorderRadius.circular(8.0),
                border: Border.all(
                    color: const Color.fromARGB(255, 54, 104, 255), width: 2),
              ),
              child: Center(
                child: Text(
                  'ECOM',
                  style: GoogleFonts.caveatBrush(
                    fontSize: 20.0,
                    fontWeight: FontWeight.bold,
                    color: const Color.fromARGB(255, 54, 104, 255),
                  ),
                  textAlign: TextAlign.center,
                ),
              ),
            ),
          ),
          const SizedBox(width: 6)
        ],
      ),
      body: SingleChildScrollView(
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 24.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: [
              const SizedBox(height: 24.0),
              // Create your account text
              Text(
                'Create your account',
                textAlign: TextAlign.left,
                style: GoogleFonts.poppins(
                  fontSize: 24.0,
                  fontWeight: FontWeight.bold,
                  color: Colors.black,
                ),
              ),
              const SizedBox(height: 24.0),
              Text(
                'Name',
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
                controller: nameController,
                decoration: InputDecoration(
                  hintText: 'ex: Jon Smith',
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

              // Name input
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

              Text(
                'Password',
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

              Text(
                'Confirm Password',
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
                controller: confirmPasswordController,
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
              // Terms and policy checkbox
              Row(
                children: [
                  Checkbox(
                    value: isChecked,
                    onChanged: (bool? value) {
                      setState(() {
                        isChecked = value!;
                      });
                    },
                    activeColor: const Color.fromARGB(255, 54, 104, 255),
                    checkColor: Colors.white,
                  ),
                  RichText(
                    text: TextSpan(
                      text: 'I understood the ',
                      style: GoogleFonts.poppins(color: Colors.black),
                      children: <TextSpan>[
                        TextSpan(
                          text: 'terms & policy.',
                          style: const TextStyle(
                            color: Color.fromARGB(255, 54, 104, 255),
                          ),
                          recognizer: TapGestureRecognizer()..onTap = () {},
                        ),
                      ],
                    ),
                  ),
                ],
              ),
              const SizedBox(height: 32.0),
              BlocProvider(
                  create: (context) =>
                      SignUpPageBloc(userRepository: getIt<UserRepository>()),
                  child: BlocListener<SignUpPageBloc, SignUpPageState>(
                    listener: (context, state) {
                      if (state is SignUpPageFailure) {
                        ScaffoldMessenger.of(context).showSnackBar(
                          SnackBar(content: Text(state.error)),
                        );
                      } else if (state is SignUpPageSuccess) {
                        Navigator.pushReplacementNamed(context, '/signin');
                      }
                    },
                    child: BlocBuilder<SignUpPageBloc, SignUpPageState>(
                        builder: (context, state) {
                      return Column(
                        children: [
                          state is SignUpPageLoading
                              ? const CircularProgressIndicator()
                              : SizedBox(
                                  width: double.infinity,
                                  child: ElevatedButton(
                                    // onPressed: () {
                                    //   if (!isChecked) {
                                    //     ScaffoldMessenger.of(context)
                                    //         .showSnackBar(
                                    //       const SnackBar(
                                    //           content: Text(
                                    //               'Please agree to the terms and conditions')),
                                    //     );
                                    //   } else {
                                    //     context.read<SignUpPageBloc>().add(
                                    //           SignUpPageButtonPressed(
                                    //             email: emailController.text,
                                    //             password: passwordController.text,
                                    //             confirmPassword:
                                    //                 confirmPasswordController
                                    //                     .text,
                                    //             name: nameController.text,
                                    //           ),
                                    //         );
                                    //   }
                                    // },
                                    onPressed: () {
                                      if (!isChecked) {
                                        ScaffoldMessenger.of(context)
                                            .showSnackBar(
                                          const SnackBar(
                                              content: Text(
                                                  'Please agree to the terms and conditions')),
                                        );
                                      } else if (passwordController.text !=
                                          confirmPasswordController.text) {
                                        ScaffoldMessenger.of(context)
                                            .showSnackBar(
                                          const SnackBar(
                                              content: Text(
                                                  'Passwords do not match')),
                                        );
                                      } else {
                                        context.read<SignUpPageBloc>().add(
                                              SignUpPageButtonPressed(
                                                email: emailController.text,
                                                password:
                                                    passwordController.text,
                                                confirmPassword:
                                                    confirmPasswordController
                                                        .text,
                                                name: nameController.text,
                                              ),
                                            );
                                      }
                                    },
                                    style: ElevatedButton.styleFrom(
                                      padding: const EdgeInsets.symmetric(
                                          vertical: 16.0),
                                      backgroundColor: const Color.fromARGB(
                                          255, 54, 104, 255),
                                      shape: RoundedRectangleBorder(
                                        borderRadius:
                                            BorderRadius.circular(8.0),
                                      ),
                                    ),
                                    child: Text(
                                      'SIGN UP',
                                      style: GoogleFonts.poppins(
                                        fontSize: 16.0,
                                        fontWeight: FontWeight.w600,
                                        color: Colors.white,
                                      ),
                                    ),
                                  ),
                                ),
                          const SizedBox(height: 30.0),
                          // Sign In text
                          Row(
                            mainAxisAlignment: MainAxisAlignment.center,
                            children: [
                              Text(
                                'Have an account?',
                                style: GoogleFonts.poppins(color: Colors.grey),
                              ),
                              TextButton(
                                onPressed: () {
                                  Navigator.pushNamed(context, '/signin');
                                },
                                child: Text(
                                  'SIGN IN',
                                  style: GoogleFonts.poppins(
                                    color:
                                        const Color.fromARGB(255, 54, 104, 255),
                                    fontWeight: FontWeight.w600,
                                  ),
                                ),
                              ),
                            ],
                          ),
                        ],
                      );
                    }),
                  ))
            ],
          ),
        ),
      ),
    );
  }
}
