
import 'package:flutter/material.dart';

class Containers extends StatelessWidget {
  final Widget child;
  final double width;
  final double height;
  final Color color;
  final Alignment alignment;
  final EdgeInsets padding;
  final EdgeInsets margin;
  final double borderRadius;
  final Color topColor;
  final Color bottomColor;
  final Color leftColor;
  final Color rightColor;
  final BoxShadow shadows;
  const Containers({
    super.key,
    required this.child,
    required this.width,
    required this.height,
    this.color = Colors.white,
    this.alignment = Alignment.center,
    this.padding = const EdgeInsets.all(0),
    this.margin = const EdgeInsets.all(0),
    this.borderRadius = 0,
    this.topColor = Colors.white,
    this.bottomColor = Colors.white,
    this.leftColor = Colors.white,
    this.rightColor = Colors.white,
    this.shadows = const BoxShadow(
      color: Colors.white,
      blurRadius: 0,
      spreadRadius: 0,
      offset: Offset(0, 0),
    ),

    });

  @override
  Widget build(BuildContext context) {
    return Container(
      width: width,
      height: height,
      
      alignment: alignment,
      padding: padding,
      margin: margin,

      decoration: BoxDecoration(
        boxShadow: [
          shadows,
        ],
        color: color,
        borderRadius: BorderRadius.circular(borderRadius),
        border:  Border(
          top: BorderSide(
            color: topColor,
            width: 1,
          ),
          bottom: BorderSide(
            color: bottomColor,
            width: 1,
          ),
          left: BorderSide(
            color: leftColor,
            width: 1,
          ),
          right: BorderSide(
            color: rightColor,
            width: 1,
          ),
        ),
        
      ),
      child: child,
    );
  }
}