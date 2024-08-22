import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/user_data.dart';
import '../repositories/auth_repo.dart';

class GetCurrentUserUsecase extends UseCase<UserEntity, NoParams> {

  final AuthRepository authRepository;
  
  GetCurrentUserUsecase({required this.authRepository});

  @override
  Future<Either<Failure, UserEntity>> call(NoParams p) async {
    return await authRepository.getCurrentUser();
  }
}
