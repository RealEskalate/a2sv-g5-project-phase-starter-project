


import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../repositories/chat_repository.dart';

class DeleteChat{
  final ChatRepository chatRepository;
  DeleteChat(this.chatRepository);
  Future<Either<Failure,void >>execute(String id){
    return chatRepository.deleteChat(id);
  }
}

