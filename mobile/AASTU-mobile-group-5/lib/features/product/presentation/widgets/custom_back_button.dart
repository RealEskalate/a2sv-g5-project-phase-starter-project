import 'package:flutter/material.dart';

class CustomBackButton extends StatelessWidget {
  final Color backgroundColor;
  final Color iconColor;
  final double size;
  final double iconSize;

  const CustomBackButton({
    super.key,
    this.backgroundColor = Colors.white,
    this.iconColor = const Color.fromARGB(255, 54, 104, 255),
    this.size = 35.0,
    this.iconSize = 15.0,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: SizedBox(
        height: size,
        width: size,
        child: Center(
          child: FloatingActionButton(
            onPressed: () {
              Navigator.pop(context);
            },
            shape: const CircleBorder(),
            backgroundColor: backgroundColor,
            child: Icon(
              Icons.arrow_back_ios_new,
              size: iconSize,
              color: iconColor,
            ),
          ),
        ),
      ),
    );
  }
}
