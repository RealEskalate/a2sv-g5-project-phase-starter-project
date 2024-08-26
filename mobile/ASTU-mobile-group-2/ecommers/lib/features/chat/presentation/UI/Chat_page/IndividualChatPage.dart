import 'package:flutter/material.dart';
import 'package:get_it/get_it.dart';

import '../../../../../core/utility/global_message_part.dart';
import '../../../../../core/utility/socket_impl.dart';
import '../../../domain/entity/chat_entity.dart';

import 'chat_app_bar.dart';
import 'single_text.dart';

class IndividualChatPage extends StatefulWidget {
  const IndividualChatPage({
    super.key,
  });

  @override
  _IndividualChatPageState createState() => _IndividualChatPageState();
}

class _IndividualChatPageState extends State<IndividualChatPage> {
  final ScrollController _scrollController = ScrollController();
  final List<ChatEntity> result = [];
  late SocketService _socketProvider;
  TextEditingController textInputControl = TextEditingController();
  List<Map<String, String>>? messageList = [];

  @override
  void initState() {
    super.initState();
    _socketProvider = GetIt.instance<SocketService>();
    
    WidgetsBinding.instance.addPostFrameCallback((_) {
      _scrollToBottom();
    });

    _initializeSocketService();
  }

  void _initializeSocketService() async {
    // Ensure socket connection
    await _socketProvider.connect();

    // Listen for incoming messages
    _socketProvider.listen('message:received', (data) {
      print(data);
      print('Message received:');
      print('=============================================================================');
      
    });
  }

  void _scrollToBottom() {
    if (_scrollController.hasClients) {
      _scrollController.animateTo(
        _scrollController.position.maxScrollExtent,
        duration: const Duration(milliseconds: 300),
        curve: Curves.easeOut,
      );
    }
  }

  void unfocusTextFields() {
    FocusScope.of(context).unfocus();
  }

  @override
  Widget build(BuildContext context) {
    final Map<String, dynamic> data =
        ModalRoute.of(context)!.settings.arguments as Map<String, dynamic>;

    messageList = GlobalMessagePart.gloablMessage[data['chatId']];

    return SafeArea(
      child: Scaffold(
        backgroundColor: Colors.white,
        appBar:  PreferredSize(
          preferredSize: const Size.fromHeight(kToolbarHeight),
          child: ChatAppBar(name: data['name']),
        ),
        body: GestureDetector(
          onTap: unfocusTextFields,
          child: Column(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              Expanded(
                child: Padding(
                  padding: const EdgeInsets.fromLTRB(20, 20, 20, 40),
                  child: ListView.separated(
                    controller: _scrollController,
                    itemBuilder: (BuildContext context, int index) {
                      final current = messageList?[index] ?? {};

                      return SingleText(
                          profile_pic: 'assets/image/profileG.jpg',
                          name: data['name'],
                          text: current['content'],
                          isMe: current['senderId'] == data['itMe'],
                          time: '22:34');
                    },
                    separatorBuilder: (BuildContext context, int index) {
                      return const SizedBox(height: 20);
                    },
                    itemCount: messageList?.length ?? 0,
                  ),
                ),
              ),
              _buildInputBar(data['chatId'], data['itMe']),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildInputBar(String chatId, senderId) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 8.0, vertical: 10.0),
      child: Row(
        children: [
          IconButton(
            icon: const Icon(
              Icons.attach_file_outlined,
              color: Colors.black,
            ),
            onPressed: () {
              // Handle attachment
            },
          ),
          Expanded(
            child: Container(
              decoration: BoxDecoration(
                color: Colors.grey[200],
                borderRadius: BorderRadius.circular(25),
              ),
              child: Row(
                children: [
                  const SizedBox(width: 10),
                  Expanded(
                    child: TextField(
                      controller: textInputControl,
                      decoration: const InputDecoration(
                        hintText: 'Write your message',
                        border: InputBorder.none,
                      ),
                      onChanged: (text) {
                        _scrollToBottom();
                      },
                    ),
                  ),
                  IconButton(
                    icon: const Icon(
                      Icons.send,
                    ),
                    onPressed: () {
                      if (textInputControl.text.isNotEmpty) {
                        _socketProvider.sendMessage(
                            chatId, textInputControl.text);
                        setState(() {
                          GlobalMessagePart.gloablMessage[chatId]?.add({
                            'content': textInputControl.text,
                            'senderId': senderId
                          });
                          _scrollToBottom();
                          textInputControl.clear();
                        });
                      }
                    },
                  ),
                ],
              ),
            ),
          ),
          IconButton(
            icon: const Icon(Icons.camera_alt_outlined),
            onPressed: () {
              // Handle camera action
            },
          ),
          IconButton(
            icon: const Icon(Icons.mic_none_outlined),
            onPressed: () {
              // Handle microphone action
            },
          ),
        ],
      ),
    );
  }

  @override
  void dispose() {
    _scrollController.dispose();
    super.dispose();
  }
}
