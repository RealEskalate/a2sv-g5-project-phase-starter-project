class ChatMessageEntity {
  final String id;
  final String senderId;
  final String receiverId;
  final String messageContent;
  final DateTime timestamp;

  ChatMessageEntity({
    required this.id,
    required this.senderId,
    required this.receiverId,
    required this.messageContent,
    required this.timestamp,
  });
}
