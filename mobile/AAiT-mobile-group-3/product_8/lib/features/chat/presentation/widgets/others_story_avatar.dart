import 'package:flutter/material.dart';

class OthersStoryAvatar extends StatelessWidget {
  final String name;
  final String avatarUrl;
  OthersStoryAvatar({
    super.key,
    required this.name,
    required this.avatarUrl,
  });
  final double radius = 22.0;
  final Color borderColor = Colors.blue.shade100;
  final Color backgroundColor = Colors.blue;

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Container(
          margin: const EdgeInsets.fromLTRB(10, 10, 20, 0),
          padding: const EdgeInsets.all(4.0),
          child: Stack(
            alignment: Alignment.center,
            children: [
              // Outer border
              Container(
                width:
                    radius * 2.3, // Adjust this value for the outer border size
                height: radius * 2.3,
                decoration: BoxDecoration(
                  shape: BoxShape.circle,
                  color: Colors.transparent,
                  border: Border.all(
                    color: borderColor, // First border color
                    width: 4.0, // Thickness of the first border
                  ),
                ),
              ),
              // Inner border
              Container(
                width: radius *
                    2.15, // Adjust this value for the inner border size
                height: radius * 2.15,
                decoration: BoxDecoration(
                  shape: BoxShape.circle,
                  color: Colors.transparent,
                  border: Border.all(
                    color: backgroundColor, // Second border color
                    width: 4.0, // Thickness of the second border
                  ),
                ),
              ),
              // CircleAvatar
              CircleAvatar(
                radius: radius,
                backgroundImage: AssetImage(avatarUrl),
              ),
            ],
          ),
        ),
        Text(name,
            style: const TextStyle(
              color: Colors.white,
            )),
      ],
    );
  }
}
