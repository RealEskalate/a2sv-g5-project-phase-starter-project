


import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../state/Login_Registration/login_registration_bloc.dart';
import '../state/Login_Registration/login_registration_event.dart';


class LoginRegister extends StatelessWidget {
  final String text;
  final String text2;
  final Color color;
  final int hight;
  final  double fontSize;
  final String navigaror;
  const LoginRegister({
    super.key,
    required this.text,
    this.color = Colors.grey,
    this.hight = 25,
    this.fontSize = 16,
    this.navigaror = '/registration',
    required this.text2

    });

  @override
  Widget build(BuildContext context) {
    
    return RichText(
      textHeightBehavior: const TextHeightBehavior(
        applyHeightToFirstAscent: false,
        applyHeightToLastDescent: true,
      ),
      overflow: TextOverflow.ellipsis,
      text: TextSpan(
        
        children: <TextSpan>[
          TextSpan(
            text: text,
            style: const TextStyle(
              color: Colors.grey,
              fontFamily: 'Poppins',
            )
            ),
          TextSpan(
              text: text2,
              style: const TextStyle(
                fontFamily: 'Poppins',
                color: Color(0xff3F51F3),
              ),
              recognizer: TapGestureRecognizer()
                ..onTap = () {
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
              
                    if(navigaror.isNotEmpty){
                     
                    Navigator.pushNamed(context, '/registration');
                    } else{
                      Navigator.pop(context);
                    }
                }),
          
        ],
      ),
    );
    
   
  }
}


