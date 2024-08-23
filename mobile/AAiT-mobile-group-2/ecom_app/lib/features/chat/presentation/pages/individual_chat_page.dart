import 'package:ecom_app/features/chat/presentation/widgets/single_text.dart';
import 'package:flutter/material.dart';

class IndividualChatPage extends StatelessWidget {
  const IndividualChatPage({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        backgroundColor: Colors.white,
        appBar: AppBar(
          title: const Text('Individual Chat'),
        ),
        body: Padding(
          padding: EdgeInsets.symmetric(horizontal: 20, vertical: 50),
          child: SingleChildScrollView(
            child: ListBody(children: [
              SingleText(
                isMe: false,
                profile_pic: 'assets/profile_picture.png',
                name: 'Annei Ellison',
                text: 'Have a great working week!! Have a great working week!!',
                time: '09:25 AM',
              ),
              const SizedBox(
                height: 25,
              ),
              SingleText(
                isMe: true,
                profile_pic: 'assets/profile_picture.png',
                name: 'Annei Ellison',
                text: 'Have a great working week!! Have a great working week!!',
                time: '10:44 PM',
              ),
              const SizedBox(
                height: 25,
              ),
              SingleText(
                isMe: false,
                profile_pic: 'assets/profile_picture.png',
                name: 'Annei Ellison',
                image_content: 'assets/blue_portrait.jpg',
                time: '10:44 PM',
              ),
            ]),
          ),
        ));
  }
}
