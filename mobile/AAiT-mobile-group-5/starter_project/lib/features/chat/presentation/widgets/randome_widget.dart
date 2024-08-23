import 'package:flutter/material.dart';

class RandomeWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(
        children: [
          Expanded(
            child: SingleChildScrollView(
              reverse: true, // Keeps the list scrolled to the bottom
              child: Column(
                children: [
                  // Replace with your chat messages
                  for (int i = 0; i < 20; i++)
                    ListTile(
                      title: Text('Message $i'),
                    ),
                ],
              ),
            ),
          ),
          Padding(
            padding: EdgeInsets.only(
              bottom: MediaQuery.of(context).viewInsets.bottom,
            ), // Adjusts padding when the keyboard is visible
            child: TextField(
              decoration: InputDecoration(
                hintText: 'Type your message',
                border: OutlineInputBorder(),
                contentPadding: EdgeInsets.all(8),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
