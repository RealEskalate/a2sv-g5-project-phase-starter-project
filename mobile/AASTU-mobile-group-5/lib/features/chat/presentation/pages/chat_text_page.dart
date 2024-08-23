import 'dart:io';

import 'package:flutter/material.dart';

import '../widgets/bottom_nav_bar/bottom_nav_bar.dart';

class TextPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Text Page'),
      ),
      body: Column(
        children: [
          Expanded(
            child: ListView(
              children: [
                // Add any content you want here
              ],
            ),
          ),
          CustomBottomNavigationBar(),
        ],
      ),
    );
  }
}