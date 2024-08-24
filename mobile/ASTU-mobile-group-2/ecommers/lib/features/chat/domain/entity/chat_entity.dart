

class ChatEntity {
  final String senderId;
  final String senderName;
  final String recieverId;
  final String recieverName;
  final String chatId;
  final List<Map<String,dynamic>> messages;
  

  const ChatEntity({
    required this.senderId,
    required this.senderName,
    required this.recieverId,
    required this.recieverName,
    required this.chatId,
    required this.messages
  });

}