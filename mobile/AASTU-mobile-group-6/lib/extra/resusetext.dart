import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

class Resusetext extends StatelessWidget {
const Resusetext({ Key? key }) : super(key: key);

  @override
  Widget build(BuildContext context){
    return Container();
  }
}
Text reusableText(String text, FontWeight weight, double size,
    [Color color = Colors.black]) {
  return Text(
    text,
    overflow: TextOverflow.clip,
    style: GoogleFonts.poppins(fontWeight: weight, fontSize: size, color: color),
  );
}

Text reusableTextpar(String text, FontWeight weight, double size,
    [Color color = Colors.black]) {
  return Text(
    text,
    maxLines: 10,
    overflow: TextOverflow.ellipsis,
    style: GoogleFonts.poppins(fontWeight: weight, fontSize: size, color: color),
  );
}