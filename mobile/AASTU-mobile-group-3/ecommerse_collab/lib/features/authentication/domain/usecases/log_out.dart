import '../repository/authentication_repository.dart';

class LogOutUseCase {
  final AuthenticationRepository _repository;

  LogOutUseCase(this._repository);

  Future<void> call() async {
    return await _repository.logOut();
  }
}