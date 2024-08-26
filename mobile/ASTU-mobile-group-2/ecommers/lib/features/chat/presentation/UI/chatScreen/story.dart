
  import 'package:flutter/material.dart';

Widget story(String name, String imagePath) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 8.0),
      child: Column(
        children: [
          CircleAvatar(
            radius: 30,
            backgroundImage: AssetImage(imagePath),
          ),
          const SizedBox(height: 8),
          Text(name, style: const TextStyle(color: Colors.white)),
        ],
      ),
    );
  }


  
Widget Mystory(String name, String imagePath) {
  return Padding(
    padding: const EdgeInsets.symmetric(horizontal: 8.0),
    child: Column(
      children: [
        Stack(
          clipBehavior: Clip.none, 
          children: [
            CircleAvatar(
              radius: 30,
              backgroundImage: AssetImage(imagePath),
            ),
            Positioned(
              bottom: 0,
              right: 2,
              child: Container(
                padding: EdgeInsets.all(4),
                decoration: const BoxDecoration(
                  shape: BoxShape.circle,
                  color: Colors.white, 
                ),
                child: const Icon(
                  Icons.add,
                  size: 10,
                  color: Colors.blue,
                ),
              ),
            ),
          ],
        ),
        SizedBox(height: 8),
        Text(name, style: TextStyle(color: Colors.white)),
      ],
    ),
  );
}