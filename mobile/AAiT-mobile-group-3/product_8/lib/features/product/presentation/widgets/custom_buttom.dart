import 'package:flutter/material.dart';

class CustomButton extends StatelessWidget {
  final void Function() onPressed;
  final Color backgroundColor;
  final Color foregroundColor;
  final Color borderColor;
  final double buttonWidth;
  final double buttonHeight;

  final Widget child;
  const CustomButton(
      {super.key,
      required this.backgroundColor,
      required this.foregroundColor,
      required this.borderColor,
      required this.onPressed,
      required this.buttonWidth,
      required this.buttonHeight,
      required this.child});

  @override
  Widget build(BuildContext context) {
    return OutlinedButton(
        onPressed: onPressed,
        style: OutlinedButton.styleFrom(
            fixedSize: Size(buttonWidth, buttonHeight),
            backgroundColor: backgroundColor,
            foregroundColor: foregroundColor,
            // padding: EdgeInsets.symmetric(vertical: 10),
            shape:
                RoundedRectangleBorder(borderRadius: BorderRadius.circular(6)),
            side: BorderSide(color: borderColor, width: 1)),
        child: child);
  }
}
