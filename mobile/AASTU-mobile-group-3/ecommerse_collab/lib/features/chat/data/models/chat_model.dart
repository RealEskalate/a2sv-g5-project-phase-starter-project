import '../../../authentication/domain/entity/user.dart';
import '../../domain/entity/chat.dart';

class ChatModel extends Chat{
  final String id;
  final User user1;
  final User user2;
  

  ChatModel({
    required this.id,
    required this.user1,  
    required this.user2,  
  }) : super(id: id, user1: user1, user2: user2);

  factory ChatModel.fromJson(Map<String, dynamic> json) {
    return ChatModel(
      id: json['id'],
      user1: json['user1'],
      user2: json['user2'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'user1': user1,
      'user2': user2,
    };
  }
}