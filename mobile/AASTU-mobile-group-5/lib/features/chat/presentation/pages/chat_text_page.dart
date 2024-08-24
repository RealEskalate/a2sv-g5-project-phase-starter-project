import 'dart:io';

import 'package:flutter/material.dart';

import '../../../product/presentation/widgets/custom_app_bar_home_page.dart';
import '../../data/datasources/websocket_service.dart';
import '../widgets/bottom_nav_bar/bottom_nav_bar.dart';
import '../widgets/chat_app_bar.dart';

class TextPage extends StatelessWidget {

  const TextPage({super.key,});
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: CustomAppBarTextPage(),
      body: Column(
        children: [
          Expanded(
            child: StreamBuilder(
              builder: (context, snapshot) {
                if (snapshot.hasData) {
                  // Handle the incoming message
                  final message = snapshot.data;
                  // Update the UI with the new message
                  return Text(message.toString());
                } else if (snapshot.hasError) {
                  return Text('Error: ${snapshot.error}');
                }
                return CircularProgressIndicator();
              },
            ),
          ),
          CustomBottomNavigationBar(),
        ],
      ),
    );
  }
}