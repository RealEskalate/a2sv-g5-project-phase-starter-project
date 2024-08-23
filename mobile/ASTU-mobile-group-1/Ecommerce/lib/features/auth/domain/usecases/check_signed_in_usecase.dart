import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../repositories/auth_repository.dart';

class CheckSignedInUsecase {
  final AuthRepository repository;

  CheckSignedInUsecase({required this.repository});

  Future<Either<Failure, bool>> call() async {
    return await repository.checkSignedIn();
  }
}
