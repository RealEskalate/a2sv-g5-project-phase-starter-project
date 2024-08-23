import 'package:flutter/material.dart';

import '../widgets/others_story_avatar.dart';
import '../widgets/own_story_avatar.dart';

class DummyPage extends StatelessWidget {
  const DummyPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        backgroundColor: Colors.blue,
        leading: IconButton(
          icon: const Icon(
            Icons.search,
            color: Colors.white,
          ),
          onPressed: () {},
        ),
      ),
      body: Container(
        color: Colors.blue,
        height: 200,
        child: ListView.builder(
          itemCount: 10,
          // make the scroll direction horizontal
          scrollDirection: Axis.horizontal,
          itemBuilder: (context, index) {
            if (index == 0) {
              return OwnStoryAvatar(
                name: 'My status',
                avatarUrl: 'assets/dummy_avatar.png',
              );
            }
            return OthersStoryAvatar(
              name: 'Dean',
              avatarUrl: 'assets/dummy_avatar.png',
            );
          },
        ),
      ),
    );
  }
}
