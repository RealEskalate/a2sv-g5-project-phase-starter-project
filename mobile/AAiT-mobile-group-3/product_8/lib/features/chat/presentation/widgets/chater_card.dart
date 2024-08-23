import 'package:flutter/material.dart';

class ChatterCard extends StatelessWidget {
  final String name;
  final String lastMessage;
  final int unreadCount;
  final String time;
  final String imageUrl;
  final bool isOnline;

  const ChatterCard({
    super.key,
    required this.name,
    required this.lastMessage,
    required this.unreadCount,
    required this.time,
    required this.imageUrl,
    required this.isOnline,
  });

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: double.infinity,
      height: 90,
      child: Padding(
        padding: const EdgeInsets.symmetric(vertical: 4.0, horizontal: 8.0),
        child: Card(
          color: Colors.white,
          elevation: 0,
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(12),
            side: const BorderSide(color: Colors.white, width: 1),
          ),
          child: Padding(
            padding:
                const EdgeInsets.symmetric(vertical: 8.0, horizontal: 12.0),
            child: Row(
              children: [
                _buildProfileImage(),
                const SizedBox(width: 12),
                Expanded(child: _buildMessageInfo()),
                _buildTimeAndUnreadCount(),
              ],
            ),
          ),
        ),
      ),
    );
  }

  Widget _buildProfileImage() {
    return Stack(
      children: [
        CircleAvatar(
          radius: 28,
          backgroundImage:
              AssetImage(imageUrl), // Replace with the correct image path
        ),

        // Add online status indicator here

        Positioned(
          bottom: 0,
          right: 0,
          top: 30,
          child: isOnline
              ? const Icon(Icons.circle, color: Colors.green, size: 10)
              : const Icon(Icons.circle, color: Colors.grey, size: 10),
        ),
      ],
    );
  }

  Widget _buildMessageInfo() {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        Text(
          name,
          style: const TextStyle(
            fontFamily: 'General Sans Variable',
            fontSize: 18,
            fontWeight: FontWeight.w500,
            color: Color(0xFF000E08),
          ),
          overflow: TextOverflow.ellipsis,
        ),
        const SizedBox(height: 4),
        Text(
          lastMessage,
          style: const TextStyle(
            fontFamily: 'General Sans Variable',
            fontSize: 12,
            color: Color(0xFF797C7B),
          ),
          maxLines: 1,
          overflow: TextOverflow.ellipsis,
        ),
      ],
    );
  }

  Widget _buildTimeAndUnreadCount() {
    return Column(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        Text(
          time,
          style: const TextStyle(
            fontSize: 12,
            fontFamily: 'General Sans Variable',
            color: Color(0xFF797C7B),
          ),
        ),
        const SizedBox(height: 8),
        if (unreadCount > 0)
          Container(
            padding: const EdgeInsets.all(6),
            decoration: const BoxDecoration(
              color: Color(0xFF3F51F3),
              shape: BoxShape.circle,
            ),
            child: Text(
              '$unreadCount',
              style: const TextStyle(
                fontSize: 11,
                color: Colors.white,
              ),
            ),
          ),
      ],
    );
  }
}
