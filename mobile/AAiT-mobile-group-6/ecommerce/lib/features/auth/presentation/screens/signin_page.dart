import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:path/path.dart';

import '../../../../core/common_widget/circular_indicator.dart';
import '../../../../core/common_widget/snack_bar.dart';
import '../../../product/domain/entitity/user.dart';
import '../../../product/presentation/widgets/text_field.dart';
import '../../data/models/auth_model.dart';
import '../bloc/auth_bloc.dart';


class SigninPage extends StatelessWidget {
  SigninPage({super.key});
  final TextEditingController emailController = TextEditingController();
  final TextEditingController passwordController = TextEditingController();


  void _logIn(BuildContext context) async {
    final newUser = LogInModel(
        id: '', email: emailController.text, password: passwordController.text);

    context.read<AuthBloc>().add(LogInEvent(logInEntity: newUser));
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: BlocConsumer<AuthBloc, AuthState>(
        listener: (context, state) {
          if (state is AuthSuccessState) {
            final snac = snackBar('User logged in successfully');
            ScaffoldMessenger.of(context).showSnackBar(snac);
            Future.delayed(const Duration(seconds: 2), () {
              Navigator.pushNamed(
                context,
                '/homepage',
              );
              // arguments: {'user': User(id: state.token)});
            });
          } else if (state is AuthErrorState) {
            final snac = errorsnackBar('Log in failed, try again');
            ScaffoldMessenger.of(context).showSnackBar(snac);
          }
        },
        builder: (context, state) {
          if (state is AuthLoadingState) {
            return const CircularIndicator();
          } else {
            return _build(context);
          }
        },
      ),
    );
  }

  Widget _build(BuildContext context) {
    return Scaffold(

      body: Center(
          child: Column(
        children: [
          Container(
            margin: const EdgeInsets.only(top: 120),
            padding: const EdgeInsets.symmetric(
              horizontal: 20,
            ),
            decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(10),
                color: Colors.white,
                border: Border.all(
                  color: const Color.fromARGB(255, 12, 62, 243),
                ),
                boxShadow: const [
                  BoxShadow(
                      color: Colors.grey, blurRadius: 10, offset: Offset(1, 5)),
                ]),
            child: const Text(
              'ECOM',
              style: TextStyle(
                fontSize: 45,
                fontWeight: FontWeight.bold,
                color: Color.fromARGB(255, 33, 68, 243),
                fontFamily: 'Caveat Brush',

              ),
            ),
          ),
          const SizedBox(
            height: 60,
          ),
          const Text(
            'Sign into your account',
            style: TextStyle(
              fontSize: 26,
              fontWeight: FontWeight.w600,
            ),
          ),
          Container(
              margin: const EdgeInsets.symmetric(horizontal: 42, vertical: 40),
              child: Column(
                children: [
                  MyTextField(
                    lable: 'Email',
                    lines: 1,
                    controller: emailController,
                    hint: 'ex: jon.smith@email.com',
                  ),
                  MyTextField(
                      lable: 'Password',
                      lines: 1,
                      controller: passwordController,
                      hint: '********',
                      obscureText: true),
                ],
              )),
          ElevatedButton(
            onPressed: () {
              _logIn(context);
            },

            style: ButtonStyle(
              backgroundColor: MaterialStateProperty.all(
                const Color.fromARGB(255, 38, 80, 232),
              ),
              minimumSize: MaterialStateProperty.all(const Size(300, 50)),
              shape: MaterialStateProperty.all(
                RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(10),
                ),
              ),
            ),
            child: const Text(
              'SIGN IN',
              style: TextStyle(color: Colors.white, fontSize: 16),
            ),
          ),
          Container(
            margin: const EdgeInsets.only(left: 60, right: 60, top: 130),
            child: Row(
              children: [
                const Text(
                  'Don\'t have an account? ',
                  style: TextStyle(
                      fontSize: 15, color: Color.fromARGB(255, 149, 148, 148)),
                ),
                GestureDetector(
                  onTap: () {
                    Navigator.pushNamed(context, '/signup_page');
                  },
                  child: const Text('SIGN UP',
                      style: TextStyle(
                        color: Color.fromARGB(255, 38, 80, 232),
                      )),
                )
              ],
            ),
          ),
        ],
      )),
    );
  }
}
