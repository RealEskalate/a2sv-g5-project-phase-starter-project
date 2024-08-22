import 'package:flutter/material.dart';

class AuthTextField extends StatelessWidget {
  final String label;
  final TextEditingController controller;
  final String hintText;
  final bool isObscure;
  final String? Function(String? text)? validator;

  const AuthTextField({
    super.key,
    required this.controller,
    required this.hintText,
    required this.label,
    this.isObscure = false,
    this.validator,
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          label,
          style: const TextStyle(color: Colors.black38, fontSize: 17),
        ),
        const SizedBox(
          height: 5,
        ),
        TextFormField(
          validator: (text) => validator?.call(text?.trim()),
          obscureText: isObscure,
          controller: controller,
          onChanged: (text) {
            controller.value = controller.value.copyWith(
              text: text.trim(),
              selection: TextSelection.fromPosition(
                TextPosition(offset: text.trim().length),
              ),
            );
          },
          decoration: InputDecoration(
            filled: true,
            fillColor: Colors.grey[100],
            hintText: hintText,
            hintStyle: const TextStyle(
              fontWeight: FontWeight.w500,
              color: Colors.black38,
            ),
            contentPadding:
                const EdgeInsets.symmetric(horizontal: 20, vertical: 15),
            border: OutlineInputBorder(
              borderRadius: BorderRadius.circular(10),
              borderSide: BorderSide.none,
            ),
          ),
        ),
      ],
    );
  }
}
