import 'package:flutter/material.dart';

class StoryCircle extends StatelessWidget {
  final String name;
  final String imageUrl;
  const StoryCircle(this.name, this.imageUrl, {super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Stack(
          children: [
            Container(
              padding: const EdgeInsets.all(2), // Outer padding
              decoration: const BoxDecoration(
                color: Colors.white, // Background color
                shape: BoxShape.circle,
              ),
              child: Container(
                padding: const EdgeInsets.all(2), // Middle padding
                decoration: const BoxDecoration(
                  color:
                      Color.fromRGBO(73, 140, 240, 1.0), // Middle circle color
                  shape: BoxShape.circle,
                ),
                child: Container(
                  padding: const EdgeInsets.all(2), // Inner padding
                  decoration: const BoxDecoration(
                    color: Colors.white, // Inner circle color
                    shape: BoxShape.circle,
                  ),
                  child: CircleAvatar(
                    radius: 24.5,
                    backgroundImage: NetworkImage(
                        imageUrl), // Image of the user who sent the message
                    backgroundColor: Colors.transparent,
                  ),
                ),
              ),
            ),
          ],
        ),
        const SizedBox(height: 5),
        Text(
          name,
          style: const TextStyle(color: Colors.white),
        ),
      ],
    );
  }
}
