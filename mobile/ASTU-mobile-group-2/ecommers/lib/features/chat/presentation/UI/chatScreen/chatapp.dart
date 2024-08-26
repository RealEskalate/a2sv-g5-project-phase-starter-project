import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:get_it/get_it.dart';

import 'package:shared_preferences/shared_preferences.dart';

import '../../../../../core/utility/global_message_part.dart';
import '../../../../../core/utility/socket_impl.dart';
import '../../../domain/entity/chat_entity.dart';
import '../../bloc/chat_bloc.dart';
import '../../bloc/chat_event.dart';
import '../../bloc/chat_state.dart';
import '../../bloc/socket/socket_bloc.dart';
import '../../bloc/socket/socket_state.dart';

class ChatListPage extends StatefulWidget {
  @override
  _ChatListPageState createState() => _ChatListPageState();
}

class _ChatListPageState extends State<ChatListPage> {
  String? userId;
  final List<String> profilePhotos = [
    'assets/image/pro1.jpg',
    'assets/image/pro2.jpg',
    'assets/image/pro3.jpg',
    'assets/image/pro4.jpg',
    'assets/image/grouppro.png',
    'assets/image/grouppro2.png',
    'assets/image/pro5.png',
    'assets/image/pro6.jpg',
    'assets/image/pro7.jpg',
  ];
  List<ChatEntity> chatEntity = [];
  late SocketService _socketProvider;
  @override
  void initState() {
    super.initState();
    _socketProvider = GetIt.instance<SocketService>();

    _loadUserId();
  }

  Future<void> _loadUserId() async {
    SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
    setState(() {
      userId = sharedPreferences.getString('user_id');
    });
  }

  @override
  void dispose() {
    // Dispose of any subscriptions or listeners here
    // Add this if your SocketService has a dispose method
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return BlocListener<SocketBloc, SocketState>(
      listener: (context, state) {
         if (state is SocketMessageReceived) {
            final messageData = state.messageData;
            // Display the received message in your chat interface
            Text('Message from ${messageData['senderId']}: ${messageData['content']}');
          } else {
            // Handle other states or default UI
            const Text('No new messages');
          }
      },
      child: Scaffold(
        backgroundColor: Colors.blue,
        appBar: AppBar(
          backgroundColor: Colors.blue,
          leading: IconButton(
            icon: const Icon(Icons.arrow_back_ios),
            onPressed: () {
              // Handle the search icon press
              Navigator.pop(context);
            },
          ),
        ),
        body: RefreshIndicator(onRefresh: () {
          context.read<ChatBloc>().add(OnGetAllChat());
          return Future.delayed(const Duration(seconds: 1));
        }, child: BlocBuilder<SocketBloc, SocketState>(
          builder: (context, state) {
            _socketProvider.listen('message:delivered', (data) {
              print('Message delivered to recipient: $data');
            
            });

           
            return BlocListener<ChatBloc, ChatState>(
              listener: (context, chatState) {
                if (chatState is ChatLoadingState) {
                  const Center(child: CircularProgressIndicator());
                }
              },
              child: Column(
                children: [
                  Container(
                    padding: const EdgeInsets.symmetric(horizontal: 3),
                    width: MediaQuery.of(context).size.width,
                    height: 150,
                    color: Colors.blue,
                    child: ListView(
                      scrollDirection: Axis.horizontal,
                      children: [
                        Mystory('My status', 'assets/image/pro.png'),
                        story('Marina', 'assets/image/pro1.jpg'),
                        story('Dean', 'assets/image/pro2.jpg'),
                        story('Max', 'assets/image/pro3.jpg'),
                        story('My status', 'assets/image/pro4.jpg'),
                        story('Adil', 'assets/image/pro5.png'),
                        story('Marina', 'assets/image/pro6.jpg'),
                        story('Dean', 'assets/image/pro7.jpg'),
                        story('My status', 'assets/image/pro.jpg'),
                        story('Adil', 'assets/image/grouppro.png'),
                        story('Marina', 'assets/image/grouppro2.png'),
                      ],
                    ),
                  ),
                  Expanded(
                    child: Container(
                      decoration: const BoxDecoration(
                          color: Colors.white,
                          borderRadius: BorderRadius.only(
                              topLeft: Radius.circular(30),
                              topRight: Radius.circular(30))),
                      child: BlocBuilder<ChatBloc, ChatState>(
                        builder: (context, state) {
                          if (state is ChatMessageGetSuccess) {
                            chatEntity = state.chatEntity;
                          }
                          return ListView.builder(
                            itemCount: chatEntity.length,
                            itemBuilder: (context, index) {
                              final current = chatEntity[index];

                              final String nameOfUser =
                                  userId == current.recieverId
                                      ? current.senderName
                                      : current.recieverName;

                              final List<dynamic> messages = GlobalMessagePart
                                      .gloablMessage[current.chatId] ??
                                  [];
                              final text = messages.isNotEmpty &&
                                      messages.last['content'] != null
                                  ? messages.last['content']
                                  : 'say hi to $nameOfUser';

                              return GestureDetector(
                                onTap: () {
                                  print(userId);
                                  Navigator.pushNamed(context, '/chat-message',
                                      arguments: {
                                        'chatId': current.chatId,
                                        'name': nameOfUser,
                                        'itMe': userId
                                      });
                                },
                                child: chatTileWithReadMessages(
                                  nameOfUser,
                                  text,
                                  '02:20',
                                  index,
                                ),
                              );
                            },
                          );
                        },
                      ),
                    ),
                  ),
                ],
              ),
            );
          },
        )),
      ),
    );
  }

