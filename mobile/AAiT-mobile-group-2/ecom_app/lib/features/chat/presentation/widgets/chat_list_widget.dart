import 'package:flutter/material.dart';

class SingleChat extends StatelessWidget {
  final String name;
  final String message;
  final String imageUrl;
  final int time;
  final int notificationCount;
  final bool isOnline;
  const SingleChat(
      {super.key,
      required this.name,
      required this.message,
      required this.imageUrl,
      required this.time,
      required this.notificationCount,
      required this.isOnline});
  Positioned status(isOnline) {
    if (isOnline) {
      return const Positioned(
        right: 0,
        bottom: 0,
        child: Icon(
          Icons.circle,
          color: Color.fromRGBO(15, 225, 109, 1),
          size: 13,
        ),
      );
    }
    return const Positioned(
      right: 0,
      bottom: 0,
      child: Icon(
        Icons.circle,
        color: Color.fromRGBO(139, 144, 141, 1),
        size: 13,
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          crossAxisAlignment: CrossAxisAlignment.end,
          children: [
            Row(
              crossAxisAlignment: CrossAxisAlignment.end,
              children: [
                Stack(
                  children: [
                    CircleAvatar(
                      radius: 26,
                      backgroundImage: NetworkImage(
                          imageUrl), // Image of the user who sent the message
                      backgroundColor: Colors.transparent,
                    ),
                    status(isOnline),
                  ],
                ),
                const SizedBox(
                  width: 7,
                ),
                Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      name,
                      style: const TextStyle(
                          fontSize: 15, fontWeight: FontWeight.w500),
                    ),
                    Text(message, style: const TextStyle(fontSize: 10)),
                  ],
                ),
              ],
            ),
            Column(
              children: [
                Text(
                  '$time min ago',
                  style: const TextStyle(fontSize: 12),
                ),
                Container(
                    padding: const EdgeInsets.fromLTRB(7, 0, 0, 0),
                    width: 22,
                    height: 22,
                    decoration: const BoxDecoration(
                      color: Color.fromRGBO(58, 107, 232, 1),
                      shape: BoxShape.circle,
                    ),
                    child: Text(
                      '$notificationCount',
                      style: const TextStyle(
                          fontFamily: 'AbhayaLibre', color: Colors.white),
                    ))
              ],
            )
          ],
        ),
        const SizedBox(
          height: 30,
        )
      ],
    );
  }
}
