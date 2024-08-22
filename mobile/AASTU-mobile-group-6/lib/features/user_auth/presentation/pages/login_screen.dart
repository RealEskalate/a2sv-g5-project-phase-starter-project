import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/home_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/home_event.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_model.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/domain/entities/user_entitiy.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/login/bloc/sign_in_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/login/bloc/sign_in_event.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/login/bloc/sign_in_state.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';

class LoginScreen extends StatefulWidget {
  const LoginScreen({super.key});

  @override
  State<LoginScreen> createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
  TextEditingController email_input = TextEditingController();

  TextEditingController password_input = TextEditingController();

  bool passwordVisible = false;
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      body: SafeArea(
        child: BlocConsumer<LoginBloc, LoginState>(
          listener: (context, state) {
            if (state is LoginLoading){
              ScaffoldMessenger.of(context).showSnackBar(
                SnackBar(
                  backgroundColor: Colors.white,
                  content: Container(
                    padding: EdgeInsets.only(left: 100),
                    child: Row(
                      children: const [
                        Text('Loading...',style: TextStyle(color: Colors.black,fontSize: 20),),
                        SizedBox(width: 10,),
                    
                        SizedBox(width:20,height:20 ,child: CircularProgressIndicator()),
                      ],
                    ),
                  ),
                  duration: Duration(seconds: 3),
                ),
              );
            }
            if (state is LoginLoaded) {
              ScaffoldMessenger.of(context).showSnackBar(
                SnackBar(
                  // width: 250,
                  // margin: EdgeInsets.only(left: 50, right: 50),
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(10),
                  ),
                  backgroundColor: Colors.green[400],
                  content: Center(child: Text('Welcome Back')),
                  duration: Duration(seconds: 5),
                ),
              );
              context.read<HomeBloc>().add(GetProductsEvent());
              Navigator.pushNamed(context, '/home',arguments: state.message);
            }
            else if (state is LoginFailure) {
              ScaffoldMessenger.of(context).showSnackBar(
                SnackBar(
                  content: Text("Error: Wrong Email or Password Entered!!!"),
                  backgroundColor: Colors.red,
                  duration: Duration(seconds: 5),
                ),
              );
            }
          },
          builder: (context, state) {
            
            return SingleChildScrollView(
              child: Column(
                // mainAxisAlignment: MainAxisAlignment.start,
                crossAxisAlignment: CrossAxisAlignment.center,
                children: [
                  Padding(
                    padding:
                        const EdgeInsets.only(top: 124, left: 124, right: 122),
                    child: Container(
                      width: 170,
                      height: 70,
                      decoration: BoxDecoration(
                        border:
                            Border.all(color: Color.fromRGBO(63, 81, 243, 1)),
                        borderRadius: BorderRadius.circular(10),
                        color: Colors.white,
                      ),
                      alignment: Alignment.center,
                      child: Text(
                        'ECOM',
                        style: GoogleFonts.caveatBrush(
                            fontSize: 48,
                            fontWeight: FontWeight.w600,
                            color: Color.fromRGBO(63, 81, 243, 1)),
                      ),
                    ),
                  ),
                  SizedBox(
                    height: 60,
                  ),
                  Column(
                    children: [
                      Text(
                        'Sign into your account',
                        style: GoogleFonts.poppins(
                            fontSize: 26,
                            fontWeight: FontWeight.w600,
                            color: Colors.black),
                      ),
                      SizedBox(
                        height: 31,
                      ),
                      Container(
                          child: Column(
                        // mainAxisAlignment: MainAxisAlignment.start,
                        crossAxisAlignment: CrossAxisAlignment.center,
                        children: [
                          Column(
                            crossAxisAlignment: CrossAxisAlignment.start,
                            children: [
                              Text(
                                'Email',
                                style: GoogleFonts.poppins(
                                    fontSize: 16,
                                    fontWeight: FontWeight.w400,
                                    color: Color.fromRGBO(111, 111, 111, 1)),
                              ),
                              SizedBox(
                                height: 12,
                              ),
                              SizedBox(
                                width: 300,
                                height: 42,
                                child: TextField(
                                  controller: email_input,
                                  decoration: InputDecoration(
                                    hintText: 'ex: jon.smith@email.com',
                                    border: OutlineInputBorder(
                                      borderRadius: BorderRadius.circular(10),
                                    ),
                                  ),
                                ),
                              ),
                            ],
                          ),
                          Column(
                            crossAxisAlignment: CrossAxisAlignment.start,
                            children: [
                              SizedBox(
                                height: 13,
                              ),
                              Text(
                                'Password',
                                style: GoogleFonts.poppins(
                                    fontSize: 16,
                                    fontWeight: FontWeight.w400,
                                    color: Color.fromRGBO(111, 111, 111, 1)),
                              ),
                              SizedBox(
                                height: 13,
                              ),
                              SizedBox(
                                width: 300,
                                height: 42,
                                child: TextField(
                                  controller: password_input,
                                  obscureText: passwordVisible,
                                  decoration: InputDecoration(
                                    hintText: '*********',
                                    border: OutlineInputBorder(
                                      borderRadius: BorderRadius.circular(10),
                                    ),
                                    suffixIcon: IconButton(
                                      icon: Icon(passwordVisible
                                          ? Icons.visibility
                                          : Icons.visibility_off),
                                      onPressed: () {
                                        setState(
                                          () {
                                            passwordVisible = !passwordVisible;
                                          },
                                        );
                                      },
                                    ),
                                  ),
                                ),
                              ),
                            ],
                          ),
                          SizedBox(
                            height: 37,
                          ),
                          SizedBox(
                            width: 300,
                            height: 42,
                            child: ElevatedButton(
                                style: ElevatedButton.styleFrom(
                                  backgroundColor:
                                      Color.fromRGBO(63, 81, 243, 1),
                                  shape: RoundedRectangleBorder(
                                    borderRadius: BorderRadius.circular(10),
                                  ),
                                ),
                                onPressed: () {
                                  var user = UserModel(
                                      name: '',
                                      email: email_input.text,
                                      password: password_input.text);
                                  var login =
                                      BlocProvider.of<LoginBloc>(context);
                                  login.add(LogUserIn(user));
                                  
                                },
                                child: Text(
                                  'SIGN IN',
                                  style: GoogleFonts.poppins(
                                      fontSize: 15,
                                      fontWeight: FontWeight.w600,
                                      color: Colors.white),
                                )),
                          ),
                          SizedBox(
                            height: 80,
                          ),
                          Row(
                            crossAxisAlignment: CrossAxisAlignment.center,
                            mainAxisAlignment: MainAxisAlignment.center,
                            children: [
                              Text(
                                'Don\'t have an account?',
                                style: GoogleFonts.poppins(
                                    fontSize: 16,
                                    fontWeight: FontWeight.w400,
                                    color: Color.fromRGBO(111, 111, 111, 1)),
                              ),
                              TextButton(
                                onPressed: () {
                                  Navigator.pushNamed(context, '/signup');
                                },
                                child: Text(
                                  'SIGN UP',
                                  style: GoogleFonts.poppins(
                                      fontSize: 16,
                                      fontWeight: FontWeight.w400,
                                      color: Color.fromRGBO(63, 81, 243, 1)),
                                ),
                              ),
                            ],
                          )
                        ],
                      ))
                    ],
                  )
                ],
              ),
            );
          },
        ), //End of ECOM
      ),
    );
  }
}

