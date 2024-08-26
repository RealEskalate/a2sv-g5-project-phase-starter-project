import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../domain/entity/chat.dart';
import '../bloc/chat_bloc.dart';
import '../bloc/chat_event.dart';
import '../bloc/chat_state.dart';

class ChatList extends StatefulWidget {
  static const String routes = '/chat_lists';

  const ChatList({super.key});

  @override
  State<ChatList> createState() => _ChatListState();
}

class _ChatListState extends State<ChatList> {
  List imagesList = [
    'assets/images/Memoji Boys 6-18.png',
    'assets/images/42 11.png',
    'assets/images/21 5.png'
  ];
  List nameList = ['My status', 'Adil', 'Marina', 'Dean', 'Max'];

  List<Color> colorList = [
    const Color(0xffC5C5C5),
    const Color(0xffFBE8A8),
    const Color(0xffFEC7D3),
    const Color(0xffC8ECFD),
    const Color(0xffC7FEE0)
  ];

  @override
  void initState() {
    super.initState();

    final chatBloc = BlocProvider.of<ChatBloc>(context);

    chatBloc.add(ConnectServerEvent());
    // chatBloc.add(CreateChatRoom("66c359e70a542ac9d961a025"));

    chatBloc.add(LoadChatRooms());
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        backgroundColor: const Color(0xff498CF0),
        body: SafeArea(
          child: Column(
              mainAxisAlignment: MainAxisAlignment.start,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                const Padding(
                  padding: EdgeInsets.only(left: 10.0, top: 15, bottom: 20),
                  child: Icon(
                    Icons.search,
                    color: Colors.white,
                  ),
                ),
                Padding(
                  padding:
                      const EdgeInsets.symmetric(horizontal: 8.0, vertical: 12),
                  child: SizedBox(
                    height: 89,
                    child: ListView.builder(
                      itemCount: 8,
                      scrollDirection: Axis.horizontal,
                      itemBuilder: (context, index) {
                        return Padding(
                          padding: const EdgeInsets.symmetric(horizontal: 8.0),
                          child: Column(
                            children: [
                              Stack(
                                alignment: Alignment.center,
                                children: [
                                  Container(
                                    height: 60,
                                    width: 60,
                                    decoration: BoxDecoration(
                                      shape: BoxShape.circle,
                                      border: Border.all(
                                        color: colorList[index % 5],
                                        width: 3.0,
                                      ),
                                    ),
                                  ),
                                  Container(
                                    height: 50,
                                    width: 50,
                                    decoration: BoxDecoration(
                                      color: colorList[index % 5],
                                      shape: BoxShape.circle,
                                    ),
                                    child: ClipRRect(
                                      borderRadius: BorderRadius.circular(50),
                                      child: Image.asset(imagesList[index % 3]),
                                    ),
                                  ),
                                ],
                              ),
                              Text(nameList[index % 5])
                            ],
                          ),
                        );
                      },
                    ),
                  ),
                ),
                BlocConsumer<ChatBloc, ChatState>(
                  listener: (context, state) {
                    // TODO: implement listener
                  },
                  builder: (context, state) {
                    if (state is ChatLoaded) {
                      return Expanded(
                        child: DecoratedBox(
                            decoration: const BoxDecoration(
                                color: Colors.white,
                                borderRadius: BorderRadius.vertical(
                                    top: Radius.circular(40))),
                            child: Padding(
                              padding: const EdgeInsets.only(top: 8.0),
                              child: ListView.builder(
                                itemCount: state.chats.length,
                                padding: const EdgeInsets.all(8),
                                scrollDirection: Axis.vertical,
                                itemBuilder: (context, index) {
                                  final item = state.chats[index];
                                  return Padding(
                                    padding: const EdgeInsets.symmetric(
                                        vertical: 8.0, horizontal: 15),
                                    child: GestureDetector(
                                      onTap: () {
                                        final chatBloc =
                                            BlocProvider.of<ChatBloc>(context);

                                        chatBloc.add(LoadMessages(item.chatId));
                                        Navigator.pushNamed(
                                            context, '/chat_room',
                                            arguments: ChatEntity(
                                                chatId: item.chatId,
                                                user1: item.user1,
                                                user2: item.user2));
                                      },
                                      child: Container(
                                        color: Colors.transparent,
                                        child: Row(
                                          children: [
                                            Container(
                                              height: 50,
                                              width: 50,
                                              decoration: BoxDecoration(
                                                color: colorList[index % 5],
                                                shape: BoxShape.circle,
                                              ),
                                              child: ClipRRect(
                                                borderRadius:
                                                    BorderRadius.circular(50),
                                                child: Image.asset(
                                                    imagesList[index % 3]),
                                              ),
                                            ),
                                            Padding(
                                              padding: const EdgeInsets.only(
                                                  left: 8.0),
                                              child: Column(
                                                crossAxisAlignment:
                                                    CrossAxisAlignment.start,
                                                children: [
                                                  Text(
                                                    item.user1.name,
                                                    // 'Alex Linderson',
                                                    style: const TextStyle(
                                                        fontWeight:
                                                            FontWeight.w600,
                                                        fontSize: 15),
                                                  ),
                                                  const Text(
                                                    'How are you today?',
                                                    style: TextStyle(
                                                        fontWeight:
                                                            FontWeight.w400,
                                                        fontSize: 12,
                                                        color: Colors.grey),
                                                  ),
                                                ],
                                              ),
                                            ),
                                            const Spacer(),
                                            const Column(
                                              children: [
                                                Text(
                                                  '2 min ago',
                                                  style: TextStyle(
                                                      fontWeight:
                                                          FontWeight.w400,
                                                      fontSize: 12,
                                                      color: Colors.grey),
                                                ),
                                                DecoratedBox(
                                                  decoration: BoxDecoration(
                                                      shape: BoxShape.circle,
                                                      color: Color(0xff498CF0)),
                                                  child: Padding(
                                                    padding:
                                                        EdgeInsets.all(8.0),
                                                    child: Center(
                                                        child: Text(
                                                      '3',
                                                      style: TextStyle(
                                                          color: Colors.white),
                                                    )),
                                                  ),
                                                )
                                              ],
                                            ),
                                          ],
                                        ),
                                      ),
                                    ),
                                  );
                                },
                              ),
                            )),
                      );
                    } else if (state is ChatLoading) {
                      return const Center(
                        child: CircularProgressIndicator(),
                      );
                    } else {
                      return Container();
                    }
                  },
                )
              ]),
        ));
  }
}
