import 'package:flutter/material.dart';

class ChatterCard extends StatefulWidget {
  final String name;
  final String lastMessage;
  final int unreadCount;
  final String time;
  final String imageUrl;

  const ChatterCard({
    super.key,
    required this.name,
    required this.lastMessage,
    required this.unreadCount,
    required this.time,
    required this.imageUrl,
  });

  @override
  _ChatterCardState createState() => _ChatterCardState();
}

class _ChatterCardState extends State<ChatterCard> {
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
            side: BorderSide(color: Colors.white, width: 1),
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
    return ClipOval(
      child: Image.asset(
        widget.imageUrl,
        width: 55,
        height: 55,
        fit: BoxFit.cover,
        errorBuilder: (context, error, stackTrace) => Icon(
          Icons.account_circle,
          size: 55,
          color: Colors.grey.shade400,
        ),
      ),
    );
  }

  Widget _buildMessageInfo() {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        Text(
          widget.name,
          style: const TextStyle(
            fontSize: 16,
            fontWeight: FontWeight.w600,
          ),
          overflow: TextOverflow.ellipsis,
        ),
        const SizedBox(height: 4),
        Text(
          widget.lastMessage,
          style: TextStyle(
            fontSize: 14,
            color: Colors.grey.shade600,
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
          widget.time,
          style: TextStyle(
            fontSize: 12,
            color: Colors.grey.shade500,
          ),
        ),
        const SizedBox(height: 8),
        if (widget.unreadCount > 0)
          Container(
            padding: const EdgeInsets.all(6),
            decoration: const BoxDecoration(
              color: Colors.blue,
              shape: BoxShape.circle,
            ),
            child: Text(
              '${widget.unreadCount}',
              style: const TextStyle(
                fontSize: 12,
                color: Colors.white,
              ),
            ),
          ),
      ],
    );
  }
}
