import 'package:flutter/material.dart';

class CustomImageContainer extends StatelessWidget {
  final double? height;
  final double? width;
  final BorderRadiusGeometry borderRadius;
  final String? imagePath;

  const CustomImageContainer({
    super.key,
    this.height,
    this.width,
    this.borderRadius = const BorderRadius.all(Radius.circular(8.0)),
    this.imagePath,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      height: height,
      width: width,
      decoration: BoxDecoration(
        image: imagePath != null
            ? DecorationImage(
                fit: BoxFit.cover,
                image: Image.asset(
                  imagePath!,
                ).image)
            : null,
        borderRadius: borderRadius,
      ),
    );
  }
}
