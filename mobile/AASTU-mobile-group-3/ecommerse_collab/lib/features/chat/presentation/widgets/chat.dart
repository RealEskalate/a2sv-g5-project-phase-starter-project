import 'dart:ui';
import 'package:flutter/material.dart';

class Chat extends StatelessWidget {
  const Chat({super.key});

  @override
  Widget build(BuildContext context) {
    return ClipRRect(
      borderRadius: const BorderRadius.only(
        topLeft: Radius.circular(20),
        topRight: Radius.circular(20),
        bottomLeft: Radius.circular(20)
      ),
  
      child: SizedBox(
        height: 20,
        width: 100,
        child: Container(
          decoration:const BoxDecoration(
            color:  Color(0xFF3E50F3),
          ),
        ),
      ),
    );
  }
}