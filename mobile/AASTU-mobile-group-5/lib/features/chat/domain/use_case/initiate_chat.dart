


import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../entities/chat_entity.dart';
import '../repositories/chat_repository.dart';

class InitiateChat{
  final ChatRepository chatRepository;
  InitiateChat(this.chatRepository);
  Future<Either<Failure, ChatEntity>>execute(String sellerId){
      return chatRepository.initiateChat(sellerId);
  }

}