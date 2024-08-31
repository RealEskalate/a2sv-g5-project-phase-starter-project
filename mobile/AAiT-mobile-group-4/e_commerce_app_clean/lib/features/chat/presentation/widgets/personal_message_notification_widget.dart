import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../authentication/domain/entities/user_data.dart';
import '../blocs/bloc/chat_bloc.dart';
import 'profile_pic_widget.dart';
class PersonalMessageNotification extends StatelessWidget {
  final String? fullName;
  final String? message;
  final String? timeSent;
  final bool? isRead;
  final bool? isOnline;
  final int? unreadMessages;
  final String imagePath;
  final Color bgColor;
  final UserEntity user;
  final String chatId;

  const PersonalMessageNotification({
    this.imagePath = 'assets/story_3.png',
    this.bgColor = Colors.blue,
    super.key,
    this.fullName,
    this.message,
    this.timeSent,
    this.isRead,
    this.isOnline,
    this.unreadMessages,
    required this.user,
    required this.chatId,
  });

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      //---------------------------------------------
      onTap: (){
        Navigator.of(context).pushNamed('/messages', arguments: user);
        BlocProvider.of<ChatBloc>(context).add(GetChatMessagesEvent(chatId: chatId));
      },
      //---------------------------------------------
      child: Container(
        padding: const EdgeInsets.symmetric(vertical: 10, horizontal: 15),
        child: Row(
          crossAxisAlignment:
              CrossAxisAlignment.center, // Align items to the center vertically
          children: [
            Stack(
              children: [
                ProfilePicWidget(
                  imagePath: imagePath,
                  bgColor: bgColor,
                  radius: 25, // Increased radius for a bigger profile picture
                ),
                Positioned(
                  bottom: 2,
                  right: 2,
                  child: CircleAvatar(
                    radius: 5,
                    backgroundColor: Colors.white,
                    child: CircleAvatar(
                      radius: 4,
                      backgroundColor:
                          isOnline == true ? Colors.green : Colors.grey,
                    ),
                  ),
                ),
              ],
            ),
            const SizedBox(width: 10),
            Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    fullName ?? 'Unknown',
                    style: const TextStyle(
                      fontWeight: FontWeight.bold,
                      fontSize: 16,
                      color: Colors.black,
                    ),
                  ),
                  const SizedBox(height: 2),
                  Text(
                    message ?? 'No message',
                    style: TextStyle(
                      color: isRead == true ? Colors.grey : Colors.black,
                      fontSize: 14,
                    ),
                    overflow: TextOverflow.ellipsis,
                    maxLines: 1,
                  ),
                ],
              ),
            ),
            Column(
              crossAxisAlignment: CrossAxisAlignment.end,
              children: [
                Text(
                  timeSent ?? 'Just now',
                  style: const TextStyle(
                    color: Colors.grey,
                    fontSize: 12,
                  ),
                ),
                const SizedBox(height: 5),
                if ((unreadMessages ?? 0) > 0)
                  Container(
                    padding:
                        const EdgeInsets.symmetric(horizontal: 6, vertical: 2),
                    decoration: BoxDecoration(
                      color: Colors.blue, // Badge color
                      borderRadius: BorderRadius.circular(12),
                    ),
                    child: Text(
                      (unreadMessages ?? 0).toString(),
                      style: const TextStyle(
                        color: Colors.white,
                        fontSize: 12,
                        fontWeight: FontWeight.bold,
                      ),
                    ),
                  ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
