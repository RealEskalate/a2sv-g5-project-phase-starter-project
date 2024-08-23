import 'package:flutter/material.dart';

class ChatListHolder extends StatefulWidget {
  final List<Widget> chats;
  const ChatListHolder({super.key, required this.chats});

  @override
  _ChatListHolder createState() => _ChatListHolder();
}

class _ChatListHolder extends State<ChatListHolder> {
  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      child: Container(
        decoration: const BoxDecoration(
          borderRadius: BorderRadius.only(
            topLeft: Radius.circular(35),
            topRight: Radius.circular(35),
          ),
          color: Colors.white,
        ),
        child: Column(
          children: widget.chats,
        ),
      ),
    );
  }
}
