import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';

import '../../../../core/Colors/colors.dart';
import '../../../../core/text/text.dart';
import '../state/Login_Registration/login_registration_bloc.dart';
import '../state/Login_Registration/login_registration_event.dart';
import '../state/Login_Registration/login_registration_state.dart';
import 'login_register.dart';
import 'login_register_buttons.dart';
import 'registration_app_bar.dart';
import 'registration_input.dart';

class RegisterBody extends StatelessWidget {
  const RegisterBody({super.key});

  @override
  Widget build(BuildContext context) {
    final width = MediaQuery.of(context).size.width;
    final height = MediaQuery.of(context).size.height;
    
    return SafeArea(
      child: Scaffold(
        body: SingleChildScrollView(
          child: BlocListener<LoginRegistrationBloc, LoginRegistrationState>(
            listener: (context, state) {
              if (state is OnLoading) {
                EasyLoading.showProgress(0.3, status: 'loading...');
              } else if (state is OnErrorState) {
                if (!state.email &
                    !state.password &
                    !state.newEmail &
                    !state.newPassword &
                    !state.confirmPassword &
                    !state.fullName &
                    !state.terms) {
                  EasyLoading.showError(state.error);
                }
              } else if (state is RegistrationSuccess) {
                EasyLoading.showSuccess('Success');
                context.read<LoginRegistrationBloc>().add(
                      OnInputChangeEvent(newEmail: '', type: 'newEmail'),
                    );
                context.read<LoginRegistrationBloc>().add(
                      OnInputChangeEvent(newPassword: '', type: 'newPassword'),
                    );
                context.read<LoginRegistrationBloc>().add(
                      OnInputChangeEvent(
                          confirmPassword: '', type: 'confirmPassword'),
                    );
                context.read<LoginRegistrationBloc>().add(
                      OnInputChangeEvent(fullName: '', type: 'fullName'),
                    );
                Navigator.pop(context);
              }
            },
            child: Container(
              padding: const EdgeInsets.only(left: 30, right: 30),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  const SizedBox(
                    height: 70,
                  ),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      GestureDetector(
                        onTap: () {
                          context.read<LoginRegistrationBloc>().add(
                                OnInputChangeEvent(
                                    newEmail: '', type: 'newEmail'),
                              );
                          context.read<LoginRegistrationBloc>().add(
                                OnInputChangeEvent(
                                    newPassword: '', type: 'newPassword'),
                              );
                          context.read<LoginRegistrationBloc>().add(
                                OnInputChangeEvent(
                                    confirmPassword: '', type: 'confirmPassword'),
                              );
                          context.read<LoginRegistrationBloc>().add(
                                OnInputChangeEvent(
                                    fullName: '', type: 'fullName'),
                              );
      
                          Navigator.pop(context);
                        },
                        child: SizedBox(
                            width: width * 0.05,
                            child: const Icon(Icons.arrow_back_ios_new,
                                color: Color(0xff3F51F3))),
                      ),
                      const RegistrationAppBar(),
                    ],
                  ),
                  const SizedBox(
                    height: 50,
                  ),
                  const SizedBox(
                    height: 35,
                    child: ConStTexts(
                      text: 'Create your account',
                      color: mainColor,
                      fontSize: 26.72,
                      fontWeight: FontWeight.bold,
                      fontFamily: 'Poppins',
                    ),
                  ),
                  const SizedBox(
                    height: 20,
                  ),
                  const RegistrationInput(),
                  BlocBuilder<LoginRegistrationBloc, LoginRegistrationState>(
                        builder: (context, state) {
                          final bool error =
                              state is OnErrorState ? state.terms : false;
                          bool checked = state is OnInputChange? state.terms:false;
                         
                      return CheckboxListTile(
                        isError: error,
                         
                        shape: BeveledRectangleBorder(
                            borderRadius: BorderRadius.circular(10)),
                        contentPadding: const EdgeInsets.all(0),
                        materialTapTargetSize: MaterialTapTargetSize.shrinkWrap,
                        controlAffinity: ListTileControlAffinity.leading,
                        title: Transform.translate(
                          offset: const Offset(-17, 0),
                          child: const LoginRegister(
                              text: 'I understood the ',
                              text2: 'terms & policy',
                              fontSize: 12,
                              hight: 35),
                        ),
                        value: checked,
                        onChanged: (checked) {
                          
                          context.read<LoginRegistrationBloc>().add(
      
                                OnInputChangeEvent(terms: checked ?? false, type: 'terms'),
                              );
                        },
                      );
                    },
                  ),
                  const SizedBox(
                    height: 10,
                  ),
                  GestureDetector(
                    onTap: () {
                      context
                          .read<LoginRegistrationBloc>()
                          .add(RegisterButtonPressed());
                    },
                    child: const LoginRegisterButtons(
                      text: 'SIGN IN',
                    ),
                  ),
                  SizedBox(
                    height: height * 0.1,
                  ),
                  const LoginRegister(
                    text: 'Have an account? ',
                    text2: 'Sign IN',
                    navigaror: '',
                    hight: 55,
                  ),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}
