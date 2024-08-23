import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

class TextFieldWidget extends StatelessWidget {
  final String hintText;
  final TextEditingController controller;
  final String obscureText;
  final String? Function(String? text)? validator;
   final bool isObscure;
  
  

  const TextFieldWidget(
      {super.key,
      required this.hintText,
      required this.controller,
      required this.obscureText,
      this.validator,
      this.isObscure = false,
      });

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.all(10),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            obscureText,
            style: const TextStyle(color: Color.fromRGBO(111, 111, 111, 15)),
          ),
          const SizedBox(
            height: 10,
          ),
          TextFormField(
            controller: controller,
            obscureText: isObscure,
            validator: (text) => validator?.call(text?.trim()),
            decoration: InputDecoration(
              label: Text(hintText,
                  style: const TextStyle(
                      color: Color.fromRGBO(111, 111, 111, 15))),
              border: const OutlineInputBorder(
                borderSide: BorderSide.none,
                // borderRadius:  BorderRadius.circular(10)
              ),
              filled: true,
              fillColor: const Color.fromRGBO(243, 243, 243, 0.7),
            ),
          ),
        ],
      ),
    );
  }
}
