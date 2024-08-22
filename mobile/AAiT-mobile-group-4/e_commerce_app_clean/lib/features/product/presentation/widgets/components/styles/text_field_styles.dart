import 'package:flutter/material.dart';

class CustomTextField extends StatelessWidget {
  final int? lines;
  final String hint;
  final bool obsecure;
  final TextEditingController controller;
  const CustomTextField({
    super.key,
    this.lines,
    this.hint = '',
    required this.controller,
    this.obsecure = false,
  });

  @override
  Widget build(BuildContext context) {
    return TextField(
      controller: controller,
      maxLines: lines,
      obscureText: obsecure,
      decoration: InputDecoration(
        contentPadding: const EdgeInsets.fromLTRB(16, -10, 0, -10),
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
