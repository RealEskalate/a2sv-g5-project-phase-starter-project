import 'package:flutter/material.dart';

import '../widgets/stories_widget.dart';

class ChatPage extends StatelessWidget {
  const ChatPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color.fromRGBO(73, 140, 240, 1),
      body: Padding(
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
                
                child: const Center(child: Text('messages go here')),

                // ------------------------------------------
              ),
            ),
          ],
        ),
      ),
    );
  }
}
