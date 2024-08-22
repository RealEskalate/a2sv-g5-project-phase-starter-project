import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

class Logo extends StatelessWidget {
  const Logo(
      {super.key,
      required this.height,
      required this.textheight,
      required this.paddingBottom,
      required this.paddingHorizontal});
  //height
  final double height;
  final double textheight;
  final double paddingBottom;
  final double paddingHorizontal;

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.only(
        left: paddingHorizontal,
        right: paddingHorizontal,
        bottom: paddingBottom,
      ),
      height: height,
      decoration: BoxDecoration(
        border: Border.all(
          color: Color.fromRGBO(63, 81, 243, 1), // Outline color
          width: 1.0, // Outline width
        ),
        borderRadius: BorderRadius.circular(8),
      ),
      child: Text(
        'ECOM',
        style: GoogleFonts.caveatBrush(
          textStyle: TextStyle(
            fontWeight: FontWeight.w400,
            fontSize: textheight,
            color: Color.fromRGBO(63, 81, 243, 1),
          ),
        ),
      ),
    );
  }
}
