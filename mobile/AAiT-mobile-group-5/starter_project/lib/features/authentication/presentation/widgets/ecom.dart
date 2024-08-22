import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

class ECOM extends StatelessWidget {
  const ECOM({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(12.0), // Add border radius here
        color: Colors.white, // Background color
        border: Border.all(
          color: Color(0xFF3F51F3), // Border color
          width: 2.0, // Border width
        ),
      ),
      child: Padding(
        padding: const EdgeInsets.symmetric(vertical: 8.0, horizontal: 16),
        child: Text(
          'ECOM',
          style: GoogleFonts.caveatBrush(
            textStyle: const TextStyle(
              fontSize: 37,
              color: Color(0xFF3F51F3),
              fontWeight: FontWeight.bold,
            ),
          ),
        ),
      ),
    );
  }
}
