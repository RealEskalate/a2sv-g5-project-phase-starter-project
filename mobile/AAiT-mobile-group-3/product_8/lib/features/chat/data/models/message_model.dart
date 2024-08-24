// ignore_for_file: overridden_fields

import '../../../auth/data/models/user_data_model.dart';
import '../../domain/entities/message_entity.dart';

class MessageModel extends MessageEntity {
  @override
  final String id;
  @override
  final UserDataModel sender;
  @override
  final String content;
  @override
  final String type;

  const MessageModel({
    required this.id,
    required this.sender,
    required this.content,
    required this.type,
  }) : super(id: id, sender: sender, content: content, type: type);

  factory MessageModel.fromJson(Map<String, dynamic> json) {
    return MessageModel(
      id: json['_id'],
      sender: UserDataModel.fromJson(json['sender']),
      content: json['content'],
      type: json['type'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      '_id': id,
      'sender': sender.toJson(),
      'content': content,
      'type': type,
    };
  }

  static MessageModel toModel(MessageEntity entity) {
    return MessageModel(
      id: entity.id,
      sender: UserDataModel.toModel(entity.sender),
      content: entity.content,
      type: entity.type,
    );
  }

  MessageEntity toEntity() {
    return MessageEntity(
      id: id,
      sender: sender.toEntity(),
      content: content,
      type: type,
    );
  }

  // to model list
  static List<MessageModel> fromJsonList(List<dynamic> json) {
    return json.map((e) => MessageModel.fromJson(e)).toList();
  }
}
