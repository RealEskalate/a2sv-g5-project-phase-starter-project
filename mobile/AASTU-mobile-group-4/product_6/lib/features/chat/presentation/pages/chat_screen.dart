import 'package:flutter/material.dart';

import '../wiget/frequent_user.dart';
import '../wiget/inbox_card.dart';

class ChatScreen extends StatefulWidget {
  @override
  _ChatScreenState createState() => _ChatScreenState();
}

class _ChatScreenState extends State<ChatScreen> {
  bool isSearching = false;

  final List<Map<String, String>> users = [
    {'image': 'images/profile.jpg', 'name': 'My status'},
    {'image': 'images/user_profile.jpg', 'name': 'Adil'},
    {'image': 'images/profile.jpg', 'name': 'Marina'},
    {'image': 'images/user_profile.jpg', 'name': 'Dean'},
    {'image': 'images/user_profile.jpg', 'name': 'Max'},
  ];

  final List<Map<String, dynamic>> conversations = [
    {
      'userImage': 'images/profile.jpg',
      'userName': 'John Borino',
      'lastMessage': 'Have a good day',
      'time': '2 min ago',
      'unreadMessages': 3,
    },
    {
      'userImage': 'images/user_profile.jpg',
      'userName': 'Alex Linderson',
      'lastMessage': 'How are you today?',
      'time': '2 min ago',
      'unreadMessages': 0,
    },
    {
      'userImage': 'images/user_profile.jpg',
      'userName': 'Team Align',
      'lastMessage': 'Don\'t miss to attend the meeting.',
      'time': '0 min ago',
      'unreadMessages': 2,
    },
    {
      'userImage': 'images/profile.jpg',
      'userName': 'John Ahrabham',
      'lastMessage': 'Hey! Can you join the meeting?',
      'time': '2 min ago',
      'unreadMessages': 1,
    },
    {
      'userImage': 'images/user_profile.jpg',
      'userName': 'Sabila Sayma',
      'lastMessage': 'How are you today?',
      'time': '2 min ago',
      'unreadMessages': 0,
    },
    {
      'userImage': 'images/user_profile.jpg',
      'userName': 'Angel Dayna',
      'lastMessage': 'How are you today?',
      'time': '0 min ago',
      'unreadMessages': 0,
    },
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Color.fromRGBO(63, 81, 243, 1),
      body: Column(
        children: [
          Container(
            padding: EdgeInsets.only(top: 10, left: 10, right: 10),
            child: Column(
              children: [
                Row(
                  children: [
                    IconButton(
                      icon: Icon(
                        isSearching ? Icons.close : Icons.search,
                        color: Colors.white,
                      ),
                      onPressed: () {
                        setState(() {
                          isSearching = !isSearching;
                        });
                      },
                    ),
                    Expanded(
                      child: isSearching
                          ? TextField(
                              decoration: InputDecoration(
                                hintText: 'Search...',
                                hintStyle: TextStyle(color: Colors.white70),
                                border: InputBorder.none,
                              ),
                              style: TextStyle(color: Colors.white),
                              autofocus: true,
                              onSubmitted: (value) {
                                setState(() {
                                  isSearching = false;
                                });
                              },
                            )
                          : Text(
                              '',
                              style: TextStyle(
                                color: Colors.white,
                                fontSize: 12,
                                fontWeight: FontWeight.bold,
                              ),
                            ),
                    ),
                  ],
                ),
                FrequentUsersCard(users: users),
              ],
            ),
          ),
          Expanded(
            child: Container(
              padding: EdgeInsets.only(top: 15, left: 10, right: 10),
              decoration: BoxDecoration(
                color: Colors.blueGrey[50],
                borderRadius: BorderRadius.only(
                  topLeft: Radius.circular(30),
                  topRight: Radius.circular(30),
                ),
              ),
              child: ListView.builder(
                itemCount: conversations.length,
                itemBuilder: (context, index) {
                  return ConversationCard(
                    userImage: conversations[index]['userImage'],
                    userName: conversations[index]['userName'],
                    lastMessage: conversations[index]['lastMessage'],
                    time: conversations[index]['time'],
                    unreadMessages: conversations[index]['unreadMessages'],
                  );
                },
              ),
            ),
          ),
        ],
      ),
    );
  }
}
