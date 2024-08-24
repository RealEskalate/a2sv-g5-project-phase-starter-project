import 'package:flutter/material.dart';
import 'package:starter_project/features/chat/presentation/widgets/bottom_navigation_bar.dart';
import 'package:starter_project/features/chat/presentation/widgets/chat_page_body.dart';

class ChatPage extends StatelessWidget {
  const ChatPage({super.key});

  static const total_member = '8 Member ,';
  static const online_member = '5 Online';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      bottomNavigationBar: const BottomNavigationBarWidget(),
      appBar: AppBar(
        toolbarHeight: 100.0,
        leading: IconButton(
          icon: const Icon(
            Icons.arrow_back,
            size: 25.0,
          ),
          onPressed: () {
            // Navigator.of(context).pop();
          },
        ),
        centerTitle: true,
        title: Row(
          children: [
            Container(
              width: 60,
              height: 60,
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(200.0 * 0.4),
                color: const Color.fromARGB(255, 254, 199, 211)
              ),
              child: ClipRRect(
                borderRadius: BorderRadius.circular(200.0 * 0.4),
                child: Image.asset(
                  'assets/avatar.png',
                  fit: BoxFit.cover,
                ),
              ),
            ),
            const SizedBox(width: 16),
            const Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                mainAxisSize: MainAxisSize.min,
                children: [
                  Text(
                    'Monkey.D Luffy',
                    style: TextStyle(fontSize: 18, fontWeight: FontWeight.w500),
                  ),
                  SizedBox(
                    height: 5.0,
                  ),
                  Row(
                    children: [
                      Text(
                        total_member,
                        style: TextStyle(fontSize: 14, color: Colors.grey),
                      ),
                      SizedBox(
                        width: 5,
                      ),
                      Text(
                        online_member,
                        style: TextStyle(fontSize: 14, color: Colors.green),
                      ),
                    ],
                  )
                ],
              ),
            ),
          ],
        ),
        actions: [
          IconButton(
            icon: const Icon(
              Icons.phone_outlined,
              size: 25.0,
            ),
            onPressed: () {
              // Add your more action here
            },
          ),
          IconButton(
            icon: const Icon(
              Icons.video_call,
              size: 25.0,
            ),
            onPressed: () {
              // Add your more action here
            },
          ),
        ],
      ),
      body: const ChatBody(),
    );
  }
}
