import 'package:flutter/material.dart';

class BigTitle extends StatelessWidget {
  final String text;
  const BigTitle({super.key, required this.text});

  @override
  Widget build(BuildContext context) {
    return Align(
      alignment: Alignment.center,
      child: Text(text, 
        style: const TextStyle(
          height: 5,
          fontFamily: 'Poppins',
          fontWeight: FontWeight.w800,
          fontSize: 27,
                  
        ),
      ),
    );
  }
}