import 'package:dartz/dartz.dart';

import '../../../../core/errors/failure.dart';
import '../../../../core/usecases/auth/usecases.dart';
import '../entity/auth_entity.dart';
import '../repository/auth_repository.dart';

class Register implements UseCase<void, SentUserEntity> {
  final AuthRepository repository;

  Register(this.repository);

  @override
  Future<Either<Failure, void>> call(SentUserEntity senduserentity) async {
    return await repository.register(senduserentity);
  }
}
