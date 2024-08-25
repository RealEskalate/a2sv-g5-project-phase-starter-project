import 'package:dartz/dartz.dart';

import '../../../../core/errors/failures/failure.dart';
import '../../../auth/domain/entities/user_entity.dart';
import '../entity/chat.dart';
import '../entity/message.dart';
import '../repository/chat_repository.dart';

class ChatUseCases{
   final ChatRepository chatRepository;

  ChatUseCases({required this.chatRepository});

  Future< List<ChatEntity>> callGetChatRooms()async {
    return await chatRepository.getChatRooms();
  }

  Future<List<MessageEntity>> callGetMessagesForChat(String chatId) async{
    return await chatRepository.getMessagesForChat(chatId);
  }

  Future<ChatEntity> callCreateChatRoom(UserEntity user1, UserEntity user2)async {
    return await chatRepository.createChatRoom(user1, user2);
  }

  Future<void> callSendMessage(String chatId, String message)async{
    return await chatRepository.sendMessage(chatId, message);
  }

  Future<void> callAcknowledgeMessageDelivery(String messageId)async{
    return await chatRepository.acknowledgeMessageDelivery(messageId);
  }

  


}

