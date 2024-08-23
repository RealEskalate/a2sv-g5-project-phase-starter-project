import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

class ECOM extends StatelessWidget {
  final double fontSize;

  const ECOM({
    super.key,
    this.fontSize = 37,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(7.0),
        color: Colors.white,
        border: Border.all(
          color: const Color(0xFF3F51F3),
          width: 2.0,
        ),
        boxShadow: [
          BoxShadow(
            color: Colors.black.withOpacity(0.2),
            spreadRadius: 2,
            blurRadius: 5,
            offset: const Offset(0, 3),
          ),
        ],
      ),
      child: Text(
        '    ECOM     ',
        style: GoogleFonts.caveatBrush(
          textStyle: TextStyle(
            fontSize: fontSize,
            color: const Color(0xFF3F51F3),
            fontWeight: FontWeight.bold,
          ),
        ),
      ),
    );
  }
}
