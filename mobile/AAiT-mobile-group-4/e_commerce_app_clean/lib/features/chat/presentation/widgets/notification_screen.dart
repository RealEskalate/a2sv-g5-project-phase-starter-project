import 'dart:math';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../authentication/domain/entities/user_data.dart';
import '../../../authentication/presentation/bloc/auth_bloc.dart';
import '../../domain/entities/chat.dart';
import '../blocs/bloc/chat_bloc.dart';
import 'personal_message_notification_widget.dart';
import 'stories_widget.dart';

class ChatPage extends StatelessWidget {
  const ChatPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        color: const Color.fromRGBO(73, 140, 240, 1),
        child: Padding(
          padding: const EdgeInsets.only(top: 36),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const IconButton(
                padding: EdgeInsets.only(left: 20),
                onPressed: null,
                iconSize: 30,
                icon: Icon(Icons.search, color: Colors.white),
              ),
              const StoriesWidget(),
              const SizedBox(height: 20),
              Expanded(
                child: Container(
                  height: MediaQuery.of(context).size.height,
                  decoration: const BoxDecoration(
                    color: Colors.white,
                    borderRadius: BorderRadius.only(
                      topLeft: Radius.circular(30),
                      topRight: Radius.circular(30),
                    ),
                  ),
                  // content goes here
                  // ------------------------------------------
                  child: BlocBuilder<AuthBloc, AuthState>(
                    builder: (context, state) {
                      UserEntity user = (state as AuthUserLoaded).userEntity;
                      return BlocBuilder<ChatBloc, ChatState>(
                        builder: (context, state) {
                          if (state is ChatsLoadedState) {
                            List<Chat> chats = state.chats; 
                            return ListView.builder(
                              itemCount: chats.length,
                              itemBuilder: (context, index) {
                                UserEntity sender = chats[index].user1.id == user.id ? chats[index].user2 :chats[index].user1;
                                return PersonalMessageNotification(
                                  chatId: chats[index].id,
                                  user: sender,
                                  imagePath: 'assets/story_${(index % 3) + 1}.png',
                                  bgColor: Color((Random().nextDouble() * 0xFFFFFF).toInt() << 0).withOpacity(1.0),
                                  fullName: sender.name,
                                  message: sender.email,
                                );
                              });
                          } else if ( state is ChatsLoadingState ) {
                            return const CircularProgressIndicator();
                          } else {
                            return Center(child: Text('Unkown State${state.runtimeType}'));
                          }
                        },
                      );
                    },
                  ),
                  // ------------------------------------------
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
