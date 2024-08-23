import 'package:flutter/material.dart';

Widget backButton({required Color iconColor, required void Function()? onTap}) {
  return GestureDetector(
    onTap: onTap,
    child: Container(
      margin: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
      padding: const EdgeInsets.only(left: 8),
      width: 40,
      height: 40,
      alignment: Alignment.center,
      decoration:
          const BoxDecoration(shape: BoxShape.circle, color: Colors.white),
      child: Icon(
        Icons.arrow_back_ios,
        color: iconColor,
        size: 20,
      ),
    ),
  );
}
