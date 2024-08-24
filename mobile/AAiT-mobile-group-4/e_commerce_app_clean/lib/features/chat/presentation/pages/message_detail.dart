import 'package:flutter/material.dart';
import '../widgets/message_contaner.dart';

class MessageDetail extends StatelessWidget {
  const MessageDetail({super.key});

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Scaffold(
        backgroundColor: Colors.white,
        appBar: AppBar(
          title: const Text('something to write'),
        ),
        body: Padding(
          padding: const  EdgeInsets.fromLTRB(10, 20, 10, 10),
          child: ListView.builder(
            itemCount: 20,
            itemBuilder: (context, index) {
              return withTime(
                text: 'Have a great working week!!', 
                isCurrentUser: index % 2 != 0 ? true : false, 
                type: index % 3 != 0 ?'image': 'audio',
                image: 'https://images.unsplash.com/photo-1557863618-9643198cb07b?w=600&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Njd8fFRveW90YSUyMFY4JTIwcGF0cm9sfGVufDB8fDB8fHww',
                time: '09:25 AM',
              );
            },
          ),
        ),
      ),
    );
  }
}