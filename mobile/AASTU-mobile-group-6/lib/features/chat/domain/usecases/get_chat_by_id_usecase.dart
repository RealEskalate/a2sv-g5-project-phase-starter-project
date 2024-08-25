import 'package:dartz/dartz.dart';
import 'package:ecommerce_app_ca_tdd/core/errors/failure/failures.dart';
import 'package:ecommerce_app_ca_tdd/core/usecases/usecases.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/chat_entity.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/repository/repository.dart';


class GetChatByIdUsecase  implements UseCase<ChatEntity, String> {
  final ChatRepository abstractChatRepository;

  GetChatByIdUsecase(this.abstractChatRepository);

  @override
  Future<Either<Failure, ChatEntity>> call(String chatId)async {
    return await abstractChatRepository.getMyChatById(chatId); 

  }
}