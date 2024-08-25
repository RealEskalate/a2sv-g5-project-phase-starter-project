import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../product/presentation/pages/home_page.dart';
import '../bloc/blocs.dart';
import '../bloc/states.dart';
import '../widgets/custom_button.dart';
import '../widgets/linked_text.dart';
import '../widgets/logo.dart';
import '../widgets/textField.dart';
import '../widgets/title.dart';
import 'sign_up.dart';

class SignIn extends StatefulWidget {
  const SignIn({super.key});

  @override
  State<SignIn> createState() => _SignInState();
}

class _SignInState extends State<SignIn> {
  TextEditingController emailController = TextEditingController(text: '');
  TextEditingController passwordController = TextEditingController(text: '');

  @override
  Widget build(BuildContext context) {
    return BlocConsumer<UserBloc, UserState>(
      listener: (context, state) {
        print(state);
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
            ),
          );
        } else if (state is UserLoggedState) {
          ScaffoldMessenger.of(context).showSnackBar(
            const SnackBar(
              content: Text("User Logged in"),
            ),
          );
          print("state in sign in ${state.user}");
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
            
              child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  const Logo(width: 144, height: 50, fontSize: 48),
                  const BigTitle(text: "Sign into your account"),
                  SingleChildScrollView(
                    child: Column(
                      children: [
                        CustomTextField(
                            name: "Email",
                            controller: emailController,
                            hintText: "ex: jon.smith@email.com"),
                        CustomTextField(
                            name: "Password",
                            controller: passwordController,
                            hintText: "********"),
                        CustomButton(
                          name: 'SIGN IN',
                          controllers: [emailController, passwordController],
                          login: true,
                        ),
                        const SizedBox(height: 40),
                        const LinkedText(
                            description: "Don’t have an account?",
                            linkedText: " SIGN UP",
                            navigateTo: SignUp()),
                      ],
                    ),
                  ),
                ],
              ),
            ),
          
        );
      },
    );
  }
}
