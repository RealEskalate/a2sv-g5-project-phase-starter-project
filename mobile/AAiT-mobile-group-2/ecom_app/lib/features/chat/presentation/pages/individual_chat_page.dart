import 'package:ecom_app/features/chat/mockData/mock_chat_data.dart';
import 'package:ecom_app/features/chat/presentation/widgets/current_user.dart';
import 'package:flutter/material.dart';

class IndividualChatPage extends StatefulWidget {
  const IndividualChatPage({
    super.key,
  });

  @override
  _IndividualChatPageState createState() => _IndividualChatPageState();
}

class _IndividualChatPageState extends State<IndividualChatPage> {
  final ScrollController _scrollController = ScrollController();

  @override
  void initState() {
    super.initState();
    WidgetsBinding.instance.addPostFrameCallback((_) {
      _scrollToBottom();
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
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        leading: const Icon(Icons.arrow_back),
        title: const CurrentUser(
          name: 'John Doe',
          image: 'assets/profile_picture.png',
          online: true,
        ),
        actions: const [
          IconButton(
            icon: Icon(
              Icons.call_outlined,
              color: Colors.black,
            ),
            onPressed: null,
          ),
          IconButton(
            icon: Icon(
              Icons.videocam_outlined,
              color: Colors.black,
            ),
            onPressed: null,
          ),
        ],
      ),
      body: GestureDetector(
        onTap: unfocusTextFields,
        child: Column(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Expanded(
              child: Padding(
                padding: EdgeInsets.fromLTRB(20, 20, 20, 40),
                child: ListView.separated(
                  controller: _scrollController,
                  itemBuilder: (BuildContext context, int index) {
                    return mock_chat_data[index];
                  },
                  separatorBuilder: (BuildContext context, int index) {
                    return const SizedBox(height: 20);
                  },
                  itemCount: mock_chat_data.length,
                ),
              ),
            ),
            _buildInputBar(),
          ],
        ),
      ),
    );
  }

  Widget _buildInputBar() {
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
                      decoration: InputDecoration(
                        hintText: "Write your message",
                        border: InputBorder.none,
                      ),
                      onChanged: (text) {
                        _scrollToBottom();
                      },
                    ),
                  ),
                  IconButton(
                    icon: const Icon(
                      Icons.copy,
                    ),
                    onPressed: () {
                      // Handle copy action
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
