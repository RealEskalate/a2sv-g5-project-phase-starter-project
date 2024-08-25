import 'package:custom_clippers/custom_clippers.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'dart:math' as math;

import '../widgets/messages.dart';
import '../widgets/reusable_textfield.dart';

class ChatRoom extends StatelessWidget {
  const ChatRoom({super.key});

  @override
  Widget build(BuildContext context) {
    TextEditingController messageController = TextEditingController();
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
        automaticallyImplyLeading: false,
        backgroundColor: Colors.white,
        title: Row(
          children: [
            const Icon(CupertinoIcons.chevron_back),
            Container(
              height: 50,
              width: 50,
              decoration: const BoxDecoration(
                color: Color(0xffC7FEE0),
                shape: BoxShape.circle,
              ),
              child: ClipRRect(
                borderRadius: BorderRadius.circular(50),
                child: Image.asset("assets/images/Memoji Boys 6-18.png"),
              ),
            ),
            const Padding(
              padding: EdgeInsets.only(left: 18.0),
              child: Column(
                children: [
                  Text(
                    "Sabila Sayma",
                    style: TextStyle(fontWeight: FontWeight.w500, fontSize: 18),
                  ),
                  Text(
                    "8 members, 5 online",
                    style: TextStyle(
                        fontWeight: FontWeight.w400,
                        fontSize: 12,
                        color: Colors.grey),
                  ),
                ],
              ),
            ),
            const Spacer(),
            const Padding(
              padding: EdgeInsets.only(right: 8.0),
              child: Icon(
                CupertinoIcons.phone,
                size: 24,
              ),
            ),
            const Icon(
              CupertinoIcons.video_camera,
              size: 32,
            ),
          ],
        ),
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
                  final isReceive = data['messageType'] == MessageType.receive;
                  return Row(
                    mainAxisAlignment: isReceive
                        ? MainAxisAlignment.start
                        : MainAxisAlignment.end,
                    crossAxisAlignment: CrossAxisAlignment.center,
                    children: [
                      if (isReceive)
                        const Padding(
                          padding: EdgeInsets.only(left: 10),
                          child: CircleAvatar(
                            backgroundImage: AssetImage(
                              "assets/images/Memoji Boys 6-18.png",
                            ),
                            radius: 25,
                          ),
                        ),
                      Flexible(
                        child: Padding(
                          padding: EdgeInsets.only(
                              left: isReceive ? 10 : 0,
                              right: isReceive ? 0 : 10,
                              top: 18),
                          child: ChatMessage(
                            messageType: data['messageType'],
                            message: data['message'],
                            timeStamp: data['timeStamp'],
                            isImage: data['isImage'],
                          ),
                        ),
                      ),
                      if (!isReceive)
                        const Padding(
                          padding: EdgeInsets.only(right: 10),
                          child: CircleAvatar(
                            backgroundImage: AssetImage(
                              "assets/images/Memoji Boys 6-18.png",
                            ),
                            radius: 25,
                          ),
                        ),
                    ],
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
      bottomSheet: Container(
        decoration: const BoxDecoration(color: Colors.white),
        height: 78,
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 8.0),
          child: Row(
            children: [
              Padding(
                padding: const EdgeInsets.only(right: 8.0),
                child: Transform.rotate(
                    angle: -math.pi / 5.7,
                    child: const Icon(Icons.attach_file_outlined)),
              ),
              Expanded(
                child: ReusableTextField(
                  hint: "Write your message",
                  textEditingController: messageController,
                  textInputType: TextInputType.text,
                ),
              ),
              const Padding(
                padding: EdgeInsets.only(left: 8.0),
                child: Icon(CupertinoIcons.camera),
              ),
              const Padding(
                padding: EdgeInsets.symmetric(horizontal: 8.0),
                child: Icon(Icons.mic_none),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
