import 'package:flutter/material.dart';

class TextBubble extends StatelessWidget {
  final String text;
  final bool isMe;
  const TextBubble({super.key, required this.text, required this.isMe});

  @override
  Widget build(BuildContext context) {
    return ConstrainedBox(
      constraints: BoxConstraints(maxWidth: 300),
      child: Container(
        decoration: BoxDecoration(
            color: isMe
                ? Color.fromRGBO(63, 81, 243, 1)
                : Color.fromRGBO(242, 247, 251, 1),
            borderRadius: BorderRadius.only(
                topLeft: isMe ? Radius.circular(30) : Radius.zero,
                topRight: isMe ? Radius.zero : Radius.circular(30),
                bottomLeft: Radius.circular(30),
                bottomRight: Radius.circular(30))),
        child: Padding(
          padding: const EdgeInsets.all(15.0),
          child: Text(
            text,
            style: TextStyle(
                color: isMe ? Colors.white : Colors.black,
                fontSize: 15,
                fontWeight: FontWeight.w500),
          ),
        ),
      ),
    );
  }
}
