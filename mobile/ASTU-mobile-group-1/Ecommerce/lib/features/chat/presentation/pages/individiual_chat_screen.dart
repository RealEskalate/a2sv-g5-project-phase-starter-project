import 'package:flutter/material.dart';

import '../widget/recieved_messages.dart';
import '../widget/sent_message.dart';
import '../widget/user_profile.dart';

class IndividiualChatScreen extends StatelessWidget {
  const IndividiualChatScreen({super.key});

  @override
  Widget build(BuildContext context) {
    final dummyMessage = [
      showRecievedMessage('This is my new 3d design'),
      showSentMessage('You did your job well!'),
      showRecievedMessage('This is my new 3d design'),
      showSentMessage('You did your job well!'),
      showRecievedMessage('This is my new 3d design'),
      showSentMessage('You did your job well!'),
      showRecievedMessage('This is my new 3d design'),
      showSentMessage('You did your job well!'),
      showRecievedMessage('This is my new 3d design'),
      showSentMessage('You did your job well!'),
      
    ];

    return Scaffold(
      
      resizeToAvoidBottomInset: true,
         bottomNavigationBar:  const Padding(
        padding: EdgeInsets.only(right: 8.0, bottom: 25.0),
        child: BottomBarElement(),
      ),
      
          appBar: AppBar(
            centerTitle: false,
            
             titleSpacing: 0,
            leading: IconButton(onPressed: (){}, icon: const Icon(Icons.arrow_back)),
            title: ListTile(
              contentPadding: const EdgeInsets.all(0),
              title: const Text('Sabila Sayima'),
              subtitle: const Text('online'),
              leading: showUser(),
            ),
            actions: [
              IconButton(onPressed: (){}, icon: const Icon(Icons.phone)),
            IconButton(onPressed: (){}, icon: const Icon(Icons.video_call))

            ],

      ),
      body: Padding(
        padding: const EdgeInsets.all(12.0),
        child: ListView.builder(
          itemCount: 10,
          shrinkWrap: true,
          itemBuilder: (context, index) => dummyMessage[index]),
      ),
    );
  }
}

class BottomBarElement extends StatelessWidget {
  const BottomBarElement({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        IconButton(
          onPressed: () {},
          icon: const Icon(Icons.attachment_sharp, color: Colors.blue),
        ),
        Expanded(
          child: TextField(
            decoration: InputDecoration(
              hintText: 'Type a message...',
              contentPadding: const EdgeInsets.symmetric(horizontal: 10),
              border: OutlineInputBorder(
                borderRadius: BorderRadius.circular(25),
                borderSide: const BorderSide(color: Colors.blue),
              ),
            ),
          ),
        ),
       
         IconButton(
          onPressed: () {},
          icon: const Icon(Icons.camera_alt, color: Colors.blue),
        ),
        IconButton(
          onPressed: () {},
          icon: const Icon(Icons.mic, color: Colors.blue),
        ),
      ],
    );
  }
}