import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../bloc/blocs.dart';
import '../bloc/events.dart';

class CustomButton extends StatelessWidget {
  final String name;
  final bool login;
  final GlobalKey<FormState> formKey;
  final List<TextEditingController> controllers;
  final Function() onPressed; // This parameter is not used in the current implementation
  final bool? isCheckboxChecked;

  const CustomButton({
    super.key,
    required this.name,
    required this.login,
    required this.formKey,
    required this.controllers,
    required this.onPressed, // This should be used in the build method
    this.isCheckboxChecked,
  });

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
      onPressed: () {
        // Currently handling the onPressed event here
        if (formKey.currentState?.validate() ?? false) {
          if (login) {
            // if (isCheckboxChecked != true) {
            //   ScaffoldMessenger.of(context).showSnackBar(
            //     const SnackBar(
            //       content: Text("You must accept the terms to log in"),
            //       backgroundColor: Colors.red,
            //     ),
            //   );
            //   return;
            // }
            context.read<UserBloc>().add(
              LogInEvent(
                email: controllers[0].text,
                password: controllers[1].text,
              ),
            );
          } else {
            if (!isCheckboxChecked!) {
              ScaffoldMessenger.of(context).showSnackBar(
                const SnackBar(
                  content: Text("You must accept the terms to register"),
                  backgroundColor: Color.fromARGB(255, 224, 81, 70),
                ),
              );
              return;
            }
            context.read<UserBloc>().add(
              RegisterUserEvent(
                email: controllers[0].text,
                password: controllers[1].text,
                username: controllers[2].text,
              ),
            );
          }
        } else {
          ScaffoldMessenger.of(context).showSnackBar(
            const SnackBar(
              content: Text("Please fix the errors in the form"),
              backgroundColor: Colors.red,
            ),
          );
        }
      },
      style: ElevatedButton.styleFrom(
        alignment: Alignment.center,
        backgroundColor: const Color(0xFF3F51F3),
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(8)),
        minimumSize: const Size(288, 42),
      ),
      child: Text(
        name,
        style: const TextStyle(
          color: Colors.white,
          fontFamily: 'Poppins',
          fontWeight: FontWeight.w600,
          fontSize: 20,
        ),
      ),
    );
  }
}
