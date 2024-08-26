import 'package:flutter/material.dart';

class CustomTextField extends StatelessWidget {
  final int? lines;
  final String hint;
  final bool obsecure;
  final TextEditingController controller;
  final TextInputType keyboardType;
  final TextInputAction textInputAction;
  const CustomTextField({
    super.key,
    this.lines,
    this.hint = '',
    required this.controller,
    this.obsecure = false,
    this.keyboardType = TextInputType.text,
    this.textInputAction = TextInputAction.next,
  });

  @override
  Widget build(BuildContext context) {
    return TextField(
      controller: controller,
      maxLines: lines,
      obscureText: obsecure,
      keyboardType: keyboardType,
      textInputAction: textInputAction,
      decoration: InputDecoration(
        contentPadding: const EdgeInsets.only(left: 16),
        hintText: hint,
        filled: true,
        fillColor: const Color.fromARGB(195, 238, 238, 238),
        border: OutlineInputBorder(
          borderRadius: BorderRadius.circular(6),
          borderSide: BorderSide.none),
      ),
    );
  }
}
