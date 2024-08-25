// import 'package:flutter/material.dart';
// import 'package:flutter_bloc/flutter_bloc.dart';

// // import '../../domain/entity/chat.dart';
// import '../../../authentication/presentation/bloc/blocs.dart';
// import '../../../product/presentation/bloc/blocs.dart';
// import '../../domain/entity/chat.dart';
// import '../bloc/blocs.dart';
// import '../bloc/event.dart';
// import '../bloc/state.dart';
// import '../widgets/chat_card.dart';
// import '../widgets/user_avater.dart';

// class ChatList extends StatefulWidget {
//   final List<Chat> chats;
//   // final List<List> infos;
//   const ChatList({super.key, required this.chats});

//   Future<void> _refreshChats(BuildContext context) async {
//     print('refreshing');
//     print(chats);
//     BlocProvider.of<ChatBloc>(context).add(const LoadChatsEvent());
//   }

//   @override
//   State<ChatList> createState() => _ChatListState();
// }

// class _ChatListState extends State<ChatList> {
//   @override
//   Widget build(BuildContext context) {
//     return BlocConsumer<ChatBloc, ChatState>(builder: (context, state) {
//       context.read<ChatBloc>().add(LoadChatsEvent());
//       return Scaffold(
//         appBar: AppBar(
//           backgroundColor: const Color(0xFF498CF0),
//           leading: IconButton(
//             icon: const Icon(Icons.menu_outlined),
//             onPressed: () {},
//             color: Colors.white,
//           ),
//         ),
//         backgroundColor: Color(0xFF498CF0),
//         body: Column(
//           children: [
//             //  const SizedBox(height: 30,),
//             SingleChildScrollView(
//               scrollDirection: Axis.horizontal,
//               child: Padding(
//                 padding: EdgeInsets.symmetric(horizontal: 8.0, vertical: 15),
//                 child: Wrap(
//                   spacing: 20.0, // space between items
//                   alignment: WrapAlignment.start,
//                   crossAxisAlignment: WrapCrossAlignment.center,
//                   children: [
//                     UserAvater(
//                       image: 'assets/images/avater.png',
//                       story: true,
//                       online: true,
//                     ),
//                     UserAvater(
//                       image: 'assets/images/avater.png',
//                       story: true,
//                       online: true,
//                     ),
//                     UserAvater(
//                       image: 'assets/images/avater.png',
//                       story: true,
//                     ),
//                     UserAvater(
//                         image: 'assets/images/avater.png',
//                         story: true,
//                         online: true),
//                     UserAvater(
//                         image: 'assets/images/avater.png',
//                         story: true,
//                         online: true),
//                     UserAvater(
//                         image: 'assets/images/avater.png',
//                         story: true,
//                         online: true),
//                   ],
//                 ),
//               ),
//             ),
//             //  const SizedBox(height: 40,),
//             //  const Text("data"),
//                 Container(
//             child: Expanded(
//                       child: RefreshIndicator(
//                         onRefresh: () => widget._refreshChats(context),
//                         child: BlocBuilder<ChatBloc, ChatState>(
//                           builder: (context, state) {
//                             if (state is ChatLoadingState) {
//                               return const CircularProgressIndicator();
//                             } else if (state is ChatLoadedState) {
//                               // print(state.products);
//                               if (state.chats.isEmpty) {
//                                 return const Text('No chats available');
//                               }
//                               return ListView.builder(
//                                 itemCount: state.chats.length,
//                                 itemBuilder: (context, index) {
//                                   return ChatCard(
//                                     topMessage: state.chats[index].topMessage,
//                                     user: state.chats[index].user,
//                                     time: state.chats[index].time,
//                                     unread: state.chats[index].unread,
//                                   );
                                
//                                 },
//                               );
//                             } else if (state is ErrorState) {
//                               print(state.message);
//                               return const Text('Failed to fetch products');
//                             } else {
//                               return Container();
//                             }
//                           },
//                         ),
//                       ),
//                     ),
                
