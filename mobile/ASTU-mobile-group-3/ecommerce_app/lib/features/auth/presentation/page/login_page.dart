import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../core/constants/constants.dart';
import '../../../../core/themes/themes.dart';
import '../../../../dependency_injection.dart';
import '../../../../landing_page.dart';
import '../../../chat/data/data_resources/socket_io_sesrvice.dart';
import '../../../product/presentation/widgets/loading_dialog.dart';
import '../bloc/auth_bloc.dart';
import '../bloc/cubit/user_input_validation_cubit.dart';
import '../widgets/auth_widgets.dart';
import '../widgets/reusable_button.dart';
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
            locator<SocketIOService>().connect();
            Navigator.of(context).pushReplacement(
                MaterialPageRoute(builder: (_) => const LandingPage()));
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
                    height: 50,
                    width: 144,
                    image: AssetImage(
                      'assets/images/Group 67.png',
                    ),
                  ),
                ),
                SizedBox(
                  height: MediaQuery.of(context).size.height / 12,
                ),
                const Text(
                  'Sign into your account',
                  style: TextStyle(fontWeight: FontWeight.w600, fontSize: 26),
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
                      child: Padding(
                        padding: const EdgeInsets.symmetric(horizontal: 18.0),
                        child: GestureDetector(
                          onTap: () {
                            String? message =
                                BlocProvider.of<UserInputValidationCubit>(
                                        context)
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
                              ScaffoldMessenger.of(context).showSnackBar(
                                  SnackBar(content: Text(message)));
                            }
                          },
                          child: const ReusableButton(
                            lable: 'SIGN IN',
                          ),
                        ),
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
                    style: const TextStyle(
                      color: MyTheme.ecGrey,
                      fontSize: 18,
                      fontWeight: FontWeight.w400,
                    ),
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
