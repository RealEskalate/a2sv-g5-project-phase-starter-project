import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:starter_project/features/chat/presentation/widgets/people_list_view_widget.dart';

import '../widgets/story_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
          backgroundColor: Colors.blue,
          // toolbarHeight: 50,
          leadingWidth: 100,
          leading: Padding(
            padding: const EdgeInsets.fromLTRB(10, 20, 0, 0),
            child: IconButton(
              onPressed: () {},
              icon: const Icon(
                color: Colors.white,
                size: 30,
                Icons.search,
              ),
            ),
          )),
      body: const Stack(
        children: [
          StoryWidget(),
          PeopleListWidget(),
        ],
      ),
      // ),
    );
  }
}
