import 'package:flutter/material.dart';

class CustomTextField extends StatefulWidget {
  final String name;
  final String placeHolder;
  final TextEditingController controller;
  final String hintText;
  final bool isPassword;
  final Widget? suffixIcon;
  final String? Function(String? value)? validator; // Make the validator nullable

  const CustomTextField({
    Key? key,
    required this.name,
    this.placeHolder = '',
    required this.controller,
    required this.hintText,
    required this.validator,
    this.isPassword = false,
    this.suffixIcon,
  }) : super(key: key);

  @override
  State<CustomTextField> createState() => _CustomTextFieldState();
}

class _CustomTextFieldState extends State<CustomTextField> {
  bool _obscureText = true;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 15, vertical: 8),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.start,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            widget.name,
            style: const TextStyle(
              fontFamily: 'Poppins',
              fontWeight: FontWeight.w400,
              fontSize: 16,
              color: Color(0xFF6F6F6F),
            ),
          ),
          TextFormField(
            obscureText: widget.isPassword ? _obscureText : false,
            controller: widget.controller,
            validator: widget.validator,
            decoration: InputDecoration(
              hintText: widget.hintText,
              hintStyle: const TextStyle(
                color: Colors.grey,
                fontFamily: 'Poppins',
                fontWeight: FontWeight.w400,
                fontSize: 15,
              ),
              fillColor: const Color(0xFFFAFAFA),
              filled: true,
              border: OutlineInputBorder(
                borderRadius: BorderRadius.circular(5),
              ),
              enabledBorder: OutlineInputBorder(
                borderRadius: BorderRadius.circular(10),
                borderSide: const BorderSide(
                  color: Color(0xFFFAFAFA),
                  width: 1.0,
                ),
              ),
              focusedBorder: const OutlineInputBorder(
                borderSide: BorderSide(
                  color: Color.fromARGB(255, 212, 224, 208),
                  width: 2,
                ),
              ),
              suffixIcon: widget.isPassword
                  ? IconButton(
                      icon: Icon(
                        _obscureText ? Icons.visibility_off : Icons.visibility,
                        color: Colors.grey,
                      ),
                      onPressed: () {
                        setState(() {
                          _obscureText = !_obscureText;
                        });
                      },
                    )
                  : widget.suffixIcon,
            ),
          ),
        ],
      ),
    );
  }
}
