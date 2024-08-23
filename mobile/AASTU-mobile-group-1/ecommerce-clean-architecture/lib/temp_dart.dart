import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';

import 'features/chat_feature/chat/domain/usecase/initialize_chat.dart';

class TempDart extends StatelessWidget {
  InitializeChat initializeChat;
  TempDart({required this.initializeChat, super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: ElevatedButton(
            onPressed: () async {
              var chatroom =
                  await initializeChat.call('66c5dc6700781697dc798f44');
              print("Chatroom: $chatroom");
            },
            child: Text('TempDart')),
      ),
    );
  }
}
