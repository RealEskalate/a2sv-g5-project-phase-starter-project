import 'package:ecom_app/features/chat/domain/entities/user_chat_entity.dart';

class UserChatModel extends UserChatEntity {
  UserChatModel({required super.ChatID, required super.name});

  UserChatEntity toEntity() => UserChatEntity(ChatID: ChatID, name: name);

  static List<UserChatEntity> toEntityList(List<UserChatModel> models){
    return models.map((model) => model.toEntity()).toList();
  }
}
