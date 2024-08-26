import 'package:custom_clippers/custom_clippers.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'dart:math' as math;

import '../../domain/entity/chat.dart';
import '../bloc/chat_bloc.dart';
import '../bloc/chat_event.dart';
import '../bloc/chat_state.dart';
import '../widgets/messages.dart';
import '../widgets/reusable_textfield.dart';

class ChatRoom extends StatefulWidget {
  static const String routes = '/chat_room';

  const ChatRoom({super.key});

  @override
  State<ChatRoom> createState() => _ChatRoomState();
}

class _ChatRoomState extends State<ChatRoom> {
  @override
  void initState() {
    final chatBloc = BlocProvider.of<ChatBloc>(context);

    chatBloc.add(ConnectServerEvent());
    chatBloc.add(SendMessage('', '', ''));
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    TextEditingController messageController = TextEditingController();

    final args = ModalRoute.of(context)!.settings.arguments as ChatEntity;

    return Scaffold(
      appBar: AppBar(
        automaticallyImplyLeading: false,
        backgroundColor: Colors.white,
        title: Row(
          children: [
            IconButton(
              icon: const Icon(CupertinoIcons.chevron_back),
              onPressed: () {
                final chatBloc = BlocProvider.of<ChatBloc>(context);

                chatBloc.add(LoadChatRooms());
                Navigator.pop(context);
              },
            ),
            Container(
              height: 50,
              width: 50,
              decoration: const BoxDecoration(
                color: Color(0xffC7FEE0),
                shape: BoxShape.circle,
              ),
              child: ClipRRect(
                borderRadius: BorderRadius.circular(50),
                child: Image.asset('assets/images/Memoji Boys 6-18.png'),
              ),
            ),
            Padding(
              padding: const EdgeInsets.only(left: 18.0),
              child: Column(
                children: [
                  Text(
                    args.user2.name,
                    // "Sabila Sayma",
                    style: const TextStyle(
                        fontWeight: FontWeight.w500, fontSize: 18),
                  ),
                  const Text(
                    '8 members, 5 online',
                    style: TextStyle(
                        fontWeight: FontWeight.w400,
                        fontSize: 12,
                        color: Colors.grey),
                  ),
                ],
              ),
            ),
            const Spacer(),
            const Padding(
              padding: EdgeInsets.only(right: 8.0),
              child: Icon(
                CupertinoIcons.phone,
                size: 24,
              ),
            ),
            const Icon(
              CupertinoIcons.video_camera,
              size: 32,
            ),
          ],
        ),
      ),
      body: Column(
        children: [
          BlocConsumer<ChatBloc, ChatState>(
            listener: (context, state) {
              // TODO: implement listener
            },
            builder: (context, state) {
              if (state is MessagesLoaded) {
                return Expanded(
                  child: Padding(
                    padding: const EdgeInsets.only(top: 20),
                    child: ListView.builder(
                      reverse: true,
                      shrinkWrap: true,
                      itemCount: state.messages.length,
                      itemBuilder: (context, index) {
                        final data = state.messages[index];

                        final isReceive = data.sender.id == args.user1.id;
                        return Row(
                          mainAxisAlignment: isReceive
                              ? MainAxisAlignment.start
                              : MainAxisAlignment.end,
                          crossAxisAlignment: CrossAxisAlignment.center,
                          children: [
                            if (isReceive)
                              const Padding(
                                padding: EdgeInsets.only(left: 10),
                                child: CircleAvatar(
                                  backgroundImage: AssetImage(
                                    'assets/images/Memoji Boys 6-18.png',
                                  ),
                                  radius: 25,
                                ),
                              ),
                            Flexible(
                              child: Padding(
                                padding: EdgeInsets.only(
                                    left: isReceive ? 10 : 0,
                                    right: isReceive ? 0 : 10,
                                    top: 18),
                                child: ChatMessage(
                                  messageType: data.sender.id == args.user1.id
                                      ? MessageType.receive
                                      : MessageType.send,
                                  message: data.content,
                                  timeStamp:
                                      '${TimeOfDay.now().hour.toString()}:${TimeOfDay.now().minute.toString()}',
                                  isImage: false,
                                ),
                              ),
                            ),
                            if (!isReceive)
                              const Padding(
                                padding: EdgeInsets.only(right: 10),
                                child: CircleAvatar(
                                  backgroundImage: AssetImage(
                                    'assets/images/Memoji Boys 6-18.png',
                                  ),
                                  radius: 25,
                                ),
                              ),
                          ],
                        );
                      },
                    ),
                  ),
                );
              } else {
                return Container(
                  child: const Center(
                      child: Text(
                    'No Chat History',
                    style: TextStyle(fontWeight: FontWeight.w600, fontSize: 18),
                  )),
                );
              }
            },
          ),
          const SizedBox(
            height: 68,
          )
        ],
      ),
      bottomSheet: Container(
        decoration: const BoxDecoration(color: Colors.white),
        height: 78,
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 8.0),
          child: Row(
            children: [
              Padding(
                padding: const EdgeInsets.only(right: 8.0),
                child: Transform.rotate(
                    angle: math.pi / 2.7,
                    child: const Icon(Icons.attach_file_outlined)),
              ),
              Expanded(
                child: ReusableTextField(
                  onsubmit: () {
                    // print("object");

                    final chatBloc = BlocProvider.of<ChatBloc>(context);

                    chatBloc.add(SendMessage(
                        args.chatId, messageController.text, 'text'));

                    messageController.clear();
                  },
                  hint: 'Write your message',
                  textEditingController: messageController,
                  textInputType: TextInputType.text,
                ),
              ),
              const Padding(
                padding: EdgeInsets.only(left: 8.0),
                child: Icon(CupertinoIcons.camera),
              ),
              const Padding(
                padding: EdgeInsets.symmetric(horizontal: 8.0),
                child: Icon(Icons.mic_none),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
