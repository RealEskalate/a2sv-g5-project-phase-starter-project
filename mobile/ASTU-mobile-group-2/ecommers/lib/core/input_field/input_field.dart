
import 'package:flutter/material.dart';




class InputFields extends StatelessWidget {
  final String text;
  final Widget? prefixIcon;
  final Widget? suffixIcon;
  final TextInputType keyboardType;
  final bool enableSuggestions;
  final bool showPassword;
  final bool error;
  final String errorMessage;
  final ValueChanged<String> onChange;

  const InputFields({
    super.key,
    required this.text,
    this.prefixIcon,
    this.suffixIcon,
    this.keyboardType = TextInputType.text,
    this.enableSuggestions = true,
    this.showPassword = false,
    this.error = false,
    this.errorMessage = '',
    required this.onChange,

    });

  @override
  Widget build(BuildContext context) {
    return TextField(
      onChanged: onChange,
      keyboardType: keyboardType,
      enableSuggestions: enableSuggestions,
      autocorrect: enableSuggestions,
      obscureText: !showPassword,
      
      decoration: InputDecoration(
        fillColor: const Color.fromARGB(255, 240, 240, 240),
        filled: true,
        error: error?Text(errorMessage):null,
        prefixIcon: prefixIcon,
        suffixIcon: suffixIcon,

        hintText:text,
        hintStyle:  const TextStyle(
          color: Color.fromARGB(255, 199, 196, 196),
        ),
        
        border: OutlineInputBorder(
          borderRadius: BorderRadius.circular(20),
          
          borderSide: const BorderSide(
            color: Color.fromARGB(255, 255, 255, 255),
            width: 1,
          ),
        ),

        focusedBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(20),
          borderSide: const BorderSide(
            color: Color.fromARGB(255, 255, 255, 255),
            width: 1,
          ),
        ),
        enabledBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(20),
          borderSide: const BorderSide(
            color: Color.fromARGB(255, 255, 255, 255),
            width: 1,
          ),
        ),

        focusColor: Colors.grey,
        
      ),
    );
  }
}