// ignore_for_file: overridden_fields

import '../../../auth/data/models/user_data_model.dart';
import '../../domain/entities/chat_entity.dart';

class ChatModel extends ChatEntity {
  @override
  final String id;
  @override
  final UserDataModel user1;
  @override
  final UserDataModel user2;

  const ChatModel({
    required this.id,
    required this.user1,
    required this.user2,
  }) : super(id: id, user1: user1, user2: user2);

  factory ChatModel.fromJson(Map<String, dynamic> json) {
    return ChatModel(
      id: json['_id'],
      user1: UserDataModel.fromJson(json['user1']),
      user2: UserDataModel.fromJson(json['user2']),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      '_id': id,
      'user1': user1.toJson(),
      'user2': user2.toJson(),
    };
  }

  static ChatModel toModel(ChatEntity entity) {
    return ChatModel(
      id: entity.id,
      user1: UserDataModel.toModel(entity.user1),
      user2: UserDataModel.toModel(entity.user2),
    );
  }

  ChatEntity toEntity() {
    return ChatEntity(
      id: id,
      user1: user1.toEntity(),
      user2: user2.toEntity(),
    );
  }

  // to model list
  static List<ChatModel> fromJsonList(List<dynamic> json) {
    return json.map((e) => ChatModel.fromJson(e)).toList();
  }
}
