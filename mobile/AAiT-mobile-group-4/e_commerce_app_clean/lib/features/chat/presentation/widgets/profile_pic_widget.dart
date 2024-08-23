import 'package:flutter/material.dart';

class ProfilePicWidget extends StatelessWidget {
  final String imagePath;
  
  final Color bgColor;

  final double radius;

  const ProfilePicWidget({super.key, this.imagePath = 'assets/story_3.png',  required this.bgColor,required this.radius});

  @override
  Widget build(BuildContext context) {
    return CircleAvatar(
        radius: radius,
        backgroundColor: bgColor,
        backgroundImage: AssetImage(imagePath),
      );
  }
}