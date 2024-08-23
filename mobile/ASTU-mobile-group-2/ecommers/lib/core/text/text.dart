

import 'package:flutter/widgets.dart';

class ConStTexts extends StatelessWidget {
  final String text;

  final double fontSize;
  final Color color;
  final FontWeight fontWeight;
  final String fontFamily;

  const ConStTexts({
    super.key,
    required this.text,
    required this.fontSize,
    required this.color,
    required this.fontWeight,
    this.fontFamily = 'Popins',


    });

  @override
  Widget build(BuildContext context) {
    return Text(
      text,
      overflow: TextOverflow.ellipsis,
      style: TextStyle(
        fontSize: fontSize,
        color: color,
        fontWeight: fontWeight,
        fontFamily: fontFamily,
        height: 0.8
       
      ),
    );
  }
}