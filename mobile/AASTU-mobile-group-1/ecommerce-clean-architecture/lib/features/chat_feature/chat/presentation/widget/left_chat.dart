import 'package:ecommerce/features/chat_feature/chat/domain/entity/message.dart';
import 'package:flutter/material.dart';

class LeftChat extends StatelessWidget {
  final Message message;
  LeftChat({required this.message, super.key});

  @override
  Widget build(BuildContext context) {
    double width = MediaQuery.of(context).size.width;
    double height = MediaQuery.of(context).size.height;

    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Row(
          mainAxisAlignment: MainAxisAlignment.start,
          children: [
           const CircleAvatar(
              radius: 25,
              backgroundImage: AssetImage('assets/images/splash.png'),
            ),
            SizedBox(
              width: width * 0.02,
            ),
           const Text(
              'Annei Ellison',
              style: TextStyle(fontSize: 15, fontWeight: FontWeight.w500),
            ),
          ],
        ),
        SizedBox(
          height: height * 0.01,
        ),
       if (message.type == 'text')
        Column(
          crossAxisAlignment: CrossAxisAlignment.end,
          children: [
            Container(
              margin: EdgeInsets.only(left: 45),
              constraints: BoxConstraints(
                maxWidth: width * 0.65,
              ),
              decoration: BoxDecoration(
                color: Color(0xFFF2F7FB),
                borderRadius: BorderRadius.only(
                  topRight: Radius.circular(12),
                  bottomLeft: Radius.circular(12),
                  bottomRight: Radius.circular(12),
                ),
              ),
              child: Padding(
                padding: const EdgeInsets.all(8.0),
                child: Text(
                  message.content!,
                  style: TextStyle(
                    fontSize: 12,
                    fontWeight: FontWeight.w500,
                  ),
                ),
              ),
            ),
           
          ],
        ),
        if (message.type == 'image')
          Column(
            crossAxisAlignment: CrossAxisAlignment.end,
            children: [
              Container(
                margin: EdgeInsets.only(left: 45, top: 8.0),
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
