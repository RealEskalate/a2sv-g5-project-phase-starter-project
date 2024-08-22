import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../product/presentation/widgets/components/styles/custom_button.dart';
import '../../../product/presentation/widgets/components/styles/snack_bar_style.dart';
import '../../../product/presentation/widgets/components/styles/text_field_styles.dart';
import '../../../product/presentation/widgets/components/styles/text_style.dart';
import '../../domain/entities/sign_up.dart';
import '../bloc/auth_bloc.dart';

class SignUpPage extends StatefulWidget {
  const SignUpPage({super.key});

  @override
  State<SignUpPage> createState() => _SignUpPageState();
}

class _SignUpPageState extends State<SignUpPage> {
  final TextEditingController usernameController = TextEditingController();

  final TextEditingController passwordController = TextEditingController();

  final TextEditingController emailController = TextEditingController();

  final TextEditingController confirmPasswordController =
      TextEditingController();

  bool isChecked = false;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: BlocConsumer<AuthBloc, AuthState>(
      listener: (context, state) {
        if (state is AuthSignedUpState) {
          ScaffoldMessenger.of(context).showSnackBar(customSnackBar('Sign up successful', Theme.of(context).primaryColor));
          Navigator.pushNamed(context, '/sign_in_page');
        } else if (state is AuthErrorState) {
          ScaffoldMessenger.of(context)
              .showSnackBar(customSnackBar(state.message, Theme.of(context).secondaryHeaderColor));
        }
      },
      builder: (context, state) {
        if (state is AuthLoadingState) {
          return const Center(
            child: CircularProgressIndicator(),
          );
        } else {
          return Padding(
            padding: const EdgeInsets.fromLTRB(32, 0, 36, 0),
            child: SingleChildScrollView(
              child: Column(
                children: [
                  const SizedBox(height: 50),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    crossAxisAlignment: CrossAxisAlignment.center,
                    children: [
                      IconButton(
                        onPressed: () {
                          Navigator.pop(context);
                        },
                        icon: Icon(Icons.arrow_back_ios_rounded,
                            size: 20, color: Theme.of(context).primaryColor),
                      ),
                      CustomTextStyle(
                        color: Theme.of(context).primaryColor,
                        name: 'ECOM',
                        weight: FontWeight.w400,
                        size: 24,
                        family: 'CaveatBrush',
                      ),
                    ],
                  ),
                  const SizedBox(height: 42),
                  Container(
                    margin: const EdgeInsets.only(left: 16),
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        const CustomTextStyle(
                            name: 'Create your account',
                            weight: FontWeight.w600,
                            size: 27),
                        const SizedBox(height: 24),
                        const CustomTextStyle(
                          name: 'Name',
                          weight: FontWeight.w400,
                          size: 16,
                          color: Color.fromRGBO(111, 111, 111, 1),
                        ),
                        const SizedBox(height: 4),
                        CustomTextField(
                            hint: 'ex: Jon Smith',
                            controller: usernameController),
                        const SizedBox(height: 12),
                        const CustomTextStyle(
                          name: 'Email',
                          weight: FontWeight.w400,
                          size: 16,
                          color: Color.fromRGBO(111, 111, 111, 1),
                        ),
                        const SizedBox(height: 4),
                        CustomTextField(
                            hint: 'ex: Jon Smith@email.com',
                            controller: emailController),
                        const SizedBox(height: 12),
                        const CustomTextStyle(
                          name: 'Password',
                          weight: FontWeight.w400,
                          size: 16,
                          color: Color.fromRGBO(111, 111, 111, 1),
                        ),
                        const SizedBox(height: 4),
                        CustomTextField(
                            lines: 1,
                            obsecure: true,
                            hint: '...........',
                            controller: passwordController),
                        const SizedBox(height: 12),
                        const CustomTextStyle(
                          name: 'Confirm Password',
                          weight: FontWeight.w400,
                          size: 16,
                          color: Color.fromRGBO(111, 111, 111, 1),
                        ),
                        const SizedBox(height: 4),
                        CustomTextField(
                            lines: 1,
                            obsecure: true,
                            hint: '...........',
                            controller: confirmPasswordController),
                        const SizedBox(height: 24),
                        Row(
                          children: [
                            Checkbox(
                              //write a code to change the size of the checkbox
                              materialTapTargetSize:
                                  MaterialTapTargetSize.shrinkWrap,
                              visualDensity: VisualDensity.compact,
                              value: isChecked,
                              onChanged: (bool? value) {
                                setState(() {
                                  isChecked = value!;
                                });
                              },
                              activeColor: Theme.of(context).primaryColor,
                            ),
                            const CustomTextStyle(
                              name: 'I understood the ',
                              weight: FontWeight.w400,
                              size: 12,
                              color: Color.fromRGBO(0, 0, 0, 1),
                            ),
                            CustomTextStyle(
                              name: 'Terms & Conditions',
                              weight: FontWeight.w400,
                              size: 12,
                              color: Theme.of(context).primaryColor,
                            ),
                          ],
                        ),
                        const SizedBox(height: 10),
                        CustomButton(
                          width: MediaQuery.of(context).size.width / 1.25,
                          height: 42,
                          name: 'SIGN UP',
                          fgcolor: Colors.white,
                          textBgColor: Colors.white,
                          bgcolor: Theme.of(context).primaryColor,
                          pressed: () {
                            if (usernameController.text.isEmpty ||
                                emailController.text.isEmpty ||
                                passwordController.text.isEmpty ||
                                confirmPasswordController.text.isEmpty) {
                              ScaffoldMessenger.of(context).showSnackBar(
                                  customSnackBar('All fields are required',
                                      Theme.of(context).secondaryHeaderColor));
                            } else if (passwordController.text !=
                                confirmPasswordController.text) {
                              ScaffoldMessenger.of(context).showSnackBar(
                                  customSnackBar('Passwords do not match',
                                      Theme.of(context).secondaryHeaderColor));
                            }
                            else if (isChecked==false){
                              ScaffoldMessenger.of(context).showSnackBar(
                                  customSnackBar('Did not agreed to terms and conditions',
                                      Theme.of(context).secondaryHeaderColor));
                            } else{
                              context.read<AuthBloc>().add(SignUpEvent(
                                  signUpEntity: SignUpEntity(
                                    email: emailController.text,
                                    password: passwordController.text,
                                    username: usernameController.text,
                                  ),
                                ));
                            }
                            
                          },
                        ),
                        const SizedBox(height: 60),
                        Row(
                          mainAxisAlignment: MainAxisAlignment.center,
                          children: [
                            const CustomTextStyle(
                              color: Color.fromRGBO(111, 111, 111, 1),
                              name: 'Have an account?',
                              weight: FontWeight.w400,
                              size: 16,
                              family: 'Poppins',
                            ),
                            const SizedBox(width: 4),
                            GestureDetector(
                              onTap: () {
                                Navigator.pushNamed(context, '/sign_in_page');
                              },
                              child: CustomTextStyle(
                                color: Theme.of(context).primaryColor,
                                name: 'SIGN IN',
                                weight: FontWeight.w400,
                                size: 16,
                                family: 'Poppins',
                              ),
                            )
                          ],
                        ),
                      ],
                    ),
                  ),
                ],
              ),
            ),
          );
        }
      },
    ));
  }
}
