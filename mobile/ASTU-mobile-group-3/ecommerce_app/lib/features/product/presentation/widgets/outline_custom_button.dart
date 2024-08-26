import 'package:flutter/material.dart';

import '../../../../core/themes/themes.dart';

class OutlineCustomButton extends StatelessWidget {
  final VoidCallback press;
  final String label;
  const OutlineCustomButton(
      {super.key, required this.press, required this.label});
  @override
  Widget build(BuildContext context) {
    return OutlinedButton(
      onPressed: press,
      style: FilledButton.styleFrom(
        foregroundColor: MyTheme.ecRed,
        side: const BorderSide(color: MyTheme.ecRed),
        padding: EdgeInsets.symmetric(vertical: 18),
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(5),
        ),
      ),
      child: Text(label,
      maxLines: 1,
      style : TextStyle( 
      fontFamily:   'Poppins',
      fontSize: 15,
      fontWeight: FontWeight.w600
      )),
    );
  }
}
