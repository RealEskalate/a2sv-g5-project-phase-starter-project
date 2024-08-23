import 'package:flutter/material.dart';

import '../widgets/chat_list_holder.dart';
import '../widgets/chater_card.dart';
import '../widgets/others_story_avatar.dart';
import '../widgets/own_story_avatar.dart';
import '../widgets/stories_card_list.dart';

class ChatPage extends StatelessWidget {
  ChatPage({super.key});
  // list builder that takes an index and returns a chat list holder
  final List<Widget> chats = List.generate(
    10,
    (index) {
      return ChatterCard(
        name: 'name',
        lastMessage: 'lastMessage',
        unreadCount: index,
        time: '2min',
        imageUrl: 'assets/dummy_avator.png',
        isOnline: false,
      );
    },
  );
  final List<Widget> stories = List.generate(
    10,
    (index) {
      if (index == 0) {
        return OwnStoryAvatar(
          name: 'name',
          avatarUrl: 'assets/dummy_avator.png',
        );
      }
      return OthersStoryAvatar(
        name: 'name',
        avatarUrl: 'assets/dummy_avator.png',
      );
    },
  );
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Theme.of(context).secondaryHeaderColor,
      appBar: AppBar(
        backgroundColor: Theme.of(context).secondaryHeaderColor,
        leading: Container(
          margin: const EdgeInsets.all(10),
          child: IconButton(
            icon: const Icon(
              Icons.search,
              color: Colors.white,
              size: 35,
            ),
            onPressed: () {
              Navigator.pop(context);
            },
          ),
        ),
      ),
      body: SingleChildScrollView(
        child: Center(
          child: Column(
            children: [
              StoriesCardList(
                stories: stories,
              ),
              ChatListHolder(
                chats: chats,
              ),
            ],
          ),
        ),
      ),
    );
  }
}
