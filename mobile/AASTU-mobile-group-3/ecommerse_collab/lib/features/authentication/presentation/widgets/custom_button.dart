import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../bloc/blocs.dart';
import '../bloc/events.dart';

class CustomButton extends StatelessWidget {
  final String name;
  final bool login;
  final List<TextEditingController> controllers;
  
  const CustomButton({super.key, required this.name, required this.login, required this.controllers, required Null Function() onPressed});

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
      onPressed: () {
        if (login == true){
          //context read
          context.read<UserBloc>().add(LogInEvent(email: controllers[0].text, password: controllers[1].text));
        
        } else {
          context.read<UserBloc>().add(RegisterUserEvent(email: controllers[0].text, password: controllers[1].text, username: controllers[2].text));
        }
      },
      style: ElevatedButton.styleFrom(
      alignment: Alignment.center,
      backgroundColor: const Color(0xFF3F51F3),
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(8)),
      minimumSize: const Size(288, 42), // Set the width and height here
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
