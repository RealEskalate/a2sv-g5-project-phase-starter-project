

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../core/input_field/input_field.dart';
import '../../../../core/text/text.dart';
import '../state/Login_Registration/login_registration_bloc.dart';
import '../state/Login_Registration/login_registration_event.dart';
import '../state/Login_Registration/login_registration_state.dart';


class RegistrationInput extends StatelessWidget {
  const RegistrationInput({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        const ConStTexts(
            text: 'Name',
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
                state is OnErrorState ? state.fullName : false;
            final String text = state is OnErrorState ? state.error : '';
            return InputFields(
              text: 'full name',
              error: error,
              errorMessage: text,
              keyboardType: TextInputType.emailAddress,
              showPassword: true,
              onChange: (value) {
                context
                    .read<LoginRegistrationBloc>()
                    .add(OnInputChangeEvent(fullName: value, type: 'fullName'));
              },
            );
          },
        ),
        const SizedBox(
          height: 10,
        ),
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
            final bool error =
                state is OnErrorState ? state.newEmail : false;
            final String text = state is OnErrorState ? state.error : '';
            return InputFields(
              text: 'email',
              error: error,
              errorMessage: text,
              keyboardType: TextInputType.emailAddress,
              showPassword: true,
              onChange: (value) {
                context
                    .read<LoginRegistrationBloc>()
                    .add(OnInputChangeEvent(newEmail: value, type: 'newEmail'));
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
                state is OnErrorState ? state.newPassword : false;
            final String text = state is OnErrorState ? state.error : '';
            return InputFields(
              text: 'Password',
              error: error,
              errorMessage: text,
              keyboardType: TextInputType.emailAddress,
              showPassword: false,
              onChange: (value) {
                context
                    .read<LoginRegistrationBloc>()
                    .add(OnInputChangeEvent(newPassword: value, type: 'newPassword'));
              },
            );
          },
        ),
        const SizedBox(
          height: 10,
        ),
        const ConStTexts(
            text: 'Confirm password',
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
                state is OnErrorState ? state.confirmPassword : false;
            final String text = state is OnErrorState ? state.error : '';
            return InputFields(
              text: 'confirm password',
              error: error,
              errorMessage: text,

              keyboardType: TextInputType.emailAddress,
              showPassword: false,
              onChange: (value) {
                context
                    .read<LoginRegistrationBloc>()
                    .add(OnInputChangeEvent(confirmPassword: value, type: 'confirmPassword'));
              },
            );
          },
        )
      ],
    );
  }
}