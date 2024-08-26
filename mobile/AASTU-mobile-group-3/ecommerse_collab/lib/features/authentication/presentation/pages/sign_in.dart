import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../product/presentation/pages/home_page.dart';
import '../bloc/blocs.dart';
import '../bloc/events.dart';
import '../bloc/states.dart';
import '../widgets/custom_button.dart';
import '../widgets/linked_text.dart';
import '../widgets/logo.dart';
import '../widgets/textField.dart';
import '../widgets/title.dart';
import 'sign_up.dart';

class SignIn extends StatefulWidget {
  const SignIn({Key? key}) : super(key: key);

  @override
  State<SignIn> createState() => _SignInState();
}

class _SignInState extends State<SignIn> {
  final TextEditingController emailController = TextEditingController();
  final TextEditingController passwordController = TextEditingController();
  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();

  String? validateEmail(String? value) {
    if (value == null || value.isEmpty) {
      return 'Email cannot be empty';
    } else if (!RegExp(r'^[^@]+@[^@]+\.[^@]+').hasMatch(value)) {
      return 'Enter a valid email';
    }
    return null;
  }

  String? validatePassword(String? value) {
    if (value == null || value.isEmpty) {
      return 'Password cannot be empty';
    } else if (value.length < 6) {
      return 'Password must be at least 6 characters';
    }
    return null;
  }

  void _submitForm(BuildContext context) {
    if (_formKey.currentState?.validate() ?? false) {
      context.read<UserBloc>().add(LogInEvent(
            email: emailController.text,
            password: passwordController.text,
          ));
    } else {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(
          content: Text("Please fix the errors in the form"),
          backgroundColor: Colors.red,
        ),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return BlocConsumer<UserBloc, UserState>(
      listener: (context, state) {
        if (state is LoginLoadingState) {
          ScaffoldMessenger.of(context).showSnackBar(
            const SnackBar(
              content: Text("Logging in"),
            ),
          );
        } else if (state is LoginErrorState) {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(
              content: Text(state.message),
              backgroundColor: const Color.fromARGB(255, 248, 92, 81),
            ),
          );
        } else if (state is UserLoggedState) {
          ScaffoldMessenger.of(context).showSnackBar(
            const SnackBar(
              content: Text(
                "User Logged in",
                style: TextStyle(color: Colors.white),
              ),
              backgroundColor: Colors.green,
            ),
          );
          Navigator.push(
            context,
            MaterialPageRoute(builder: (context) => HomePage(user: state.user)),
          );
        }
      },
      builder: (context, state) {
        return Scaffold(
          backgroundColor: Colors.white,
          body: SingleChildScrollView(
            child: Padding(
              padding: const EdgeInsets.all(16.0),
              child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  const SizedBox(height: 100),
                  const Logo(width: 120, height: 50, fontSize: 30),
                  const BigTitle(text: "Sign into your account"),
                  Form(
                    key: _formKey,
                    child: Column(
                      children: [
                        CustomTextField(
                          name: "Email",
                          controller: emailController,
                          hintText: "ex: jon.smith@email.com",
                          isPassword: false,
                          validator: validateEmail,
                        ),
                        CustomTextField(
                          name: "Password",
                          controller: passwordController,
                          hintText: "********",
                          isPassword: true,
                          validator: validatePassword,
                        ),
                        CustomButton(
                          name: 'SIGN IN',
                          login: true,
                          formKey: _formKey,
                          controllers: [emailController, passwordController],
                          onPressed: () => _submitForm(context),
                        ),
                        const SizedBox(height: 40),
                        LinkedText(
                          description: "Don’t have an account?",
                          linkedText: " SIGN UP",
                          navigateTo: const SignUp(),
                        ),
                      ],
                    ),
                  ),
                ],
              ),
            ),
          ),
        );
      },
    );
  }
}
