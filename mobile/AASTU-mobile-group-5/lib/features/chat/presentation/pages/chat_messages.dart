import 'package:flutter/material.dart';

import '../widgets/avatar_with_name.dart';
import '../widgets/user_conversation_widget.dart';


class Messages extends StatelessWidget {
  const Messages({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Color.fromRGBO(73, 140, 240, 1),
      body: Stack(
        children: [
          Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Padding(
                padding: const EdgeInsets.only(
                  top: 32,
                  left: 15,
                ),
                child: IconButton(
                  icon: const Icon(
                    Icons.search,
                    color: Colors.white,
                  ),
                  onPressed: () {},
                ),
              ),
              Padding(
                padding: const EdgeInsets.only(
                  left: 15,
                ),
                child: SizedBox(
                  height: 69,
                  child: ListView(
                    scrollDirection: Axis.horizontal,
                    children: const [
                      avatar_with_name(name: 'Aryam'),
                      avatar_with_name(name: 'Afomia'),
                      avatar_with_name(name: 'Daniel'),
                      avatar_with_name(name: 'Makda'),
                      avatar_with_name(name: 'Aryam'),
                      avatar_with_name(name: 'Afomia'),
                      avatar_with_name(name: 'Daniel'),
                      avatar_with_name(name: 'Makda'),
                    ],
                  ),
                ),
              ),
            ],
          ),
          Positioned(
            top: 190,
            bottom: 0,
            child: Container(
              padding: const EdgeInsets.symmetric(vertical: 40),
              height: 568,
              width: MediaQuery.of(context).size.width,
              decoration: const BoxDecoration(
                color: Colors.white,
                borderRadius: BorderRadius.only(
                  topLeft: Radius.circular(40),
                  topRight: Radius.circular(40),
                ),
              ),
              child: ListView(
                padding: const EdgeInsets.only(
                  left: 15,
                  right: 15,
                ),
                children: const [
                  user_conversation_widget(name: 'Aryam',),
                  user_conversation_widget(name: 'Afomia',),
                  user_conversation_widget(name: 'Daniel',),
                  user_conversation_widget(name: 'Makda',),
                  user_conversation_widget(name: 'Aryam',),
                  user_conversation_widget(name: 'Afomia',),
                  user_conversation_widget(name: 'Daniel',),
                  user_conversation_widget(name: 'Makda',),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}

