import 'package:flutter/material.dart';

import '../../../../core/themes/themes.dart';

class SearchInput extends StatelessWidget {
  final String hint;
  final TextEditingController control;
  final VoidCallback search;
  final Function(String) onChange;
  const SearchInput(
      {super.key,
      required this.hint,
      required this.control,
      required this.search,
      required this.onChange});

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 20),
      child: TextField(
        controller: control,
        onChanged: onChange,
        decoration: InputDecoration(
            suffixIcon: IconButton(
                onPressed: search,
                icon: const Icon(
                  Icons.arrow_forward,
                  color: MyTheme.ecBlue,
                )),
            fillColor: Colors.white,
            filled: true,
            border: OutlineInputBorder(
              borderRadius: BorderRadius.circular(10),
            ),
            hintText: hint,
            enabledBorder: const OutlineInputBorder(
              borderSide: BorderSide(color: MyTheme.ecGrey),
            ),
            focusedBorder: const OutlineInputBorder(
                borderSide:
                    BorderSide(color: Color.fromARGB(255, 196, 196, 196)))),
      ),
    );
  }
}
