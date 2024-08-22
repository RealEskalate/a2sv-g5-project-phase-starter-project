import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../repositories/auth_repo.dart';

class LogOutUsecase extends UseCase<void, NoParams> {

  final AuthRepository authRepository;
  
  LogOutUsecase({required this.authRepository});

  @override
  Future<Either<Failure, void>> call(NoParams p) async {
    return await authRepository.logOut();
  }
}