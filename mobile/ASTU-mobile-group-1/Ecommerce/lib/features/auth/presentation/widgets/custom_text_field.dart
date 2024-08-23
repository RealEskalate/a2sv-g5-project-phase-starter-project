import 'package:flutter/material.dart';

import 'custom_text.dart';

class CustomTextField extends StatelessWidget {
  const CustomTextField(
      {super.key,
      required this.field,
      this.fillColor,
      this.hintText,
      this.suffixIcon,
      this.obscureText = false,
      TextEditingController? controller})
      : _controller = controller;
  final String? hintText;
  final Icon? suffixIcon;
  final Color? fillColor;
  final String field;
  final bool obscureText;
  final TextEditingController? _controller;

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        CustomText(
          text: field,
          fontSize: 16,
          color: const Color(0xFF6F6F6F),
        ),
        Padding(
          padding: const EdgeInsets.all(8.0),
          child: TextField(
            controller: _controller,
            obscureText: obscureText,
            maxLines: 1,
            minLines: null,
            onTapOutside: (event) {
              FocusManager.instance.primaryFocus?.unfocus();
            },
            decoration: InputDecoration(
              border: const OutlineInputBorder(
                borderSide: BorderSide.none,
                borderRadius: BorderRadius.all(
                  Radius.circular(
                    6,
                  ),
                ),
              ),
              focusedBorder: const OutlineInputBorder(
                borderSide: BorderSide.none,
                borderRadius: BorderRadius.all(
                  Radius.circular(
                    10,
                  ),
                ),
              ),
              fillColor: fillColor,
              filled: true,
              suffixIcon: suffixIcon,
              hintText: hintText,
              contentPadding:
                  const EdgeInsets.symmetric(vertical: 0, horizontal: 12),
            ),
          ),
        ),
      ],
    );
  }
}
