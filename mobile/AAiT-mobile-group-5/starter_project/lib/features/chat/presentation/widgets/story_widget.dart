import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';

final List<String> names = [
  "Alex Anderson",
  "Alex Anderson",
  "Alex Anderson",
  "Alex Anderson",
  "Alex Anderson",
  "Alex Anderson",
  "Alex Anderson",
  "Alex Anderson",
  "Alex Anderson",
  "Alex Anderson",
  "Alex Anderson",
  "Alex Anderson",
];

class StoryWidget extends StatelessWidget {
  const StoryWidget({super.key});

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      child: Container(
        padding: const EdgeInsets.fromLTRB(20, 20, 0, 30),
        width: MediaQuery.of(context).size.width,
        height: 230,
        color: Colors.blue,
        child: SingleChildScrollView(
          physics: const AlwaysScrollableScrollPhysics(),
          scrollDirection: Axis.horizontal,
          child: Row(
            children: generateListIcons(names),
          ),
        ),
      ),
    );
  }
}

List<Widget> generateListIcons(List<String> names) {
  final List<Widget> widgets = [myAvatar()];

  widgets.addAll(names.map((name) {
    return Padding(
      padding: const EdgeInsets.fromLTRB(0, 10, 20, 10),
      child: Column(
        children: [
          const CircleAvatar(
            radius: 32,
            backgroundColor: Color.fromARGB(255, 33, 243, 93),
            child: CircleAvatar(
              radius: 31,
              backgroundColor: Colors.blue,
              child: CircleAvatar(
                  radius: 28,
                  backgroundColor: Color.fromARGB(255, 33, 243, 93),
                  child: Icon(
                    Icons.person,
                    size: 20,
                  )),
            ),
          ),
          const SizedBox(height: 10),
          Text(
            name,
            style: const TextStyle(
              fontSize: 14,
              color: Colors.white,
            ),
          ),
        ],
      ),
    );
  }).toList());

  return widgets;
}

Widget myAvatar() {
  return Padding(
      padding: const EdgeInsets.fromLTRB(0, 10, 20, 10),
      child: Column(children: [
        Stack(children: [
          const CircleAvatar(
            radius: 32,
            backgroundColor: Color.fromARGB(255, 150, 241, 65),
            child: CircleAvatar(
              radius: 31,
              backgroundColor: Colors.blue,
              child: CircleAvatar(
                  radius: 28,
                  backgroundColor: Color.fromARGB(255, 150, 241, 65),
                  child: Icon(
                    Icons.person,
                    size: 20,
                  )),
            ),
          ),
          Positioned(
            right: 0,
            bottom: 0,
            child: IconButton(
              color: Colors.red,
              onPressed: () {},
              icon: const Icon(
                weight: 15,
                Icons.add,
                size: 20,
                color: Colors.white,
              ),
            ),
          ),
        ]),
        const SizedBox(height: 10),
        const Text(
          "My Status",
          style: TextStyle(
            fontSize: 14,
            color: Colors.white,
          ),
        ),
      ]));
}
