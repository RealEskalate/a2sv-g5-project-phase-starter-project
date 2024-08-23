import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

import 'custom_text.dart';

class EcomLogo extends StatelessWidget {
  final double? width;
  final double? height;
  final Color? backgroundColor;
  final double boarderRadius;

  final double fontSize;
  const EcomLogo({
    super.key,
    this.fontSize = 14,
    this.width,
    this.height,
    this.backgroundColor,
    this.boarderRadius = 7,
  });

  @override
  Widget build(BuildContext context) {
    const color = Color(0xFF3F51F3);
    return Align(
      alignment: Alignment.center,
      child: Container(
        alignment: Alignment.center,
        height: height,
        width: width,
        decoration: BoxDecoration(
          color: backgroundColor,
          border: Border.all(color: color),
          borderRadius: BorderRadius.all(
            Radius.circular(boarderRadius),
          ),
        ),
        child: Text(
          'ECOM',
          style: GoogleFonts.caveatBrush(fontSize: fontSize, color: color),
        ),
      ),
    );
  }
}
