import 'package:dartz/dartz.dart';


import '../../../../core/base_usecase.dart';
import '../../../../core/error/failures.dart';
import '../entities/sign_up_entity.dart';
import '../repository/auth_repo.dart';

class SignUpUseCase implements UseCase<String, UseCaseParams> {
  final AuthRepository authRepository;
  SignUpUseCase(this.authRepository);

  @override
  Future<Either<Failure, String>> call(UseCaseParams params) async {
    return await authRepository.signUP(params.signUpEntity);
  }
  
  @override
  Future<Either<Failure, String>> execute(UseCaseParams params) {
    // TODO: implement execute
    throw UnimplementedError();
  }
}

class UseCaseParams {
  final SignUpEntity signUpEntity;
  UseCaseParams(this.signUpEntity);
}
