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



class ChatPage extends StatefulWidget {
  final ChatEntity chat;


  ChatPage({super.key, required this.chat});

  @override
  State<ChatPage> createState() => _ChatPageState();
}

class _ChatPageState extends State<ChatPage> {
  final TextEditingController _messageController = TextEditingController();

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

  @override
  Widget build(BuildContext context) {
    final route = ModalRoute.of(context);

    final String seller = widget.chat.user2.name;
    final chatID = widget.chat.chatid;

    
    // context.read<MessageBloc>().add(MessageConnection(chat));
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      home: Scaffold(
          appBar:
              ChatAppBar(seller, 'last seen yesterday',''),
          body: Container(
            color: Colors.white,
            child: Column(
              children: [
                Expanded(
            child:  ChatBody(ownerId: widget.chat.user2.id,senderIds: [widget.chat.user2.id,widget.chat.user1.id],)
              
            
          ),
                ChatBottomAppBar(
                  messageController: _messageController,
                  onSend: () {
                    setState(() {
                      if (_messageController.text.isNotEmpty){
                        chats.add(MessageType(content: _messageController.text,type: 'text'));
                      }
                      
                    });
                    context.read<MessageBloc>().add(MessageSent(chatID, _messageController.text, 'text'));
                    BlocListener<MessageBloc,MessageState>(
                      listener: (context, state){
                        if (state is MessageInitial){
                         setState(() {
                            if (_messageController.text.isNotEmpty){
                              chats.add(MessageType(content: _messageController.text,type: 'text'));
                            }
                            
                          });
                          
    
                        }
                      }
                      );
                      

                    
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

