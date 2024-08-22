import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_model.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/domain/entities/user_entitiy.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/signup/bloc/sign_up_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/signup/bloc/sign_up_event.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/signup/bloc/sign_up_state.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';

import '../bloc/login/bloc/sign_in_bloc.dart';

class SignUpScreen extends StatefulWidget {
  const SignUpScreen({super.key});

  @override
  State<SignUpScreen> createState() => _SignUpScreenState();
}

class _SignUpScreenState extends State<SignUpScreen> {
  TextEditingController name_input = TextEditingController();

  TextEditingController email_input = TextEditingController();

  TextEditingController password_input = TextEditingController();

  TextEditingController confirm_password_input = TextEditingController();

  bool passwordVisible = true;

  bool confirmpasswordVisible = true;
  bool isChecked = false;
  @override 
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        backgroundColor: Colors.white,
        automaticallyImplyLeading: false,
        title: Row(
          crossAxisAlignment: CrossAxisAlignment.center,
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            IconButton(
                onPressed: () {
                  Navigator.pop(context);
                },
                icon: Icon(
                  Icons.arrow_back_ios_new,
                  color: Color.fromRGBO(63, 81, 243, 1),
                  size: 27,
                )),
            Card(
              elevation: 5,
              child: Container(
                width: 60,
                height: 30,
                decoration: BoxDecoration(
                  // boxShadow: [BoxShadow(color: Colors.black,blurRadius: 2)],
                  border: Border.all(color: Color.fromRGBO(63, 81, 243, 1)),
                  borderRadius: BorderRadius.circular(10),
                  color: Colors.white,
                ),
                alignment: Alignment.center,
                child: Text(
                  'ECOM',
                  style: GoogleFonts.caveatBrush(
                      fontSize: 22,
                      fontWeight: FontWeight.w600,
                      color: Color.fromRGBO(63, 81, 243, 1)),
                ),
              ),
            )
          ],
        ),
      ),
      body: SafeArea(
        child: BlocConsumer<SignUpBloc, SignUpState>(
          listener: (context, state) {
            if (state is SignUpLoaded && state.message != 'Authentication Failed') {
              ScaffoldMessenger.of(context).showSnackBar(
                SnackBar(
                  backgroundColor: Colors.green,
                  content: Text(state.message),
                  duration: Duration(seconds: 5),
                ),
              );
              Navigator.pushNamed(context, '/login');
            }
            if (state is SignUpLoading) {
              ScaffoldMessenger.of(context).showSnackBar(
                SnackBar(
                  backgroundColor: Colors.white,
                  content: Container(
                    padding: EdgeInsets.only(left: 50),
                    child: Row(
                      children: const [
                        Text('Loading...',style: TextStyle(color: Colors.black,fontSize: 30),),
                        SizedBox(width: 10,),
                    
                        SizedBox(width:30,height:30 ,child: CircularProgressIndicator()),
                      ],
                    ),
                  ),
                  duration: Duration(seconds: 2),
                ),
              );
            }
            else if (state is SignUpLoaded && state.message == 'Authentication Failed') {
              ScaffoldMessenger.of(context).showSnackBar(
                SnackBar(
                  backgroundColor: Colors.red,
                  content: Text(state.message),
                  duration: Duration(seconds: 5),
                ),
              );
            }
          },
          builder: (context, state) {
            return SingleChildScrollView(
              child: Container(
                padding: EdgeInsets.only(top: 10, left: 50, right: 50),
                child: Column(
                  children: [
                    Text(
                      'Create your account',
                      style: GoogleFonts.poppins(
                          fontSize: 26, fontWeight: FontWeight.w600),
                    ),
                    SizedBox(
                      height: 30,
                    ),
                    Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        Text(
                          'Name',
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
                            controller: name_input,
                            decoration: InputDecoration(
                              hintText: 'ex: jon smith',
                              border: OutlineInputBorder(
                                borderRadius: BorderRadius.circular(10),
                              ),
                            ),
                          ),
                        ),
                      ],
                    ),
                    SizedBox(
                      height: 15,
                    ),
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
                    SizedBox(
                      height: 17,
                    ),
                    Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        Text(
                          'Password',
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
                            controller: password_input,
                            // helperText:"Password must contain special character",
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
                      height: 17,
                    ),
                    Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        Text(
                          'Confirm password',
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
                            controller: confirm_password_input,
                            obscureText: confirmpasswordVisible,
                            decoration: InputDecoration(
                              hintText: '*********',
                              border: OutlineInputBorder(
                                borderRadius: BorderRadius.circular(10),
                              ),
                              suffixIcon: IconButton(
                                icon: Icon(confirmpasswordVisible
                                    ? Icons.visibility
                                    : Icons.visibility_off),
                                onPressed: () {
                                  setState(
                                    () {
                                      confirmpasswordVisible =
                                          !confirmpasswordVisible;
                                    },
                                  );
                                },
                              ),
                            ),
                          ),
                        ),
                      ],
                    ),
                    Row(
                      children: [
                        SizedBox(
                          child: Checkbox(
                              activeColor: Colors.blue,
                              value: isChecked,
                              onChanged: (value) {
                                setState(() {
                                  isChecked = value!;
                                });
                              }),
                        ),
                        SizedBox(
                          width: 2,
                        ),
                        Row(
                          mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                          children: [
                            Text(
                              "I understood the ",
                              style: GoogleFonts.poppins(
                                fontSize: 12,
                                fontWeight: FontWeight.w400,
                              ),
                            ),
                            TextButton(
                                onPressed: () {},
                                child: Text(
                                  'terms & policy.',
                                  style: GoogleFonts.poppins(
                                      color: Color.fromRGBO(63, 81, 243, 1),
                                      fontSize: 12),
                                ))
                          ],
                        )
                      ],
                    ),
                    SizedBox(
                      height: 10,
                    ),
                    SizedBox(
                      width: 300,
                      height: 42,
                      child: ElevatedButton(
                          style: ElevatedButton.styleFrom(
                            backgroundColor: Color.fromRGBO(63, 81, 243, 1),
                            shape: RoundedRectangleBorder(
                              borderRadius: BorderRadius.circular(10),
                            ),
                          ),
                          onPressed: () {
                            var user = UserModel(
                                name: name_input.text,
                                email: email_input.text,
                                password: password_input.text);
                                if (name_input.text.isEmpty || email_input.text.isEmpty || password_input.text.isEmpty || confirm_password_input.text.isEmpty ) {
                                      ScaffoldMessenger.of(context).showSnackBar(
                                        SnackBar(
                                          content: Text('All fields are required'),
                                          duration: Duration(seconds: 10),
                                        ),
                                );}else if (!email_input.text.contains('@')|| !email_input.text.contains('.com')){
                                      ScaffoldMessenger.of(context).showSnackBar(
                                            SnackBar(
                                              content: Text('Please Enter Valid Email Address'),
                                              duration: Duration(seconds: 5),
                                            ),
                                    );
                                  
                                  
                                }
                                else if (password_input.text.length < 8) {
                                ScaffoldMessenger.of(context).showSnackBar(
                                  SnackBar(
                                    content: Text('Password must be at least 8 characters'),
                                    duration: Duration(seconds: 10),
                                  ),
                                );
                                }else if (isChecked!=true){
                                  ScaffoldMessenger.of(context).showSnackBar(
                                  SnackBar(
                                    content: Text('Please Check The Checkbox'),
                                    duration: Duration(seconds: 4),
                                  ),
                                );

                                }
                              else if (password_input.text != confirm_password_input.text) {
                                ScaffoldMessenger.of(context).showSnackBar(
                                  SnackBar(
                                    content: Text('Password does not match'),
                                    duration: Duration(seconds: 10),
                                  ),
                                );
                                return;
                              }
                              else{
                              var signUp = BlocProvider.of<SignUpBloc>(context);
                                signUp.add(RegisterUserEvent(user));
                              }
                            
                          },
                          child: Text(
                            'SIGN UP',
                            style: GoogleFonts.poppins(
                                fontSize: 15,
                                fontWeight: FontWeight.w600,
                                color: Colors.white),
                          )),
                    ),
                    SizedBox(
                      height: 80,
                    ),
                    Container(
                      padding: EdgeInsets.only(
                        left: 35,
                      ),
                      alignment: Alignment.center,
                      child: Row(
                        children: [
                          Text(
                            'Have an account?',
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
                              'SIGN IN',
                              style: GoogleFonts.poppins(
                                  fontSize: 16,
                                  fontWeight: FontWeight.w400,
                                  color: Color.fromRGBO(63, 81, 243, 1)),
                            ),
                          ),
                        ],
                      ),
                    ),
                  ],
                ),
              ),
            );
          },
        ),
      ),
    );
  }
}
