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
                  height: 100, // Adjust height to accommodate both profile and name
                  child: AvatarWithName(
                    names: ['Aryam', 'Afomia', 'Daniel', 'Makda', 'Aryam', 'Afomia', 'Daniel', 'Makda'],
                  ),
                ),
              ),
            ],
          ),
          Positioned(
            top: 200,
            bottom: 0,
            child: Container(
              
              padding: const EdgeInsets.only(top: 15),
              height: 568,
              width: MediaQuery.of(context).size.width,
              decoration: const BoxDecoration(
                color: Colors.white,
                borderRadius: BorderRadius.only(
                  topLeft: Radius.circular(45),
                  topRight: Radius.circular(45),
                ),
              ),
              child: ListView(
                padding: const EdgeInsets.only(
                  left: 15,
                  right: 15,
                  
                ),
                children: const [
                  UserConversationWidget(name: 'Aryam',),
                  UserConversationWidget(name: 'Afomia',),
                  UserConversationWidget(name: 'Daniel',),
                  UserConversationWidget(name: 'Makda',),
                  UserConversationWidget(name: 'Aryam',),
                  UserConversationWidget(name: 'Afomia',),
                  UserConversationWidget(name: 'Daniel',),
                  UserConversationWidget(name: 'Makda',),
                  UserConversationWidget(name: 'Makda',),
                  UserConversationWidget(name: 'Makda',),
                  UserConversationWidget(name: 'Makda',),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}
