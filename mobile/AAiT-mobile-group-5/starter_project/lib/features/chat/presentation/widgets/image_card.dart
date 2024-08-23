import 'package:flutter/material.dart';

class ImageCard extends StatelessWidget {
  final String imageUrl;
  final double borderRadius;

  const ImageCard({
    super.key,
    required this.imageUrl,
    this.borderRadius = 12, // Default to a circular shape
  });

  @override
  Widget build(BuildContext context) {
    // Get the screen width
    final screenWidth = MediaQuery.of(context).size.width;

    // Calculate the width of the image
    final imageWidth = screenWidth * 0.5;

    return ClipRRect(
      borderRadius: BorderRadius.circular(borderRadius),
      child: Image.asset(
        imageUrl, // Assuming you're using AssetImage
        width: imageWidth,
        height: 200, // To make it square and round
        fit: BoxFit.cover,
      ),
    );
  }
}
