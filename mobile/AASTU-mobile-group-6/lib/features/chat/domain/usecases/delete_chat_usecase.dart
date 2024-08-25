import 'package:dartz/dartz.dart';
import 'package:ecommerce_app_ca_tdd/core/errors/failure/failures.dart';
import 'package:ecommerce_app_ca_tdd/core/usecases/usecases.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/repository/repository.dart';


class DeleteChatUsecase  implements UseCase<String, String> {
  final ChatRepository abstractChatRepository;

  DeleteChatUsecase(this.abstractChatRepository);

  @override
  Future<Either<Failure, String>> call(String chatId)async {
    return await abstractChatRepository.deleteChat(chatId); 

  }
}