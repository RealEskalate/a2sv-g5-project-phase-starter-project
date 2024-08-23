import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entities/user_entity.dart';
import '../repositories/auth_repository.dart';

class GetUserUsecase {
  final AuthRepository repository;

  GetUserUsecase({required this.repository});

  Future<Either<Failure, UserEntity>> call() async {
    return await repository.getUser();
  }
}
