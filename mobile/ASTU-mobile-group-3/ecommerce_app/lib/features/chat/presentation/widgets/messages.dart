import 'package:custom_clippers/custom_clippers.dart';
import 'package:flutter/material.dart';

enum messageType {
  send,
  receive,
}

class ChatMessage extends StatelessWidget {
  final MessageType messageType;
  final String message;
  final String timeStamp;
  final bool isImage;

  const ChatMessage({
    super.key,
    required this.messageType,
    required this.message,
    required this.timeStamp,
    required this.isImage,
  });

  @override
  Widget build(BuildContext context) {
    return Align(
      alignment: messageType == MessageType.receive
          ? Alignment.centerLeft
          : Alignment.centerRight,
      child: Padding(
        padding: messageType == MessageType.receive
            ? const EdgeInsets.only(right: 80, bottom: 12)
            : const EdgeInsets.only(left: 80, bottom: 12),
        child: ClipPath(
          clipper: UpperNipMessageClipperTwo(messageType),
          child: Container(
            width: MediaQuery.sizeOf(context).width * 0.4,
            padding: const EdgeInsets.all(10),
            color: messageType == MessageType.receive
                ? Color(0xff3F51F3)
                : const Color(0xffF2F7FB),
            alignment: Alignment.center,
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Padding(
                  padding: const EdgeInsets.only(left: 12),
                  child: isImage
                      ? Image.network(message)
                      : Text(
                          message,
                          style: TextStyle(
                            fontSize: 12,
                            fontWeight: FontWeight.w400,
                            color: messageType == MessageType.receive
                                ? Colors.white
                                : const Color(0xff222222),
                          ),
                        ),
                ),
                Padding(
                  padding: const EdgeInsets.only(left: 8, top: 10),
                  child: Align(
                    alignment: Alignment.bottomLeft,
                    child: Text(
                      timeStamp,
                      style: TextStyle(
                        fontSize: 10,
                        fontWeight: FontWeight.w400,
                        color: messageType == MessageType.receive
                            ? Colors.white
                            : const Color(0xff222222),
                      ),
                    ),
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
