import 'package:flutter/material.dart';
import 'package:starter_project/features/chat/presentation/widgets/image_card.dart';

class MessageSendReceive extends StatelessWidget {
  const MessageSendReceive({super.key});

  @override
  Widget build(BuildContext context) {
    // Define a list of messages
    final messages = [
      _Message(
        isSentByMe: true,
        profileUrl: 'assets/images/flutter1.jpg',
        name: 'You',
        message: 'Hey! How are you?',
        timestamp: '10:15 AM',
      ),
      _Message(
        isSentByMe: false,
        profileUrl: 'assets/images/flutter1.jpg',
        name: 'Annel Ellison',
        message: 'I am good, thanks! How about you?',
        timestamp: '10:16 AM',
      ),
      _Message(
        isSentByMe: true,
        profileUrl: 'assets/images/flutter1.jpg',
        name: 'You',
        message: 'Hey! How are you?',
        timestamp: '10:15 AM',
      ),
      _Message(
        isSentByMe: false,
        profileUrl: 'assets/images/flutter1.jpg',
        name: 'Annel Ellison',
        message: 'I am good, thanks! How about you?',
        timestamp: '10:16 AM',
      ),
      _Message(
        isSentByMe: true,
        profileUrl: 'assets/images/flutter1.jpg',
        name: 'You',
        message: 'Hey! How are you?',
        timestamp: '10:15 AM',
      ),
      _Message(
        isSentByMe: false,
        profileUrl: 'assets/images/flutter1.jpg',
        name: 'Annel Ellison',
        message: 'I am good, thanks! How about you?',
        timestamp: '10:16 AM',
      ),
      _Message(
        isSentByMe: true,
        profileUrl: 'assets/images/flutter1.jpg',
        name: 'You',
        message: 'Hey! How are you?',
        timestamp: '10:15 AM',
      ),
      _Message(
        isSentByMe: false,
        profileUrl: 'assets/images/flutter1.jpg',
        name: 'Annel Ellison',
        message: 'I am good, thanks! How about you?',
        timestamp: '10:16 AM',
      ),
      // Add more messages here
    ];

    return Scaffold(
      body: Column(
        children: [
          Expanded(
            child: ListView.builder(
              padding: const EdgeInsets.symmetric(vertical: 8, horizontal: 16),
              itemCount: messages.length,
              itemBuilder: (context, index) {
                final message = messages[index];
                return _buildMessage(
                  context, // Pass context here
                  isSentByMe: message.isSentByMe,
                  profileUrl: message.profileUrl,
                  name: message.name,
                  message: message.message,
                  timestamp: message.timestamp,
                );
              },
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildMessage(
    BuildContext context, {
    // Accept context here
    required bool isSentByMe,
    required String profileUrl,
    required String name,
    required String message,
    required String timestamp,
  }) {
    return Container(
      margin: const EdgeInsets.symmetric(vertical: 8),
      child: Row(
        crossAxisAlignment: CrossAxisAlignment.start,
        mainAxisAlignment:
            isSentByMe ? MainAxisAlignment.end : MainAxisAlignment.start,
        children: [
          if (!isSentByMe) ...[
            CircleAvatar(
              backgroundImage: AssetImage(profileUrl),
              radius: 20,
            ),
            const SizedBox(width: 8),
          ],
          ConstrainedBox(
            constraints: BoxConstraints(
              maxWidth: MediaQuery.of(context).size.width * 0.5, // Uses context
            ),
            child: Column(
              crossAxisAlignment: isSentByMe
                  ? CrossAxisAlignment.end
                  : CrossAxisAlignment.start,
              children: [
                Text(
                  name,
                  style: const TextStyle(
                    fontWeight: FontWeight.bold,
                  ),
                ),
                Container(
                  padding: const EdgeInsets.all(12),
                  decoration: BoxDecoration(
                    color: isSentByMe ? Colors.blue : Colors.grey[300],
                    borderRadius: BorderRadius.circular(12),
                  ),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(
                        message,
                        style: TextStyle(
                          fontSize: 14.0,
                          color: isSentByMe ? Colors.white : Colors.black,
                        ),
                      ),
                      Align(
                        alignment: Alignment.bottomRight,
                        child: Text(
                          timestamp,
                          style:
                              const TextStyle(fontSize: 10, color: Colors.grey),
                        ),
                      ),
                    ],
                  ),
                ),
                const SizedBox(height: 20.0),
                ImageCard(imageUrl: 'assets/send_image.png'),
              ],
            ),
          ),
          if (isSentByMe) ...[
            const SizedBox(width: 8),
            CircleAvatar(
              backgroundImage: AssetImage(profileUrl),
              radius: 20,
            ),
          ],
        ],
      ),
    );
  }
}

class _Message {
  final bool isSentByMe;
  final String profileUrl;
  final String name;
  final String message;
  final String timestamp;

  _Message({
    required this.isSentByMe,
    required this.profileUrl,
    required this.name,
    required this.message,
    required this.timestamp,
  });
}
