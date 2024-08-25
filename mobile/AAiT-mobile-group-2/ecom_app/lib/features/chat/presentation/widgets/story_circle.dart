import 'package:flutter/material.dart';

class StoryCircle extends StatelessWidget {
  final bool hasStory;
  final int numberOfStories;
  final String profileImageUrl;
  final double size; // Added size parameter for flexibility

  const StoryCircle({
    super.key,
    required this.hasStory,
    required this.numberOfStories,
    required this.profileImageUrl,
    this.size = 50.0, // Default size if not provided
  });

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: size,
      height: size,
      child: CustomPaint(
        painter: StoryCirclePainter(hasStory: hasStory, numberOfStories: numberOfStories),
        child: Container(
          decoration: BoxDecoration(
            shape: BoxShape.circle,
            image: DecorationImage(
              image: NetworkImage(profileImageUrl),
              fit: BoxFit.cover,
            ),
          ),
        ),
      ),
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
        ..strokeWidth = 4.0;

      final center = Offset(size.width / 2, size.height / 2);
      final radius = size.width / 2;

      if (numberOfStories > 1) {
        const gapAngle = 5 * 3.141592653589793 / 180; // 5 degrees in radians
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
