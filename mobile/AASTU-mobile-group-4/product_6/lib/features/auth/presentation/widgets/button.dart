import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

class PrimaryButton extends StatelessWidget {
  final String text;
  
  const PrimaryButton({super.key, required this.text});

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
      style: ElevatedButton.styleFrom(
          minimumSize: const Size(308, 45),
          padding: const EdgeInsets.symmetric(horizontal: 25),
          backgroundColor: Color.fromRGBO(63, 81, 243, 1),
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(8),
          )),
      onPressed: () {
        
      },
      child: Text(
        '$text',
        style: GoogleFonts.poppins(
          textStyle: const TextStyle(
            fontWeight: FontWeight.w600,
            fontSize: 15,
            color: Color.fromRGBO(255, 255, 255, 1),
          ),
        ),
      ),
    );
  }
}
