import 'package:ecommerce/features/auth/data/model/user_model.dart';
import 'package:ecommerce/features/auth/presentation/bloc/authbloc/auth_bloc.dart';
import 'package:ecommerce/features/auth/presentation/bloc/authbloc/auth_event.dart';
import 'package:ecommerce/features/auth/presentation/bloc/authbloc/auth_state.dart';
import 'package:ecommerce/features/auth/presentation/register.dart';
import 'package:ecommerce/service_locator.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';

import '../../product/presentation/pages/homepage.dart';
import '../data/data_sources/remote_data_source.dart';

class login extends StatefulWidget {
  const login({super.key});

  @override
  State<login> createState() => _loginState();
}

class _loginState extends State<login> {
  TextEditingController emailController = TextEditingController();
  TextEditingController passwordController = TextEditingController();
  String? token;
  var user_bloc = getIt<UserBloc>();
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      body: SingleChildScrollView(
        child: BlocListener<UserBloc, UserState>(
          listener: (context, state) {
            if (state is logging) {
              ScaffoldMessenger.of(context).showSnackBar(
                SnackBar(
                  content: Text('Logging in...'),
                ),
              );
            } else if (state is logged) {
              UserModel user = state.user;
              ScaffoldMessenger.of(context).showSnackBar(
                SnackBar(
                  content: Text('Logged in as ${user.name}'),
                ),
              );
              print("curr user ${user.id} ${user.name} ${user.email}");
              
              print(user.name);
              print(user.email);
              print(user);
              Navigator.push(
                  context,
                  MaterialPageRoute(
                      builder: (context) => MyHomePage(title: '', user: user)));
            } else if (state is logginfailure) {
              ScaffoldMessenger.of(context).showSnackBar(
                SnackBar(
                  content: Text(
                      'Loggin failed. Please ensure you used the correct email and password.'),
                ),
              );
            }
          },
          child: Padding(
              padding:
                  const EdgeInsets.symmetric(horizontal: 40, vertical: 125),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Center(
                    child: Container(
                      width: 120,
                      height: 50,
                      decoration: BoxDecoration(
                        color: Colors.white,
                        borderRadius: BorderRadius.circular(10),
                        border: Border.all(color: Color(0xFF3F51F3), width: 1),
                        boxShadow: const [
                          BoxShadow(
                            color: Colors.grey,
                            offset: Offset(0.0, 1.0), //(x,y)
                            blurRadius: 6.0,
                          ),
                        ],
                      ),
                      child: Center(
                        child: Text(
                          'ECOM',
                          style: GoogleFonts.caveatBrush(
                            textStyle: const TextStyle(
                              color: Color(0xFF3F51F3),
                              fontSize: 30,
                              fontWeight: FontWeight.w900,
                            ),
                          ),
                        ),
                      ),
                    ),
                  ),
                  SizedBox(height: 60),
                  Center(
                    child: Text(
                      'Sign into your account',
                      style: GoogleFonts.poppins(
                        textStyle: const TextStyle(
                          color: Colors.black,
                          fontSize: 23,
                          fontWeight: FontWeight.w600,
                        ),
                      ),
                    ),
                  ),
                  const SizedBox(height: 20),
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
                      controller: emailController,
                    ),
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
                      hinttext: '*******',
                      controller: passwordController,
                    ),
                    padding: const EdgeInsets.only(left: 20),
                  ),
                  const SizedBox(height: 30),
                  GestureDetector(
                      child: Container(
                        width: double.infinity,
                        height: 50,
                        decoration: BoxDecoration(
                          borderRadius: BorderRadius.circular(10),
                          color: Color(0xFF3F51F3),
                        ),
                        child: Center(
                          child: Text(
                            'SIGN IN',
                            style: GoogleFonts.poppins(
                                textStyle: TextStyle(
                              color: Colors.white,
                              fontSize: 17,
                              fontWeight: FontWeight.w500,
                            )),
                          ),
                        ),
                      ),
                      onTap: () {
                        if (emailController.text.isEmpty ||
                            passwordController.text.isEmpty) {
                          ScaffoldMessenger.of(context).showSnackBar(
                            const SnackBar(
                              content: Text('Please fill all fields'),
                            ),
                          );
                        } else {
                          user_bloc.add(LoginEvent(
                              email: emailController.text,
                              password: passwordController.text));
                        }
                      }),
                  const SizedBox(height: 180),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Text(
                        'Don\'t have an account?',
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
                                  builder: (context) => Register()));
                        },
                        child: Text(
                          'SIGN UP',
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
              )),
        ),
      ),
    );
  }
}

class CustomTextField extends StatelessWidget {
  TextEditingController controller = TextEditingController();

  final hinttext;
  CustomTextField({
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
