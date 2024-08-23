import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../core/input_field/input_field.dart';
import '../../../../core/text/text.dart';
import '../state/Login_Registration/login_registration_bloc.dart';
import '../state/Login_Registration/login_registration_event.dart';
import '../state/Login_Registration/login_registration_state.dart';

class LoginBody extends StatelessWidget {
  const LoginBody({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        const ConStTexts(
            text: 'Email',
            fontSize: 16,
            fontFamily: 'Poppins',
            color: Color(0xff6F6F6F),
            fontWeight: FontWeight.w400),
        const SizedBox(
          height: 10,
        ),
        BlocBuilder<LoginRegistrationBloc, LoginRegistrationState>(
          builder: (context, state) {
            final bool email = state is OnErrorState ? state.email : false;
            final String error = state is OnErrorState ? state.error : '';
            return InputFields(
              text: 'Email Address',
              keyboardType: TextInputType.emailAddress,
              showPassword: true,
              error: email,
              errorMessage: error,
              onChange: (value) {
                context
                    .read<LoginRegistrationBloc>()
                    .add(OnInputChangeEvent(email: value, type: 'email'));
              },
            );
          },
        ),
        const SizedBox(
          height: 10,
        ),
        const ConStTexts(
            text: 'Password',
            fontSize: 16,
            fontFamily: 'Poppins',
            color: Color(0xff6F6F6F),
            fontWeight: FontWeight.w400),
        const SizedBox(
          height: 10,
        ),
        BlocBuilder<LoginRegistrationBloc, LoginRegistrationState>(
          builder: (context, state) {
            final bool error =
                state is OnErrorState ? state.password : false;
            final String text = state is OnErrorState ? state.error : '';
            return InputFields(
              text: '********',
              error: error,
              errorMessage: text,
              keyboardType: TextInputType.emailAddress,
              showPassword: false,
              onChange: (value) {
                context
                    .read<LoginRegistrationBloc>()
                    .add(OnInputChangeEvent(password: value, type: 'password'));
              },
            );
          },
        )
      ],
    );
  }
}
