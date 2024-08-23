import 'package:ecom_app/features/chat/presentation/widgets/current_user.dart';
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
        leading: const Icon(Icons.arrow_back),
        title: const CurrentUser(
          name: 'John Doe',
          image: 'assets/images/user.png',
          online: true,
        ),
        actions: const [
          IconButton(
            icon: Icon(
              Icons.call_outlined,
              color: Colors.black,
            ),
            onPressed: null,
          ),
          IconButton(
            icon: Icon(
              Icons.videocam_outlined,
              color: Colors.black,
            ),
            onPressed: null,
          ),
        ],
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
              _buildInputBar(),
            ]),
          ),
          
        ));
  }
}

Widget _buildInputBar() {
  return Padding(
    padding: const EdgeInsets.symmetric(horizontal: 8.0, vertical: 10.0),
    child: Row(
      children: [
        IconButton(
          icon: const Icon(Icons.attach_file_outlined,
            color: Colors.black,
          
          ),
          onPressed: () {
            // Handle attachment
          },
        ),
        Expanded(
          child: Container(
            decoration: BoxDecoration(
              color: Colors.grey[200],
              borderRadius: BorderRadius.circular(25),
            ),
            child: Row(
              children: [
                const SizedBox(width: 10),
                const Expanded(
                  child: TextField(
                    decoration: InputDecoration(
                      hintText: "Write your message",
                      border: InputBorder.none,
                    ),
                  ),
                ),
                IconButton(
                  icon: const Icon(Icons.copy,),
                  onPressed: () {
                    // Handle copy action
                  },
                ),
              ],
            ),
          ),
        ),
        IconButton(
          icon: const Icon(Icons.camera_alt_outlined),
          onPressed: () {
            // Handle camera action
          },
        ),
        IconButton(
          icon: const Icon(Icons.mic_none_outlined),
          onPressed: () {
            // Handle microphone action
          },
        ),
      ],
    ),
  );
}

