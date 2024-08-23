

import 'package:flutter/material.dart';


class Bordrs extends StatelessWidget {
  final int hight;
  final int width;
  final Color color;
 
  const Bordrs({
    super.key,
    required this.hight,
    required this.width,
    required this.color
  });

  @override
  Widget build(BuildContext context) {
    return Container(
              width: width.toDouble(),
              height: hight.toDouble(),
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(8),
                border: Border.all(
                  color: Colors.grey
                ),
                color: Colors.white
              ),
              
              
            );
  }
}