import 'package:flutter/material.dart';

import 'image_message.dart';
import 'text_message.dart';

class MessageCard extends StatelessWidget {
  const MessageCard({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(),
      body: Container(
        width: 285,
        child: Center(
          child: Row(
            verticalDirection: VerticalDirection.down,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Card(
                elevation: 10,
                color: const Color(0xFFF2F7FB).withOpacity(0),
                shadowColor: const Color(0xFFF2F7FB).withOpacity(0.5),
                child: const CircleAvatar(
                  foregroundImage: AssetImage(
                    'assets/image.png',
                  ),
                ),
              ),
              const SizedBox(
                width: 9,
              ),
              const Expanded(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    const Text(
                      'Annei Ellison',
                      style: TextStyle(
                        fontWeight: FontWeight.w500,
                        fontSize: 15,
                        color: Color.fromARGB(212, 0, 14, 8),
                      ),
                    ),
                    TextMessage(),
                    ImageMessage()
                  ],
                ),
              )
            ],
          ),
        ),
      ),
    );
  }
}
