import 'package:flutter/material.dart';

// import '../../domain/entity/chat.dart';
import '../widgets/chat_card.dart';
import '../widgets/user_avater.dart';

class ChatList extends StatefulWidget {
  // final List<Chat> chats;
  final List<List> infos;
  const ChatList({super.key, required this.infos});

  @override
  State<ChatList> createState() => _ChatListState();
}

class _ChatListState extends State<ChatList> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: const Color(0xFF498CF0),
        leading: IconButton(icon: const Icon(Icons.menu_outlined), onPressed: (){}, color: Colors.white,),
      ),
      backgroundColor: Color(0xFF498CF0),
      body: Column(
        children: [
        //  const SizedBox(height: 30,),
           SingleChildScrollView(
            scrollDirection: Axis.horizontal,
            child: Padding(
              padding: EdgeInsets.symmetric(horizontal: 8.0, vertical: 15),
              child: Wrap(
                 spacing: 20.0, // space between items
                alignment: WrapAlignment.start,
                crossAxisAlignment: WrapCrossAlignment.center,
                children: [
                  UserAvater(image: 'assets/images/avater.png', story: true,online: true,),
                  UserAvater(image: 'assets/images/avater.png', story: true, online: true,),
                  UserAvater(image: 'assets/images/avater.png', story: true, ),
                  UserAvater(image: 'assets/images/avater.png', story: true,online: true),
                  UserAvater(image: 'assets/images/avater.png', story: true,online: true),
                  UserAvater(image: 'assets/images/avater.png', story: true,online: true),
                ],
              ),
            ),
          ),
        //  const SizedBox(height: 40,),
        //  const Text("data"),
          Expanded(
            
              
              child: Container(
                   padding: const EdgeInsets.only(top: 50),
            decoration: const BoxDecoration(color: Colors.white,
            borderRadius: BorderRadius.only(topLeft: Radius.circular(50), topRight: Radius.circular(50))),
            child: ListView.builder(
                  itemCount: widget.infos.length,
                  itemBuilder: (context, index) {
                    return ChatCard(
                      topMessage: widget.infos[index][0],
                      user: widget.infos[index][1],
                      time: widget.infos[index][2],
                      unread: widget.infos[index][3],
                      online: widget.infos[index][4],
                    );
                  }),
            ),
          ),
        ],
      ),
    );
  }
}
