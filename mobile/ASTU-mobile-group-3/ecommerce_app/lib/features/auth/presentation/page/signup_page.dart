import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../core/constants/constants.dart';
import '../../../../core/themes/themes.dart';
import '../../../../core/validator/validator.dart';
import '../../../product/presentation/widgets/fill_custom_button.dart';
import '../../../product/presentation/widgets/loading_dialog.dart';
import '../bloc/auth_bloc.dart';
import '../bloc/cubit/user_input_validation_cubit.dart';
import '../widgets/auth_widgets.dart';

// ignore: must_be_immutable
class SignUpPage extends StatelessWidget {
  static const String routes = '/sign_up_page';
  SignUpPage({super.key});
  TextEditingController nameController = TextEditingController();
  TextEditingController emailController = TextEditingController();
  TextEditingController passwordController = TextEditingController();
  TextEditingController confirmPassController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: MyTheme.ecWhite,
        leading: IconButton(
          onPressed: () {
            Navigator.pop(context);
          },
          icon: const Icon(Icons.chevron_left),
        ),
        actions: const [
          Padding(
            padding: EdgeInsets.all(15.0),
            child: Image(
              image: AssetImage('assets/images/logo.png'),
            ),
          ),
        ],
      ),
      body: BlocListener<AuthBloc, AuthState>(
        listener: (context, state) {
          if (state is RegisterSuccessState) {
            Navigator.pop(context);

            Navigator.pop(context);
          } else if (state is SignupErrorState) {
            Navigator.pop(context);

            ScaffoldMessenger.of(context).showSnackBar(
              SnackBar(
                content: Text(state.message),
              ),
            );
          }
        },
        child: SingleChildScrollView(
          child: Padding(
            padding: const EdgeInsets.symmetric(horizontal: 20.0),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                const SizedBox(
                  height: 40,
                ),
                const Text(
                  'Sign into your account',
                  style: TextStyle(fontWeight: FontWeight.bold, fontSize: 35),
                ),
                const SizedBox(
                  height: 40,
                ),
                CostumInput(
                  hint: 'ex: jon smith',
                  control: nameController,
                  text: 'Name',
                  textColor: MyTheme.ecGrey,
                  fromWhere: AppData.signup,
                ),
                CostumInput(
                  hint: 'ex: json.smith@email.com',
                  control: emailController,
                  text: 'Email',
                  textColor: MyTheme.ecGrey,
                  fromWhere: AppData.signup,
                ),
                CostumInput(
                  hint: '***********',
                  control: passwordController,
                  text: 'Password',
                  textColor: MyTheme.ecGrey,
                  fromWhere: AppData.signup,
                  obscure: true,
                ),
                CostumInput(
                  hint: '***********',
                  control: confirmPassController,
                  text: 'Confirm Password',
                  textColor: MyTheme.ecGrey,
                  fromWhere: AppData.signup,
                  obscure: true,
                ),
                Padding(
                  padding:
                      const EdgeInsets.symmetric(horizontal: 15, vertical: 10),
                  child: Row(
                    children: [
                      BlocBuilder<UserInputValidationCubit,
                          UserInputValidationState>(
                        builder: (context, state) {
                          return Checkbox(
                            key: const Key(InputDataValidator.checkBox),
                            value: state.checkbox == AppData.strValidated,
                            onChanged: (val) {
                              BlocProvider.of<UserInputValidationCubit>(context)
                                  .changeCheckbox(AppData.signup, val!);
                            },
                          );
                        },
                      ),
                      RichText(
                        text: const TextSpan(
                          text: 'I undrestood the ',
                          style: TextStyle(color: MyTheme.ecGrey, fontSize: 15),
                          children: [
                            TextSpan(
                                text: 'term & policy.',
                                style: TextStyle(color: MyTheme.ecBlue))
                          ],
                        ),
                      )
                    ],
                  ),
                ),
                Row(
                  children: [
                    Expanded(
                      child: FillCustomButton(
                        press: () {
                          String? message =
                              BlocProvider.of<UserInputValidationCubit>(context)
                                  .state
                                  .validate();
                          if (message == null) {
                            showDialog(
                                context: context,
                                builder: (context) {
                                  return const LoadingDialog();
                                });
                            BlocProvider.of<AuthBloc>(context).add(
                              SignUpEvent(
                                name: nameController.text.trim(),
                                email: emailController.text.trim(),
                                password: passwordController.text.trim(),
                              ),
                            );
                          } else {
                            ScaffoldMessenger.of(context)
                                .showSnackBar(SnackBar(content: Text(message)));
                          }
                        },
                        label: 'SIGN UP',
                      ),
                    ),
                  ],
                ),
                SizedBox(
                  height: MediaQuery.of(context).size.height / 8,
                ),
                RichText(
                  text: TextSpan(
                    text: 'Have an account?',
                    style: const TextStyle(color: MyTheme.ecGrey, fontSize: 20),
                    children: [
                      TextSpan(
                        text: ' SIGN IN',
                        style: const TextStyle(color: MyTheme.ecBlue),
                        recognizer: TapGestureRecognizer()
                          ..onTap = () {
                            BlocProvider.of<UserInputValidationCubit>(context)
                                .reset();
                            Navigator.pop(context);
                          },
                      )
                    ],
                  ),
                ),
                const SizedBox(
                  height: 40,
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
