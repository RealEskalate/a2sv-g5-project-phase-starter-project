import 'package:flutter/material.dart';

class MyTextField extends StatelessWidget {
  final String lable;
  final dynamic suffIcon;
  final int lines;
  final TextEditingController controller;
  final bool obscureText;
  final String hint;
  const MyTextField({
    super.key,
    required this.lable,
    required this.lines,
    this.suffIcon,
    required this.controller,
    this.obscureText = false,
    this.hint = '',
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.start,
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          lable,
          style: const TextStyle(fontSize: 16),
          textAlign: TextAlign.left,
        ),
        const SizedBox(
          height: 5,
        ),
        TextField(
          maxLines: lines,
          controller: controller,
          decoration: InputDecoration(
            border: const OutlineInputBorder(
              borderSide: BorderSide.none,
            ),
            suffixIcon: suffIcon,
            filled: true,
            fillColor: const Color.fromARGB(155, 232, 229, 229),
            hintText: hint,
            hintStyle: const TextStyle(
                fontSize: 16, color: Color.fromARGB(255, 140, 138, 138)),
          ),
          obscureText: obscureText,
        ),
        const SizedBox(
          height: 5,
        ),
      ],
    );
  }
}
