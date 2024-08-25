import 'package:dartz/dartz.dart';
import 'package:ecommerce/core/error/failure.dart';
import 'package:ecommerce/features/chat_feature/chat/domain/repository/chat_repository.dart';

class DeleteChat {
  final ChatRepository repository;

  DeleteChat({required this.repository});

  Future<Either<Failure, void>> call(String chatId) async {
    return await repository.deleteMessage(chatId);
  }
}