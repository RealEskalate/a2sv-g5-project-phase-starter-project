import 'package:custom_clippers/custom_clippers.dart';
import 'package:flutter/material.dart';

import '../widgets/messages.dart';

class ChatRoom extends StatelessWidget {
  const ChatRoom({super.key});

  @override
  Widget build(BuildContext context) {
    // Static list of chat messages data
    final List<Map<String, dynamic>> messagesData = [
      {
        'messageType': MessageType.receive,
        'message': "Hey there!",
        'timeStamp': "10:00 AM",
        'isImage': false,
      },
      {
        'messageType': MessageType.send,
        'message': "Hello! How are you?",
        'timeStamp': "10:02 AM",
        'isImage': false,
      },
      {
        'messageType': MessageType.receive,
        'message': "I'm good, thanks! How about you?",
        'timeStamp': "10:05 AM",
        'isImage': false,
      },
      {
        'messageType': MessageType.send,
        'message': "Doing well! Just working on a project.",
        'timeStamp': "10:07 AM",
        'isImage': false,
      },
      {
        'messageType': MessageType.receive,
        'message': "That sounds interesting! What's it about?",
        'timeStamp': "10:10 AM",
        'isImage': false,
      },
    ];

    return Scaffold(
      appBar: AppBar(
        title: const Text("Chat Room"),
      ),
      body: Column(
        children: [
          Expanded(
            child: Padding(
              padding: const EdgeInsets.only(top: 20),
              child: ListView.builder(
                shrinkWrap: true,
                itemCount: messagesData.length,
                itemBuilder: (context, index) {
                  final data = messagesData[index];
                  return ChatMessage(
                    messageType: data['messageType'],
                    message: data['message'],
                    timeStamp: data['timeStamp'],
                    isImage: data['isImage'],
                  );
                },
              ),
            ),
          ),
          Padding(
            padding: const EdgeInsets.all(8.0),
            child: Row(
              children: [
                const Expanded(
                  child: TextField(
                    decoration: InputDecoration(
                      hintText: 'Type a message',
                    ),
                  ),
                ),
                IconButton(
                  icon: const Icon(Icons.send),
                  onPressed: () {},
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
