import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../domain/entities/message.dart';
import '../blocs/bloc/chat_bloc.dart';

class MessageBox extends StatelessWidget {
  const MessageBox({super.key});

  @override
  Widget build(BuildContext context) {
    final TextEditingController messageController = TextEditingController();
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 10),
      height: 90,
      color: Colors.white,
      width: double.infinity,
      child: Center(
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Transform.rotate(
              angle: 0.5,
              child: const Icon(Icons.attach_file, color: Color.fromARGB(255, 14, 8, 1),)
            ),
            const SizedBox(width: 5),
            Expanded(
              child: Container(
                decoration: BoxDecoration(
                  color: const Color.fromARGB(255, 243, 246, 246),
                  borderRadius: BorderRadius.circular(12),
                ),
                child: TextField(
                  controller: messageController,
                  decoration: InputDecoration(
                    hintText: 'Type a message',
                    hintStyle: const TextStyle(
                      color: Color.fromARGB(255, 121, 124, 123),
                    ),
                    contentPadding: const EdgeInsets.symmetric(horizontal: 8, vertical: 15),
                    border: InputBorder.none,
                    suffixIcon: IconButton(
                      icon: const Icon(Icons.content_copy),
                      color: const Color.fromARGB(255, 121, 124, 123),
                      onPressed: () {
                        if(messageController.text.isEmpty) {
                          return;
                        }
                        var message = Message(content: messageController.text, type: MessageType.text);
                        BlocProvider.of<ChatBloc>(context).add(SendMessageEvent(message: message));
                      },
                    ),
                  ),
                ),
              ),
            ),
            IconButton(
              icon: const Icon(Icons.camera_alt_outlined),
              onPressed: () {},
              color: const Color.fromARGB(255, 14, 8, 1),
            ),
            IconButton(
              icon: const Icon(Icons.mic_rounded),
              color: const Color.fromARGB(255, 14, 8, 1),
              onPressed: () {},
            ),
          ],
        ),
      ),
    );
  }
}
  