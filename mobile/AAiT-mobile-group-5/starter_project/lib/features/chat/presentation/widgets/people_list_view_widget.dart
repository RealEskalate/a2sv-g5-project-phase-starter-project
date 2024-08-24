import 'package:flutter/material.dart';
import 'package:starter_project/features/chat/presentation/message_entity_template.dart';
import 'package:starter_project/features/chat/presentation/widgets/static_card.dart';

const message1 = Message(
  sender: "Alex Linderson",
  timeStamp: 2,
  unRead: 3,
  body: "How are you today",
);
const messages = [
  message1,
  message1,
  message1,
  message1,
  message1,
  message1,
  message1,
];

class PeopleListWidget extends StatelessWidget {
  const PeopleListWidget({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.fromLTRB(0, 20, 0, 0),
      margin: const EdgeInsets.fromLTRB(0, 150, 0, 0),

      // ),
      height: 550,
      decoration: const BoxDecoration(
        borderRadius: BorderRadius.only(
          topLeft: Radius.circular(70),
          topRight: Radius.circular(70),
          bottomLeft: Radius.circular(0),
          bottomRight: Radius.circular(0),
        ),
        color: Colors.white,
      ),
      child: Column(
        children: [
          Padding(
            padding: const EdgeInsets.only(
              top: 10,
              bottom: 10,
            ),
            child: Center(
              child: MouseRegion(
                cursor: SystemMouseCursors.click,
                child: InkWell(
                  hoverColor: Colors.blue.withOpacity(0.1),
                  child: Container(
                    height: 5,
                    width: 60,
                    margin: const EdgeInsets.fromLTRB(45, 20, 0, 20),
                    decoration: BoxDecoration(
                      borderRadius: BorderRadius.circular(11),
                      border: Border.all(
                        color: const Color.fromARGB(255, 204, 201, 201),
                        width: 2.0,
                      ),
                      color: const Color.fromARGB(255, 204, 201, 201),
                    ),
                  ),
                ),
              ),
            ),
          ),
          Expanded(
            child: SingleChildScrollView(
              scrollDirection: Axis.vertical,
              // physics: const AlwaysScrollableScrollPhysics(),
              child: Column(
                children: messages.map((mess) {
                  return StaticCard(message: mess);
                }).toList(),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
