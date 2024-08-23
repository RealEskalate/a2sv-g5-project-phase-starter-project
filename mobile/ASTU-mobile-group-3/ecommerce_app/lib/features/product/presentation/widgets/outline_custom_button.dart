import 'package:flutter/material.dart';

import '../../../../core/themes/themes.dart';

class OutlineCustomButton extends StatelessWidget {
  final VoidCallback press;
  final String label;
  const OutlineCustomButton(
      {super.key, required this.press, required this.label});
  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 10),
      child: OutlinedButton(
        onPressed: press,
        style: FilledButton.styleFrom(
          foregroundColor: MyTheme.ecRed,
          side: const BorderSide(color: MyTheme.ecRed),
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(5),
          ),
        ),
        child: Text(label),
      ),
    );
  }
}
