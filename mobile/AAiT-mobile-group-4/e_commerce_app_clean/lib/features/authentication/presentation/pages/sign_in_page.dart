import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../product/presentation/widgets/components/styles/custom_button.dart';
import '../../../product/presentation/widgets/components/styles/snack_bar_style.dart';
import '../../../product/presentation/widgets/components/styles/text_field_styles.dart';
import '../../../product/presentation/widgets/components/styles/text_style.dart';
import '../../domain/entities/log_in.dart';
import '../bloc/auth_bloc.dart';

class SignInPage extends StatelessWidget {
  SignInPage({super.key});

  final TextEditingController emailController = TextEditingController();
  final TextEditingController passwordController = TextEditingController();
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: BlocConsumer<AuthBloc, AuthState>(
        listener: (context, state) {
          if (state is AuthSignedInState) {
            ScaffoldMessenger.of(context).showSnackBar( customSnackBar('Logged In successfully',Theme.of(context).primaryColor));
            Navigator.pushNamed(context, '/home_page');
          } else if (state is AuthErrorState) {
            ScaffoldMessenger.of(context)
                .showSnackBar(customSnackBar(state.message,Theme.of(context).primaryColor));
          } else if (state is AuthLogOutState) {
            ScaffoldMessenger.of(context).showSnackBar(customSnackBar('Logged Out Successfully',Theme.of(context).primaryColor));
          }
        },
        builder: (context, state) {
          if (state is AuthLoadingState) {
            return const Center(
              child: CircularProgressIndicator(),
            );
          } else {
            return SingleChildScrollView(
              child: Column(children: [
                const SizedBox(height: 116),
                Container(
                  padding: const EdgeInsets.symmetric(horizontal: 24),
                  decoration: BoxDecoration(
                      border: Border.all(
                          color:Theme.of(context).primaryColor,
                          width: 2.0,
                          style: BorderStyle.solid),
                      borderRadius: BorderRadius.circular(6)),
                  child: CustomTextStyle(
                    color: Theme.of(context).primaryColor,
                    name: 'ECOM',
                    weight: FontWeight.w400,
                    size: 48,
                    family: 'CaveatBrush',
                  ),
                ),
                const SizedBox(height: 46),
                Padding(
                  padding: const EdgeInsets.fromLTRB(50, 0, 0, 0),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      const CustomTextStyle(
                        color: Colors.black,
                        name: 'Sign into your account',
                        weight: FontWeight.w600,
                        size: 27,
                        family: 'Poppins',
                      ),
                      const SizedBox(height: 36),
                      const CustomTextStyle(
                        color: Color.fromRGBO(111, 111, 111, 1),
                        name: 'Email',
                        weight: FontWeight.w400,
                        size: 16,
                        family: 'Poppins',
                      ),
                      const SizedBox(height: 8),
                      SizedBox(
                        width: MediaQuery.of(context).size.width / 1.25,
                        child: CustomTextField(
                          lines: 1,
                          controller: emailController,
                          hint: 'ex.jon.smith@email.com',
                        ),
                      ),
                      const SizedBox(height: 8),
                      const CustomTextStyle(
                        color: Color.fromRGBO(111, 111, 111, 1),
                        name: 'Password',
                        weight: FontWeight.w400,
                        size: 16,
                        family: 'Poppins',
                      ),
                      const SizedBox(height: 8),
                      SizedBox(
                        width: MediaQuery.of(context).size.width / 1.25,
                        child: CustomTextField(
                          obsecure: true,
                          lines: 1,
                          controller: passwordController,
                          hint: '................',
                        ),
                      ),
                      const SizedBox(height: 20),
                      CustomButton(
                        width: MediaQuery.of(context).size.width / 1.25,
                        height: 42,
                        name: 'SIGN IN',
                        fgcolor: Colors.white,
                        textBgColor: Colors.white,
                        bgcolor: Theme.of(context).primaryColor,
                        pressed: () {
                          context.read<AuthBloc>().add(
                                LogInEvent(
                                  logInEntity: LogInEntity(
                                    email: emailController.text,
                                    password: passwordController.text,
                                  ),
                                ),
                              );
                        },
                      ),
                      const SizedBox(height: 100),
                      Row(
                        children: [
                          const SizedBox(width: 24),
                          const CustomTextStyle(
                            color: Color.fromRGBO(111, 111, 111, 1),
                            name: 'Don\'t have an account?',
                            weight: FontWeight.w400,
                            size: 16,
                            family: 'Poppins',
                          ),
                          const SizedBox(width: 4),
                          GestureDetector(
                            onTap: () {
                              Navigator.pushNamed(context, '/sign_up_page');
                            },
                            child: CustomTextStyle(
                              color: Theme.of(context).primaryColor,
                              name: 'SIGN UP',
                              weight: FontWeight.w400,
                              size: 16,
                              family: 'Poppins',
                            ),
                          )
                        ],
                      )
                    ],
                  ),
                ),
              ]),
            );
          }
        },
      ),
    );
  }
}
