import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../bloc/blocs.dart';
import '../bloc/states.dart';
import '../widgets/custom_button.dart';
import '../widgets/linked_text.dart';
import '../widgets/logo.dart';
import '../widgets/textField.dart';
import '../widgets/title.dart';
import 'sign_in.dart';

class SignUp extends StatelessWidget {
  const SignUp({super.key});

  @override
  Widget build(BuildContext context) {
    TextEditingController nameController = TextEditingController();
    TextEditingController emailController = TextEditingController(text: '');
    TextEditingController passwordController = TextEditingController(text: '');
    TextEditingController confirmPasswordController =
        TextEditingController(text: '');
    return BlocConsumer<UserBloc, UserState>(
      listener: (context, state) {
        // TODO: implement listener
        print(state);
            if (state is RegisterLoadingState) {
              ScaffoldMessenger.of(context).showSnackBar(
                const SnackBar(
                  content: Text("Registering User"),
                ),
                );
                
              
            } else if (state is RegisterErrorState) {

              ScaffoldMessenger.of(context).showSnackBar(
                SnackBar(
                  content: Text(state.message),
                  backgroundColor: Colors.red,
                  
                ),
              );
            } else if (state is UserRegisteredState) {
              ScaffoldMessenger.of(context).showSnackBar(
                const SnackBar(
                  content: Text("User Registered"),
                  backgroundColor: Colors.green,
                  
                ),
                
              );
              Navigator.push(
                context,
                MaterialPageRoute(builder: (context) => SignIn()),
                );
            }
      },
      builder: (context, state) {
        
            return Scaffold(
              appBar: AppBar(
                actions: const [Logo(width: 60, height: 25, fontSize: 24)],
              ),
              body: Column(
                // mainAxisAlignment: MainAxisAlignment.center,
                crossAxisAlignment: CrossAxisAlignment.center,
                children: [
                  const BigTitle(text: "Create an account"),
                  Expanded(
                    child: SingleChildScrollView(
                      child: Column(
                        children: [
                          CustomTextField(
                              name: 'Name',
                              controller: nameController,
                              hintText: 'ex: jon smith'),
                          CustomTextField(
                              name: 'Email',
                              controller: emailController,
                              hintText: 'ex: jon.smith@email.com'),
                          CustomTextField(
                              name: 'Password',
                              controller: passwordController,
                              hintText: '*********'),
                          CustomTextField(
                              name: 'Confirm Password',
                              controller: confirmPasswordController,
                              hintText: '*********'),
                          CustomButton(
                              name: "SIGN UP",
                              controllers: [
                                emailController,
                                passwordController,
                                nameController,
                                confirmPasswordController
                              ],
                              login: false),
                          const Row(
                              mainAxisAlignment: MainAxisAlignment.start,
                              children: [
                                // Checkbox(value: false, onChanged: (bool? value) {
                                //   value = !value;
                                // }),
                                LinkedText(
                                    description: "I understood the ",
                                    linkedText: "terms & policy.",
                                    navigateTo: SignUp()),
                              ]),
                          const LinkedText(
                              description: "Have an account? ",
                              linkedText: "SIGN IN",
                              navigateTo: SignIn())
                        ],
                      ),
                    ),
                  )
                ],
              ),
            );
          },
        );
     
  }
}
