import 'package:flutter/material.dart';

class ChatListTile extends StatefulWidget {
  ChatListTile({super.key});

  @override
  State<ChatListTile> createState() => _ChatListTileState();
}

class _ChatListTileState extends State<ChatListTile> {
  bool isOnline = true;

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Row(mainAxisAlignment: MainAxisAlignment.spaceBetween, children: [
          Row(mainAxisAlignment: MainAxisAlignment.start, children: [
            Stack(
              children: [
                const CircleAvatar(
                  radius: 30,
                  backgroundImage: AssetImage('assets/images/user.png'),
                ),
                Positioned(
                    left: 50,
                    top: 47,
                    child: CircleAvatar(
                      radius: 5,
                      backgroundColor: isOnline ? Colors.green : Colors.grey,
                    ))
              ],
            ),
            SizedBox(
              width: 20,
            ),
            const Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  'User Name',
                  style: TextStyle(
                    color: Color.fromARGB(255, 17, 16, 16),
                    fontSize: 20,
                  ),
                ),
                Text(
                  'Last Message',
                  style: TextStyle(
                    color: Color.fromARGB(255, 94, 90, 90),
                    fontSize: 13,
                  ),
                ),
              ],
            ),
          ]),
          const SizedBox(
            width: 10,
          ),
          const Column(
            children: [
              Text(
                '2 min ago',
                style: TextStyle(
                  color: Color.fromARGB(255, 149, 145, 145),
                  fontSize: 12,
                ),
              ),
              SizedBox(
                height: 10,
              ),
              CircleAvatar(
                radius: 10,
                backgroundColor: Color.fromARGB(255, 66, 104, 217),
                child: Text(
                  '2',
                  style: TextStyle(
                    color: Colors.white,
                    fontSize: 10,
                  ),
                ),
              ),
            ],
          ),
        ]),
        SizedBox(
          height: 20,
        ),
      ],
    );
  }
}
