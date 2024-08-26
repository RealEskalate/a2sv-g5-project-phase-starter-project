import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/message.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/widgets/chat_container.dart';
import 'package:flutter/material.dart';


class ChatBody extends StatelessWidget {
  final List<MessageType> chats;
  final String ownerId;
  final List<String> senderIds;

  ChatBody({
    required this.ownerId,
    required this.senderIds,
    this.chats = const [],
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.start,
      children: [
        Expanded(
          child: ListView.builder(
            itemCount: chats.length,
            itemBuilder: (context, index) {
              bool isOwner = ownerId == senderIds[index];
              return Align(
                alignment: isOwner ? Alignment.centerRight : Alignment.centerLeft,
                child: Container(
                  padding: EdgeInsets.symmetric(vertical: 8.0, horizontal: 12.0),
                  margin: EdgeInsets.symmetric(vertical: 5.0),
                  decoration: BoxDecoration(
                    color: isOwner ? Colors.blue : Colors.grey[300],
                    borderRadius: BorderRadius.circular(12.0),
                  ),
                  child: SingleText(chat: chats[index].content, parentColor: isOwner ? 'blue' : 'grey'),
                ),
              );
            },
          ),
        ),
      ],
    );
  }
}
