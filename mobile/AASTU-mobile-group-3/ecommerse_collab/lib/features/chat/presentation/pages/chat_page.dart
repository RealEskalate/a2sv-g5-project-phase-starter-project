import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';

import '../widgets/text_inputter.dart';
import '../widgets/user_avater.dart';

class ChatPage extends StatefulWidget {
  const ChatPage({super.key});

  @override
  State<ChatPage> createState() => _ChatPageState();
}

class _ChatPageState extends State<ChatPage> {
  bool isOnline = false;
  DateTime lastSeen = DateTime.now().subtract(const Duration(minutes: 5));

  bool isSelf = false;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        backgroundColor: Colors.white,
        toolbarHeight: 100,
        leading: IconButton(
          icon: const Icon(Icons.arrow_back),
          onPressed: () {
            Navigator.pop(context);
          },
          
        ),
        
        actions: <Widget>[
          UserAvater(image: 'assets/images/avater.png'),
          Column(
            // crossAxisAlignment: CrossAxisAlignment.start,
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
            Text('Sabila Sayma', style: TextStyle(color: Colors.black, fontSize: 16),),
            Text (
              isOnline 
             ? 'Online' 
             : 'Last seen at ${lastSeen.hour} : ${lastSeen.minute}', 
             style: TextStyle(
              color: Colors.green,
              fontSize: 10),),
            ]
          ),
          Spacer(),
          IconButton(
            icon: const Icon(Icons.phone_outlined),
            onPressed: () {},
          ),
          IconButton(
            icon: const Icon(Icons.videocam_outlined),
            onPressed: () {},
          ),
        ],),

      body: Column(
        // mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Expanded(
            child: SingleChildScrollView(
              child: Container(

              ),
            ),
          ),
        Positioned(
          bottom: 0,
          child: TextInputter(),
        ),
        ],
      ),
    );
  }
}