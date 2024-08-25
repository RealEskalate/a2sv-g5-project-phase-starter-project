import 'package:dartz/dartz.dart';
import 'package:ecommerce_app_ca_tdd/core/errors/failure/failures.dart';
import 'package:ecommerce_app_ca_tdd/core/usecases/usecases.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/chat_entity.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/repository/repository.dart';


class GetChatsUsecase  implements UseCase<List<ChatEntity>, NoParams> {
  final ChatRepository abstractChatRepository;

  GetChatsUsecase(this.abstractChatRepository);

  @override
  Future<Either<Failure, List<ChatEntity>>> call(NoParams params)async {
    return await abstractChatRepository.getAllChats(); 

  }
}