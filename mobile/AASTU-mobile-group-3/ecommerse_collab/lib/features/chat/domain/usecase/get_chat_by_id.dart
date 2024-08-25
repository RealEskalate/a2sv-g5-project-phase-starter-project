import 'package:dartz/dartz.dart';
import '../../../../core/error/failure.dart';
import '../entity/chat.dart';
import '../repository/chat_repository.dart';

class GetMessageByIdUseCase {
  late ChatRepository chatRepository;
  GetMessageByIdUseCase(this.chatRepository);

  Future<Either<Failure,Chat>> execute(String id) async{
    final chat = await chatRepository.chatById(id);
    return Right(chat);

    
  }
}
