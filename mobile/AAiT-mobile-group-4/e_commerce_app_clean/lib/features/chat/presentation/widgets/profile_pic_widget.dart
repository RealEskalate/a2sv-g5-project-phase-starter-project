import 'package:dotted_border/dotted_border.dart';
import 'package:flutter/material.dart';

class ProfilePicWidget extends StatelessWidget {
  final String imagePath;
  final bool isMystory;
  final bool isStory;
  final Color bgColor;

  final double radius;

  const ProfilePicWidget(
      {this.isMystory = false,
      this.isStory = false,
      super.key,
      this.imagePath = 'assets/story_3.png',
      required this.bgColor,
      required this.radius});

  @override
  Widget build(BuildContext context) {
    if (isMystory && isStory) {
      return Center(
        child: Stack(
          children: [
            DottedBorder(
              borderType: BorderType.Circle,
              color: Colors.white,
              strokeWidth: 2,
              dashPattern: const [52, 3],
              child: Container(
                decoration: BoxDecoration(
                  shape: BoxShape.circle,
                  border: Border.all(
                    color: Colors.blue,
                    width: 2,
                  ),
                ),
                child: CircleAvatar(
                  radius: radius,
                  backgroundColor: bgColor,
                  backgroundImage: AssetImage(imagePath),
                ),
              ),
            ),
            const Positioned(
              bottom: 2,
              right: 2,
              child: CircleAvatar(
                radius: 7,
                backgroundColor: Colors.grey,
                child: CircleAvatar(
                  radius: 6,
                  backgroundColor: Colors.white,
                  child: Icon(Icons.add, color: Colors.grey, size: 10),
                ),
              ),
            ),
          ],
        ),
      );
    } else if (isStory) {
      return Container(
        decoration: BoxDecoration(
          shape: BoxShape.circle,
          border: Border.all(
            color: bgColor,
            width: 2,
          ),
        ),
        child: Container(
          decoration: BoxDecoration(
            shape: BoxShape.circle,
            border: Border.all(
              color: Colors.blue,
              width: 2,
            ),
          ),
          child: CircleAvatar(
            radius: radius,
            backgroundColor: bgColor,
            backgroundImage: AssetImage(imagePath),
          ),
        ),
      );
    } else {
      return CircleAvatar(
        radius: radius,
        backgroundColor: bgColor,
        backgroundImage: AssetImage(imagePath),
      );
    }
  }
}
