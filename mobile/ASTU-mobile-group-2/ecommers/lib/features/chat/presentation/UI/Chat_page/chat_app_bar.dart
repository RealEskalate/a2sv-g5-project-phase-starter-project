import 'package:flutter/material.dart';

import 'currentUser.dart';

class ChatAppBar extends StatelessWidget {
  final String name;
  const ChatAppBar({
    super.key,
    required this.name
    });

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Colors.white,
        elevation: 0,
        leading: IconButton(
          icon: const Icon(Icons.arrow_back, color: Colors.black),
          onPressed: () {
            Navigator.pop(context);
          },
        ),
        title: Transform.translate(
          offset: const Offset(-15, 0),
          child:  CurrentUser(
            name: name,
            image: 'assets/image/avator.png',
            online: true,
          ),
        ),
        
        // Call and video call buttons as actions
        
      ),
      body: const SizedBox(),
    );
  }
}

