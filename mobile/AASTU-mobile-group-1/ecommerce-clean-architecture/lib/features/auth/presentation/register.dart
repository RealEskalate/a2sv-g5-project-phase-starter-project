import 'dart:ui';
import 'package:ecommerce/features/auth/presentation/bloc/authbloc/auth_bloc.dart';
import 'package:ecommerce/features/auth/presentation/bloc/authbloc/auth_state.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter/widgets.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';

import '../../../service_locator.dart';
import 'bloc/authbloc/auth_event.dart';
import 'login.dart';

class Register extends StatefulWidget {
  const Register({super.key});

  @override
  State<Register> createState() => _RegisterState();
}

class _RegisterState extends State<Register> {
  TextEditingController nameController = TextEditingController();
  TextEditingController emailController = TextEditingController();
  TextEditingController passwordController = TextEditingController();
  TextEditingController confirmPasswordController = TextEditingController();
  var user_bloc = getIt<UserBloc>();
  bool _isChecked = true;
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      body: SingleChildScrollView(
        child: BlocListener<UserBloc, UserState>(
          listener: (context, state) {
            print(state);
            print('from register');
            if (state is registered) {
              ScaffoldMessenger.of(context).showSnackBar(
                const SnackBar(
                  content: Text('User Registered Successfully'),
                ),
              );
              Navigator.push(
                  context, MaterialPageRoute(builder: (context) => const login()));
            }else if(state is registering){
              ScaffoldMessenger.of(context).showSnackBar(
                const SnackBar(
                  content: Text('Registering User'),
                ),
              );
              }else if (state is registerfailure) {
              ScaffoldMessenger.of(context).showSnackBar(
                SnackBar(
                  content: Text(state.message + ' Please try again'),
                ),
              );
            }
          },
          child: Padding(
            padding: const EdgeInsets.symmetric(horizontal: 30, vertical: 80),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Container(
                  // padding: const EdgeInsets.symmetric(horizontal: 30, vertical: 10),
                  color: Colors.white,
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      IconButton(
                        onPressed: () {
                          Navigator.pop(context);
                        },
                        icon: const Icon(Icons.arrow_back_ios,
                            color: Color(0xFF3F51F3), size: 17),
                      ),
                      Container(
                        width: 80,
                        height: 30,
                        decoration: BoxDecoration(
                          color: Colors.white,
                          borderRadius: BorderRadius.circular(10),
                          border:
                              Border.all(color: Color(0xFF3F51F3), width: 1),
                        ),
                        child: Center(
                          child: Text(
                            'ECOM',
                            style: GoogleFonts.caveatBrush(
                              textStyle: const TextStyle(
                                color: Color(0xFF3F51F3),
                                fontSize: 20,
                                fontWeight: FontWeight.w900,
                              ),
                            ),
                          ),
                        ),
                      ),
                    ],
                  ),
                ),
                const SizedBox(height: 20),
                Center(
                  child: Text(
                    'Create your account',
                    style: GoogleFonts.poppins(
                      textStyle: const TextStyle(
                        color: Colors.black,
                        fontSize: 27,
                        fontWeight: FontWeight.w500,
                      ),
                    ),
                  ),
                ),
                const SizedBox(height: 20),
                Padding(
                  padding: const EdgeInsets.only(left: 20.0),
                  child: Text(
                    'Name',
                    style: GoogleFonts.poppins(
                      textStyle: TextStyle(
                        color: Colors.grey,
                        fontSize: 17,
                      ),
                    ),
                  ),
                ),
                const SizedBox(height: 7),
                Padding(
                  child: CustomTextField(
                    hinttext: 'ex: jon smith',
                    controller: nameController,
                  ),
                  padding: const EdgeInsets.only(left: 20),
                ),
                const SizedBox(height: 7),
                Padding(
                  padding: const EdgeInsets.only(left: 20.0),
                  child: Text(
                    'Email',
                    style: GoogleFonts.poppins(
                      textStyle: TextStyle(
                        color: Colors.grey,
                        fontSize: 17,
                      ),
                    ),
                  ),
                ),
                const SizedBox(height: 7),
                Padding(
                  child: CustomTextField(
                      hinttext: 'ex: jon.smith@gmail.com',
                      controller: emailController),
                  padding: const EdgeInsets.only(left: 20),
                ),
                const SizedBox(height: 7),
                Padding(
                  padding: const EdgeInsets.only(left: 20.0),
                  child: Text(
                    'Password',
                    style: GoogleFonts.poppins(
                      textStyle: TextStyle(
                        color: Colors.grey,
                        fontSize: 17,
                      ),
                    ),
                  ),
                ),
                const SizedBox(height: 7),
                Padding(
                  child: CustomTextField(
                    hinttext: '*********',
                    controller: passwordController,
                  ),
                  padding: const EdgeInsets.only(left: 20),
                ),
                const SizedBox(height: 7),
                Padding(
                  padding: const EdgeInsets.only(left: 20.0),
                  child: Text(
                    'Confirm Password',
                    style: GoogleFonts.poppins(
                      textStyle: TextStyle(
                        color: Colors.grey,
                        fontSize: 17,
                      ),
                    ),
                  ),
                ),
                const SizedBox(height: 7),
                Padding(
                  child: CustomTextField(
                    hinttext: '********',
                    controller: confirmPasswordController,
                  ),
                  padding: const EdgeInsets.only(left: 20),
                ),
                const SizedBox(height: 7),
                Row(
                  mainAxisSize:
                      MainAxisSize.min, // Minimize row size to fit content
                  children: [
                    Checkbox(
                      value: _isChecked,
                      activeColor: Color(0xFF3F51F3), // Color when checked
                      checkColor: Colors.white, // Color of the checkmark
                      onChanged: (bool? value) {
                        setState(() {
                          _isChecked = value ?? false;
                        });
                      },
                    ),
                    SizedBox(width: 7), // Space between checkbox and text
                    Text(
                      'I understand the ',
                      style: GoogleFonts.poppins(
                          textStyle: TextStyle(
                        color: Colors.black,
                      )),
                    ),
                    Text(
                      'terms & policy.',
                      style: GoogleFonts.poppins(
                          textStyle: TextStyle(
                        color: Color(0xFF3F51F3),
                      )),
                    )
                  ],
                ),
                const SizedBox(
                  height: 10,
                ),
                GestureDetector(
                  onTap: () {
                    if (nameController.text.isEmpty  ||
                        emailController.text.isEmpty ||
                        passwordController.text.isEmpty ||
                        confirmPasswordController.text.isEmpty) {
                      ScaffoldMessenger.of(context).showSnackBar(
                        const SnackBar(
                          content: Text('Please fill all fields'),
                        ),
                      );
                    } else if (passwordController.text !=
                        confirmPasswordController.text) {
                      ScaffoldMessenger.of(context).showSnackBar(
                        const SnackBar(
                          content: Text('Password does not match'),
                        ),
                      );
                    } else if (!_isChecked) {
                      ScaffoldMessenger.of(context).showSnackBar(
                        const SnackBar(
                          content: Text('Please accept terms & policy'),
                        ),
                      );
                    } else {
                      user_bloc.add(RegisterEvent(
                          name: nameController.text,
                          email: emailController.text,
                          password: passwordController.text));
                    }
                  },
                  child: Container(
                    width: double.infinity,
                    height: 50,
                    decoration: BoxDecoration(
                      borderRadius: BorderRadius.circular(10),
                      color: Color(0xFF3F51F3),
                    ),
                    child: Center(
                      child: Text(
                        'SIGN UP',
                        style: GoogleFonts.poppins(
                            textStyle: TextStyle(
                          color: Colors.white,
                          fontSize: 17,
                          fontWeight: FontWeight.w500,
                        )),
                      ),
                    ),
                  ),
                ),
                SizedBox(height: 80),
                Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Text(
                      'Have an account?',
                      style: GoogleFonts.poppins(
                          textStyle: TextStyle(
                        color: Colors.grey,
                        fontSize: 16,
                      )),
                    ),
                    GestureDetector(
                      onTap: () {
                        Navigator.push(
                            context,
                            MaterialPageRoute(
                                builder: (context) => const login()));
                      },
                      child: Text(
                        'SIGN IN',
                        style: GoogleFonts.poppins(
                            textStyle: TextStyle(
                          color: Color(0xFF3F51F3),
                          fontSize: 16,
                        )),
                      ),
                    ),
                  ],
                )
              ],
            ),
          ),
        ),
      ),
    );
  }
}

class CustomTextField extends StatelessWidget {
  final TextEditingController controller;
  final hinttext;
  const CustomTextField({
    Key? key,
    this.hinttext,
    required this.controller,
  }) : super(key: key);
  @override
  Widget build(BuildContext context) {
    return Container(
      width: double.infinity,
      height: 50,
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(10),
        color: Color(0xFFFAFAFA),
      ),
      child: TextField(
        controller: controller,
        decoration: InputDecoration(
          border: InputBorder.none,
          hintText: hinttext,
          contentPadding: const EdgeInsets.symmetric(horizontal: 10),
          hintStyle: GoogleFonts.poppins(
            textStyle: const TextStyle(
              color: Color(0xFFC1C1C1),
              fontSize: 15,
            ),
          ),
        ),
      ),
    );
  }
}
