import 'package:flutter/material.dart';
import 'package:starter_project/features/authentication/presentation/widgets/authentication_text_field.dart';
import 'package:starter_project/features/authentication/presentation/widgets/ecom.dart';

void main() {
  runApp(MaterialApp(
    home: Scaffold(
      body: Column(
        children: [
          const ECOM(),
          Center(
            child: Padding(
              padding: const EdgeInsets.all(16.0),
              child: Column(
                children: [
                  AuthenticationTextField(
                    labelText: 'Name',
                    hintText: 'ex: jon smith',
                    controller: TextEditingController(),
                  ),
                  AuthenticationTextField(
                    labelText: 'Password',
                    hintText: '***********',
                    controller: TextEditingController(),
                    isPassword: true,
                  ),
                  FilledButton(onPressed:(){}, child: Text("SIGN IN"),style: ButtonStyle(),)
                ],
              ),
            ),
          ),
        ],
      ),
    ),
  ));
}
