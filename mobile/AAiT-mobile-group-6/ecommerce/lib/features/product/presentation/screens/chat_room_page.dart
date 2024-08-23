import 'package:flutter/material.dart';

import '../widgets/widgets.dart';

class ChatPage extends StatefulWidget {
  const ChatPage({super.key});

  @override
  State<ChatPage> createState() => _ChatPageState();
}

class _ChatPageState extends State<ChatPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: const UserInfo(name: 'John', members: 8, onlines: 5),
          actions: [
            IconButton(onPressed: () {}, icon: const Icon(Icons.call_outlined)),
            IconButton(
                onPressed: () {}, icon: const Icon(Icons.videocam_outlined)),
          ],
        ),
        bottomNavigationBar: const BottomChatBar(),
        body: const ListOfMessages());
  }
}

class UserInfo extends StatelessWidget {
  final String name;
  final int onlines, members;
  const UserInfo({
    super.key,
    required this.name,
    required this.onlines,
    required this.members,
  });

  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        const Padding(
          padding: EdgeInsets.all(8.0),
          child: UserImage(
            size: 50,
            imageUrl: 'someimage',
          ),
        ),
        Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(name,
                style: const TextStyle(
                    fontFamily: 'Sora',
                    fontSize: 20,
                    fontWeight: FontWeight.bold)),
            Text(
              '$members members, $onlines onlines',
              style: TextStyle(
                  color: Colors.grey.shade600,
                  fontFamily: 'Sora',
                  fontSize: 12,
                  fontWeight: FontWeight.w400),
            )
          ],
        ),
      ],
    );
  }
}

class UserImage extends StatelessWidget {
  final double size;
  final String imageUrl;
  const UserImage({
    super.key,
    required this.size,
    required this.imageUrl,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      width: size,
      height: size,
      decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(size / 2),
          color: Colors.grey.shade400),
      child: CachedNetworkImgeWithshimmerWaiter(
        imageUrl: imageUrl,
      ),
    );
  }
}

class BottomChatBar extends StatelessWidget {
  const BottomChatBar({super.key});

  @override
  Widget build(BuildContext context) {
    return BottomAppBar(
        color: Colors.white,
        shadowColor: Colors.black,
        surfaceTintColor: Colors.white,
        elevation: 50,
        child: Row(
          children: [
            IconButton(
                onPressed: () {},
                icon: Transform.rotate(
                  //
                  angle: 45,
                  child: const Icon(
                    Icons.attach_file,
                    textDirection: TextDirection.rtl,
                  ),
                )),
            Container(
              width: MediaQuery.of(context).size.width / 1.9,
              decoration: BoxDecoration(
                color: Colors.grey.shade200,
                borderRadius: BorderRadius.circular(10),
              ),
              clipBehavior: Clip.hardEdge,
              child: TextField(
                decoration: InputDecoration(
                  hintText: 'Write your message',
                  fillColor: Colors.grey.shade200,
                  filled: true,
                  border: InputBorder.none,
                ),
              ),
            ),
            IconButton(
                onPressed: () {},
                icon: const Icon(Icons.photo_camera_outlined)),
            IconButton(onPressed: () {}, icon: const Icon(Icons.mic_outlined)),
          ],
        ));
  }
}

class ListOfMessages extends StatelessWidget {
  const ListOfMessages({super.key});

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      height: MediaQuery.of(context).size.height,
      child: ListView.builder(
          itemCount: 9,
          reverse: true,
          itemBuilder: (context, index) {
            return const Padding(
              padding: EdgeInsets.all(15.0),
              child: SenderMessage(
                  imageUrl: 'someImage',
                  message: 'what ever the message is!',
                  time: '2 : 49 PM'),
            );
          }),
    );
  }
}

class SenderMessage extends StatelessWidget {
  final String imageUrl;
  final String message;
  final String time;
  const SenderMessage(
      {super.key,
      required this.imageUrl,
      required this.message,
      required this.time});

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Row(
          children: [
            const Spacer(),
            const Text('you'),
            const SizedBox(width: 3),
            UserImage(size: 40, imageUrl: imageUrl)
          ],
        ),
        Container(
            padding: const EdgeInsets.all(8),
            width: MediaQuery.of(context).size.width / 1.5,
            decoration: const BoxDecoration(
                color: Colors.blue,
                borderRadius: BorderRadius.only(
                  topLeft: Radius.circular(10),
                  bottomLeft: Radius.circular(10),
                  bottomRight: Radius.circular(10),
                )),
            child: Text(
              message,
              style: const TextStyle(color: Colors.white),
            )),
        Text(time)
      ],
    );
  }
}
