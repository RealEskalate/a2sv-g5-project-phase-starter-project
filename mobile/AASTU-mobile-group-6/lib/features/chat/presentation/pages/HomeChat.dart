import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/chat_entity.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/bloc/chat_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/bloc/chat_event.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/bloc/chat_state.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/pages/chatView.dart';
import 'package:flutter/material.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/widgets/bottomnavbar.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class Chat extends StatelessWidget {
  const Chat({super.key});

  @override
  Widget build(BuildContext context) {
    final List<Map<String, String>> activeUsers = [
      {'name': 'Adil', 'profilePic': 'assets/av1.png'},
      {'name': 'Marina', 'profilePic': 'assets/av2.png'},
      {'name': 'Dean', 'profilePic': 'assets/av3.png'},
      {'name': 'Max', 'profilePic': 'assets/av1.png'},
      {'name': 'Me', 'profilePic': 'assets/av3.png'},
      {'name': 'Adil', 'profilePic': 'assets/av1.png'},
      {'name': 'Marina', 'profilePic': 'assets/av2.png'},
      {'name': 'Dean', 'profilePic': 'assets/av3.png'},
      {'name': 'Max', 'profilePic': 'assets/av1.png'},
      {'name': 'Me', 'profilePic': 'assets/av3.png'},
    ];

    final List<Color> userColors = [
      Colors.red,
      Colors.green,
      Colors.blue,
      Colors.orange,
      Colors.purple,
      Colors.teal,
      Colors.yellow,
      Colors.pink,
      Colors.cyan,
      Colors.amber,
    ];

    return Scaffold(
      appBar: AppBar(
        title: Text(
          'Search Users',
          style: TextStyle(
            // color: Theme.of(context).colorScheme.surface,
            color: Colors.white,
            fontWeight: FontWeight.w500,
          ),
        ),
        leading: Icon(
          Icons.search,
          color: Colors.white,
          // color: Theme.of(context).colorScheme.onSurface,
        ),
        backgroundColor: Theme.of(context).colorScheme.primary,
      ),
      body: Column(
        children: [
          Container(
            // margin: EdgeInsets.only(bottom: 10),
            height: 109, // Set a fixed height for the horizontal list
            color: Theme.of(context).colorScheme.primary,
            child: ListView.builder(
              scrollDirection: Axis.horizontal,
              itemCount: activeUsers.length,
              itemBuilder: (context, index) {
                final user = activeUsers[index];
                return GestureDetector(
               onTap: () {
  // Handle user tap
                Navigator.push(
                  context,
                  MaterialPageRoute(
                    builder: (context) => ChatView(user: user,),
                  ),
                );
              },
                  child: Padding(
                    padding: const EdgeInsets.all(8.0),
                    child: Column(
                      children: [
                        CircleAvatar(
                          backgroundImage: AssetImage(user['profilePic']!),
                          radius: 30,
                          backgroundColor:
                              userColors[index % userColors.length],
                        ),
                        const SizedBox(height: 8),
                        Text(
                          user['name']!,
                          style: TextStyle(
                              color: Theme.of(context).colorScheme.onSurface),
                        ),
                      ],
                    ),
                  ),
                );
              },
            ),
          ),
          Expanded(
                child: BlocBuilder<ChatBloc, ChatState>(
                  builder: (context, state) {
                    
                    if (state is ChatLoaded && state.messages.length>0){
                    return RefreshIndicator(
                      onRefresh: ()async{
                        context.read<ChatBloc>().add(ListAllMessagesEvent());},
                        
                      child: Container(
                        width: MediaQuery.of(context).size.width,
                        height: MediaQuery.of(context).size.height,
                        child: ListView.builder(
                          itemCount: state.messages.length,
                          itemBuilder: (context,index){
                            final item = state.messages[index];
                            return _duplicate(
                              context, 
                              item,
                              'assets/av1.png',
                               'How are you today?');
                                                
                          }),
                      )
                    );}else if (state is ChatLoaded && state.messages.length==0){
                      return Center(child: Text("No Chats Availble"),);
                
                    }
                    else{
                      return RefreshIndicator(child:  Container(), 
                      onRefresh:()async{
                        context.read<ChatBloc>().add(ListAllMessagesEvent());},
                      );
                     
                    }
                  },
                ),
              )
          
        ],
      ),
      bottomNavigationBar: Bottomnavbar(),
    );
  }
}

