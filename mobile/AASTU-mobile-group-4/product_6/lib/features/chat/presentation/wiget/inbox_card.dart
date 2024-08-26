import 'package:flutter/material.dart';

import '../pages/dm_page.dart';

class ConversationCard extends StatelessWidget {
  final String userImage;
  final String userName;
  final String lastMessage;
  final String time;
  final int unreadMessages;

  const ConversationCard({
    required this.userImage,
    required this.userName,
    required this.lastMessage,
    required this.time,
    this.unreadMessages = 0,
  });

  @override
  Widget build(BuildContext context) {
    return ListTile(
      contentPadding: EdgeInsets.all(5.0),
      leading: Stack(
        children: [
          Container(
            width: 45,
            height: 45,
            decoration: const BoxDecoration(
              shape: BoxShape.circle,
              color: Colors.blue,
            ),
            clipBehavior: Clip.hardEdge,
            child: Image.asset(userImage, fit: BoxFit.cover),
          ),
          Positioned(
            bottom: 0,
            right: 0,
            child: Container(
              width: 12,
              height: 12,
              decoration: BoxDecoration(
                color: onlineChecker(time)
                    ? Colors.green
                    : Colors.grey, // Green if online, gray if offline
                shape: BoxShape.circle,
                border: Border.all(
                  color:
                      Colors.white, // Optional: Add a white border to the dot
                  width: 2.0,
                ),
              ),
            ),
          ),
        ],
      ),
      title: GestureDetector(
        onTap: () {
          print('Tapped');
          Navigator.push(
            context,
            MaterialPageRoute(
              builder: (context) => DmPage(),
            ),
          );
        },
        child: Text(
          userName,
          style: TextStyle(
            fontWeight: FontWeight.bold,
            fontSize: 16,
          ),
        ),
      ),
      subtitle: Text(
        lastMessage,
        style: TextStyle(
          color: Colors.grey,
          fontSize: 14,
        ),
      ),
      trailing: Column(
        crossAxisAlignment: CrossAxisAlignment.end,
        children: [
          Text(
            displayOnline(time),
            style: TextStyle(
              color: Colors.grey,
              fontSize: 12,
            ),
          ),
          if (unreadMessages > 0)
            Container(
              margin: EdgeInsets.only(top: 4.0),
              padding: EdgeInsets.symmetric(horizontal: 6.0, vertical: 2.0),
              decoration: BoxDecoration(
                color: Colors.blue[700],
                borderRadius: BorderRadius.circular(12.0),
              ),
              child: Text(
                '$unreadMessages',
                style: TextStyle(
                  color: Colors.white,
                  fontSize: 12,
                ),
              ),
            ),
        ],
      ),
      onTap: () {
        // Handle tap event
      },
    );
  }
}

bool onlineChecker(String time) {
  RegExp regExp = RegExp(r'\d+');

  Match? match = regExp.firstMatch(time);

  if (match != null) {
    int number = int.parse(match.group(0)!);
    if (number != 0) {
      return false;
    } else {
      return true;
    }
  } else {
    return true;
  }
}

String displayOnline(String timeStamp) {
  if (onlineChecker(timeStamp)) {
    return 'Online';
  } else {
    return timeStamp;
  }
}
