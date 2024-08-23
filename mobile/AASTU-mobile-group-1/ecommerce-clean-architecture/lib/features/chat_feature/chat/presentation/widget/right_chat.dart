import 'package:ecommerce/features/chat_feature/chat/domain/entity/message.dart';
import 'package:flutter/material.dart';

// ignore: must_be_immutable
class RightChat extends StatelessWidget {
  Message message;
  RightChat({required this.message, super.key});

  @override
  Widget build(BuildContext context) {
    double width = MediaQuery.of(context).size.width;
    double height = MediaQuery.of(context).size.height;
    return Column(
      crossAxisAlignment: CrossAxisAlignment.end,
      children: [
        Row(
          mainAxisAlignment: MainAxisAlignment.end,
          children: [
            const Text('You',
                style: TextStyle(fontSize: 15, fontWeight: FontWeight.w500)),
            SizedBox(
              width: width * 0.02,
            ),
            const CircleAvatar(
              radius: 25,
              backgroundImage: AssetImage("assets/images/splash.png"),
            )
          ],
        ),
        SizedBox(
          height: height * 0.01,
        ),
        Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            if (message.type == "text")
              Container(
                margin: EdgeInsets.only(right: 45),
                constraints: BoxConstraints(
                  maxWidth: width * 0.65,
                ),
                decoration: BoxDecoration(
                    color: Color(0xFF3F51F3),
                    borderRadius: BorderRadius.only(
                        topLeft: Radius.circular(12),
                        bottomLeft: Radius.circular(12),
                        bottomRight: Radius.circular(12))),
                child: Padding(
                  padding: const EdgeInsets.all(8.0),
                  child: Text(
                    message.content!,
                    style: TextStyle(
                        color: Colors.white,
                        fontSize: 12,
                        fontWeight: FontWeight.w500),
                  ),
                ),
              ),
            
          ],
        ),
        if (message.type == "image")
          Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Container(
                margin: EdgeInsets.only(right: 45, top: 8.0),
                width: width * 0.4,
                height: height * 0.2,
                decoration: BoxDecoration(
                  borderRadius: BorderRadius.circular(12),
                  image: DecorationImage(
                    image: AssetImage(message.content!),
                    fit: BoxFit.cover,
                  ),
                ),
              ),
            ],
          )
      ],
    );
  }
}
