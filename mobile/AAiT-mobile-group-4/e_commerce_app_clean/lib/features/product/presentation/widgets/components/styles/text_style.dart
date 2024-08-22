import 'package:flutter/material.dart';

class CustomTextStyle extends StatelessWidget {
  final String name;
  final FontWeight weight;
  final double size;
  final Color color;
  final String family;
  const CustomTextStyle({
    super.key,
    required this.name,
    required this.weight,
    required this.size,
    this.color = Colors.black,
    this.family = 'Poppins',
  });

  @override
  Widget build(BuildContext context) {
    return Text(
      name,
      style: TextStyle(
        fontFamily: family,
        fontWeight: weight,
        fontSize: size,
        color: color,
      ),
    );
  }
}
