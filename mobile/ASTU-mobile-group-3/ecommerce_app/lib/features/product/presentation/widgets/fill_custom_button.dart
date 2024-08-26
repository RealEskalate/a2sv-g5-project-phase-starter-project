import 'package:flutter/material.dart';

import '../../../../core/themes/themes.dart';

class FillCustomButton extends StatelessWidget {
  final VoidCallback press;
  final String label;
  const FillCustomButton({super.key, required this.press, required this.label});
  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
      key: Key(label),
      onPressed: press,
      style: FilledButton.styleFrom(
          padding: const EdgeInsets.symmetric(vertical: 16),
          backgroundColor: MyTheme.ecBlue,
          shape:
              RoundedRectangleBorder(borderRadius: BorderRadius.circular(5))),
      child: Text(label,
      style: TextStyle(
        color: Colors.white,
        fontFamily: 'Poppins',
        fontSize: 15,
        fontWeight: FontWeight.w600),),
    );
  }
}
