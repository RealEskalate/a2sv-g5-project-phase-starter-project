import 'dart:ffi';
import 'package:flutter/material.dart';
import 'package:image_picker/image_picker.dart';

class ChatView extends StatefulWidget {
  final Map<String, String> user;

  const ChatView({required this.user, super.key});

  @override
  State<ChatView> createState() => _ChatViewState();
}

class _ChatViewState extends State<ChatView> {
  final List<ChatMessage> _messages = [];
  final TextEditingController _controller = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Row(
          children: [
            Padding(
              padding: const EdgeInsets.all(10),
              child: CircleAvatar(
                backgroundImage: AssetImage(widget.user['profilePic']!),
                radius: 20,
                backgroundColor: Colors.green,
              ),
            ),
            Text(widget.user['name']!),
          ],
        ),
        actions: [
          IconButton(
            icon: const Icon(Icons.call),
            onPressed: () {
              // Handle call button press
            },
          ),
          IconButton(
            icon: const Icon(Icons.video_call),
            onPressed: () {
              // Handle video call button press
            },
          ),
        ],
      ),
      body: Column(
        children: [
          Expanded(
            child: ListView.builder(
              reverse: true,
              itemCount: _messages.length,
              itemBuilder: (context, index) {
                final message = _messages[index];
                return Align(
                  alignment: message.isSentByUser
                      ? Alignment.centerRight
                      : Alignment.centerLeft,
                  child: Container(
                    margin:
                        const EdgeInsets.symmetric(vertical: 4, horizontal: 8),
                    padding: const EdgeInsets.all(12),
                    decoration: BoxDecoration(
                      color: message.isSentByUser
                          ? Colors.blueAccent
                          : Colors.grey[300],
                      borderRadius: BorderRadius.circular(12),
                    ),
                    child: message.isFile
                        ? Column(
                            children: [
                              Icon(Icons.attach_file, color: Colors.grey),
                              Text(
                                message.text,
                                style: TextStyle(
                                  color: message.isSentByUser
                                      ? Colors.white
                                      : Colors.black,
                                ),
                              ),
                            ],
                          )
                        : Text(
                            message.text,
                            style: TextStyle(
                              color: message.isSentByUser
                                  ? Colors.white
                                  : Colors.black,
                            ),
                          ),
                  ),
                );
              },
            ),
          ),
          CustomChatInput(
            onSendPressed: (text) {
              if (text.isNotEmpty) {
                setState(() {
                  _messages.insert(
                    0,
                    ChatMessage(text: text, isSentByUser: true),
                  );
                });
                _controller.clear();
              }
            },
            onImageSelected: (filePath) {
              setState(() {
                _messages.insert(
                  0,
                  ChatMessage(text: filePath, isSentByUser: true, isFile: true),
                );
              });
            },
            onRecordingComplete: (audioPath) {
              // Handle sending the audio file
            },
          ),
        ],
      ),
    );
  }
}

class ChatMessage {
  final String text;
  final bool isSentByUser;
  final bool isFile;

  ChatMessage({
    required this.text,
    required this.isSentByUser,
    this.isFile = false,
  });
}

class CustomChatInput extends StatefulWidget {
  final Function(String) onSendPressed;
  final Function(String) onImageSelected;
  final Function(String) onRecordingComplete;

  const CustomChatInput({
    required this.onSendPressed,
    required this.onImageSelected,
    required this.onRecordingComplete,
  });

  @override
  _CustomChatInputState createState() => _CustomChatInputState();
}

class _CustomChatInputState extends State<CustomChatInput> {
  final TextEditingController _controller = TextEditingController();

  Future<void> _pickImage() async {
    final ImagePicker picker = ImagePicker();
    final XFile? image = await picker.pickImage(source: ImageSource.gallery);

    if (image != null) {
      widget.onImageSelected(image.path);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 8.0),
      child: Row(
        children: [
          // IconButton(
          //   icon: const Icon(Icons.image, color: Colors.grey),
          //   onPressed: () async {
          //     await _pickImage();
          //   },
          // ),
          IconButton(
            icon: Icon(Icons.attach_file),
            onPressed: () {},
          ),
          IconButton(
            icon: Icon(
              Icons.keyboard_voice_rounded,
            ),
            onPressed: () {},
          ),
          Expanded(
            child: TextField(
              controller: _controller,
              decoration: InputDecoration(
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(12.0),
                  borderSide: BorderSide.none,
                ),
                contentPadding: const EdgeInsets.all(10),
                hintText: 'Type a message...',
              ),
              onSubmitted: (text) {
                _sendMessage(text);
              },
            ),
          ),
          IconButton(
            icon: const Icon(Icons.send, color: Colors.blueAccent),
            onPressed: () {
              _sendMessage(_controller.text);
            },
          ),
        ],
      ),
    );
  }

  void _sendMessage(String text) {
    if (text.isNotEmpty) {
      widget.onSendPressed(text);
      _controller.clear();
    }
  }
}
