import 'package:flutter/material.dart';

class StoriesCardList extends StatelessWidget {
  const StoriesCardList({
    super.key,
    required this.stories,
  });

  final List<Widget> stories;

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      scrollDirection: Axis.horizontal,
      child: Container(
        margin: const EdgeInsets.fromLTRB(10, 10, 10, 40),
        child: Row(
          children: stories,
        ),
      ),
    );
    // return ListView.builder(
    //   scrollDirection: Axis.horizontal,
    //   itemCount: stories.length,
    //   itemBuilder: (context, index) {
    //     return stories[index];
    //   },
    // );
  }
}
