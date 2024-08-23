import 'dart:ui';
import 'package:flutter/material.dart';

class Chat extends StatefulWidget {
  final String message;
  const Chat({super.key, required this.message});

  @override
  State<Chat> createState() => _ChatState();
}

class _ChatState extends State<Chat> {
    bool isSelf = false;

  @override
  Widget build(BuildContext context) {
    return ClipRRect(
      borderRadius: BorderRadius.only(
        topRight: isSelf ? Radius.circular(0) : Radius.circular(18),
        topLeft: isSelf ? Radius.circular(18) : Radius.circular(0),
        bottomLeft: Radius.circular(18),
        bottomRight: Radius.circular(18)
      ),
  
      child: Container(
        width: 280,
        child: Container(
          
          decoration:BoxDecoration(
            color: isSelf ? Color(0xFF3E50F3): Color.fromARGB(255, 225, 238, 249),
          ),
          child: Padding(
            padding: const EdgeInsets.symmetric(horizontal: 8, vertical:  12.0),
            child: Text(
              widget.message,
              style: TextStyle(
                color: isSelf ? Colors.white : Color.fromARGB(255, 36, 36, 36),
                fontSize: 12,
                fontFamily: 'Poppins',
                fontWeight: FontWeight.w400
              ),
            ),
          ),
        ),
      ),
    );
  }
}