//               ],
//             ),
//           ),
//               // child: Container(
//               //   padding: const EdgeInsets.only(top: 50),
//               //   decoration: const BoxDecoration(
//               //       color: Colors.white,
//               //       borderRadius: BorderRadius.only(
//               //           topLeft: Radius.circular(50),
//               //           topRight: Radius.circular(50))),
//               //   child: ListView.builder(
//               //       itemCount: widget.infos.length,
//               //       itemBuilder: (context, index) {
//               //         return ChatCard(
//               //           topMessage: widget.infos[index][0],
//               //           user: widget.infos[index][1],
//               //           time: widget.infos[index][2],
//               //           unread: widget.infos[index][3],
//               //           online: widget.infos[index][4],
//               //         );
//               //       }),
//               // ),
//             ),
//           ],
//         ),
//       );
//     }, listener: (BuildContext context, ChatState state) {  },);
//   }
// }

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../authentication/domain/entity/user.dart';
import '../../../authentication/presentation/bloc/blocs.dart';
import '../../../product/presentation/bloc/blocs.dart';
import '../../domain/entity/chat.dart';
import '../bloc/blocs.dart';
import '../bloc/event.dart';
import '../bloc/state.dart';
import '../widgets/chat_card.dart';
import '../widgets/user_avater.dart';

class ChatList extends StatefulWidget {
  final User user;

  const ChatList(this.user, {super.key});

  Future<void> _refreshChats(BuildContext context) async {
    print('refreshing');
    print(user);
    BlocProvider.of<ChatBloc>(context).add(LoadChatsEvent(Chat as Chat));
  }

  @override
  State<ChatList> createState() => _ChatListState();
}

class _ChatListState extends State<ChatList> {
  @override
  Widget build(BuildContext context) {
    return BlocConsumer<ChatBloc, ChatState>(
      listener: (context, state) {},
      builder: (context, state) {
        context.read<ChatBloc>().add(LoadChatsEvent(Chat as Chat));
        return Scaffold(
          appBar: AppBar(
            backgroundColor: const Color(0xFF498CF0),
            leading: IconButton(
              icon: const Icon(Icons.menu_outlined),
              onPressed: () {},
              color: Colors.white,
            ),
          ),
          backgroundColor: const Color(0xFF498CF0),
          body: Column(
            children: [
              SingleChildScrollView(
                scrollDirection: Axis.horizontal,
                child: Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 8.0, vertical: 15),
                  child: Wrap(
                    spacing: 20.0,
                    alignment: WrapAlignment.start,
                    crossAxisAlignment: WrapCrossAlignment.center,
                    children: [
                      UserAvater(
                        image: 'assets/images/avater.png',
                        story: true,
                        online: true,
                      ),
                      UserAvater(
                        image: 'assets/images/avater.png',
                        story: true,
                        online: true,
                      ),
                      UserAvater(
                        image: 'assets/images/avater.png',
                        story: true,
                      ),
                      UserAvater(
                          image: 'assets/images/avater.png',
                          story: true,
                          online: true),
                      UserAvater(
                          image: 'assets/images/avater.png',
                          story: true,
                          online: true),
                      UserAvater(
                          image: 'assets/images/avater.png',
                          story: true,
                          online: true),
                    ],
                  ),
                ),
              ),
              Expanded(
                child: Container(
                  child: RefreshIndicator(
                    onRefresh: () => widget._refreshChats(context),
                    child: BlocBuilder<ChatBloc, ChatState>(
                      builder: (context, state) {
                        if (state is ChatLoadingState) {
                          return const Center(child: CircularProgressIndicator());
                        } else if (state is ChatLoadedState) {
                          if (state.chats.isEmpty) {
                            return const Center(child: Text('No chats available'));
                          }
                          return ListView.builder(
                            itemCount: state.chats.length,
                            itemBuilder: (context, index) {
                              return ChatCard(
                                user1: state.chats[index].user1,
                                user2: state.chats[index].user2,
                                // topMessage: state.chats[index].topMessage, // if available
                                // time: state.chats[index].time, // if available
                                // unread: state.chats[index].unread, 
                              );
                            },
                          );
                        } else if (state is ChatErrorState) {
                          print(state.message);
                          return const Center(child: Text('Failed to fetch chats'));
                        } else {
                          return Container();
                        }
                      },
                    ),
                  ),
                ),
              ),
            ],
          ),
        );
      },
    );
  }
}

