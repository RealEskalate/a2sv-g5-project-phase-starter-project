import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../state/Login_Registration/login_registration_bloc.dart';
import '../state/Login_Registration/login_registration_event.dart';
import '../state/Login_Registration/login_registration_state.dart';
import 'header.dart';
import 'login_body.dart';
import 'login_register.dart';
import 'login_register_buttons.dart';

class LoginHomePage extends StatelessWidget {
  const LoginHomePage({super.key});

  @override
  Widget build(BuildContext context) {
    
    final hight = MediaQuery.of(context).size.height;
    return SafeArea(
      child: Scaffold(
        backgroundColor: Colors.white,
        body: BlocListener<LoginRegistrationBloc, LoginRegistrationState>(
          listener: (context, state) {
            if (state is OnErrorState) {
              if (!state.email &
                  !state.password &
                  !state.newEmail &
                  !state.newPassword &
                  !state.confirmPassword &
                  !state.fullName) {
                
                ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(
                      content: Text(state.error),
                    ),
                  );
              }
            } else if (state is LoginSuccess) {
              context
              .read<LoginRegistrationBloc>()
              .add(OnInputChangeEvent(email: '', type: 'email'));
            context
              .read<LoginRegistrationBloc>()
              .add(OnInputChangeEvent(email: '', type: 'password'));
              context
                .read<LoginRegistrationBloc>()
                .add(OnInputChangeEvent(newEmail: '', type: 'newEmail'),);
                context
                .read<LoginRegistrationBloc>()
                .add(OnInputChangeEvent(newPassword: '', type: 'newPassword'),);
                context
                .read<LoginRegistrationBloc>()
                .add(OnInputChangeEvent(confirmPassword: '', type: 'confirmPassword'),);
                context
                .read<LoginRegistrationBloc>()
                .add(OnInputChangeEvent(fullName: '', type: 'fullName'),);
              context.read<LoginRegistrationBloc>().add(
      
                  OnInputChangeEvent(terms: false, type: 'terms'),
                );
                    
              Navigator.pushReplacementNamed(context, '/home');
            }
          },
          child: SingleChildScrollView(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Container(
                  padding: const EdgeInsets.all(16),
                  margin: const EdgeInsets.only(top: 35),
                  child: Center(
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.center,
                      // mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        // header part
                        const SizedBox(
                          height: 60,
                        ),
                        const Header(),
                        const SizedBox(
                          height: 40,
                        ),
      
                        const LoginBody(),
                        const SizedBox(
                          height: 27,
                        ),
                        BlocBuilder<LoginRegistrationBloc, LoginRegistrationState>(
                          builder: (context, state) {
                            return GestureDetector(
                              onTap: state is OnLoading?null:() {
                                context
                                    .read<LoginRegistrationBloc>()
                                    .add(LoginButtonPressed());
                              },
                              child: LoginRegisterButtons(
                                text: state is OnLoading?'Loading...':'SIGN IN',
                              ),
                            );
                          },
                        ),
                        SizedBox(
                          height: hight*0.12,
                        ),
                      ],
                    ),
                  ),
                ),
      
                const Padding(
                  padding:  EdgeInsets.all(16),
                  child: 
                      LoginRegister(
                        text: 'Don\'t have an account? ',
                        text2: 'Sign Up',
                        navigaror: '/registration',
                        hight: 55,
                      ),
                    
                )
                // const RegisterBody(),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
