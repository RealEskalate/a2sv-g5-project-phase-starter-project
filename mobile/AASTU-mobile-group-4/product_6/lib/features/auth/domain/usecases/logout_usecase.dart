



//logout usecase
import 'package:dartz/dartz.dart';

import '../../../../core/errors/failure.dart';
import '../../../../core/usecases/auth/usecases.dart';
import '../repository/auth_repository.dart';

class LogoutUseCase extends UseCase<void, void> {
  final AuthRepository repository;

  LogoutUseCase(this.repository);
   @override
  Future<Either<Failure, void>> call(void params) async {
    return repository.logout();
  }
}
