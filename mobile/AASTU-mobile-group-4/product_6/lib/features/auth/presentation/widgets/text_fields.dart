import 'package:flutter/material.dart';

class Fields extends StatelessWidget {
  final TextEditingController controller;
  final String hintText;
  final bool obsecureText;
  final String? Function(String?) validator;

  const Fields({
    super.key,
    required this.controller,
    required this.hintText,
    required this.obsecureText,
    required this.validator,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 25),
      child: TextFormField(
        controller: controller,
        obscureText: obsecureText,
        decoration: InputDecoration(
          enabledBorder: const OutlineInputBorder(
            borderSide: BorderSide(color: Colors.white),
          ),
          focusedBorder: const OutlineInputBorder(
            borderSide: BorderSide(color: Colors.grey),
          ),
          fillColor: const Color.fromRGBO(250, 250, 250, 1),
          filled: true,
          hintText: hintText,
        ),
        validator: validator,
      ),
    );
  }
}
