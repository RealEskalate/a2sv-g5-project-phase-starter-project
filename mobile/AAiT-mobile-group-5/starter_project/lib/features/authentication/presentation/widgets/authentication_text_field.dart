import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

class AuthenticationTextField extends StatelessWidget {
  final String labelText;
  final String hintText;
  final TextEditingController controller;
  final bool isPassword;

  AuthenticationTextField({
    required this.labelText,
    required this.hintText,
    required this.controller,
    this.isPassword = false,
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          labelText,
          style: GoogleFonts.poppins(
            textStyle: const TextStyle(
              color: Color(0xFF6F6F6F),
              fontWeight: FontWeight.normal,
              fontSize: 15,
            ),
          ),
        ),
        const SizedBox(height: 5.0),
        TextFormField(
          controller: controller,
          obscureText: isPassword, // Always obscure the text if isPassword is true
          decoration: InputDecoration(
            hintText: hintText,
            hintStyle: TextStyle(
              color: Colors.grey.shade400,
              fontStyle: FontStyle.italic,
            ),
            filled: true,
            fillColor: Colors.grey.shade100,
            border: OutlineInputBorder(
              borderRadius: BorderRadius.circular(8.0),
              borderSide: BorderSide.none,
            ),
          ),
        ),
      ],
    );
  }
}
