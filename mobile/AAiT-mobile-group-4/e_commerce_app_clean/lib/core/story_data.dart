import 'dart:ui';

import 'package:flutter/material.dart';

class StoryData {
  String storyImage;
  String storyTitle;
  Color storyColor;

  StoryData(
      {required this.storyImage,
      required this.storyTitle,
      required this.storyColor});
}

List<StoryData> stories = [
  StoryData(storyImage: 'assets/story_1.png', storyTitle: 'My status', storyColor: const Color.fromARGB(255, 233, 148, 216)),
  StoryData(storyImage: 'assets/story_2.png', storyTitle: 'Beth', storyColor: Colors.green),
  StoryData(storyImage: 'assets/story_3.png', storyTitle: 'Bereket', storyColor: Colors.brown),
  StoryData(storyImage: 'assets/story_3.png', storyTitle: 'Imran', storyColor: Colors.redAccent),
  StoryData(storyImage: 'assets/story_2.png', storyTitle: 'Leykun', storyColor: Colors.purple),
  StoryData(storyImage: 'assets/story_1.png', storyTitle: 'Felmeta', storyColor: Colors.orange),
];
