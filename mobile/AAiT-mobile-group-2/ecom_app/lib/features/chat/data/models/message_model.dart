import 'package:ecom_app/features/chat/data/models/message_sender_model.dart';
import 'package:ecom_app/features/chat/domain/entities/message_entity.dart';

class MessageModel extends MessageEntity {
  final MessageSenderModel messageSenderModel;
  final String content;

  MessageModel({
    required this.messageSenderModel,
    required this.content,
  }) : super(messageSenderEntity: messageSenderModel, content: content);

  // Factory constructor to create a MessageModel from JSON
  factory MessageModel.fromJson(Map<String, dynamic> json) {
    return MessageModel(
      messageSenderModel: MessageSenderModel.fromJson(json['sender']),
      content: json['content'],
    );
  }

  // Method to convert MessageModel to JSON (if needed)
  Map<String, dynamic> toJson() {
    return {
      'sender': messageSenderModel.toJson(),
      'content': content,
    };
  }

  MessageEntity toEntity() => MessageEntity(
        messageSenderEntity: messageSenderModel,
        content: content,
      );
  
  static List<MessageEntity> toEntityList(List<MessageModel> models){
    return models.map((model) => model.toEntity()).toList();
  }
}
