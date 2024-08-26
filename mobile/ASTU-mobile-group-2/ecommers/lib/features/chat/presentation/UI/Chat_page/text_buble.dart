import 'package:flutter/material.dart';

class TextBubble extends StatelessWidget {
  final String text;
  final bool isMe;
  const TextBubble({super.key, required this.text, required this.isMe});

  @override
  Widget build(BuildContext context) {
    final width = MediaQuery.of(context).size.width;
    return ConstrainedBox(
      constraints: const BoxConstraints(maxWidth: 300),
      child: Container(
        width: width * 0.5,
        decoration: BoxDecoration(
            color: isMe
                ? const Color.fromRGBO(63, 81, 243, 1)
                : const Color.fromRGBO(242, 247, 251, 1),
            borderRadius: BorderRadius.only(
                topLeft: isMe ? const Radius.circular(30) : Radius.zero,
                topRight: isMe ? Radius.zero : const Radius.circular(30),
                bottomLeft: const Radius.circular(30),
                bottomRight: const Radius.circular(30))),
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