  Widget chatTileWithReadMessages(
      String name, String message, String time, int index) {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 5),
      child: ListTile(
        leading: CircleAvatar(
          backgroundImage:
              AssetImage(profilePhotos[index % profilePhotos.length]),
          radius: 30,
        ),
        title: Text(
          name,
          style: const TextStyle(
            fontWeight: FontWeight.bold,
          ),
        ),
        subtitle: Text(message),
        trailing: Text(time,
            style: const TextStyle(fontSize: 12, color: Colors.grey)),
      ),
    );
  }

  Widget Mystory(String name, String image) {
    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: Column(
        children: [
          Stack(
            children: [
              Container(
                width: 60,
                height: 60,
                decoration: BoxDecoration(
                  shape: BoxShape.circle,
                  border: Border.all(
                    color: Colors.blue,
                    width: 3,
                  ),
                ),
                child: CircleAvatar(
                  backgroundImage: AssetImage(image),
                ),
              ),
              Positioned(
                bottom: 0,
                right: 0,
                child: Container(
                  width: 20,
                  height: 20,
                  decoration: const BoxDecoration(
                    color: Colors.blue,
                    shape: BoxShape.circle,
                  ),
                  child: const Icon(
                    Icons.add,
                    size: 14,
                    color: Colors.white,
                  ),
                ),
              ),
            ],
          ),
          const SizedBox(height: 8),
          Text(
            name,
            style: const TextStyle(
              color: Colors.white,
              fontWeight: FontWeight.bold,
              fontSize: 12,
            ),
          ),
        ],
      ),
    );
  }

  Widget story(String name, String image) {
    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: Column(
        children: [
          Container(
            width: 60,
            height: 60,
            decoration: BoxDecoration(
              shape: BoxShape.circle,
              border: Border.all(
                color: Colors.blue,
                width: 3,
              ),
            ),
            child: CircleAvatar(
              backgroundImage: AssetImage(image),
            ),
          ),
          const SizedBox(height: 8),
          Text(
            name,
            style: const TextStyle(
              color: Colors.white,
              fontSize: 12,
            ),
          ),
        ],
      ),
    );
  }
}
