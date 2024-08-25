import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../core/constants/constants.dart';
import '../../../../core/themes/themes.dart';
import '../../../../core/validator/validator.dart';
import '../../../product/presentation/pages/product_list_page.dart';
import '../../../product/presentation/widgets/fill_custom_button.dart';
import '../../../product/presentation/widgets/loading_dialog.dart';
import '../bloc/auth_bloc.dart';
import '../bloc/cubit/user_input_validation_cubit.dart';
import '../widgets/auth_widgets.dart';
import 'signup_page.dart';

// ignore: must_be_immutable
class LoginPage extends StatelessWidget {
  LoginPage({super.key});
  static const String routes = '/log_in_page';
  TextEditingController emailController = TextEditingController();
  TextEditingController passwordController = TextEditingController();
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: BlocListener<AuthBloc, AuthState>(
        listener: (context, state) {
          if (state is LogInSuccessState) {
            Navigator.pop(context);
            Navigator.of(context).pushReplacement(
                MaterialPageRoute(builder: (_) => const ProductListPage()));
          } else if (state is LoginErrorState) {
            Navigator.pop(context);
            ScaffoldMessenger.of(context)
                .showSnackBar(SnackBar(content: Text(state.message)));
          }
        },
        child: SingleChildScrollView(
          child: Padding(
            padding: const EdgeInsets.symmetric(horizontal: 20.0),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                SizedBox(
                  height: MediaQuery.of(context).size.height / 8,
                ),
                const Center(
                  child: Image(
                    height: 80,
                    image: AssetImage(
                      'assets/images/logo.png',
                    ),
                  ),
                ),
                SizedBox(
                  height: MediaQuery.of(context).size.height / 8,
                ),
                const Text(
                  'Sign into your account',
                  style: TextStyle(fontWeight: FontWeight.bold, fontSize: 35),
                ),
                const SizedBox(
                  height: 40,
                ),
                CostumInput(
                  hint: 'ex: json.smith@email.com',
                  control: emailController,
                  text: 'Email',
                  textColor: MyTheme.ecGrey,
                  fromWhere: AppData.login,
                ),
                CostumInput(
                  hint: '***********',
                  control: passwordController,
                  text: 'Password',
                  textColor: MyTheme.ecGrey,
                  fromWhere: AppData.login,
                  obscure: true,
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
                            BlocProvider.of<AuthBloc>(context).add(LogInEvent(
                                email: emailController.text.trim(),
                                password: passwordController.text.trim()));
                          } else {
                            ScaffoldMessenger.of(context)
                                .showSnackBar(SnackBar(content: Text(message)));
                          }
                        },
                        label: 'SIGN IN',
                      ),
                    ),
                  ],
                ),
                SizedBox(
                  height: MediaQuery.of(context).size.height / 8,
                ),
                RichText(
                  key: const Key('click_sing'),
                  text: TextSpan(
                    text: 'Don\'t have an account?',
                    style: const TextStyle(color: MyTheme.ecGrey, fontSize: 20),
                    children: [
                      TextSpan(
                        text: ' SIGN UP',
                        style: const TextStyle(color: MyTheme.ecBlue),
                        recognizer: TapGestureRecognizer()
                          ..onTap = () {
                            BlocProvider.of<UserInputValidationCubit>(context)
                                .reset();
                            Navigator.pushNamed(context, SignUpPage.routes);
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
