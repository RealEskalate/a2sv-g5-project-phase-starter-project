import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';

import '../widgets/chat.dart';
import '../widgets/text_inputter.dart';
import '../widgets/user_avater.dart';

class ChatPage extends StatefulWidget {
  const ChatPage({super.key});

  @override
  State<ChatPage> createState() => _ChatPageState();
}

class _ChatPageState extends State<ChatPage> {
  bool isOnline = true;
  DateTime lastSeen = DateTime.now().subtract(const Duration(minutes: 5));
  bool isSelf = false;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        backgroundColor: Colors.white,
        toolbarHeight: 70,
        automaticallyImplyLeading: false,
        title: Row(
          children: [
            // IconButton(onPressed: onPressed, icon: icon)
            IconButton(
              icon: Icon(Icons.west),
            onPressed: () {
              Navigator.pop(context);
            },),
            SizedBox(width: 2),
            UserAvater(image: 'assets/images/avater.png'),
            const SizedBox(width: 8), // Add space between avatar and text
            Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Text(
                  'Sabila Sayma',
                  style: TextStyle(
                    color: Colors.black,
                    fontSize: 16,
                    fontFamily: 'Poppins'),
                ),
                Text(
                  isOnline
                      ? 'Online'
                      : 'Last seen at ${lastSeen.hour}:${lastSeen.minute}',
                  style: TextStyle( 
                    color: isOnline 
                  ? Colors.green 
                  : Colors.grey, 
                  fontSize: 10, 
                  fontFamily: 'Poppins', 
                  fontWeight: FontWeight.w400),
                ),
              ],
            ),
          ],
        ),
        actions: <Widget>[
          IconButton(
            icon: const Icon(Icons.call),
            onPressed: () {},
          ),
          IconButton(
            icon: const Icon(Icons.videocam_outlined),
            onPressed: () {},
          ),
        ],
      ),
      body: Column(
        children: [
          Expanded(
            child: SingleChildScrollView(
              child: Container(
                  // Chat content goes here

                  child :Column(children: const [
                  
                  ],)
                    // Chat messages go here
                  ),
            ),
          ),
          
          SafeArea(
              child: Container(
                constraints: BoxConstraints(
                maxHeight: 70, // Limit the height to avoid overflow
              ),
              decoration: BoxDecoration(
          color: Colors.white,
          boxShadow: [
            BoxShadow(
              color: Colors.grey.withOpacity(0.2),
              spreadRadius: 1,
              blurRadius: 7,
              offset: Offset(0, -3), // Offset upwards for a shadow at the top
            ),
          ],
        ),
              child: TextInputter(),  // Your TextInputter widget
            ),
          ),
          // TextInputter(),  // Your TextInputter widget
          // TextInputter(),
        ],
      ),
    );
  }
}
