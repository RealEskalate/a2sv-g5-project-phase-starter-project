import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/chat_entity.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/models/seller_model.dart';

class ChatModel{
  final String id;
  final String id;
  final ChatEntity chat;
  final User sender;
  final String type;
  final String content;

  get receiver => chat.user1 == sender ? chat.user2 : chat.user1;

  const ChatModel({
    required this.id,
    required this.chat,
    required this.sender,
    required this.type,
    required this.content,
  });

  @override
  List<Object?> get props => [id, chat, sender, type, content];
}
  factory ChatModel.fromJson(Map<String, dynamic> json) {
    return ChatModel(
      
    );
  }

  Map<String, dynamic> toJson() => {
        "name": name,
        "description": description,
        "price": price,
        
      };
}