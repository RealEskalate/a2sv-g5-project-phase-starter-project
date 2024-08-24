import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/chat_entity.dart';
import '../repositories/chat_repository.dart';

class MyChatsUseCase extends UseCase<List<ChatEntity>, NoParams> {
  final ChatRepository chatRepository;

  MyChatsUseCase(this.chatRepository);

  @override
  Future<Either<Failure, List<ChatEntity>>> call(NoParams params) async {
    return await chatRepository.myChats();
  }
}
