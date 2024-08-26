import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../repositories/chat_repository.dart';

class SendMessageUsecase {
  final ChatRepository chatRepository;

  SendMessageUsecase({required this.chatRepository});

  Future<Either<Failure,void>>send(String chatId,String message,String type){
    return chatRepository.sendMessage(chatId,message,type);
  }
}