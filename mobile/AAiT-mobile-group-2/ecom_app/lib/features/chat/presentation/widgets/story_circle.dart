import 'package:flutter/material.dart';

class StoryCircle extends StatelessWidget {
  final String name;
  final bool hasStory;
  final int numberOfStories;
  final String profileImageUrl;
  final double size; // Added size parameter for flexibility

  const StoryCircle({
    super.key,
    required this.name,
    required this.hasStory,
    required this.numberOfStories,
    required this.profileImageUrl,
    this.size = 50.0, // Default size if not provided
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        SizedBox(
          width: size,
          height: size,
          child: CustomPaint(
            painter: StoryCirclePainter(hasStory: hasStory, numberOfStories: numberOfStories),
            child: Container(
              padding: const EdgeInsets.all(10),
              decoration: BoxDecoration(
                shape: BoxShape.circle,
                image: DecorationImage(
                  image: NetworkImage(profileImageUrl),
                  fit: BoxFit.cover,
                ),
              ),
            ),
          ),
        ),
        const SizedBox(height: 5),
        Text(
          name,
          style: const TextStyle(
            fontSize: 12,
            fontWeight: FontWeight.bold,
            color: Colors.white
          ),
        ),
      ],
    );
  }
}

class StoryCirclePainter extends CustomPainter {
  final bool hasStory;
  final int numberOfStories;

  StoryCirclePainter({required this.hasStory, required this.numberOfStories});

  @override
  void paint(Canvas canvas, Size size) {
    if (hasStory) {
      final paint = Paint()
        ..color = Colors.blue
        ..style = PaintingStyle.stroke
        ..strokeWidth = 2;

      final center = Offset(size.width / 2, size.height / 2);
      final radius = size.width / 2 + 3;

      if (numberOfStories > 1) {
        const gapAngle = 10 * 3.141592653589793 / 180; // 5 degrees in radians
        final angleStep = (2 * 3.141592653589793 - gapAngle * numberOfStories) / numberOfStories;

        for (int i = 0; i < numberOfStories; i++) {
          final startAngle = i * (angleStep + gapAngle);
          final sweepAngle = angleStep;
          canvas.drawArc(
            Rect.fromCircle(center: center, radius: radius),
            startAngle,
            sweepAngle,
            false,
            paint,
          );
        }
      } else {
        canvas.drawCircle(center, radius, paint);
      }
    }
  }

  @override
  bool shouldRepaint(CustomPainter oldDelegate) {
    return false;
  }
}
