import 'package:dotted_border/dotted_border.dart';
import 'package:flutter/material.dart';

class OwnStoryAvatar extends StatelessWidget {
  final String name;
  final String avatarUrl;
  OwnStoryAvatar({
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
              DottedBorder(
                color: borderColor,
                // divide the border into 4 parts
                dashPattern: [radius * 1.5, 3],

                strokeWidth: 1,
                borderType: BorderType.Circle,
                child: CircleAvatar(
                  radius: radius,
                  backgroundImage: AssetImage(avatarUrl),
                ),
              ),
              Positioned(
                bottom: 0,
                right: 0,
                child: Container(
                  width: radius * 0.6,
                  height: radius * 0.6,
                  decoration: const BoxDecoration(
                    shape: BoxShape.circle,
                    color: Colors.white,
                  ),
                  child:
                      Icon(Icons.add, size: radius * 0.6, color: borderColor),
                ),
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
