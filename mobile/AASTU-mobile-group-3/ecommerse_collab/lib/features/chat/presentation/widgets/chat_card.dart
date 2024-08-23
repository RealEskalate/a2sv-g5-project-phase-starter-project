import 'package:flutter/material.dart';

import '../../../authentication/domain/entity/user.dart';
import 'user_avater.dart';

class ChatCard extends StatefulWidget {
  final String topMessage;
  final User user;
  // final DateTime time;
  final int time;
  final int unread;
  final bool online;

  const ChatCard(
      {super.key,
      required this.topMessage,
      required this.user,
      required this.time,
      required this.unread,
      this.online = false
      });

  @override
  State<ChatCard> createState() => _ChatCardState();
}

class _ChatCardState extends State<ChatCard> {
  @override
  Widget build(BuildContext context) {
    return Container( 
      padding: const EdgeInsets.symmetric(horizontal: 10, vertical:2),
      margin: const EdgeInsets.all(2),
      width: 370,
      height: 70,
      decoration: const BoxDecoration(
        color: Colors.white
      ),
        child: Row(
          crossAxisAlignment: CrossAxisAlignment.center,
      children: [
         UserAvater(image: 'assets/images/avater.png', online: widget.online),
        Expanded(
          child: Padding(
            padding: const EdgeInsets.all(10.0),
            child: Column(

              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: [
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Text(
                      widget.user.username,
                      style: const TextStyle(
                          fontWeight: FontWeight.w500, fontSize: 17),
                    ),
                    Text("${widget.time} min ago",style: const TextStyle(
                          fontWeight: FontWeight.w400, fontSize: 12, color: Color(0xFF797C7B)),)
                  ],
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Text(
                      widget.topMessage,
                      style: const TextStyle(
                          fontWeight: FontWeight.w400,
                          color: Color(0xFF797C7B),
                          fontSize: 12),
                    ),
                    Container(
                      width: 22, 
                      height: 22,
                      decoration: BoxDecoration(borderRadius: BorderRadius.circular(22), color: const Color(0xFF3F51F3)),
                      child: Align(alignment: Alignment.center,child: Text(widget.unread.toString(), style: const TextStyle(
                        fontFamily: "AbhayaLibre",
                        fontWeight: FontWeight.w700,
                        fontSize: 15,
                        color: Colors.white)),),
                    )
                  ],
                )
              ],
            ),
          ),
        ),
      ],
    ));
  }
}
