import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/error/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/message.dart';
import '../repositories/chat_repository.dart';

class SendMessageUsecase extends UseCase<void, SendMessageParams> {
  final ChatRepository repository;
  SendMessageUsecase(this.repository);

  @override
  Future<Either<Failure, void>> call(SendMessageParams p) async {
    return repository.sendMessage(p.message);
  }
}

class SendMessageParams extends Equatable {
  final Message message;

  const SendMessageParams(this.message);
  
  @override
  List<Object?> get props => [message];
}