import 'package:flutter/material.dart';

class Logo extends StatelessWidget {
  final double? width;
  final double? height;
  final double? fontSize;
  
  const Logo({super.key, this.width, this.height, this.fontSize});

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
        color: Colors.white, // Background color
        border: Border.all(
            color: const Color(0xFF3F51F3), width: 1.0), // Border
        borderRadius: BorderRadius.circular(12.0),
        boxShadow: [
          BoxShadow(
            color: Color(0xFF3F51F3).withOpacity(0.3),
            blurRadius: 4,
            offset: const Offset(0, 2),
          ),
        ],
      ),
      padding: const EdgeInsets.symmetric(vertical : 10.0, horizontal: 15.0 ),
      child: Image.asset(
        'assets/images/Ecom.png',
        width: width,   // Set the width of the image
        height: height, // Set the height of the image
        fit: BoxFit.contain, // Ensure the image fits within the given dimensions
      ),
    );
  }
}
