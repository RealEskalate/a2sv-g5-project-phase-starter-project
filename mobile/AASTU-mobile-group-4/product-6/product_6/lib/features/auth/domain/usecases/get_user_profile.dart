import 'package:dartz/dartz.dart';

import '../../../../core/errors/failure.dart';
import '../../../../core/usecases/auth/usecases.dart';
import '../entity/auth_entity.dart';
import '../repository/auth_repository.dart';


class GetUserProfile implements UseCase<UserEntity, void> {
  final AuthRepository repository;

  GetUserProfile(this.repository);

  @override
  Future<Either<Failure, UserEntity>> call(void params) async {
    return await repository.getUserProfile();
  }
}


