import 'package:flutter/material.dart';

import '../widgets/chat_box.dart';
import '../widgets/profile_pic_widget.dart';

class ChatPage extends StatelessWidget {
  const ChatPage({super.key});

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Scaffold(
        appBar: PreferredSize(
          preferredSize: const Size.fromHeight(70),
          child: Padding(
            padding: const EdgeInsets.only(top: 8.0),
            child: AppBar(
              titleSpacing: 8,
              leading: IconButton(
                icon: const Icon(Icons.arrow_back),
                onPressed: () {
                  Navigator.of(context).pop();
                },
              ),
              title: const Row(
                mainAxisAlignment: MainAxisAlignment.start,
                crossAxisAlignment: CrossAxisAlignment.center,
                children: [
                  ProfilePicWidget(bgColor: Color.fromARGB(255, 228, 146, 119), radius: 30),
                  SizedBox(width: 5),
                  Expanded(
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        Text(
                          'Jhon Doe',
                          style: TextStyle(
                            fontSize: 18,
                            color: Color.fromRGBO(14, 8, 1, 1),
                          ),
                        ),
                        Text(
                          '8 members, 5 Online',
                          style: TextStyle(
                            color: Color.fromARGB(255, 121, 124, 123),
                            fontSize: 12,
                          ),
                        ),
                                  ],
                    ),
                  )
                ],
              ),
              actions: [
                IconButton(
                  icon: const Icon(Icons.call_outlined),
                  onPressed: () {},
                ),
                IconButton(
                  icon: const Icon(Icons.videocam_outlined),
                  onPressed: () {},
                ),
              ],
              ),
          ),
        ),
        body: const Stack(
          children: [
          Positioned(
            bottom: 0,
            left: 0,
            right: 0,
            child: MessageBox(),
            
          ),
          ]
        ),
      ),
    );
  }
}