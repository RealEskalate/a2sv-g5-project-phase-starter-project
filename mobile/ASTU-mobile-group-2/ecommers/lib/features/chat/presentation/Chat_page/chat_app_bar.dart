import 'package:flutter/material.dart';

class ChatAppBar extends StatelessWidget {
  const ChatAppBar({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Colors.white,
        elevation: 0,
        leading: IconButton(
          icon: const Icon(Icons.arrow_back, color: Colors.black),
          onPressed: () {
            Navigator.pop(context);
          },
        ),
        title: Transform.translate(
          offset: const Offset(-15, 0),
          child: const Row(
            children: [
              CurrentUser(
                name: 'Sabila Sayma',
                image: 'assets/image/avator.png',
                online: true,
              ),
              SizedBox(width: 8),
              // For the group chat add ( "10 members 5 online")
            ],
          ),
        ),
        // Call and video call buttons as actions
        actions: [
          IconButton(
            icon: const Icon(Icons.call_outlined, color: Colors.black),
            onPressed: () {
              // Add call action
            },
          ),
          IconButton(
            icon: const Icon(Icons.videocam_outlined, color: Colors.black),
            onPressed: () {
              // Add video call action
            },
          ),
        ],
      ),
      body: const SizedBox(),
    );
  }
}

class CurrentUser extends StatelessWidget {
  final String name;
  final String image;
  final bool online;
  final Color? statusColor;

  const CurrentUser({
    Key? key,
    required this.name,
    required this.image,
    required this.online,
    this.statusColor,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        Stack(
          children: [
            CircleAvatar(
              radius: 20.0,
              backgroundImage: AssetImage(image),
            ),
            Positioned(
              bottom: 0,
              right: 0,
              child: Container(
                width: 12,
                height: 12,
                decoration: BoxDecoration(
                  color: online ? Colors.green : Colors.red,
                  shape: BoxShape.circle,
                  border: Border.all(
                    color: Colors.white,
                    width: 2,
                  ),
                ),
              ),
            ),
          ],
        ),
        const SizedBox(width: 8),
        Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              name,
              style: const TextStyle(
                fontSize: 16,
                fontWeight: FontWeight.bold,
                color: Colors.black,
              ),
            ),
            Text(
              online ? 'Online' : 'Offline',
              style: TextStyle(
                fontSize: 12,
                color: statusColor ?? Colors.black54,
              ),
            ),
          ],
        ),
      ],
    );
  }
}
