import 'package:dotted_border/dotted_border.dart';
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
          child: DottedBorder(
            color: borderColor,
            dashPattern: const [1, 0],
            strokeWidth: 1,
            borderType: BorderType.Circle,
            child: CircleAvatar(
              radius: radius,
              backgroundImage: AssetImage(avatarUrl),
            ),
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
