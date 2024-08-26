import 'package:dartz/dartz.dart';
import 'package:ecommerce_app_ca_tdd/core/errors/failure/failures.dart';
import 'package:ecommerce_app_ca_tdd/core/usecases/usecases.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/data/models/chat_models.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/chat_entity.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/repository/repository.dart';
import 'package:equatable/equatable.dart';

class SendUseCase implements UseCase<Unit,SendUseCaseParams>{
  final ChatRepository repository;
  const SendUseCase(this.repository);
  @override
  Future<Either<Failure, Unit>> call(SendUseCaseParams params) async {
    return await repository.sendChat(params.chat, params.content, params.type);
  }
}

class SendUseCaseParams extends Equatable {
  final String chat;
  final String content;
  final String type;

  const SendUseCaseParams(this.chat, this.content, this.type);

  @override
  List<Object?> get props => [chat, content, type];
}
