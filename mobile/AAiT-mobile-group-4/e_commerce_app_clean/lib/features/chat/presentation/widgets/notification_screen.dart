import 'dart:math';

import 'package:flutter/material.dart';

import '../../../../core/personal_message_notification_data.dart';
import 'personal_message_notification_widget.dart';
import 'stories_widget.dart';

class ChatPage extends StatelessWidget {
  const ChatPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      color: const Color.fromRGBO(73, 140, 240, 1),
      child: Padding(
        padding: const EdgeInsets.only(top: 36),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            const IconButton(
              padding: EdgeInsets.only(left: 20),
              onPressed: null,
              iconSize: 30,
              icon: Icon(Icons.search, color: Colors.white),
            ),
            const StoriesWidget(),
            const SizedBox(height: 20),
            Expanded(
              child: Container(
                height: MediaQuery.of(context).size.height,
                decoration: const BoxDecoration(
                  color: Colors.white,
                  borderRadius: BorderRadius.only(
                    topLeft: Radius.circular(30),
                    topRight: Radius.circular(30),
                  ),
                ),
                // content goes here
                // ------------------------------------------

                child: ListView.builder(
                    itemCount: personalNotifications.length,
                    itemBuilder: (context, index) {
                      return PersonalMessageNotification(
                        imagePath: 'assets/story_${(index%3)+1}.png',
                        bgColor: Color((Random().nextDouble() * 0xFFFFFF).toInt() << 0).withOpacity(1.0),
                        fullName: personalNotifications[index].fullName,
                        message: personalNotifications[index].message,
                        timeSent: personalNotifications[index].timeSent,
                        isRead: personalNotifications[index].isRead,
                        isOnline: personalNotifications[index].isOnline,
                        unreadMessages: personalNotifications[index].unreadMessages,
                      );
                    }),

                // ------------------------------------------
              ),
            ),
          ],
        ),
      ),
    );
  }
}
