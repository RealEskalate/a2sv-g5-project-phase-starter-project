import 'package:flutter/material.dart';

mixin AppBars {
  AppBar normalAppBar(String title, VoidCallback onPressed) {
    return AppBar(
      backgroundColor: Colors.white,
      leading: IconButton(
        onPressed: onPressed,
        icon: const Icon(
          Icons.chevron_left,
        ),
      ),
      title: Text(
        title,
        style: const TextStyle(fontSize: 20),
        textDirection: TextDirection.rtl,
      ),
    );
  }
}