Widget _duplicate(
  BuildContext context,
  ChatEntity chat,
  String imageUrl,
  String peekMessage,
) {
  return GestureDetector(
    onTap: () => Navigator.pushNamed(context, '/chatPage',arguments: chat,),
    child: Container(
      margin: EdgeInsets.only(bottom: 30),
      width: MediaQuery.of(context).size.width * 0.95,
      height: MediaQuery.of(context).size.height * 0.08,
      child: Row(
        children: [
          CircleAvatar(
            backgroundColor: Colors.yellow,
            radius: 30,
            child: Icon(
              Icons.person,
              size: 35,
            ),
          ),
          Container(
            margin: EdgeInsets.only(left: 10),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Text(
                  chat.user2.name,
                  style: TextStyle(
                    fontWeight: FontWeight.w800,
                    fontSize: 20,
                  ),
                ),
                Text(
                  peekMessage,
                  style: TextStyle(
                    fontSize: 15,
                    color: Colors.grey,
                  ),
                )
              ],
            ),
          ),
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.end,
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Text(
                  '2 min ago',
                  style: TextStyle(
                    color: Colors.grey,
                  ),
                ),
                Container(
                  child: CircleAvatar(
                    backgroundColor: Colors.purple,
                    radius: 13,
                    child: Text(
                      '4',
                      style: TextStyle(
                        color: Colors.white60,
                      ),
                    ),
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




/*
import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/chat_entity.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/bloc/chat_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/bloc/chat_event.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/bloc/chat_state.dart';
import 'package:flutter/material.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/widgets/bottomnavbar.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class Chat extends StatelessWidget {
  const Chat({super.key});

  @override
  Widget build(BuildContext context) {
    // context.read<ChatBloc>().add(ListAllMessagesEvent());
    final List<Map<String, String>> activeUsers = [
      {'name': 'Adil', 'profilePic': 'assets/av1.png'},
      {'name': 'Marina', 'profilePic': 'assets/av2.png'},
      {'name': 'Dean', 'profilePic': 'assets/av3.png'},
      {'name': 'Max', 'profilePic': 'assets/av1.png'},
      {'name': 'Me', 'profilePic': 'assets/av3.png'},
      {'name': 'Adil', 'profilePic': 'assets/av1.png'},
      {'name': 'Marina', 'profilePic': 'assets/av2.png'},
      {'name': 'Dean', 'profilePic': 'assets/av3.png'},
      {'name': 'Max', 'profilePic': 'assets/av1.png'},
      {'name': 'Me', 'profilePic': 'assets/av3.png'},
    ];

    final List<Color> userColors = [
      Colors.red,
      Colors.green,
      Colors.blue,
      Colors.orange,
      Colors.purple,
      Colors.teal,
      Colors.yellow,
      Colors.pink,
      Colors.cyan,
      Colors.amber,
    ];

    return BlocListener<ChatBloc, ChatState>(
      listener: (context, state) {
         if (state is ChatInitateFailure ) {
          showError(context, state.error);
        }
      },
      child: Scaffold(
          appBar: AppBar(
            title: const Text(
              'Search Users',
              style: TextStyle(color: Colors.white),
            ),
            leading: const Icon(
              Icons.search,
              color: Colors.white,
            ),
            backgroundColor: const Color.fromARGB(255, 72, 168, 246),
          ),
          body: Column(
            children: [
              Container(
                margin: EdgeInsets.only(bottom: 10),
                height: 109, // Set a fixed height for the horizontal list
                color: const Color.fromARGB(255, 72, 168, 246),
                child: ListView.builder(
                  scrollDirection: Axis.horizontal,
                  itemCount: activeUsers.length,
                  itemBuilder: (context, index) {
                    final user = activeUsers[index];
                    return GestureDetector(
                      onTap: () {
                        // Handle user tap
                      },
                      child: Padding(
                        padding: const EdgeInsets.all(8.0),
                        child: Column(
                          children: [
                            CircleAvatar(
                              backgroundImage: AssetImage(user['profilePic']!),
                              radius: 30,
                              backgroundColor:
                                  userColors[index % userColors.length],
                            ),
                            const SizedBox(height: 8),
                            Text(
                              user['name']!,
                              style: const TextStyle(color: Colors.white),
                            ),
                          ],
                        ),
                      ),
                    );
                  },
                ),
              ),
              
              Expanded(
                child: BlocBuilder<ChatBloc, ChatState>(
                  builder: (context, state) {
                    
                    if (state is ChatLoaded && state.messages.length>0){
                    return RefreshIndicator(
                      onRefresh: ()async{
                        context.read<ChatBloc>().add(ListAllMessagesEvent());},
                        
                      child: Container(
                        width: MediaQuery.of(context).size.width,
                        height: MediaQuery.of(context).size.height,
                        child: ListView.builder(
                          itemCount: state.messages.length,
                          itemBuilder: (context,index){
                            final item = state.messages[index];
                            return _duplicate(
                              context, 
                              item,
                              'assets/av1.png',
                               'How are you today?');
                                                
                          }),
                      )
                    );}else if (state is ChatLoaded && state.messages.length==0){
                      return Center(child: Text("No Chats Availble"),);
                
                    }
                    else{
                      return RefreshIndicator(child:  Container(), 
                      onRefresh:()async{
                        context.read<ChatBloc>().add(ListAllMessagesEvent());},
                      );
                     
                    }
                  },
                ),
              ),
            ],
          ),
          // bottomNavigationBar: Bottomnavbar(),
        ),
    );
  }
}



Widget _duplicate(
  BuildContext context,
  ChatEntity chat,
  String imageUrl,
  String peekMessage,
) {
  return GestureDetector(
    onTap: () => Navigator.pushNamed(context, '/chatPage',arguments: {'',chat.chatid},),
    child: Container(
      margin: EdgeInsets.only(bottom: 30),
      width: MediaQuery.of(context).size.width * 0.95,
      height: MediaQuery.of(context).size.height * 0.08,
      child: Row(
        children: [
          CircleAvatar(
            backgroundColor: Colors.white,
            radius: 30,
            child: Icon(
              Icons.person,
              size: 35,
            ),
          ),
          Container(
            margin: EdgeInsets.only(left: 10),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Text(
                  chat.user1.name,
                  style: TextStyle(
                    fontWeight: FontWeight.w800,
                    fontSize: 20,
                  ),
                ),
                Text(
                  peekMessage,
                  style: TextStyle(
                    fontSize: 15,
                    color: Colors.grey,
                  ),
                )
              ],
            ),
          ),
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.end,
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Text(
                  '2 min ago',
                  style: TextStyle(
                    color: Colors.grey,
                  ),
                ),
                Container(
                  child: CircleAvatar(
                    backgroundColor: Colors.purple,
                    radius: 13,
                    child: Text(
                      '4',
                      style: TextStyle(
                        color: Colors.white60,
                      ),
                    ),
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
void showError(BuildContext context, String message) {
  ScaffoldMessenger.of(context).showSnackBar(
    SnackBar(
      content: Text(message),
      backgroundColor: Theme.of(context).colorScheme.error,
    ),
  );
}*/