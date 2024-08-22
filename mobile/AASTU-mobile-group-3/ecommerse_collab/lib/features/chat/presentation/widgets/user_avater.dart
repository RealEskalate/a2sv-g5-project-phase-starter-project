import 'package:flutter/material.dart';

class UserAvater extends StatelessWidget {
  final String image;
  const UserAvater({super.key, required this.image});

  @override
  Widget build(BuildContext context) {
    return Container(
      width: 100,
      height: 100,
      decoration: BoxDecoration(
        color: Colors.amber[100],
        borderRadius: BorderRadius.circular(100),
      ),
      child: Image.asset(image),
    );
  }
}