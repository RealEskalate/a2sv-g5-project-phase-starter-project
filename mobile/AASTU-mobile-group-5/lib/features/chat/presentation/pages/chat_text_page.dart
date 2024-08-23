import 'dart:io';

import 'package:flutter/material.dart';

import '../../../product/presentation/widgets/custom_app_bar_home_page.dart';
import '../widgets/bottom_nav_bar/bottom_nav_bar.dart';
import '../widgets/chat_app_bar.dart';

class TextPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: CustomAppBarTextPage(),
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