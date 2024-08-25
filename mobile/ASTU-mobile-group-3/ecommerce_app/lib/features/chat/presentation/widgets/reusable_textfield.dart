import 'package:flutter/material.dart';

class ReusableTextField extends StatelessWidget {
  final String hint;
  final TextEditingController textEditingController;
  final TextInputType textInputType;
  final String? Function(String?)? validator;

  const ReusableTextField({
    Key? key,
    required this.hint,
    required this.textEditingController,
    required this.textInputType,
    this.validator, // Optional validator parameter
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 50,
      padding: const EdgeInsets.symmetric(horizontal: 4),
      decoration: BoxDecoration(
        color: const Color.fromARGB(255, 234, 244, 227),
        borderRadius: BorderRadius.circular(6),
      ),
      child: Center(
        child: Padding(
          padding: const EdgeInsets.only(top: 20),
          child: TextFormField(
            keyboardType: textInputType,
            controller: textEditingController,
            validator: validator, // Apply the validator here
            decoration: InputDecoration(
              suffixIcon: Padding(
                padding: const EdgeInsets.only(bottom: 8.0),
                child: Icon(Icons.file_copy),
              ),
              hintText: hint,
              contentPadding: const EdgeInsetsDirectional.only(start: 4),
              hintStyle: const TextStyle(color: Color(0xffC1C1C1)),
              border: InputBorder.none,
              isDense: true,
            ),
          ),
        ),
      ),
    );
  }
}
