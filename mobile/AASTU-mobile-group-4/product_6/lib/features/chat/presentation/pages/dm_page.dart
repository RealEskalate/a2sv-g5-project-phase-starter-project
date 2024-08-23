import 'package:flutter/material.dart';

import '../widget/message.dart';

class DmPage extends StatelessWidget {
   DmPage({super.key});

  final TextEditingController _Messagecontroller = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leading: IconButton(
          icon: const Icon(Icons.arrow_back),
          onPressed: () {
            //
          },
        ),
        title: Row(
          children: [
            Container(
              width: 40,
              height: 40,
              decoration: const BoxDecoration(
                shape: BoxShape.circle,
                color: Colors.blue,
              ),
              child: Center(
                child: Image.asset('images/profile.png'),
              ),
            ),
            const SizedBox(width: 10),
            const Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  'Sabila Sayma',
                  style: TextStyle(
                    color: Colors.black,
                    fontWeight: FontWeight.bold,
                    fontSize: 16,
                  ),
                ),
                Text(
                  'Online',
                  style: TextStyle(
                    color: Colors.black,
                    fontSize: 12,
                  ),
                ),
              ],
            ),
          ],
        ),
        actions: [
          IconButton(
            icon: const Icon(Icons.phone_outlined),
            onPressed: () {
              // Add your phone call logic here
            },
          ),
          IconButton(
            icon: const Icon(Icons.videocam_outlined),
            onPressed: () {
              // Add your video call logic here
            },
          ),
        ],
      ),
      body: ListView.builder(
        itemCount: 15,
        itemBuilder: (context, index) {
          return MessageBubble();
        },
      ),
      bottomSheet: Container(
        padding: const EdgeInsets.all(10),
        color: Colors.white,
        child: Row(
          children: [
            IconButton(
              onPressed: () {},
              icon: Icon(Icons.attach_file),
            ),
            Flexible(
              child: TextField(
                controller: _Messagecontroller,
                decoration:  InputDecoration(
                  fillColor: Color.fromRGBO(243, 246, 246, 1),
                  hintText: 'Type a message',
                  border: InputBorder.none,
                  filled: true,
                  suffixIcon: IconButton(
                    onPressed: () {
                      print('hello');
                    },
                    icon: Icon(Icons.send),
                  ),
                ),
              ),
            ),
            IconButton(
              onPressed: () {},
              icon: Icon(Icons.camera_alt),
            ),
            IconButton(
              onPressed: () {},
              icon: Icon(Icons.mic),
            ),
          ],
        ),
      ),
    );
  }
}
