import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:starter_project/features/chat/presentation/message_entity_template.dart';

class StaticCard extends StatefulWidget {
  final Message message;

  const StaticCard({Key? key, required this.message}) : super(key: key);

  @override
  State<StaticCard> createState() => _StaticCardState();
}

class _StaticCardState extends State<StaticCard> {
  bool isHovered = false;

  @override
  Widget build(BuildContext context) {
    return MouseRegion(
      onEnter: (_) => _onHover(true),
      onExit: (_) => _onHover(false),
      child: AnimatedContainer(
        duration: const Duration(milliseconds: 200),
        decoration: BoxDecoration(
          color: isHovered ? Color.fromARGB(26, 247, 244, 244) : Colors.white,
        ),
        child: Card(
          color: Colors.white,
          elevation: 0,
          margin: const EdgeInsets.only(left: 5, right: 5, bottom: 15),
          child: Padding(
            padding: const EdgeInsets.all(5.0),
            child: Row(
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                // First column with icon and sender + lastMessage
                Expanded(
                  flex: 3,
                  child: Row(
                    children: [
                      // Icon
                      Stack(
                        children: [
                          const CircleAvatar(
                            radius: 35,
                            child: Icon(
                              Icons.people_outline_sharp,
                              size: 25.0,
                            ),
                          ),
                          Positioned(
                            right: 0,
                            bottom: 15,
                            child: Container(
                              height: 10,
                              width: 10,
                              decoration: BoxDecoration(
                                borderRadius: BorderRadius.circular(20),
                                border: Border.all(
                                  color: Color.fromARGB(255, 155, 220, 220),
                                  width: 2.0,
                                ),
                                color: Color.fromARGB(255, 145, 221, 221),
                              ),
                            ),
                          )
                        ],
                      ),
                      const SizedBox(width: 10),

                      // Name and LastMessage column
                      Expanded(
                        child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Text(
                              widget.message.sender,
                              style: const TextStyle(
                                fontWeight: FontWeight.bold,
                                fontSize: 16.0,
                              ),
                            ),
                            const SizedBox(height: 5),
                            Text(
                              widget.message.body,
                              style: const TextStyle(
                                color: Colors.grey,
                                fontSize: 14.0,
                              ),
                              maxLines: 1,
                              overflow: TextOverflow.ellipsis,
                            ),
                          ],
                        ),
                      ),
                    ],
                  ),
                ),
                // Second column with timestamp and unread count
                Expanded(
                  flex: 1,
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.center,
                    // mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Text(
                        _formatTimeStamp(widget.message.timeStamp),
                        style: const TextStyle(
                          color: Colors.grey,
                          fontSize: 10.0,
                        ),
                      ),
                      const SizedBox(height: 5),
                      if (widget.message.unRead > 0)
                        CircleAvatar(
                          radius: 11,
                          backgroundColor: Colors.blue[700],
                          child: Text(
                            widget.message.unRead.toString(),
                            style: const TextStyle(
                              color: Colors.white,
                              fontSize: 11.0,
                            ),
                          ),
                        ),
                    ],
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }

  String _formatTimeStamp(int timeStamp) {
    // Replace this with actual time formatting logic
    final dateTime = DateTime.fromMillisecondsSinceEpoch(timeStamp);

    return '${dateTime.hour} minutes ago';
  }

  void _onHover(bool hover) {
    setState(() {
      isHovered = hover;
    });
  }
}
