import 'package:e_commerce_app/features/auth/presentation/view/widgets.dart';
import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';

import '../bloc/auth_bloc.dart';
import '../bloc/auth_event.dart';
import '../bloc/auth_state.dart';

class LoginScreen extends StatelessWidget {
  const LoginScreen({super.key});

  @override
  Widget build(BuildContext context) {
    TextEditingController emailController = TextEditingController();
    TextEditingController passwordController = TextEditingController();
    void login() {
      context.read<AuthBloc>().add(LoginEvent(
          email: emailController.text, password: passwordController.text));
    }

    return SafeArea(
      child: Scaffold(
        backgroundColor: Colors.white,
        body: Center(
          child: SingleChildScrollView(
            padding: EdgeInsets.only(left: 25, right: 25),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Container(
                  
                  alignment: Alignment.center,
                  
                  decoration: BoxDecoration(
                    border: Border.all(
                      color: Color.fromRGBO(63, 81, 243, 1),
                      width: 1,
                    ),
                    borderRadius: BorderRadius.all(Radius.circular(10)),
                  ),
                  height: 50,
                  width: 143,
                  child:  Center(
                    child: FittedBox(
                      fit: BoxFit.cover,
                      child: Padding(
                        padding: EdgeInsets.all(0),
                        child: Text(
                          "ECOM",
                          style: GoogleFonts.caveatBrush(
                              color: Color.fromRGBO(63, 81, 243, 1),
                                fontSize: 48

                          ),
                        ),
                      ),
                    ),
                  ),
                ),
                const SizedBox(
                  height: 20,
                ),
                 Text(
                  "Sign into your account",
                  style: GoogleFonts.poppins(
                    fontWeight: FontWeight.w600,
                    fontSize: 26
                  ),
                ),
                const SizedBox(
                  height: 10,
                ),
                TextFieldTitle(title: "Email", controller: emailController),
                TextFieldTitle(
                  title: "Password",
                  controller: passwordController,
                  pass: true,
                ),
                const SizedBox(
                  height: 10,
                ),
                BlocConsumer<AuthBloc, AuthState>(
                  listener: (context, state) {
                    if (state is LoginSuccess) {
                      Navigator.pushReplacementNamed(context, '/home');
                    } else if (state is AuthFailure) {
                      ScaffoldMessenger.of(context).showSnackBar(
                        SnackBar(
                          backgroundColor: Colors.red,
                          content: Text(
                            "failed to login,try again",
                            style: TextStyle(color: Colors.white),
                          ),
                        ),
                      );
                    }
                  },
                  builder: (context, state) {
                    if (state is AuthInitial) {
                      return BackgroundButton(title: "LOGIN", callback: login);
                    } else if (state is LoginLoading) {
                      return Center(child: CircularProgressIndicator());
                    } else {
                      return Column(
                        children: [
                          BackgroundButton(title: "LOGIN", callback: login)
                        ],
                      );
                    }
                  },
                ),
                SizedBox(height: 40),
                Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Text("Don't have an account? "),
                    TextButton(
                      onPressed: () {
                        Navigator.pushReplacementNamed(context, '/signup');
                      },
                      child: Text("Sign Up",style: GoogleFonts.poppins(
                        fontWeight: FontWeight.bold,
                        color: Color.fromARGB(255, 63, 81, 243)
                      ),),
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
