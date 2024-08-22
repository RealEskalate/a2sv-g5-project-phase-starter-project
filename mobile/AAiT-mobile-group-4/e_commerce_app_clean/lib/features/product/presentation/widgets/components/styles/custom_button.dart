import 'package:flutter/material.dart';

class CustomButton extends StatelessWidget {
  final double width;
  final double height;
  final Color fgcolor;
  final Color bgcolor;
  final String name;
  final void Function()? pressed;
  final Color textBgColor;
  const CustomButton({
    super.key,
    required this.width,
    required this.height,
    required this.name,
    required this.fgcolor,
    required this.bgcolor,
    this.pressed,
    this.textBgColor = Colors.black,
  });

  @override
  Widget build(BuildContext context) {
    return OutlinedButton(
      onPressed: pressed,
      style: OutlinedButton.styleFrom(
          fixedSize: Size(width, height),
          foregroundColor: fgcolor, //const Color.fromARGB(230, 255, 19, 19),
          backgroundColor: bgcolor,
          side: BorderSide(
            color: fgcolor, //Color.fromARGB(230, 255, 19, 19),
            width: 1.0,
          ),
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(8),
          )),
      child: Center(
        child: Text(
          name,
          style: TextStyle(
            color: textBgColor,
          ),
        ),
      ),
    );
  }
}
