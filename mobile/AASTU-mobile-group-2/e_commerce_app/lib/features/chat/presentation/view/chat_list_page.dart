import 'package:e_commerce_app/features/chat/presentation/view/chat_list_widget.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'dart:math' as math;

import '../bloc/chat_bloc.dart';

class ChatPeopleList extends StatefulWidget {
  const ChatPeopleList({super.key});

  @override
  State<ChatPeopleList> createState() => _ChatPeopleListState();
}

class _ChatPeopleListState extends State<ChatPeopleList>
    with SingleTickerProviderStateMixin {
  late AnimationController _controller;

  @override
  void initState() {
    super.initState();
    _controller = AnimationController(
      vsync: this,
      duration: const Duration(seconds: 5),
    )..repeat();
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final screenWidth = MediaQuery.of(context).size.width;
    final screenHeight = MediaQuery.of(context).size.height;

    return Scaffold(
      backgroundColor: Colors.blue,
      body: SafeArea(
        child: Column(
          children: [
            Padding(
              padding: const EdgeInsets.fromLTRB(20, 10, 20, 0),
              child: AppBar(
                backgroundColor: Colors.blue,
                elevation: 0,
                automaticallyImplyLeading: false,
                leading: IconButton(
                  icon: const Icon(Icons.search_outlined),
                  color: Colors.white,
                  iconSize: screenWidth * 0.07,
                  onPressed: () {
                    // To be implemented
                  },
                ),
              ),
            ),
            Expanded(
              child: Stack(
                children: [
                  Positioned(
                    top: screenHeight * 0.18,
                    left: 0,
                    right: 0,
                    child: Container(
                      height: screenHeight * 0.85,
                      width: double.infinity,
                      decoration: const BoxDecoration(
                        color: Colors.white,
                        borderRadius: BorderRadius.only(
                          topLeft: Radius.circular(50),
                          topRight: Radius.circular(50),
                        ),
                      ),
                      child: Padding(
                        padding: const EdgeInsets.all(20),
                        child: BlocConsumer<ChatBloc, ChatState>(
                          listener: (context, chatState) {
                            // TODO: implement listener
                          },
                          builder: (context, chatState) {
                            // return ListView.builder(
                            //   itemCount: 20,
                            //   itemBuilder: (context, index) {
                            //     return chatListItem(
                            //       context,
                            //       imageName: 'smile.png',
                            //       title: 'Contact $index',
                            //       subtitle: 'Hey, how are you?',
                            //       time: '12:0${index} PM',
                            //       unreadMessages: index % 3,
                            //     );
                            //   },
                            // );
                            if (chatState is LoadingCurrentChats) {
                              return const Center(
                                  child: CircularProgressIndicator());
                            } else if (chatState is CurrentChatsLoaded) {
                              final chatListItems = chatState.chats.map((chat) {
                                return chatListItem(context,
                                  
                                  chat: chat,
                                  
                                );
                              }).toList();

                              return ListView(
                                children: chatListItems,
                              );
                            } else {
                              return const Center(
                                  child: Text('No chats found'));
                            }
                          },
                        ),
                      ),
                    ),
                  ),
                  Padding(
                    padding: const EdgeInsets.all(20),
                    child: Column(
                      children: [
                        Container(
                          height: screenHeight * 0.2,
                          child: Row(
                            children: [
                              Column(
                                children: [
                                  Stack(
                                    alignment: Alignment.center,
                                    children: [
                                      AnimatedBuilder(
                                        animation: _controller,
                                        builder: (context, child) {
                                          return Transform.rotate(
                                            angle:
                                                _controller.value * 2 * math.pi,
                                            child: ShaderMask(
                                              shaderCallback: (rect) {
                                                return SweepGradient(
                                                  colors: [
                                                    Colors.blue[200]!,
                                                    const Color.fromRGBO(
                                                        199, 254, 224, 1),
                                                    Colors.blue[200]!,
                                                  ],
                                                  stops: [0.0, 0.10, 5.0],
                                                ).createShader(rect);
                                              },
                                              child: CircleAvatar(
                                                radius: screenWidth * 0.109,
                                                backgroundColor: Colors.white,
                                              ),
                                            ),
                                          );
                                        },
                                      ),
                                      CircleAvatar(
                                        radius: screenWidth * 0.1,
                                        backgroundColor: Colors.blue,
                                      ),
                                      CircleAvatar(
                                        radius: screenWidth * 0.086,
                                        backgroundImage: const AssetImage(
                                            'assets/smile.png'),
                                      ),
                                      Positioned(
                                        bottom: 0,
                                        right: 0,
                                        left: 52,
                                        child: Container(
                                          width: screenWidth * 0.3,
                                          height: screenHeight * 0.026,
                                          decoration: BoxDecoration(
                                            color: Colors.green,
                                            shape: BoxShape.circle,
                                            border: Border.all(
                                              color: Colors.white,
                                              width: 2,
                                            ),
                                          ),
                                        ),
                                      ),
                                    ],
                                  ),
                                  const SizedBox(height: 5),
                                  const Text(
                                    "My Status",
                                    style: TextStyle(color: Colors.white),
                                  ),
                                ],
                              ),
                              SizedBox(width: screenWidth * 0.02),
                              Expanded(
                                child: SingleChildScrollView(
                                  scrollDirection: Axis.horizontal,
                                  child: Row(
                                    children: List.generate(10, (index) {
                                      return Padding(
                                        padding: const EdgeInsets.symmetric(
                                            horizontal: 5.0),
                                        child: Column(
                                          children: [
                                            Stack(
                                              alignment: Alignment.center,
                                              children: [
                                                AnimatedBuilder(
                                                  animation: _controller,
                                                  builder: (context, child) {
                                                    return Transform.rotate(
                                                      angle: _controller.value *
                                                          2 *
                                                          math.pi,
                                                      child: ShaderMask(
                                                        shaderCallback: (rect) {
                                                          return SweepGradient(
                                                            colors: [
                                                              Colors.blue[200]!,
                                                              const Color
                                                                  .fromRGBO(199,
                                                                  254, 224, 1),
                                                              Colors.blue[200]!,
                                                            ],
                                                            stops: [
                                                              0.0,
                                                              0.10,
                                                              5.0
                                                            ],
                                                          ).createShader(rect);
                                                        },
                                                        child: CircleAvatar(
                                                          radius: screenWidth *
                                                              0.109,
                                                          backgroundColor:
                                                              Colors.white,
                                                        ),
                                                      ),
                                                    );
                                                  },
                                                ),
                                                CircleAvatar(
                                                  radius: screenWidth * 0.1,
                                                  backgroundColor: Colors.blue,
                                                ),
                                                CircleAvatar(
                                                  radius: screenWidth * 0.086,
                                                  backgroundImage:
                                                      const AssetImage(
                                                          'assets/smile.png'),
                                                ),
                                              ],
                                            ),
                                            const SizedBox(height: 5),
                                            const Text(
                                              'Names',
                                              style: TextStyle(
                                                  color: Colors.white),
                                            ),
                                          ],
                                        ),
                                      );
                                    }),
                                  ),
                                ),
                              ),
                            ],
                          ),
                        ),
                      ],
                    ),
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
