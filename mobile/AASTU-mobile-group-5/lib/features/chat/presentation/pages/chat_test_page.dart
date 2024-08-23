import 'package:flutter/material.dart';

import '../widgets/profileWidget.dart';

class TextPage extends StatelessWidget {
  const TextPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leading: Padding(
          padding: EdgeInsets.only(left: 15),
          
          ),
        title: const Row(
          children: [
            CircleAvatar(
              backgroundImage: AssetImage('assets/profile.jpg'), // Example image
            ),
            SizedBox(width: 10),
            Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                ProfileWidget(),
                Text('Sabila Sayma', style: TextStyle(fontSize: 18)),
                Text('8 members, 5 online', style: TextStyle(fontSize: 12)),
              ],
            ),
          ],
        ),
        actions: [
          IconButton(
            icon: const Icon(Icons.call),
            onPressed: () {},
          ),
          IconButton(
            icon: const Icon(Icons.video_call),
            onPressed: () {},
          ),
        ],
      ),
      body: Column(
        children: [
          Expanded(
            child: ListView(
              padding: const EdgeInsets.all(8),
              children: const [
                // Incoming Message
                Text('Annei Ellison'),
                Text('Annei Ellison'),
                Text('You did your job well!'),
        
              ],
            ),
          ),
          
        ],
      ),
    );
  }

}