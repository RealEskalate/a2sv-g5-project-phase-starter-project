import 'package:flutter/material.dart';

import 'text_message.dart';


class MessageCard extends StatelessWidget {
  final bool isLeft;
  MessageCard({super.key, required this.isLeft});

  Widget avator = Card(
    elevation: 10,
    color: const Color(0xFFF2F7FB).withOpacity(0),
    shadowColor: const Color(0xFFF2F7FB).withOpacity(0.5),
    child: const CircleAvatar(
      foregroundImage: AssetImage(
        'assets/img1.jpg',
      ),
    ),
  );

  Widget message(bool isLeft) {
    return Expanded(
      child: Column(
        crossAxisAlignment:
            isLeft ? CrossAxisAlignment.start : CrossAxisAlignment.end,
        children: [
          Text(
            isLeft ? 'Annei Ellison' : 'You',
            style: const TextStyle(
              fontWeight: FontWeight.w500,
              fontSize: 15,
              color: Color.fromARGB(212, 0, 14, 8),
            ),
          ),
          TextMessage(
            isLeft: isLeft,
          ),
        ],
      ),
    );
  }

  late Widget right;

  late Widget left;

  void intiate() {
    if (isLeft == true) {
      left = avator;
      right = message(isLeft);
    } else {
      left = message(isLeft);
      right = avator;
    }
  }

  @override
  Widget build(BuildContext context) {
    intiate();
    return Scaffold(
      appBar: AppBar(),
      body: Row(
        mainAxisAlignment:
            isLeft ? MainAxisAlignment.start : MainAxisAlignment.end,
        children: [
          SizedBox(
            width: 300,
            child: Row(
              crossAxisAlignment: CrossAxisAlignment.start,
              mainAxisAlignment:
                  isLeft
                  ? MainAxisAlignment.start
                  : MainAxisAlignment.end,
              children: [
                left,
                const SizedBox(
                  width: 9,
                ),
                right,
              ],
            ),
          ),
        ],
      ),
    );
  }
}
