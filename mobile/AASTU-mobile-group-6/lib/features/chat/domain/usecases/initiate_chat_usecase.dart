import 'package:dartz/dartz.dart';
import 'package:ecommerce_app_ca_tdd/core/errors/failure/failures.dart';
import 'package:ecommerce_app_ca_tdd/core/usecases/usecases.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/repository/repository.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/entities/product_entity.dart';

class InitiateChatUsecase  implements UseCase<String, String> {
  final ChatRepository abstractChatRepository;

  InitiateChatUsecase(this.abstractChatRepository);

  @override
  Future<Either<Failure, String>> call(String userId)async {
    return await abstractChatRepository.initiateChat(userId); 

  }
}
