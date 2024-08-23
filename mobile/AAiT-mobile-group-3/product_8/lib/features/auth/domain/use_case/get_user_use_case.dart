import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/user_data_entity.dart';
import '../repositories/auth_repository.dart';

class GetUserUseCase extends UseCase<UserDataEntity, NoParams> {
  final AuthRepository authRepository;

  GetUserUseCase(this.authRepository);

  @override
  Future<Either<Failure, UserDataEntity>> call(NoParams params) async {
    return await authRepository.getUser();
  }
}