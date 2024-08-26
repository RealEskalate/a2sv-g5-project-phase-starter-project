import 'dart:developer';

import 'package:flutter/widgets.dart';

import '../../../../core/constants/colors.dart';
import '../../../auth/presentation/pages/pages.dart';
import '../bloc/chat_bloc.dart';
import '../widget/user_profile.dart';
import 'individiual_chat_screen.dart';

import '../bloc/chat_bloc.dart';

class ChatHomeScreen extends StatefulWidget {
  const ChatHomeScreen({super.key});

  @override
  State<ChatHomeScreen> createState() => _ChatHomeScreenState();
}

class _ChatHomeScreenState extends State<ChatHomeScreen> {
  @override
  void initState() {
    context.read<ChatBloc>().add(GetAllChatEvent());
    // context.read<ChatBloc>().add(ListOfMessageEvent(chatId: '66cc462cdab43c1a2e9803cd'));

    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return BlocConsumer<ChatBloc, ChatState>(
      listener: (context, state) {
        // log('from the home $state');
      },
      builder: (context, state) {
        if (state is ChatFailureState) {
          return Center(child: Text('Error: ${state.message}'));
        } else if (state is! LoadedAllChatState) {
          return Center(child: CircularProgressIndicator());
        }

        return Scaffold(
            backgroundColor: ChatColors.lightBlueColor,
            appBar: AppBar(
              leading: IconButton(
                  onPressed: () {},
                  icon: const Icon(
                    Icons.search,
                    color: Colors.white,
                  )),
              backgroundColor: Colors.transparent,
            ),
            body: Column(children: [
              SizedBox(
                height: 130,
                child: ListView.builder(
                  padding: const EdgeInsets.all(10),
                  shrinkWrap: true,
                  itemCount: 5,
                  scrollDirection: Axis.horizontal,
                  itemBuilder: (context, index) => Padding(
                    padding: const EdgeInsets.all(5.0),
                    child: Column(
                      children: [
                        showUser(onClicked: () {
                          // Navigator.of(context).push(MaterialPageRoute(
                          //   builder: (context) => const IndividiualChatScreen(),
                          // ));
                        }),
                        const Text(
                          'Marina',
                          style: TextStyle(color: Colors.white),
                        )
                      ],
                    ),
                  ),
                ),
              ),
              const SizedBox(
                height: 20,
              ),
              Expanded(
                child: Container(
                  padding: const EdgeInsets.only(top: 30),
                  decoration: const BoxDecoration(
                      borderRadius: BorderRadius.only(
                          topLeft: Radius.circular(50),
                          topRight: Radius.circular(50)),
                      color: Colors.white),
                  child: ListView.builder(
                    scrollDirection: Axis.vertical,
                    shrinkWrap: true,
                    itemCount: state.allChats.length,
                    itemBuilder: (context, index) => ListTile(
                      onTap: () {
                        Navigator.of(context).push(MaterialPageRoute(
                          // builder: (context) =>  IndividiualChatScreen(chatId:state.allChats[index].chatId ,),
                          builder: (context) => const IndividiualChatScreen(),
                        ));
                      },
                      title: Text(state.allChats[index].user2.name),
                      subtitle: const Text('How are you?'),
                      leading: showUser(),
                      trailing: Column(
                        children: [
                          const Text('2 min ago'),
                          CircleAvatar(
                            maxRadius: 10,
                            backgroundColor: ChatColors.lightBlueColor,
                            child: const Text('3'),
                          )
                        ],
                      ),
                    ),
                  ),
                ),
              )
            ]));
      },
    );
  }
}
