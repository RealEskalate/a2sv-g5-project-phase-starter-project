import 'package:flutter/material.dart';

import 'text_buble.dart';

class SingleText extends StatelessWidget {
  const SingleText(
      {super.key,
      required this.profile_pic,
      required this.name,
      this.text,
      required this.isMe,
      required this.time,
      this.image_content});

  final String profile_pic;
  final String name;
  final String? text;
  final String? image_content;
  final String time;
  final bool isMe;

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: isMe ? MainAxisAlignment.end : MainAxisAlignment.start,
      crossAxisAlignment: CrossAxisAlignment.start,
      children: isMe
          ? [
              Column(
                crossAxisAlignment: CrossAxisAlignment.end,
                children: [
                  Text(
                    'You',
                    style: TextStyle(fontSize: 18, fontWeight: FontWeight.w600),
                  ),
                  const SizedBox(height: 8),
                  _buildChatContent(
                      isMe: isMe, image_content: image_content, text: text),
                  const SizedBox(height: 8),
                  Padding(
                    padding: const EdgeInsets.only(right: 15.0),
                    child: Text(
                      time,
                      style: TextStyle(color: Colors.grey),
                    ),
                  )
                ],
              ),
              const SizedBox(width: 10),
              CircleAvatar(
                backgroundImage: AssetImage(profile_pic),
                radius: 30,
              ),
              const SizedBox(
                height: 5,
              ),
            ]
          : [
              CircleAvatar(
                backgroundImage: AssetImage(profile_pic),
                radius: 30,
              ),
              const SizedBox(width: 10),
              Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    name,
                    style: TextStyle(fontSize: 18, fontWeight: FontWeight.w600),
                  ),
                  const SizedBox(height: 8),
                  _buildChatContent(
                      isMe: isMe, image_content: image_content, text: text),
                  const SizedBox(height: 8),
                  Padding(
                    padding: const EdgeInsets.only(left: 15.0),
                    child: Text(
                      time,
                      style: TextStyle(color: Colors.grey),
                    ),
                  )
                ],
              ),
            ],
    );
  }
}

class _buildChatContent extends StatelessWidget {
  final String? image_content;
  final String? text;
  final bool isMe;

  const _buildChatContent(
      {super.key, this.image_content, this.text, required this.isMe});

  @override
  Widget build(BuildContext context) {
    if (image_content != null) {
      return ConstrainedBox(
        constraints: BoxConstraints(
          maxWidth: 300,
          maxHeight: 300,
        ),
        child: Container(
          decoration: BoxDecoration(
            borderRadius: BorderRadius.circular(15),
            image: DecorationImage(
              image: AssetImage(image_content!),
              fit: BoxFit.cover,
            ),
          ),
        ),
      );
    } else {
      return TextBubble(text: text!, isMe: isMe);
    }
  }
}
