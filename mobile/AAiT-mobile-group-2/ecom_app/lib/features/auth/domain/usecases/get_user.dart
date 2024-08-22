import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/user_data_entity.dart';
import '../repositories/auth_repository.dart';

class GetUserUsecase extends UseCase<UserDataEntity, NoParams> {
  final AuthRepository authRepository;
  GetUserUsecase(this.authRepository);


  @override
  Future<Either<Failure, UserDataEntity>> call(NoParams params) {
    return authRepository.getUser();
  }

}