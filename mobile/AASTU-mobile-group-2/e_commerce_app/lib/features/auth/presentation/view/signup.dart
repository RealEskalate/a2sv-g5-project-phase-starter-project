import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/features/auth/presentation/bloc/auth_bloc.dart';
import 'package:e_commerce_app/features/auth/presentation/bloc/auth_event.dart';
import 'package:e_commerce_app/features/auth/presentation/bloc/auth_state.dart';
import 'package:e_commerce_app/features/auth/presentation/view/widgets.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';

class SignUpScreen extends StatelessWidget {
  const SignUpScreen({super.key});

  @override
  Widget build(BuildContext context) {
    TextEditingController nameController = TextEditingController();
    TextEditingController emailController = TextEditingController();
    TextEditingController passwordController = TextEditingController();
    TextEditingController confirmPassController = TextEditingController();
    void signup() {
      context.read<AuthBloc>().add(SignUpEvent(
          name: nameController.text,
          email: emailController.text,
          password: passwordController.text));
    }

    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        automaticallyImplyLeading: false,
        leading: GoBack(),
        actions: [
          Container(
            // padding: EdgeInsets.all(0),
            decoration: BoxDecoration(
              boxShadow: [
              
                         BoxShadow(
          color: Color.fromRGBO(0, 0, 0, 0.25),
          blurRadius: 200,
          offset: Offset(4, 4), 
          spreadRadius: 0,
        ),      
                 
              ],
              border: Border.all(
                color: Color.fromRGBO(63, 81, 243, 1),
                
              ),
              borderRadius: BorderRadius.all(Radius.circular(9)),
            ),
            height: 25,
            width: 60,
            // color: Colors.black,
            
                child: FittedBox(
                  fit: BoxFit.cover,

                child: Padding(
                  padding: EdgeInsets.all(5),
                  child: Text(
                    "ECOM",
                    style: GoogleFonts.caveatBrush(
                      color: Color.fromRGBO(63, 81, 243, 1),
                      fontSize: 23,
                      fontWeight: FontWeight.w400,
                    ),
                  ),
                ),
              ),
            
            
          ),
          SizedBox(
            width: 35,
          )
        ],
      ),
      body: Center(
        child: SingleChildScrollView(
          padding: EdgeInsets.only(left: 25, right: 25),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              Text(
                "Create your account",
                style: GoogleFonts.poppins(
                  fontWeight: FontWeight.w600,
                  fontSize: 26
                ),
              ),
              SizedBox(
                height: 10,
              ),

              TextFieldTitle(
                title: "Name",
                controller: nameController,
                hint: "ex: jon smith",
              ),
              TextFieldTitle(
                controller: emailController,
                title: "Email",
                hint: "ex: jon.smith@email.com",
              ),

              TextFieldTitle(
                controller: passwordController,
                pass: true,
                title: "Password",
                hint: "********",
              ),
              TextFieldTitle(
                controller: confirmPassController,
                pass: true,
                title: "Confirm Password",
                hint: "********",
              ),

              SizedBox(
                height: 25,
              ),
              Row(
                mainAxisAlignment: MainAxisAlignment.start,
                children: [
                  Checkbox(
                    
                    value: false, onChanged: null),
                  Text("I understand the  ",
                  style: GoogleFonts.poppins(
                    fontWeight: FontWeight.w400,
                    
                  ),
                  ),
                  TextButton(
                    onPressed: () {
                      Navigator.pushReplacementNamed(context, '/login');
                    },
                    child: Text("terms & policy", style: GoogleFonts.poppins(
                    fontWeight: FontWeight.w400,
                    color: Color.fromRGBO(63, 81, 243, 1)

                    
                  ),),
                  ),
                ],
              ),
              BlocConsumer<AuthBloc, AuthState>(
                listener: (context, state) {
                  if (state is SignUpSuccess) {
                    Navigator.pushReplacementNamed(context, '/login');
                  }
                },
                builder: (context, state) {
                  if (state is AuthInitial) {
                    return BackgroundButton(title: "SIGN UP", callback: signup);
                  } else if (state is SignUpLoading) {
                    return Center(child: CircularProgressIndicator());
                  } else {
                    return Column(
                      children: [
                        Text(
                          "failed to signup,try again",
                          style: TextStyle(color: Colors.red),
                        ),
                        BackgroundButton(title: "SIGN UP", callback: signup)
                      ],
                    );
                  }
                },
              ),
              // ElevatedButton(
              //   onPressed: () {

              //     Navigator.pushNamed(context, '/login');
              //   },
              //   child: Text("Login"),
              // ),
              // BackgroundButton(
              //   title: "LOGIN",
              //   callback: () {
              //     Navigator.pushNamed(context, '/login');
              //   },
              // )

              // ,
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Text("You have an account? ",  style: GoogleFonts.poppins(
                    fontWeight: FontWeight.w400,
                    fontSize: 16,
                    color: Color.fromRGBO(136, 136, 136, 1),),),
                  TextButton(
                    onPressed: () {
                      Navigator.pushReplacementNamed(context, '/login');
                    },
                    child: Text("LOGIN",
                    
                    style: GoogleFonts.poppins(
                    fontWeight: FontWeight.w400,
                    fontSize: 16,
                    color: Color.fromRGBO(63, 81, 243, 1))
                    ),
                  ),
                ],
              )
            ],
          ),
        ),
      ),
    );
  }
}
