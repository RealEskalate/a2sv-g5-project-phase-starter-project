import 'dart:math';

import 'package:ecommerce_app_ca_tdd/features/chat/data/models/chat_models.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/chat_entity.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/message.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/bloc/chat_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/bloc/chat_event.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/bloc/chat_state.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/message_bloc/message_bloc_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/message_bloc/message_bloc_event.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/message_bloc/message_bloc_state.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/socket_n/chatbox.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/models/seller_model.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/pages/HomeChat.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/widgets/chat_appbar.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/widgets/chat_body.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/widgets/chat_bottom_appbar.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:get/get_navigation/get_navigation.dart';
import 'package:shared_preferences/shared_preferences.dart';



class ChatPage extends StatelessWidget {
  final TextEditingController _messageController = TextEditingController();
  // final ChatEntity chat;
  // Dummy data for chat
  final String ownerId = '1';
  final List<MessageType> chats = [
  ];

  final List<String> senderIds = [
    '2', // Sender 1
    '1', // Owner
    '2',
    '1',
    '2',
    '1',
    '2',
    '1',
    '2',
    '1',
  ];

  ChatPage({super.key});
  @override
  Widget build(BuildContext context) {
    final route = ModalRoute.of(context);
    if (route == null || route.settings.arguments == null) {
      return Scaffold(
        body: Center(
          child: Text("No chat Availbale"),
        ),
      );
    }
    final arguments = route.settings.arguments as Map<String, dynamic>;
    print(arguments);
    final SellerModel seller = arguments[''];
    final String chatId = arguments['chatId'];

    
    // context.read<MessageBloc>().add(MessageConnection(chat));
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      home: Scaffold(
          appBar:
              ChatAppBar(seller.name, 'last seen yesterday', "image.jpg"),
          body: Container(
            color: Colors.white,
            child: Column(
              children: [
                // ChatBody to display the chat messages
                Expanded(
                  child: ChatBody(
                    ownerId: ownerId,
                    senderIds: senderIds,
                    chats: chats,
                  ),
                ),
    
                ChatBottomAppBar(
                  messageController: _messageController,
                  onSend: () {
                    BlocListener<ChatBloc,ChatState>(
                      listener: (context, state){
                        if (state is ChatInitateLoaded){
                          print('Chat initiated');
                          chats.add(MessageType(content: _messageController.text,type: 'text'));
                          context.read<MessageBloc>().add(
                            MessageSent(
                              state.chat.chatid,
                              _messageController.text,
                              'text',
                            ),
                          );
    
                        }
                      }
                      );
                      _messageController.clear();
                    
                    print('Send button pressed');
                  },
                  onFile: () {
                    print('File button pressed');
                  },
                  onCamera: () {
                    print('Camera button pressed');
                  },
                  onGallery: () {
                    print('Gallery button pressed');
                  },
                  onAudio: () {
                    print('Audio button pressed');
                  },
                ),
              ],
            ),
          )),
    );
  }
}
void showError(BuildContext context, String message) {
  ScaffoldMessenger.of(context).showSnackBar(
    SnackBar(
      content: Text(message),
      backgroundColor: Theme.of(context).colorScheme.error,
    ),
  );
}
Future<String> getSenderId() async {
  var prefs = await SharedPreferences.getInstance();
  return prefs.getString('id')!;
}

